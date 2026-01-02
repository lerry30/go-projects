package utils

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"strconv"
	"io"
)

func LogFile(fileName string, textLog string, autoIncrement bool) {
	if fileName == "" || textLog == "" {
		fmt.Fprintln(os.Stderr, "LogFile function: arguments error")
		return
	}

	logF, err := os.OpenFile(fileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Fprintln(os.Stderr, "File error: ", fileName, err)
		return
	}

	defer logF.Close()

	if autoIncrement {
		reader, err := os.Open(fileName)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error file reader: ", fileName, err)
			return
		}

		defer reader.Close()

		scanner := bufio.NewScanner(reader)
		currentLine := ""

		for scanner.Scan() {
			tempLine := scanner.Text()
			if tempLine != "" {
				currentLine = tempLine
			}
		}

		index := "1"
		if currentLine != "" {
			parts := strings.Split(currentLine, "|")
			if len(parts) > 0 {
				index = strings.TrimSpace(parts[0])
				num, err := strconv.Atoi(index)
				if err != nil {
					fmt.Fprintln(os.Stderr, "string to number", err)
					return
				}
				num++
				index = strconv.Itoa(num)
			}
		}

		textLog = index + " | " + textLog
	}
	
	_, err = logF.WriteString(textLog + "\n")
	if err != nil {
		fmt.Fprintln(os.Stderr, "log into file error: ", fileName, err)
		return
	}

	fmt.Println("✔️  Data successfully inserted")

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

func FileContentOverride(fileName string, textLog string) {
	if fileName == "" || textLog == "" {
		fmt.Fprintln(os.Stderr, "LogFile function: arguments error")
		return
	}

	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error: opening file to override content", err)
		return
	}

	defer file.Close()

	_, err = file.WriteString(textLog)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error: log into file to override content", err)
		return
	}

	fmt.Println("✔️  Data successfully inserted")
}