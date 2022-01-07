package plugin

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

//完成后应
//将域名存进[]Domains
//字典存进[]Lists
func Parse(Info *DomainInfo) {
	//ParseHelp(Info)
	//判断是否为.txt
	//如果不是，则单域名解析，解析完成压进字符串组
	//如果是，则将文件逐行写入临时字符串组，然后进行单域名解析，如果没问再压进字符串组
	ParseDomain(Info)
	//ParseDomains(Info)
	ParseList(Info)

}

func ParseHelp(Info *DomainInfo) {
	flag.Usage()
	os.Exit(0)
}

//解析输入的域名
func ParseDomain(Info *DomainInfo) {
	//判断字符串长度是否合法
	if len(Info.Domain) < 4 {
		fmt.Println("输入的命令有误，请重新输入")
		os.Exit(0)
	}

	if strings.EqualFold(Info.Domain[len(Info.Domain)-4:], ".txt") {
		//多域名模式
		ReadDomains(Info)
	} else {
		if strings.EqualFold("http://", Info.Domain[7:]) || strings.EqualFold("https:/", Info.Domain[7:]) {
			fmt.Println("\n!!!!!!!!!!当前域名文件有错误!!!!!!!!!!\n")
			os.Exit(0)
		} else {
			Info.Domains = append(Info.Domains, Info.Domain)
		}
	}

	fmt.Print("域名数量为：")
	fmt.Print(len(Info.Domains))
	fmt.Println(" , 正在录入字典...")

}

//读多域名文件 , 返回一个字符串组 。
//组内可能有非法域名，此处进行了匹配http字符处理
func ReadDomains(Info *DomainInfo) {
	//fmt.Println("读取目标位置文件，读取完成后存入[]Domains")
	file, err := os.Open(Info.Domain)
	if err != nil {
		fmt.Println("\n!!!!!!!!!!当前域名文件有错误!!!!!!!!!!\n")
		os.Exit(0)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		if strings.EqualFold("http://", scanner.Text()[7:]) || strings.EqualFold("https:/", scanner.Text()[7:]) {
			fmt.Println("\n!!!!!!!!!!当前域名文件有错误!!!!!!!!!!\n")
			os.Exit(0)
		}
		Info.Domains = append(Info.Domains, scanner.Text())
		//fmt.Println(scanner.Text())
	}
	Info.Domains = RemoveDuplication(Info.Domains)
}

//解析字典
func ParseList(Info *DomainInfo) {
	ReadTxt(Info)
	fmt.Print("字典长度为：")
	fmt.Print(len(Info.Lists))
	fmt.Println(" , 正在拼接子域...")
}

//处理字典，读文件
func ReadTxt(Info *DomainInfo) {
	//fmt.Println("读取目标位置文件，读取完成后存入[]Lists")
	file, err := os.Open(Info.List)
	if err != nil {
		fmt.Println("\n!!!!!!!!!!当前字典有错误!!!!!!!!!!\n")
		os.Exit(0)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		Info.Lists = append(Info.Lists, scanner.Text())
		//fmt.Println(scanner.Text())
	}
	Info.Lists = RemoveDuplication(Info.Lists)
}

//数组去重
func RemoveDuplication(arr []string) []string {
	set := make(map[string]struct{}, len(arr))
	j := 0
	for _, v := range arr {
		_, ok := set[v]
		if ok {
			continue
		}
		set[v] = struct{}{}
		arr[j] = v
		j++
	}

	return arr[:j]
}
