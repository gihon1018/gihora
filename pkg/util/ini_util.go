package util

import (
	"fmt"

	"gopkg.in/ini.v1"
)

type IniFile struct {
	file *ini.File
	path string
}

func NewIniFile(path string) (*IniFile, error) {
	file, err := ini.Load(path)
	if err != nil {
		return nil, fmt.Errorf("加载 ini 文件 %s 失败: %w", path, err)
	}
	return &IniFile{
		file: file,
		path: path,
	}, nil
}

func (i *IniFile) Path() string {
	return i.path
}

func (i *IniFile) HasSection(section string) bool {
	return i.file.HasSection(section)
}

func (i *IniFile) HasKey(section, key string) bool {
	s, err := i.file.GetSection(section)
	if err != nil {
		return false
	}
	return s.HasKey(key)
}

func (i *IniFile) GetString(section, key, def string) string {
	return i.file.Section(section).Key(key).MustString(def)
}

func (i *IniFile) GetInt(section, key string, def int) int {
	return i.file.Section(section).Key(key).MustInt(def)
}

func (i *IniFile) GetBool(section, key string, def bool) bool {
	return i.file.Section(section).Key(key).MustBool(def)
}

func (i *IniFile) Set(section, key, value string) {
	i.file.Section(section).Key(key).SetValue(value)
}

func (i *IniFile) DeleteKey(section, key string) error {
	s, err := i.file.GetSection(section)
	if err != nil {
		return fmt.Errorf("获取 section 区块 %s 失败: %w", section, err)
	}
	s.DeleteKey(key)
	return nil
}

func (i *IniFile) MapTo(section string, obj any) error {
	err := i.file.Section(section).MapTo(obj)
	if err != nil {
		return fmt.Errorf("映射 section 区块 %s 失败: %w", section, err)
	}
	return nil
}

func (i *IniFile) Save() error {
	return i.file.SaveTo(i.path)
}

func (i *IniFile) SaveTo(savePath string) error {
	return i.file.SaveTo(savePath)
}
