package config

import (
	"fmt"
	"net"
	"os"
	"prometheusd/utils"
	"strings"

	"github.com/google/uuid"
)

type Option struct {
	Server   string
	UUID     string
	Hostanem string
	Addr     string
}

func GetUUID() string {

	// 先判断是否存在,存在就用现在的,不存在就生成一个然后写入文件
	path := "promeagent.uuid"
	if agentID := utils.ReadFile(path); agentID != "" {
		return agentID
	}

	agentID := strings.ReplaceAll(uuid.New().String(), "-", "")
	utils.WiteFile(path, agentID)
	return agentID
}

func GetAddr() (string, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "", err
	}
	for _, addr := range addrs {
		// 判断是否为IPV4
		if (strings.Index(addr.String(), ":")) >= 0 {
			continue
		}
		// 再用斜线分割成N份,拿到ip地址
		nodes := strings.SplitN(addr.String(), "/", 2)
		if len(nodes) != 2 {
			continue
		}
		return nodes[0], nil
	}
	return "", fmt.Errorf("not net addr")
}

// new函数,实例化对象用的
func NewOption(server string) (*Option, error) {
	uuid := GetUUID()
	hostname, err := os.Hostname()
	if err != nil {
		return nil, err
	}
	addr, err := GetAddr()
	if err != nil {
		return nil, err
	}
	return &Option{
		Server:   server,
		UUID:     uuid,
		Hostanem: hostname,
		Addr:     addr,
	}, nil
}
