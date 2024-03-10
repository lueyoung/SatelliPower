package main

type Config struct {
	StartTime   int64  `json:startTime,int64`
	BindUdpPort int64  `json:bindUdpPort,int64`
	DstUdpHost  string `json:dstUdpHost,string`
	DstUdpPort  int64  `json:dstUdpPort,int64`
}
