package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// GetFileContentAsStringLines real file contents
func GetFileContentAsStringLines(filePath string) ([]string, error) {
	fmt.Println("get file content as lines: %v", filePath)
	result := []string{}
	b, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("read file: %v error: %v", filePath, err)
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
	fmt.Println("get file content as lines: %v, size: %v", filePath, len(result))
	return result, nil
}

func WriteToFile(contents []string, filePath string) {
	var f *os.File
	defer f.Close()
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		f, _ = os.Create(filePath)
	} else {
		f, _ = os.OpenFile(filePath, os.O_APPEND, 0666)
	}
	w := bufio.NewWriter(f)

	for ind, str := range contents {
		w.WriteString(str + "\n")
		w.Flush()
		if ind%100 == 0 {
			fmt.Println("aready write: %v", ind)
		}
	}
}

func main() {
	lines, _ := GetFileContentAsStringLines("/home/SENSETIME/heyulin/result.txt")
	WriteToFile(lines, "1.txt")
	for _, line := range lines {
		splited := strings.Fields(line)
		fmt.Println(splited[0], splited[1])
	}
}
