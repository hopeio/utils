package main

import "github.com/hopeio/utils/net/http/proxy"

func main() {
	proxy.DirectorServer(":8080")
}
