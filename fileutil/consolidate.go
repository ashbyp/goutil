package fileutil

import (
	"log"
	"os"
	"path/filepath"
	"strings"
)

func Consolidate(targpath string, otherpath string, extensions ...string) (targetsfound, othersfound int, err error) {
	exclude := []string{targpath, otherpath}
	sep := string(filepath.Separator)

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
			} else {
				for _, e := range extensions {
					if strings.HasSuffix(path, e) {
						err = CopyFileWithClashDetection(path, targpath+sep+filepath.Base(path))
						if err != nil {
							return err
						}
						targetsfound += 1
						return err
					}
				}
				err = CopyFileWithClashDetection(path, otherpath+sep+filepath.Base(path))
				if err != nil {
					return err
				}
				othersfound += 1
				return err
			}
		}
		return err
	}

	if !FileExists(targpath) {
		err = os.Mkdir(targpath, 0777)
		if err != nil {
			log.Print("Failed to create " + targpath)
			return
		}
	}
	if !FileExists(otherpath) {
		err = os.Mkdir(otherpath, 0777)
		if err != nil {
			log.Print("Failed to create " + otherpath)
			return
		}
	}

	err = filepath.Walk(".", f)
	return
}
