package test02

import (
	//"net/http"
	"io/ioutil"
	"os"
)

func ListDir(dirPath string) (files []string, err error) {
	files = make([]string, 0, 10)// 初始化数组，长度为0,初始容量为10,容量会随实际增加
	dir, err := ioutil.ReadAll(dirPath)
	if err != nil {
		return nil,err
	}
	PathSep := string(os.PathSeparator) //系统路径分割符
	for _,file := range dir {
		files = append(files, dirPath + PathSep + file.Name() + "\n")

	}
	return files,nil
}

func OpenFile(filePath string) (datas string) {
	if err != nil {
		return ""
	}
	return string(data)
}
