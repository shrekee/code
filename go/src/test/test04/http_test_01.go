package main

import "net/http"

/**
 * author       : liwenqiang
 *creating_time : 19-7-18 下午3:38
 * file_name    : http_test_01.py
 * IDE          : GoLand
**/
type Client struct {
	Transport http.RoundTripper

}