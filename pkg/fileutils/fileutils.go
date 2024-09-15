package fileutils

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path"
)

func CleanAndCopyFromSourceDirToDestinationDir(sourceDir, destinationDir string) error {
	err := deleteContentsOfDestinationDir(destinationDir)
	if err != nil {
		return err
	}

	return copyFromSourceDirToDestinationDir(sourceDir, destinationDir)
}

func deleteContentsOfDestinationDir(destinationDir string) error {
	dirEntries, err := os.ReadDir(destinationDir)
	if err != nil {
		if errors.Is(err, fs.ErrNotExist) {
			return createDestinationDir(destinationDir)
		}
		return fmt.Errorf("failed to read %s directory: %v", destinationDir, err)
	}

	fmt.Printf("Deleting contents of destination directory (%s)...\n", destinationDir)

	for _, dirEntry := range dirEntries {
		subPath := path.Join(destinationDir, dirEntry.Name())

		if err := os.RemoveAll(subPath); err != nil {
			return fmt.Errorf(
				"deleting contents of destination directory (%s) failed: failed to remove %s file/directory: %v",
				destinationDir, subPath, err,
			)
		}
	}

	fmt.Printf("Deleted contents of destination directory (%s) successfully!\n", destinationDir)

	return nil
}

func createDestinationDir(destinationDir string) error {
	fmt.Printf("Creating destination directory (%s)...\n", destinationDir)

	err := os.Mkdir(destinationDir, 0o755)
	if err != nil {
		return fmt.Errorf("creating destination directory (%s) failed: %v", destinationDir, err)
	}

	fmt.Printf("Created destination directory (%s) successfully!\n", destinationDir)

	return nil
}

func copyFromSourceDirToDestinationDir(sourceDir, destinationDir string) error {
	fmt.Printf(
		"Copying contents of source directory (%s) into destination directory (%s)...\n",
		sourceDir, destinationDir,
	)

	err := os.CopyFS(destinationDir, os.DirFS(sourceDir+"/"))
	if err != nil {
		return fmt.Errorf(
			"copying contents of source directory (%s) into destination directory (%s) failed: %v",
			sourceDir, destinationDir, err,
		)
	}

	fmt.Printf(
		"Copied contents of source directory (%s) into destination directory (%s) successfully!\n",
		sourceDir, destinationDir,
	)

	return nil
}
