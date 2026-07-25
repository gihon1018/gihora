package util

import (
	"bufio"
	"log"
	"os"
)

func ReadFileByLine(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		log.Printf("[ERROR] 打开文件失败: %v\n", err)
		return nil, err
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Printf("[ERROR] 关闭文件失败: %v\n", err)
		}
	}(file)

	var lineSlice []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lineSlice = append(lineSlice, line)
		if err := scanner.Err(); err != nil {
			log.Printf("[ERROR] 读取文件失败: %v\n", err)
			return nil, err
		}
	}

	return lineSlice, nil
}

func WriteToFile(path, content string) error {
	if err := os.WriteFile(path, []byte(content), 0644); err != nil {
		log.Printf("[ERROR] 写入文件失败: %v\n", err)
		return err
	}

	return nil
}
