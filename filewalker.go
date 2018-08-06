package main

import (
	"os"
	"path/filepath"
	"strings"
)

type FileConf struct {
	Filename string
	Filepath string
}

func WalkFiles(path string) []FileConf {
	filelist := []FileConf{}
	err := filepath.Walk(
		path,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if strings.HasSuffix(info.Name(), ".go") {
				filelist = append(filelist, FileConf{
					Filename: path,
					Filepath: filepath.Dir(path),
				})
			}
			return nil
		})

	if err != nil {
		panic(err)
	}

	return filelist
}
