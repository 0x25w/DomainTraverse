package plugin

import (
	"fmt"
	"os"
)

// 							path + filename
func WriteFile(str string, filename string) {
	var text = []byte(str + "\n")
	fl, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		fmt.Printf("Open %s error, %v\n", filename, err)
		return
	}
	_, err = fl.Write(text)
	fl.Close()
	if err != nil {
		fmt.Printf("Write %s error, %v\n", filename, err)
	}
	fmt.Println(str)
}

func Output(Info *DomainInfo) {
	for _, str := range Info.Dst {
		WriteFile(str, "result.txt")
	}
	fmt.Println("Success!")
}
