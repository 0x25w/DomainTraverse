package plugin

import (
	"fmt"
	"github.com/miekg/dns"
	"time"
)

// cname1 返回所有层的cname
// src: 域名
// dnsService: dns服务器, 如114.114.114.114
func Cname1(Info *DomainInfo) {

	cname2(Info)
}

func cname2(Info *DomainInfo) {
	c := dns.Client{
		Timeout: 5 * time.Second,
	}

	var lastErr error
	// retry 3 times
	for i := 0; i < 3; i++ {

		m := dns.Msg{}
		// 最终都会指向一个ip 也就是typeA, 这样就可以返回所有层的cname.
		m.SetQuestion(Info.SubDomains[0]+".", dns.TypeA)
		Info.Server = "8.8.8.8"
		r, _, err := c.Exchange(&m, Info.Server+":53")
		if err != nil {
			lastErr = err
			time.Sleep(1 * time.Second * time.Duration(i+1))
			continue
		}

		for _, ans := range r.Answer {
			record, isType := ans.(*dns.CNAME)
			if isType {
				Info.Dst = append(Info.Dst, record.Target)
			}
		}
		lastErr = nil
		break
	}
	fmt.Println(lastErr)
}
