package plugin

import (
	"net"
	"sync"
)

func CNAME(Info *DomainInfo) {
	wg := sync.WaitGroup{}
	wg.Add(len(Info.SubDomains))
	dolookup := func() {
		a, _ := net.LookupCNAME(Info.TmpStr)

		Info.Dst = append(Info.Dst, a)
		wg.Done()
	}
	for _, Info.TmpStr = range Info.SubDomains {
		go dolookup()

		//	if dst == "" {
		//		continue
		//	}
		//	dst = dst[:len(dst)-1]
		//	Info.Dst = append(Info.Dst, dst)
		//}
		//fmt.Println(len(Info.Dst))
	}
	wg.Wait()
}
