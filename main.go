package main

import "DomainTraverse/plugin"

func main() {

	var Info plugin.DomainInfo

	plugin.Flag(&Info)
	//解析
	plugin.Parse(&Info)
	//运行
	plugin.Running(&Info)
}
