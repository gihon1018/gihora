package util

import (
	"bufio"
	"fmt"
	"os"
)

func ReadFileByLine(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("打开文件失败: %w", err)
	}
	defer file.Close()

	var lineSlice []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lineSlice = append(lineSlice, line)
		if err := scanner.Err(); err != nil {
			return nil, fmt.Errorf("读取文件错误: %w", err)
		}
	}

	return lineSlice, nil
}
