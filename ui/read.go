package ui

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	AbortKeyword = "abort"
)

func ReadLineOrAbort(msg, abortKeyword string) (*string, error) {
	s, err := ReadLine(fmt.Sprintf("%s (or type '%s' to abort): ", msg, abortKeyword))
	if err != nil {
		return nil, err
	}
	if strings.ToLower(s) == abortKeyword {
		PrintfInfo("Aborting...\n")
		return nil, nil
	}
	return &s, nil
}

// Enter rss link (or type 'abort' to exit): "
func ReadLine(msg string) (string, error) {
	var s string
	PrintfInfo(msg)
	_, err := fmt.Scanf("%s", &s)
	return s, err
}

func ReadInt64OrAbort(msg, keyword string) (*int64, error) {
	s, err := ReadLineOrAbort(msg, keyword)
	if err != nil {
		return nil, err
	}
	if s == nil {
		return nil, nil
	}
	i, err := strconv.ParseInt(*s, 10, 64)
	return &i, err
}
