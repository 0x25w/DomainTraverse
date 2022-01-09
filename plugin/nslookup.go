package plugin

import (
	"net"
)

func CNAME(Info *DomainInfo) {
	for _, i := range Info.SubDomains {
		dst, _ := net.LookupCNAME(i)
		if dst == "" {
			continue
		}
		dst = dst[:len(dst)-1]
		Info.Dst = append(Info.Dst, dst)
	}
	//fmt.Println(len(Info.Dst))
}
