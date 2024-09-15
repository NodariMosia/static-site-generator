package generator

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"static-site-generator/pkg/adapters"
	"static-site-generator/pkg/markdown"
)

func GeneratePagesRecursive(contentDir, templatePath, destinationDir string) error {
	fmt.Println("Generating pages...")

	handleWalkDirEntry := func(path string, entry os.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if entry.IsDir() {
			return nil
		}

		if !strings.HasSuffix(entry.Name(), ".md") {
			return nil
		}

		relativePath, err := filepath.Rel(contentDir, path)
		if err != nil {
			return err
		}

		destinationPath := filepath.Join(destinationDir, relativePath)
		destinationPath = strings.TrimSuffix(destinationPath, ".md")
		destinationPath += ".html"

		return GeneratePage(path, templatePath, destinationPath)
	}

	err := filepath.WalkDir(contentDir, handleWalkDirEntry)
	if err != nil {
		return fmt.Errorf("generating pages failed: %v", err)
	}

	fmt.Println("Generated pages successfully!")

	return nil
}

func GeneratePage(sourcePath, templatePath, destinationPath string) error {
	fmt.Printf(
		"Generating page from source (%s) to destination (%s) using template (%s)...\n",
		sourcePath, destinationPath, templatePath,
	)

	sourceBytes, err := os.ReadFile(sourcePath)
	if err != nil {
		return fmt.Errorf(
			"generating page failed, couldn't read source file (%s): %v", sourcePath, err,
		)
	}

	sourceMarkdown := string(sourceBytes)

	title, err := markdown.ExtractMarkdownTitle(sourceMarkdown)
	if err != nil {
		return fmt.Errorf(
			"generating page failed, couldn't extract title from markdown: %v", err,
		)
	}

	contentHTMLNode, err := adapters.MarkdownToHTMLNode(sourceMarkdown)
	if err != nil {
		return fmt.Errorf(
			"generating page failed, couldn't transform markdown to HTML node: %v", err,
		)
	}

	content, err := contentHTMLNode.ToHTML()
	if err != nil {
		return fmt.Errorf(
			"generating page failed, couldn't transform HTML Node to HTML: %v", err,
		)
	}

	templateBytes, err := os.ReadFile(templatePath)
	if err != nil {
		return fmt.Errorf(
			"generating page failed, couldn't read template file (%s): %v", templatePath, err,
		)
	}

	template := string(templateBytes)
	template = strings.Replace(template, "{{ Title }}", title, 1)
	template = strings.Replace(template, "{{ Content }}", content, 1)

	destinationDir := filepath.Dir(destinationPath)

	err = os.MkdirAll(destinationDir, 0o755)
	if err != nil {
		return fmt.Errorf(
			"generating page failed, couldn't create destination file's directory (%s): %v",
			destinationDir, err,
		)
	}

	file, err := os.Create(destinationPath)
	if err != nil {
		return fmt.Errorf(
			"generating page failed, couldn't open or create destination file (%s): %v",
			destinationPath, err,
		)
	}

	defer file.Close()

	_, err = file.WriteString(template)
	if err != nil {
		return fmt.Errorf(
			"generating page failed, couldn't write to destination file (%s): %v",
			destinationPath, err,
		)
	}

	fmt.Printf(
		"Generated page from source (%s) to destination (%s) using template (%s) Successfully!\n",
		sourcePath, destinationPath, templatePath,
	)

	return nil
}
