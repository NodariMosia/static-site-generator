package main

import "static-site-generator/pkg/fileutils"

const (
	sourceDir      = "./static"
	destinationDir = "./public"
)

func main() {
	err := fileutils.CleanAndCopyFromSourceDirToDestinationDir(sourceDir, destinationDir)
	if err != nil {
		panic(err)
	}
}
