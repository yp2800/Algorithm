package leetcode

import (
	"bytes"
	"net"
	"sort"
	"strings"
)

//冒泡
func ipv4Sort(ips []string) []string {
	for i := 0; i < len(ips)-1; i++ {
		for j := i + 1; j < len(ips); j++ {
			if strings.Compare(ips[i], ips[j]) == -1 {
				ips[i], ips[j] = ips[j], ips[i]
			}
		}
	}
	return ips
}

// sort排序
func ipv4Sort2(ips []string) []string {
	sort.Strings(ips)
	for i, j := 0, len(ips)-1; i < j; i, j = i+1, j-1 {
		ips[i], ips[j] = ips[j], ips[i]
	}
	return ips
}

// sort排序，net.ip 库使用
func ipv4Sort3(ips []string) []string {
	realIPs := make([]net.IP, 0, len(ips))
	for _, ip := range ips {
		realIPs = append(realIPs, net.ParseIP(ip))
	}
	sort.Slice(realIPs, func(i, j int) bool {
		return bytes.Compare(realIPs[i], realIPs[j]) > 0
	})

	var result []string
	for _, ip := range realIPs {
		result = append(result,ip.String())
	}
	return result
}

// 比较四次，每段比较一次
func ipv4Sort4(ips []string) []string {
	for i:=0;i<3;i++{
		for j := 0;j < len(ips)-1;j++{
			if strings.Compare(ips[j],ips[j+1]) < 0 {
				ips[j],ips[j+1] = ips[j+1],ips[j]
			}
		}
	}
	return ips
}