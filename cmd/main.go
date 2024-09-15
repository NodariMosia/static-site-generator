package main

import (
	"static-site-generator/pkg/fileutils"
	"static-site-generator/pkg/generator"
	"static-site-generator/pkg/server"
)

const (
	staticDir      = "./static"
	contentDir     = "./content"
	destinationDir = "./public"

	templatePath = "./template.html"

	serverPort = ":8888"
)

func main() {
	panicIfError(fileutils.CleanAndCopyFromSourceDirToDestinationDir(staticDir, destinationDir))

	panicIfError(generator.GeneratePagesRecursive(contentDir, templatePath, destinationDir))

	panicIfError(server.ServeDir(destinationDir, serverPort))
}

func panicIfError(err error) {
	if err != nil {
		panic(err)
	}
}
