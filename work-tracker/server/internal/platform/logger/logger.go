package logger

import (
	"fmt"
	"log"
	"strings"

	"tracker/internal/shared/utils"
)

func LogError(message string, err error) {
	fileName := fmt.Sprintf("logs/%s.txt", utils.DateNow())
	absPath := utils.AppendRootDir(fileName)

	separatorLine1 := strings.Repeat("=", 60)
	separatorLine2 := strings.Repeat("*", 60)

	if err != nil {
		message = fmt.Sprintf("%s: %v", message, err)
	}

	dateLog := fmt.Sprintf("%s\n%s\n%s", separatorLine1, utils.Now(), separatorLine1)
	textLog := fmt.Sprintf("%s\n%s\n%s\n", dateLog, message, separatorLine2)

	utils.LogFile(absPath, textLog)
	
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("\n⚠️error: %v\n", message)
}