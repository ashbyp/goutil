package fileutil

import (
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

