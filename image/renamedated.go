package image

import (
    "os"
    "log"
    "strconv"
    "github.com/rwcarlsen/goexif/exif"
    "path/filepath"
    "github.com/ashbyp/goutil/fileutil"
)

func RenameAllDated(targpath string, patterns ...string) (results map[string]string, err error) {
    results = make(map[string]string)
    var matches []string

    matches, err = fileutil.Glob(patterns...)
    if err != nil {
        log.Print("Failed to glob")
        return
    }

    if len(matches) > 0 {
        if ! fileutil.FileExists(targpath) {
            err = os.Mkdir(targpath, 0777)
            if err != nil {
                log.Print("Failed to create targpath " + targpath)
                return
            }
        }
    }

    for _, filename := range matches {
        var newname string
        newname, err = RenameDated(targpath, filename)

        if err == nil {
            results[filename] = newname
        }
    }
    return
}

func RenameDated(targpath, filename string) (newname string, err error) {
    var f *os.File
    f, err = os.Open(filename)
    
    if err != nil {
        log.Print(err)
    } else {
        defer f.Close()

        var x *exif.Exif
        x, err = exif.Decode(f)

        if err != nil {
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

            err = fileutil.CopyFile(filename, newname)
            if err != nil {
                log.Print(err)
            }
        }
    }
    return
}











