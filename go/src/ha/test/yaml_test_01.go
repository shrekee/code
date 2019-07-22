package main

/**
 * author       : liwenqiang
 *creating_time : 19-7-22 下午4:15
 * file_name    : yaml_test_01.py
 * IDE          : GoLand
**/
import (
	"github.com/spf13/viper"
	//"github.com/fsnotify/fsnotify"
	"fmt"
	//"os"
	"time"
)

func main()  {
	viper.SetConfigName("config")
	viper.AddConfigPath("/home/lsing/code/code/go/src/ha")
	err := viper.ReadInConfig()

	if err!= nil{
		fmt.Printf("Fatal errors: config file %s\n",err.Error())
	}else{
		fmt.Println("Done: read config.yaml file content.")
	}
	s :=viper.AllKeys()
	fmt.Println(s)
	fmt.Println(viper.AllSettings(),viper.AllSettings()["name"])
	//fmt.Println(viper.)
	fmt.Println("================")
	time.Time{}


}
