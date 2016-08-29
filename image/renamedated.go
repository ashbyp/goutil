package image

import (
	"github.com/ashbyp/goutil/fileutil"
	"github.com/rwcarlsen/goexif/exif"
	"log"
	"os"
	"path/filepath"
	"strconv"
)

func RenameAllDated(targpath string, patterns ...string) (results map[string]string, err error) {
	results = make(map[string]string)
	var matches []string

	if matches, err = fileutil.Glob(patterns...); err != nil {
		log.Print("Failed to glob")
		return
	}

	if len(matches) > 0 {
		if !fileutil.FileExists(targpath) {
			if err = os.Mkdir(targpath, 0777); err != nil {
				log.Print("Failed to create targpath " + targpath)
				return
			}
		}
	}

	for _, filename := range matches {
		var newname string
		if newname, err = RenameDated(targpath, filename); err == nil {
			results[filename] = newname
		}
	}
	return
}

func RenameDated(targpath, filename string) (newname string, err error) {
	var f *os.File
	if f, err = os.Open(filename); err != nil {
		log.Print(err)
	} else {
		defer f.Close()

		var x *exif.Exif

		if x, err = exif.Decode(f); err != nil {
			log.Printf("%s %s", filename, err)
		} else {
			sep := string(filepath.Separator)
			tm, _ := x.DateTime()
			if tm.Year() == 1 {
				log.Printf("Found year == 1 for %s", filename)

			}

			newname = targpath + sep + tm.Format("20060102_15-04-05") + ".jpg"

			for i := 1; fileutil.FileExists(newname); i++ {
				newname = targpath + sep + tm.Format("20060102_15-04-05") + "_" + strconv.Itoa(i) + ".jpg"
			}

			if err = fileutil.CopyFile(filename, newname); err != nil {
				log.Print(err)
			}
		}
	}
	return
}
