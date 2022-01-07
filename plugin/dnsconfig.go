package plugin

import (
	"fmt"
)

//生成一个新的[]SubDomains
func MixUp(Info *DomainInfo) {
	for n, _ := range Info.Domains {
		for _, l := range Info.Lists {
			str := l + "." + Info.Domains[n]
			Info.SubDomains = append(Info.SubDomains, str)
		}
	}
	fmt.Print("子域名数为：")
	fmt.Print(len(Info.SubDomains))
	fmt.Println(" , 正在解析子域...")
}
