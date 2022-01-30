package fileutil

import (
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// AddToLogFile adds line to log-file
func AddToLogFile(path string, fn string, line string) error {
	var dir string
	if path != "" {
		dir = strings.TrimSpace(path)
	}
	dir, err := resolveLogPath(dir)

	if err != nil {
		return errors.New("could not find a directory to save the log file")
	}

	f, err := os.OpenFile(dir+string(os.PathSeparator)+fn, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		return err
	}
	defer f.Close()

	if _, err = f.WriteString("\n" + line); err != nil {
		return err
	}
	return nil
}

// resolveLogPath check log path from config file
// if path is empty, log file will be save in workdir
func resolveLogPath(path string) (string, error) {
	if filepath.IsAbs(path) {
		return path, nil
	}

	wd, err := os.Getwd()
	if err != nil {
		return "", err
	}

	if path == "" {
		return wd, nil
	}

	return filepath.Join(wd, path), nil
}

// IsLogExist check log-file existance with `fn` name in `path` directory
func IsLogExist(fn string, path *string) (bool, error) {
	var dir string
	if fn == "" {
		return false, errors.New("log name not passed")
	}
	if path != nil {
		dir = strings.TrimSpace(*path)
	}

	dir, err := resolveLogPath(dir)
	if err != nil {
		return false, errors.New("could not find directory to check log file")
	}

	_, err = os.Stat(dir + fn)
	if os.IsNotExist(err) {
		return false, nil
	}

	return true, nil
}

// SaveLog saves log file
func SaveLog(file []byte, path *string, fn string) error {
	if file == nil {
		return errors.New("log file not passed")
	}

	var dir string
	if path != nil {
		dir = strings.TrimSpace(*path)
	}

	dir, err := resolveLogPath(dir)
	if err != nil {
		return errors.New("could not find directory to save log file")
	}

	if err := os.MkdirAll(dir, 0644); err != nil {
		return errors.New("could not create directory to check log file")
	}

	err = ioutil.WriteFile(filepath.Join(dir, fn), file, 0644)
	if err != nil {
		return errors.New("could not save log file")
	}

	return nil
}
