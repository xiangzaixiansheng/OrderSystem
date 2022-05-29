package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type Server struct {
	ServerName string
	ServerIp   string
}

type Serverslice struct {
	Servers []Server
}

type Server1 struct {
	ServerName string `json:"serverName,string"`
	ServerIP   string `json:"serverIP,omitempty"`
}
type Serverslice1 struct {
	Servers1 []Server1 `json:"servers"`
}

type ExpireTime struct {
	Start string `json:"start"`
	End   string `json:"end"`
}

//Json解析到结构体
func main() {
	var s Serverslice
	str := `{"servers":
	[{"serverName":"Guangzhou_Base","serverIP":"127.0.0.1"},
	{"serverName":"Beijing_Base","serverIP":"127.0.0.2"}]}`
	err := json.Unmarshal([]byte(str), &s)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(s)
	fmt.Println(s.Servers[0].ServerName)

	//struct生成字符串
	var s2 Serverslice1
	s2.Servers1 = append(s2.Servers1, Server1{ServerName: "Guangzhou_Base", ServerIP: "127.0.0.1"})
	s2.Servers1 = append(s2.Servers1, Server1{ServerName: "Beijing_Base", ServerIP: "127.0.02"})
	b, err := json.Marshal(s2)
	if err != nil {
		fmt.Println("JSON ERR:", err)
	}
	fmt.Println(string(b))

	//json转map
	data := []byte(`{"IP": "127.0.0.1", "name": "SKY"}`)

	m := make(map[string]string)

	err2 := json.Unmarshal(data, &m)
	if err2 != nil {
		fmt.Println("Umarshal failed:", err)
		return
	}

	fmt.Println("m:", m)

	for k, v := range m {
		fmt.Println(k, ":", v)
	}

	et := &ExpireTime{}
	et.Start = time.Now().Format("2006-01-02 15:04:05")
	et.End = time.Now().Format("2006-01-02 15:04:05")
	et_1, err := json.Marshal(et)
	fmt.Printf("测试", string(et_1))
}
