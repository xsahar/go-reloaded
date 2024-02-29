package main

import (
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
)

func editTextFile(filename string) error {
	// Open the input file
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	// Read the content of the file
	var builder strings.Builder
	_, err = io.Copy(&builder, file)
	if err != nil {
		return err
	}
	content := []byte(builder.String())
	text := string(content)

	re := regexp.MustCompile(`([.,!?;:])\s*`)
	text = re.ReplaceAllString(text, "$1 ")

	re = regexp.MustCompile(`\s+([.,!?;:])|([.!?]){3,}\s*`)
	text = re.ReplaceAllString(text, "$1")
	re = regexp.MustCompile(`' (.*?) '`)
	text = re.ReplaceAllString(text, "'$1'")

	s := strings.TrimSpace(text)

	f := []byte(s)
	// Close the input file
	if err := file.Close(); err != nil {
		return err
	}

	// Open the output file for writing
	outputFile, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer outputFile.Close()

	// Write the modified content to the output file
	_, err = outputFile.Write(f)
	if err != nil {
		return err
	}

	fmt.Printf("The file '%s' has been successfully edited.\n", filename)
	return nil
}
