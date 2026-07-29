package model

import (
	"gihora/pkg/util"
)

const (
	shardIniFileName = "server.ini"
)

type Shard struct {
	dirPath   string
	iniFile   *util.IniFile
	shardName string
	// TODO
}

//func NewShard(dirPath string) (*Shard, error) {
//	shardIniFilePath := filepath.Join(dirPath, shardIniFileName)
//	iniFile, err := loadShardIniFile(shardIniFilePath)
//	if err != nil {
//		return nil, err
//	}
//
//	shard := &Shard{
//		dirPath: dirPath,
//		iniFile: iniFile,
//		// TODO
//	}
//
//	return shard, nil
//}
//
//func loadShardIniFile(path string) (*util.IniFile, error) {
//	iniFile, err := util.NewIniFile(path)
//	if err != nil {
//		return nil, err
//	}
//
//	if iniFile.GetString(section: "server", key: "shardId") == "" {
//		return nil, fmt.Errorf(format: "无效 ini 对象：%s", path)
//	}
//
//	return iniFile, nil
//}
