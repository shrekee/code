package main

/**
 * author       : liwenqiang
 *creating_time : 19-7-12 下午2:35
 * file_name    : myplayer.py
 * IDE          : GoLand
**/
import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"pkg/myplayer/mlib"
	"pkg/myplayer/mp"
	"player/library"
)

var lib *library.MusicManager
var id int = 1
var ctr1, signal chan int

func handleLibCommands(tokens []string)  {
	switch tokens[1] {
	case "list":
		for i := 0; i < lib.Len(); i++ {
			e, _ := lib.Get(i)
			fmt.Println(i+1, ":", e.Name, e.Artist, e.Source, e.Type)
		}
	case "add":
		if len(tokens) == 6 {
			id++
			lib.Add(&library.MusicEntry{strconv.Itoa(id),
			tokens[2], tokens[3], tokens[4], tokens[5]})
		} else {
			fmt.Println("USAGE: lib add <name> <artist> <source> <type>")
		}
	case "remove":
		if len(tokens) == 3 {
			v,err := strconv.Atoi(tokens[2])
			if err != nil {
				fmt.Println("Failed to convert str to int.")
				return
			}
			lib.Remove(v)
		} else {
			fmt.Println("USAGE: lib remove <index>")
		}
	default:
		fmt.Println("Unrecognized lib command:", tokens[1])
	}
}

func handlePlayerCommand(tokens []string)  {
	if len(tokens) != 2 {
		fmt.Println("USAGE: play <name>")
		return
	}

	e := lib.Find(tokens[1])
	if e != nil {
		fmt.Println("The music ", tokens[1], "does not exist.")
		return
	}
	mp.Play

}














