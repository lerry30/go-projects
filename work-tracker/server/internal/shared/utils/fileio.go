package utils

import (
	"fmt"
	"os"
	"bufio"
	"io"

	"path/filepath"
)

func AppendRootDir(rel string) string {
	// 1. Explicit config wins (set this in production)
	if dir := os.Getenv("ROOT_DIR"); dir != "" {
		return filepath.Join(dir, rel)
	}

	// 2. Fall back to rel to where the process was started
	wd, err := os.Getwd()
	if err != nil {
		return rel
	}

	return filepath.Join(wd, rel)
}

func LogFile(fileName string, textLog string) {
	if fileName == "" || textLog == "" {
		fmt.Fprintln(os.Stderr, "LogFile function: arguments error")
		return
	}

	// Ensure the parent directory exists
	if err := os.MkdirAll(filepath.Dir(fileName), 0755); err != nil {
		fmt.Fprintln(os.Stderr, "Directory error: ", filepath.Dir(fileName), err)
		return
	}

	logF, err := os.OpenFile(fileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Fprintln(os.Stderr, "File error: ", fileName, err)
		return
	}

	defer logF.Close()
	
	_, err = logF.WriteString(textLog + "\n")
	if err != nil {
		fmt.Fprintln(os.Stderr, "log into file error: ", fileName, err)
		return
	}

	fmt.Println("✔️  Logged into file.")
}

func ReadFile(fileName string) (func() (string, error), func() error, error) {
	// O_RDONLY: read only
    // O_CREATE: create if doesn't exist
    // 0644: standard permissions for new files
	file, err := os.OpenFile(fileName, os.O_RDONLY|os.O_CREATE, 0644)
	if err != nil {
		return nil, nil, err
	}
	scanner := bufio.NewScanner(file)

	next := func() (string, error) {
		if scanner.Scan() {
			return scanner.Text(), nil
		}

		if err := scanner.Err(); err != nil {
			return "", err
		}

		return "", io.EOF
	}

	close := func() error {
		return file.Close()
	}

	return next, close, nil
}