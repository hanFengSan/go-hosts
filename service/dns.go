package service

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net"
)

// GetSubjectHosts 获取主题Hosts文件
func GetSubjectHosts(subjectName string) (filePath string, err error) {
	records, err := GetSubjectDNSRecords(subjectName)
	filePath, err = GenerateHosts(subjectName, records)
	return
}

// GetSubjectDNSRecords 获取主题DNS记录
func GetSubjectDNSRecords(subjectName string) (records []string, err error) {
	subjectMap, err := readSubjectCfg()
	subjectDomains, _ := subjectMap[subjectName]
	for _, domain := range subjectDomains {
		ips, err := LookupDNS(domain)
		if err != nil {
			return nil, err
		}
		for _, ip := range ips {
			records = append(records, ip+" "+domain)
		}
	}
	return
}

// LookupDNS 解析dns
func LookupDNS(domain string) (result []string, err error) {
	ips, err := net.LookupIP(domain)
	if err != nil {
		log.Printf(domain+" could not get IPs: %v\n", err)
		return
	}
	for _, ip := range ips {
		result = append(result, ip.String())
	}
	return
}

func readSubjectCfg() (subjectMap map[string][]string, err error) {
	b, err := ioutil.ReadFile(AssetPath + "/subject.json")
	if err != nil {
		log.Print(err)
	}
	json.Unmarshal(b, &subjectMap)
	return
}
