package main

import (
	"fmt"
	//"io"
	"bytes"
	"io/ioutil"
	"os"
	"strconv"
	//"path/filepath"
	//"strings"
)

const file = "├───"
const lastFile = "└───"
const fileLevelSep = "│\t"

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
	return dirTree1(out, path, printFiles, "")
}

func dirTree1(out *bytes.Buffer, path string, printFiles bool, st string) error {
	separator := string(os.PathSeparator)

	files, err := ioutil.ReadDir(path)
	if err != nil {
		return err
	}

	for i, f := range files {
		startText := file
		isLast := i+1 == len(files)
		nextStart := st
		if isLast {
			startText = lastFile
			nextStart = st + "\t"
		} else {
			nextStart = st + fileLevelSep
		}
		fi, err := os.Stat(path + separator + f.Name())
		if err != nil {
			return err
		}

		if f.IsDir() {
			print(out, st, startText, f.Name(), fi.Size(), printFiles, true)
			err := dirTree1(out, path+separator+f.Name(), printFiles, nextStart)
			if err != nil {
				return err
			}
		} else if printFiles {
			print(out, st, startText, f.Name(), fi.Size(), printFiles, false)
		}

	}
	return nil
}

func print(out *bytes.Buffer, start string, startText string, name string, size int64, printFiles bool, isFolder bool) {
	if printFiles {
		sizeText := " (empty)"
		if size > 0 {
			sizeText = " (" + strconv.FormatInt(size, 10) + "b)"
		}
		if isFolder {
			sizeText = ""
		}
		out.Write([]byte("\n"))
		out.Write([]byte(fmt.Sprintf("%s%s%s%s", start, startText, name, sizeText)))
	} else {
		out.Write([]byte("\n"))
		out.Write([]byte(fmt.Sprintf("%s%s%s", start, startText, name)))
	}
}
