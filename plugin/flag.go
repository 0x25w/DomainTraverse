package plugin

import (
	"flag"
	"fmt"
	"os"
)

func Flag(Info *DomainInfo) {
	Banner()
	//flag.StringVar(&Info.Help, "h", "", "-h help     |    获取帮助")
	//自动识别输入的是域名还是.txt
	//然后开始解析，存入[]Domains
	flag.StringVar(&Info.Domain, "d", "", "-d xxx.com     |    单域名爆破 \n-d domains.txt |    多域名爆破")
	//flag.StringVar(&Info.Domains, "D", "", "-D domain.txt | 多域名爆破")
	flag.StringVar(&Info.List, "l", "", "-l dic.txt     |    导入词典")
	//flag.IntVar(&Threads, "t", 600, "Thread nums")
	flag.Parse()
	//判断空值
	if Info.Domain == "" || Info.List == "" {
		fmt.Println("\n	!!!!!!!Args is NULL!!!!!!!\n")
		flag.Usage()
		os.Exit(0)
	}

	//fmt.Println(*Info)
	//flag
	//fmt.Println(*Info)
	//判断输入正误or为空

	//解析

	//返回域名字符串组
}
