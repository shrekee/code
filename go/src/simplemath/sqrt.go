package simplemath

/**
 * author       : liwenqiang
 *creating_time : 19-7-10 下午5:09
 * file_name    : sqrt.py
 * IDE          : GoLand
**/
import "math"

func Sqrt(i int) int  {
	v := math.Sqrt(float64(i))
	return int(v)

}