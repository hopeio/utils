package main

import "github.com/hopeio/gox/net/http/proxy"

func main() {
	proxy.DirectorServer(":8080")
}
