package util

import (
	"fmt"
	"os"

	"github.com/goccy/go-yaml"
)

func ReadYAML(path string, out any) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("读取 YAML 文件 %s 失败: %w", path, err)
	}

	if err := yaml.Unmarshal(data, out); err != nil {
		return fmt.Errorf("解析 YAML 文件 %s 失败: %w", path, err)
	}

	return nil
}
