package net

import (
	"fmt"
	"net"
)

func IPStrToUint32(ipStr string) (uint32, error) {
	ip := net.ParseIP(ipStr)
	if ip == nil {
		return 0, fmt.Errorf("invalid ip addr")
	}

	return IPv4ToUint32(ip)
}

func IPv4ToUint32(ip net.IP) (uint32, error) {
	if ip.To4() == nil {
		return 0, fmt.Errorf("invalid ip addr")
	}

	ipBytes := ip.To4()
	return uint32(ipBytes[0])<<24 | uint32(ipBytes[1])<<16 | uint32(ipBytes[2])<<8 | uint32(ipBytes[3]), nil
}

func Uint32ToIPStr(ipInt uint32) (string, error) {
	ip, err := Uint32ToIPv4(ipInt)
	if err != nil {
		return "", fmt.Errorf("invalid ip addr")
	}

	return ip.String(), nil
}

func Uint32ToIPv4(ipInt uint32) (net.IP, error) {
	ip := make(net.IP, 4)
	ip[0] = byte(ipInt >> 24)
	ip[1] = byte(ipInt >> 16)
	ip[2] = byte(ipInt >> 8)
	ip[3] = byte(ipInt)

	if ip.IsUnspecified() {
		return nil, fmt.Errorf("invalid ip addr")
	}
	return ip, nil
}
