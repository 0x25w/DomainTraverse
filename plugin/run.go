package plugin

import "fmt"

func Running(Info *DomainInfo) {
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

	//stdin
	//for _,i := range Info.Dst{
	//	fmt.Println(i)
	//}

	//put into file
	Output(Info)
}
