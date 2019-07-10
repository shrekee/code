package simplemath

/**
 * author       : liwenqiang
 *creating_time : 19-7-10 下午3:39
 * file_name    : add_test.py
 * IDE          : GoLand
**/
import (
	"testing"
)

func TestAdd1(t *testing.T)  {
	r := Add(1,2)
	if r != 3 {
		t.Errorf("Add(1, 2) failed. Got %d, expected 3.", r)
	}

}