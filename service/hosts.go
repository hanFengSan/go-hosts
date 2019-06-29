package service

import (
	"io/ioutil"
	"strings"
	"sync"
)

var mutex sync.RWMutex

// GenerateHosts 生成hosts文件
func GenerateHosts(subjectName string, records []string) (filePath string, err error) {
	// 追加通用解析
	records = append([]string{"127.0.0.1 localhost", "255.255.255.255	broadcasthost", "::1 localhost"}, records...)
	hostStr := strings.Join(records[:], "\n")
	filePath = OutputPath + "/" + subjectName + ".hosts"
	mutex.Lock()
	defer mutex.Unlock()
	err = ioutil.WriteFile(filePath, []byte(hostStr), 0644)
	return
}
