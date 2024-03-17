package main

import (
	"fmt"
	"github.com/Henry-GongZY/bilibili_ytb_utils/utils"
)

func main() {
	available, latency, statusCode := utils.NetworkAvailable("https://www.google.com", 10000, true, "127.0.0.1:7897")
	fmt.Println(available)
	fmt.Println(latency)
	fmt.Println(statusCode)
}
