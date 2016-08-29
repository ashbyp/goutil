package main

import (
	"flag"
	"fmt"
	"github.com/ashbyp/goutil/fileutil"
	"github.com/ashbyp/goutil/image"
)

func main() {
	cmdptr := flag.String("cmd", "", "command to execute <(r)ename>|<(c)onsolidate>")
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
	case "consolidate", "c":
		fmt.Println("Consolidating all images to ./all_pix and others to ./all_not_pix")
		pix, not_pix, err := fileutil.Consolidate("all_pix", "all_not_pix", ".JPG", ".jpg")
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("Found %d pictures and %d non-pictures", pix, not_pix)
	default:
		flag.Usage()
	}
}
