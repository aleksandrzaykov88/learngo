package fileutil

import (
	"io/fs"
	"path/filepath"

	"github.com/aleksandrzaykov88/learngo/Other/arrayutil"

	"github.com/gabriel-vasile/mimetype"
)

// GetFileList goes through catalog structure and seraching files
func GetFileList(root string, validExtesions []string) ([]string, error) {
	var files []string
	err := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if !d.IsDir() { // Skip dirs
			mtype, err := mimetype.DetectFile(path) // Get fileinfo
			if err != nil {
				return err
			}

			// Validates file's extension according config file
			isExtValid := arrayutil.StringSliceContains(validExtesions, mtype.Extension())

			if isExtValid {
				files = append(files, path)
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return files, nil
}
