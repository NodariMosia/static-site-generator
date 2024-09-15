package main

import (
	"static-site-generator/pkg/fileutils"
	"static-site-generator/pkg/generator"
	"static-site-generator/pkg/server"
)

func main() {
	const (
		sourceDir      = "./static"
		destinationDir = "./public"

		sourcePath      = "./content/index.md"
		templatePath    = "./template.html"
		destinationPath = "./public/index.html"

		serverPort = ":8888"
	)

	panicIfError(fileutils.CleanAndCopyFromSourceDirToDestinationDir(sourceDir, destinationDir))

	panicIfError(generator.GeneratePage(sourcePath, templatePath, destinationPath))

	panicIfError(server.ServeDir(destinationDir, serverPort))
}

func panicIfError(err error) {
	if err != nil {
		panic(err)
	}
}
