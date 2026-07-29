package model

import (
	"fmt"
	"gihora/pkg/util"
	"log"
	"path/filepath"
)

const (
	clusterIniFileName = "cluster.ini"
)

var sectionArr = [4]string{"GAMEPLAY", "NETWORK", "MISC", "SHARD"}

type Cluster struct {
	dirPath   string
	iniFile   *util.IniFile
	shardList []Shard

	// GAMEPLAY
	GameMode       string `ini:"game_mode" json:"game_mode"`
	MaxPlayers     int    `ini:"max_players" json:"max_players"`
	Pvp            bool   `ini:"pvp" json:"pvp"`
	PauseWhenEmpty bool   `ini:"pause_when_empty" json:"pause_when_empty"`

	// NETWORK
	LanOnlyCluster     string `ini:"lan_only_cluster" json:"lan_only_cluster"`
	ClusterPassword    string `ini:"cluster_password" json:"cluster_password"`
	ClusterDescription string `ini:"cluster_description" json:"cluster_description"`
	ClusterName        string `ini:"cluster_name" json:"cluster_name"`
	OfflineCluster     bool   `ini:"offline_cluster" json:"offline_cluster"`
	ClusterLanguage    string `ini:"cluster_language" json:"cluster_language"`
	ClusterCloudId     string `ini:"cluster_cloud_id" json:"cluster_cloud_id"`

	// MISC
	ConsoleEnabled bool `ini:"console_enabled" json:"console_enabled"`

	// SHARD
	ShardEnabled bool   `ini:"shard_enabled" json:"shard_enabled"`
	BindIp       string `ini:"bind_ip" json:"bind_ip"`
	MasterIp     string `ini:"master_ip" json:"master_ip"`
	MasterPort   int    `ini:"master_port" json:"master_port"`
	ClusterKey   string `ini:"cluster_key" json:"cluster_key"`
}

func NewCluster(path string) (*Cluster, error) {
	clusterIniFilePath := filepath.Join(path, clusterIniFileName)
	iniFile, err := loadClusterIniFile(clusterIniFilePath)
	if err != nil {
		log.Println("[ERROR] 加载 ini 配置文件失败")
		return nil, err
	}

	var cluster Cluster

	for _, section := range sectionArr {
		if err := iniFile.MapTo(section, &cluster); err != nil {
			log.Println("[ERROR] 映射 ini 配置文件失败")
			return nil, err
		}
	}

	return &cluster, nil
}

func loadClusterIniFile(path string) (*util.IniFile, error) {
	iniFile, err := util.NewIniFile(path)
	if err != nil {
		return nil, err
	}

	if iniFile.GetString("NETWORK", "cluster_cloud_id", "") == "" {
		return nil, fmt.Errorf("加载的 cluster ini 对象 %s 无效", path)
	}

	return iniFile, nil
}

func IsValidClusterDirPath(path string) bool {
	clusterIniFilePath := filepath.Join(path, clusterIniFileName)
	_, err := loadClusterIniFile(clusterIniFilePath)
	if err != nil {
		return false
	}
	return true
}
