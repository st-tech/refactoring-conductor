package internal

import (
	"io/ioutil"
	"path/filepath"
)

// DirectoryInternalFiles is to find files inside the directory.
func DirectoryInternalFiles(dir string) []string {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(err)
	}

	var paths []string
	for _, file := range files {
		if file.IsDir() {
			paths = append(paths, DirectoryInternalFiles(filepath.Join(dir, file.Name()))...)
			continue
		}
		paths = append(paths, filepath.Join(dir, file.Name()))
	}

	return paths
}
