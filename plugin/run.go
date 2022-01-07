package plugin

func Running(Info *DomainInfo) {
	//开始拼接
	MixUp(Info)
	SetDnsConfig()
	SendDnsReq()
	RecAndOutput()
}
