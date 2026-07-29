package util

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func ReadFileByLine(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("打开文件 %s 失败: %w\n", path, err)
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Printf("[ERROR] 关闭文件 %s 失败: %v\n", path, err)
		}
	}(file)

	var lineSlice []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lineSlice = append(lineSlice, line)
		if err := scanner.Err(); err != nil {
			return nil, fmt.Errorf("读取文件 %s 失败: %w\n", path, err)
		}
	}

	return lineSlice, nil
}

func WriteToFile(path, content string) error {
	if err := os.WriteFile(path, []byte(content), 0644); err != nil {
		return fmt.Errorf("写入文件 %s 失败: %w\n", path, err)
	}

	return nil
}

func ListSubDirPaths(path string) ([]string, error) {
	stat, err := os.Stat(path)
	if err != nil {
		return nil, fmt.Errorf("获取文件 %s 的信息失败: %w", path, err)
	}
	if !stat.IsDir() {
		return nil, fmt.Errorf("文件 %s 不是目录", path)
	}

	entries, err := os.ReadDir(path)
	if err != nil {
		return nil, fmt.Errorf("读取目录 %s 失败: %w", path, err)
	}

	var paths []string
	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}

		paths = append(paths, filepath.Join(path, entry.Name()))

	}

	return paths, nil
}
