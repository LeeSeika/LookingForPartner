package discovery

import (
	"encoding/json"
	"errors"
	"fmt"
	"google.golang.org/grpc/resolver"
	"strings"
)

type Server struct {
	Name string `json:"name"`
	Addr string `json:"addr"` // 地址
	//Version string `json:"version"` // 版本
	//Weight  int64  `json:"weight"`  // 权重
}

func BuildEndPointKey(server Server) string {
	//fmt.Sprintf("/%s/%s/%s", server.Name, server.Version, server.Addr)
	return fmt.Sprintf("/%s/%s", server.Name, server.Addr)
}

func ParseValue(v []byte) (Server, error) {
	server := Server{}
	err := json.Unmarshal(v, &server)
	if err != nil {
		return server, err
	}
	return server, nil
}

func SplitPath(path string) (Server, error) {
	server := Server{}
	strs := strings.Split(path, "/")
	if len(strs) == 0 {
		return server, errors.New("invalid path")
	}

	server.Addr = strs[len(strs)-1]

	return server, nil
}

func Exist(l []resolver.Address, addr resolver.Address) bool {
	for i := range l {
		if l[i].Addr == addr.Addr {
			return true
		}
	}

	return false
}

func Remove(s []resolver.Address, addr resolver.Address) ([]resolver.Address, bool) {
	for i := range s {
		if s[i].Addr == addr.Addr {
			s[i] = s[len(s)-1]
			return s[:len(s)-1], true
		}
	}
	return nil, false
}
