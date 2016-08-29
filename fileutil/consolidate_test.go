package fileutil

import (
    "testing"
    "path/filepath"
    "os"
)

func TestConsolidate(t *testing.T) {
    targetsfound, _, _ := Consolidate("go_files", "not_go_files", ".go")
    defer os.RemoveAll("go_files")
    defer os.RemoveAll("not_go_files")
    if targetsfound == 0  {
         t.Error("No targets were found")
    }
    if !FileExists("go_files" + string(filepath.Separator) + "consolidate_test.go") {
         t.Error("consolidate_test.go not copied")
    }
}
