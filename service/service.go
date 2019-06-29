package service

import (
	"os"
	"github.com/hanFengSan/go-hosts/util"
)

// AssetPath 资源路径
var AssetPath string

// OutputPath 输出路径
var OutputPath string

func init() {
	InitPath()
}

// InitPath 初始化路径, 内存盘加速
func InitPath() {
	basePath := "."
	assetSrcPath := "./asset"
	if _, err := os.Stat("/Volumes/RAMDisk"); !os.IsNotExist(err) {
		// Mac 存在内存盘
		basePath = "/Volumes/RAMDisk/go-hosts"
	}
	if _, err := os.Stat("/dev/shm"); !os.IsNotExist(err) {
		// Linux 存在内存盘
		basePath = "/dev/shm/go-hosts"
	}
	OutputPath = basePath + "/output"
	AssetPath = basePath + "/asset"
	os.RemoveAll(OutputPath)
	os.MkdirAll(OutputPath, 0777)
	if assetSrcPath != AssetPath {
		util.Copy(assetSrcPath, AssetPath)
	}
}
