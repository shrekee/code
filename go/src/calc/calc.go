package main

/**
 * author       : liwenqiang
 *creating_time : 19-7-10 下午1:53
 * file_name    : calc.py
 * IDE          : GoLand
**/
import (
	"os"
	"fmt"
	"simplemath"
	"strconv"
)

var Usage = func() {
	fmt.Println("USAGE: calc command [arguments] ...")
	fmt.Println("\nThe  commands are:\n\tadd\tAddition of two values.\n\tsqrt\tSquare" +
		" root of a non-negative value.")
	
}

func main()  {
	args := os.Args
	//fmt.Println("len is ", len(args),args)
	if args == nil || len(args) < 3 {
		Usage()
		return
	}

	switch args[1] {
	case "add":
		if len(args) != 4 {
			fmt.Println("USAGE: calc add <integer1><integer2>")
			fmt.Println("len is (in add): ", len(args))
			return
		}
		v1, err1 := strconv.Atoi(args[2])
		v2, err2 := strconv.Atoi(args[3])
		if err1 != nil || err2 != nil {
			fmt.Println("USAGE: calc add <integer1><integer2>")
			return
		}
		ret := simplemath.Add(v1, v2)
		fmt.Println("Result: ", ret)
	case "sqrt":
		if len(args) != 3 {
			fmt.Println("USAGE: calc sqrt<integer>")
			return
		}
		v, err := strconv.Atoi(args[2])
		if err != nil {
			fmt.Println("USAGE: calc sqrt <integer>")
			return
		}
		ret := simplemath.Sqrt(v)
		fmt.Println("Result: ", ret)
	default:
		//fmt.Println(args[1])
		Usage()
	}

}
