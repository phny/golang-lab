package iotest

import (
	"fmt"
	"github.com/phachon/go-logger"
	"io/ioutil"
	"strings"
	"testing"
)

func GetFileContentAsStringLines(filePath string) ([]string, error) {
	logger := go_logger.NewLogger()
	logger.Infof("get file content as lines: %v", filePath)
	result := []string{}
	b, err := ioutil.ReadFile(filePath)
	if err != nil {
		logger.Errorf("read file: %v error: %v", filePath, err)
		return result, err
	}
	s := string(b)
	for _, lineStr := range strings.Split(s, "\n") {
		lineStr = strings.TrimSpace(lineStr)
		if lineStr == "" {
			continue
		}
		result = append(result, lineStr)
	}
	logger.Infof("get file content as lines: %v, size: %v", filePath, len(result))
	return result, nil
}

func TestReadTextFile(t *testing.T) {
	file := "/home/SENSETIME/heyulin/format_transform.py"
	lines, _ := GetFileContentAsStringLines(file)
    for _, line := range lines {
	    fmt.Println(line)
    }
}
