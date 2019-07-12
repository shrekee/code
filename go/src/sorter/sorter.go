package main

/**
 * author       : liwenqiang
 *creating_time : 19-7-11 下午5:50
 * file_name    : sorter.py
 * IDE          : GoLand
**/
import "flag"
import "fmt"
import "bufio"
import "io"
import "strconv"
import "os"

var (
	infile *string = flag.String("i","infile", "File contains values for sorting")
	outfile *string = flag.String("o", "outfile", "File to receive sorted values")
	algorithm *string = flag.String("a", "qsort", "Sort algorithm")
)

func readValues(infile string) (values []int, err error) {
	file, err := os.Open(infile)
	if err != nil {
		fmt.Println("Failed to open the input file ", infile)
		return
	}
	defer file.Close()

	br := bufio.NewReader(file)
	values = make([]int, 0)
	for {
		line, isPrefix, err1 := br.ReadLine()
		if err1 != nil {
			if err1 != io.EOF {
				err =  err1
			}
			break
		}
		if isPrefix {
			fmt.Println("A so long line, seems unexpected.")
			return
		}
		str := string(line)
		fmt.Println("line is ", line, "\tstr is ", str)
		value, err1 := strconv.Atoi(str)
		if err1 != nil {
			err = err1
			return
		}

		values = append(values, value)
	}
	return
}


func main()  {
	flag.Parse()
	if infile != nil {
		fmt.Println("infile =", *infile, "outfile =", *outfile, "algorithm = ", *algorithm)
	}
	//fmt.Println(*infile, *outfile, *algorithm)

	values, err := readValues(*infile)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Read values: ",values, "len is ", len(values))
	}
	err = writeVaules(values, *outfile)
	if err != nil {
		fmt.Println("Failed to write to outfile ", *outfile)
	} else {
		fmt.Println("Wrote to outfile ", *outfile)
	}
}

func writeVaules(vaules []int, outfile string) error  {
	file, err := os.Create(outfile)
	if err != nil {
		fmt.Println("Failed to create the output file ", outfile)
		return err
	}

	defer file.Close()
	for _, value := range vaules {
		str := strconv.Itoa(value)
		file.WriteString(str + "\n")
	}
	return nil
}