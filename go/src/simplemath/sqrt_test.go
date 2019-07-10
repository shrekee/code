package simplemath

/**
 * author       : liwenqiang
 *creating_time : 19-7-10 下午5:11
 * file_name    : sqrt_test.py
 * IDE          : GoLand
**/
import "testing"

func TestSqrt(t *testing.T) {
	v := Sqrt(16)
	if v != 4 {
		t.Errorf("Sqrt(16) failed. Got %v, expected 4.", v)
	}
}