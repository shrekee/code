package main
import (
	"fmt"
	"net/http"
	"strings"
	"./craw01"
	"os/exec"
)

type Hello struct {}

func (h Hello) ServerHTTP(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	var filePath string// = "/Users/vince"
	var fileName string// = "/Users/vince/test.xml"

	var cmds string

	for k,v := range r.Form {
		if strings.ContainsAny(k, "ls") {
			filePath = strings.join(v, "")
		}
		if strings.ContainsAny(k, "pt") {
			fileName = strings.Join(v, "")
		}
		if strings.ContainsAny(k, "cmd") {
			cmds = strings.Join(v, "")
			if cmds == "reboot" {
				fmt.Println("reboot")
				cmd := exec.Command("shutdown","-r","now")//生成重启命令
				cmd.Start()
			}
		}
	}
	if filePath != "" {//输出文件列表
		fmt.Fprint(w,"the path is :")
		fmt.Fprint(w, filePath)
		fmt.Fprint(w, "\n")
		files,err := craw01
	}
}
