package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"
	"text_changer/fileutil"
	clog "text_changer/log"
	"time"

	"github.com/spf13/viper"
)

var (
	toChange       string   // text to be replaced
	replacer       string   // replacement text
	root           string   // path to file's catalog
	workers        int      // pool of workers
	validExtesions []string // permissible file extensions
)

func init() {
	t := time.Now()

	if err := initConfig(); err != nil { // get params from config file
		log.Fatalf("error initializing configs: %s", err.Error())
	}

	log := clog.NewConfig() // initialize log-object
	log.Path = viper.GetString("log.path")
	log.Name = "log_" + t.Format("20060102150405") + ".log"

	validExtesions = viper.GetStringSlice("ext")

	workers = runtime.NumCPU()
}

// initConfig reads the configuration file
func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")

	return viper.ReadInConfig()
}

func main() {
	clog.Print(clog.INFO, "Text Changer started")

	runtime.GOMAXPROCS(runtime.NumCPU())

	// initialization of global variables
	if len(os.Args) >= 4 {
		root = os.Args[1]
		toChange = os.Args[2]
		replacer = os.Args[3]
	}

	if root == "" || toChange == "" || replacer == "" {
		// If programs starts in Docker container
		root = os.Getenv("ROOT")
		toChange = os.Getenv("TOCHANGE")
		replacer = os.Getenv("REPLACER")
	}

	if root == "" || toChange == "" || replacer == "" {
		clog.Print(clog.FATAL, "error while getting args")
	}

	// creating jobs object
	jobs := make(chan string, workers*16)
	done := make(chan struct{}, workers)

	files, err := fileutil.GetFileList(root, validExtesions) // Getting all filenames
	if err != nil {
		clog.Print(clog.FATAL, "error while getting files list. Check root variable.")
	}

	go addJobs(files, jobs) // Pushing filenames to job-channel
	for i := 0; i < workers; i++ {
		go doJobs(done, jobs) // Work is start
	}
	waitUntil(done) // waiting for work to finish

	clog.Print(clog.INFO, "Text Changer finished")
}

// addJobs pushes filenames to jobs-channel and after closes it
func addJobs(files []string, jobs chan<- string) {
	for _, filename := range files {
		jobs <- filename
	}
	close(jobs)
}

// doJobs calls replaceText func for each filename and signals the end of work
func doJobs(done chan<- struct{}, jobs <-chan string) {
	for job := range jobs {
		replaceText(job)
	}
	done <- struct{}{}
}

// replaceText replace text in file from `toChange` to `Replacer`
func replaceText(filename string) {
	raw, err := ioutil.ReadFile(filename)
	if err != nil {
		clog.Print(clog.WARNING, "failed to read: ", filename, " error: ", err)
		return
	}

	fileText := string(raw) // Getting file's content

	err = backupOldFile(filename, &fileText) // Creates old file backup
	if err != nil {
		clog.Print(clog.WARNING, "error while backuping ", "old_"+filename, " error: ", err)
	}

	newText := strings.ReplaceAll(fileText, toChange, replacer) // Replacing `toChange` to `Replacer` in whole file

	file, err := os.Create(filename) // Create (and rewrite) target file
	if err != nil {
		clog.Print(clog.WARNING, "couldn't update ", filename, " error: ", err)
		return
	}
	defer file.Close()
	if _, err := file.WriteString(newText); err != nil {
		clog.Print(clog.WARNING, "error when updating ", filename, " error: ", err)
	}

	err = logger(filename, fileText, toChange, replacer) // Writing logs
	if err != nil {
		clog.Print(clog.WARNING, "couldn't log ", filename, " error: ", err)
		return
	}
}

// waitUntil checks finished jobs
func waitUntil(done <-chan struct{}) {
	for i := 0; i < workers; i++ {
		<-done
	}
}

// backupOldFile backups old file if there is corresponding condition in config
func backupOldFile(filename string, fileText *string) error {
	var extension = filepath.Ext(filename)
	var newName = filename[0:len(filename)-len(extension)] + "_old" + extension
	if viper.GetBool("backup") {
		file, err := os.Create(newName)
		if err != nil {
			return err
		}
		defer file.Close()
		if _, err := file.WriteString(*fileText); err != nil {
			return err
		}
	}

	return nil
}

// logger writes logs
func logger(filename, fileText, toChange, replacer string) error {
	lines := strings.Split(fileText, "\n") // Splitting file on strings

	for i, l := range lines { // Iterate through strings
		re, err := regexp.Compile(`.{0,5}` + toChange + `.{0,5}`) // Replacement text and a few characters before and after
		if err != nil {
			return err
		}

		res := re.FindAllString(l, -1)            // Finding all strings by regexp
		positions := re.FindAllStringIndex(l, -1) // Finding corresponding positions

		if len(res) > 0 {
			for j, matchLine := range res { // Iterating through matches and creating new log-message for each one
				border := "\n============================================="
				fileInfo := fmt.Sprintf("\n File: %s\n", filename)
				lineInfo := fmt.Sprintf("Line: %d\n", i+1)
				positionInfo := fmt.Sprintf("Position: %d\n", positions[j][0]+1)
				oldText := fmt.Sprintf("Old text: %s\n", matchLine)

				// Since all lines in the file change together at once, this little trick allows to get the `after text`
				replaceLine := strings.ReplaceAll(l, toChange, replacer)
				re, err := regexp.Compile(`.{0,5}` + replacer + `.{0,5}`)
				if err != nil {
					return err
				}
				changeRes := re.FindAllString(replaceLine, -1)

				newText := fmt.Sprintf("New text: %s", changeRes[j])
				clog.Print(clog.INFO, border, fileInfo, lineInfo, positionInfo, oldText, newText, border)
			}
		}
	}

	return nil
}
