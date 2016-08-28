package main

import (
    "fmt"
    "flag"
    "github.com/ashbyp/goutil/image"
)

func main() {
    cmdptr := flag.String("cmd", "", "command to execute <(r)ename>")
    flag.Parse()

    switch *cmdptr {
    case "rename", "r":
        fmt.Println("Renaming image files according to date picture was taken")
        results, err := image.RenameAllDated("renamed_pix", "*.JPG", "*.jpg")
        if err != nil {
            fmt.Println(err)
        }
        for k, v := range results {
            fmt.Printf("%s\t--->\t%s\n", k, v)
        }
    default:
        flag.Usage()
    }
}
