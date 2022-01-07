package plugin

//只需要解析两个参数，然后分别存进字符串
type DomainInfo struct {
	//Help	  string
	Domain     string
	List       string
	Domains    []string
	Lists      []string
	SubDomains []string
	Timeout    int64
}
