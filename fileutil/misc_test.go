package fileutil

import (
    "testing"
    "path/filepath"
)

func TestFileExists(t *testing.T) {
    exists := FileExists("misc_test.go")
    if !exists {
        t.Error("misc_test.go does not exist")
    }    
}


func TestGlob(t *testing.T) {
    files, _ := Glob("*.go")
    if len(files) == 0 {
        t.Error("No .go files in current dir found")
    }
}

func TestFindDirs(t *testing.T) {
    dirs, _ := FindDirs("..", ".git")
    var found bool = false
    for _, d := range dirs {
        if filepath.Base(d) == "fileutil" {
            found = true
            break
        }
    }
    if !found {
        t.Error("fileutil dir not found")
    }
}
