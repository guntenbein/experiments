package main

import (
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"os/exec"
	"fmt"
)

func main() {
	//wd, err := os.Getwd()
	//if err != nil {
	//	panic(err)
	//}
	//dirs := DirsWith(wd, ".go")
	currentDir := "/Users/anovikau/go/src/bitbucket.org/inturnco/documentservice"
	dirs := DirsWith(currentDir, ".go")

	for _,d := range dirs {
		fmt.Println("start building " + d)
		cmd := exec.Command("go", "build", strings.Replace(d, currentDir, ".",1))
		out, err := cmd.Output()
		if err != nil {
			fmt.Println(fmt.Printf("building %s finished with error %s", d, err.Error()))
			fmt.Printf(string(out))
			os.Exit(1)
		} else {
			fmt.Println(fmt.Printf("building %s finished ok", d))
		}
	}
}

func DirsWith(root, mask string) []string {
	var dirs []string
	filepath.Walk(root, func(path string, f os.FileInfo, _ error) error {
		if !f.IsDir() && !strings.Contains(path,"/vendor/") {
			r, err := regexp.MatchString(mask, f.Name())
			if err == nil && r {
				d := filepath.Dir(path)
				if !contains(dirs, d) {
					fmt.Println(d)
					dirs = append(dirs, filepath.Dir(path))
				}
			}
		}
		return nil
	})
	return dirs
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}