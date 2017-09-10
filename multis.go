package main

import (
	"fmt"
	"strings"
	"os"
	"strconv"
	"io/ioutil"
	//"MultiTool/gifsMaker"
	//"MultiTool/gifsMaker"
	"MultiTool/network"
)

func main() {
	//gifsMaker.MakeGif()
	network.Fetch()
	//gifsMaker.Ec()
	//dup()
}

//print the arguments passed to the program
func echo (){
	fmt.Println(strings.Join(os.Args [1:], " "))
}
//print the name of the path to the program and the arguments with the index
//print the name of the path to the program and the arguments with the index
func echo2 (){
	s,sep := "",""

	for inx, arg := range os.Args[0:]{
		s+=sep+strconv.Itoa(inx)+sep+arg
		sep=" "
	}
	fmt.Println(s)
}
//print the occurence of duplicates line in the files passed as args
func dup (){
	counts:= make (map[string]int)
	for _, filename:=range os.Args [1:]{
		data,err:=ioutil.ReadFile(filename)
		if err !=nil {
			fmt.Printf("dup : %v",err)
			continue
		}
		for _,line:=range strings.Split(string(data), "\n"){
			counts[line]++
		}
	}
	for line, n:=range counts{
		if n>1{
			fmt.Printf("%d\t%s\n", n, line)

		}
	}
}


