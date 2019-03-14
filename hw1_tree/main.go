package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

const (
	middleFile = "├───"
	lastFile   = "└───"
	fileLine   = "│\t"
	separator  = string(os.PathSeparator)
	tab        = "\t"
)

func main() {
	out := new(bytes.Buffer)
	if !(len(os.Args) == 2 || len(os.Args) == 3) {
		panic("usage go run main.go . [-f]")
	}
	path := os.Args[1]
	printFiles := len(os.Args) == 3 && os.Args[2] == "-f"
	err := dirTree(out, path, printFiles)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(fmt.Sprintf("%s", out))
}

func dirTree(out *bytes.Buffer, path string, printFiles bool) error {
	return dirTreeInner(out, path, printFiles, "")
}

func dirTreeInner(out *bytes.Buffer, path string, printFiles bool, previousStart string) error {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return err
	}

	for i, file := range files {
		//prepare text
		text := previousStart + middleFile + file.Name()
		isLast := isLastFile(i, files, printFiles)
		if isLast {
			text = previousStart + lastFile + file.Name()
		}

		fi, err := os.Stat(path + separator + file.Name())
		if err != nil {
			return err
		}

		if file.IsDir() {
			//print file
			print(out, text, fi.Size(), printFiles, true)

			//calculate next line start
			nextStart := previousStart
			if !isLast {
				nextStart += fileLine
			} else {
				nextStart += tab
			}
			//print childs
			err := dirTreeInner(out, path+separator+file.Name(), printFiles, nextStart)
			if err != nil {
				return err
			}
		} else if printFiles {
			//print file
			print(out, text, fi.Size(), printFiles, false)
		}

	}
	return nil
}

func print(out *bytes.Buffer, text string, size int64, printFiles bool, isFolder bool) {
	if printFiles {
		sizeText := " (empty)"
		if size > 0 {
			sizeText = fmt.Sprintf(" (%sb)", strconv.FormatInt(size, 10))
		}
		if isFolder {
			sizeText = ""
		}
		out.Write([]byte(fmt.Sprintf("%s%s\n", text, sizeText)))
	} else {
		out.Write([]byte(fmt.Sprintf("%s\n", text)))
	}
}

func isLastFile(currentIndex int, files []os.FileInfo, printFiles bool) bool {
	if currentIndex+1 == len(files) {
		return true
	}
	if printFiles {
		return false
	}
	for i := currentIndex + 1; i < len(files); i++ {
		file := files[i]
		if file.IsDir() {
			return false
		}
	}
	return true
}
