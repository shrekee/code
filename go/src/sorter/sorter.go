package main

/**
 * author       : liwenqiang
 *creating_time : 19-7-11 下午5:50
 * file_name    : sorter.py
 * IDE          : GoLand
**/
import "flag"
import "fmt"

var (
	infile *string = flag.String("i","infile", "File contains values for sorting")
	outfile *string = flag.String("o", "outfile", "File to receive sorted values")
	algorithm *string = flag.String("a", "qsort", "Sort algorithm")
)


func main()  {
	flag.Parse()
	if infile != nil {
		fmt.Println("infile =", *infile, "outfile =", *outfile, "algorithm = ", *algorithm)
	}
	//fmt.Println(*infile, *outfile, *algorithm)

}
