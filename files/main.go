package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "filemanager",
		Usage: "A CLI tool for managing JSON files",
		Commands: []*cli.Command{
			{
				Name:      "upload",
				Usage:     "Upload a JSON file",
				ArgsUsage: "[source_file] [destination_file]",
				Action:    uploadFile,
			},
			{
				Name:      "download",
				Usage:     "Download a JSON file",
				ArgsUsage: "[source_file]",
				Action:    downloadFile,
			},
			{
				Name:      "delete",
				Usage:     "Delete a file",
				ArgsUsage: "[file_path]",
				Action:    deleteFile,
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Println("Error:", err)
	}
}

func uploadFile(c *cli.Context) error {
	if c.NArg() != 2 {
		return fmt.Errorf("please provide source and destination file paths")
	}
	source := c.Args().Get(0)
	destination := c.Args().Get(1)

	srcFile, err := os.Open(source)
	if err != nil {
		return fmt.Errorf("error opening source file: %v", err)
	}
	defer srcFile.Close()

	var data interface{}
	decoder := json.NewDecoder(srcFile)
	err = decoder.Decode(&data)
	if err != nil {
		return fmt.Errorf("error decoding JSON: %v", err)
	}

	dstFile, err := os.Create(destination)
	if err != nil {
		return fmt.Errorf("error creating destination file: %v", err)
	}
	defer dstFile.Close()

	encoder := json.NewEncoder(dstFile)
	encoder.SetIndent("", "  ")
	err = encoder.Encode(data)
	if err != nil {
		return fmt.Errorf("error encoding JSON: %v", err)
	}

	fmt.Printf("File uploaded successfully: %s to %s\n", source, destination)
	return nil
}

func downloadFile(c *cli.Context) error {
	if c.NArg() != 1 {
		return fmt.Errorf("please provide the source file path")
	}
	source := c.Args().Get(0)

	homeDir, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("could not get home directory: %v", err)
	}

	destination := filepath.Join(homeDir, "Downloads", filepath.Base(source))

	srcFile, err := os.Open(source)
	if err != nil {
		return fmt.Errorf("error opening source file: %v", err)
	}
	defer srcFile.Close()

	dstFile, err := os.Create(destination)
	if err != nil {
		return fmt.Errorf("error creating destination file: %v", err)
	}
	defer dstFile.Close()

	_, err = io.Copy(dstFile, srcFile)
	if err != nil {
		return fmt.Errorf("error copying file: %v", err)
	}

	fmt.Printf("File downloaded successfully to: %s\n", destination)
	return nil
}

func deleteFile(c *cli.Context) error {
	if c.NArg() != 1 {
		return fmt.Errorf("please provide a file path to delete")
	}
	filePath := c.Args().Get(0)

	err := os.Remove(filePath)
	if err != nil {
		return fmt.Errorf("error deleting file: %v", err)
	}

	fmt.Printf("File deleted successfully: %s\n", filePath)
	return nil
}
