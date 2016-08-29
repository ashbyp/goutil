package fileutil

import (
	"log"
	"os"
	"path/filepath"
)

func FileExists(path string) (exists bool) {
	_, err := os.Stat(path)
	return err == nil
}

func Glob(patterns ...string) (matches []string, err error) {
	var m []string
	for _, pattern := range patterns {
		m, err = filepath.Glob(pattern)
		if err != nil {
			return
		}
		matches = append(matches, m...)
	}
	return
}

func FindDirs(root string, exclude ...string) (dirs []string, err error) {
	f := func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Print(err)
		} else {
			if info.IsDir() {
				dir := filepath.Base(path)
				for _, d := range exclude {
					if d == dir {
						return filepath.SkipDir
					}
				}
				dirs = append(dirs, path)
			}
		}
		return err
	}
	err = filepath.Walk(root, f)
	return dirs, err
}
