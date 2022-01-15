package main

import (
	"archive/tar"
	"archive/zip"
	"compress/bzip2"
	"compress/gzip"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

func main() {
	log.SetFlags(0)
	if len(os.Args) == 1 || os.Args[1] == "-h" || os.Args[1] == "--help" {
		fmt.Printf("usage: %s archive.{zip,tar,tar.gz,tar.bz2}\n",
			filepath.Base(os.Args[0]))
		os.Exit(1)

	}
	filename := os.Args[1]
	if !validSuffix(filename) {
		log.Fatalln("unrecognized archive suffix")
	}
	if err := unpackArchive(filename); err != nil {
		log.Fatalln(err)
	}
}

// validSuffix validates archive extension
func validSuffix(filename string) bool {
	for _, suffix := range []string{".zip", ".tar", ".tar.gz", ".tar.bz2"} {
		if strings.HasSuffix(filename, suffix) {
			return true
		}
	}
	return false
}

// unpackArchive calls unpack functions for both types of archives
func unpackArchive(filename string) error {
	if strings.HasSuffix(filename, ".zip") {
		return unpackZip(filename)
	}
	return unpackTar(filename)
}

// unpackZip unzips file's structure
func unpackZip(filename string) error {
	reader, err := zip.OpenReader(filename)
	if err != nil {
		return err
	}
	defer reader.Close()
	for _, zipFile := range reader.Reader.File {
		name := sanitizedName(zipFile.Name)
		mode := zipFile.Mode()
		if mode.IsDir() {
			if err = os.MkdirAll(name, 0755); err != nil {
				return err
			}
		} else {
			if err = unpackZippedFile(name, zipFile); err != nil {
				return err
			}
		}
	}
	return nil
}

// unpackZippedFile unzips file
func unpackZippedFile(filename string, zipFile *zip.File) error {
	writer, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer writer.Close()
	reader, err := zipFile.Open()
	if err != nil {
		return err
	}
	defer reader.Close()
	if _, err = io.Copy(writer, reader); err != nil {
		return err
	}
	if filename == zipFile.Name {
		fmt.Println(filename)
	} else {
		fmt.Printf("%s [%s]\n", filename, zipFile.Name)
	}
	return nil
}

// unpackTar unpacks tar-files with inner structure
func unpackTar(filename string) (err error) {
	var file *os.File
	if file, err = os.Open(filename); err != nil {
		return err
	}
	defer file.Close()
	var fileReader io.Reader = file
	var decompressor *gzip.Reader
	if strings.HasSuffix(filename, ".gz") {
		if decompressor, err = gzip.NewReader(file); err != nil {
			return err
		}
		defer decompressor.Close()
	} else if strings.HasSuffix(filename, ".bz2") {
		fileReader = bzip2.NewReader(file)
	}
	var reader *tar.Reader
	if decompressor != nil {
		reader = tar.NewReader(decompressor)
	} else {
		reader = tar.NewReader(fileReader)
	}
	return unpackTarFiles(reader)
}

// unpackTarFiles going through catalog's structure, creates folders if they are and unpacks tar-files
func unpackTarFiles(reader *tar.Reader) error {
	for {
		header, err := reader.Next()
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
		filename := sanitizedName(header.Name)
		switch header.Typeflag {
		case tar.TypeDir:
			if err = os.MkdirAll(filename, 0755); err != nil {
				return err
			}
		case tar.TypeReg:
			if err = unpackTarFile(filename, header.Name, reader); err != nil {
				return err
			}
		}
	}
	return nil
}

// unpackTarFile unpacks tar-file
func unpackTarFile(filename, tarFilename string, reader *tar.Reader) (err error) {
	var writer *os.File
	if writer, err = os.Create(filename); err != nil {
		return err
	}
	defer writer.Close()
	if _, err = io.Copy(writer, reader); err != nil {
		return err
	}
	if filename == tarFilename {
		fmt.Println(filename)
	} else {
		fmt.Printf("%s [%s]\n", filename, tarFilename)
	}
	return nil
}

// sanitizedName returns pure filepath
func sanitizedName(filename string) string {
	if len(filename) > 1 && filename[1] == ':' &&
		runtime.GOOS == "windows" {
		filename = filename[2:]
	}
	filename = filepath.ToSlash(filename)
	filename = strings.TrimLeft(filename, "/.")
	return strings.Replace(filename, "../", "", -1)
}
