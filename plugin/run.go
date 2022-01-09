package plugin

import (
	"fmt"
	"time"
)

func Running(Info *DomainInfo) {
	tim := time.Now()
	//开始拼接
	MixUp(Info)
	////配置DNS
	//SetDnsConfig(Info)
	////发送请求
	//SendDnsReq()
	////收到请求
	//RecAndOutput()

	//exec
	//Cname1(Info)
	CNAME(Info)

	fmt.Println("解析中，请勿关闭...")

	Info.Dst = RemoveDuplication(Info.Dst)
	//stdin
	//for _,i := range Info.Dst{
	//	fmt.Println(i)
	//}

	//put into file
	Output(Info)
	fmt.Println("\n", "开始时间", tim, "\n", "结束时间", time.Now())
	fmt.Printf("共发现%d个子域名", len(Info.Dst))
}
