package main

import (
	"fmt"
	"strconv"
	"os"
	"bufio"
	"regexp"
)

type Bag struct {
	color string
	ctype string
	contains int
}

func str2int( s string ) (int64, error) {
	u, err := strconv.ParseInt( s, 10, 64)
	if err != nil {
		return 0, err
	}
	return u, nil
}

func readFile( filename string ) ([]string, error) {
	var ret []string = []string{}

	fd, err := os.Open( filename )

    if err != nil {
		return  nil, err
	}

    defer fd.Close()

    scanner := bufio.NewScanner(fd)
    for scanner.Scan() {
		v := scanner.Text()
		if err != nil {
			return nil, err
		}
		ret = append( ret, v )
	}

    if err := scanner.Err(); err != nil {
		return nil, err 
	}
	return ret, nil
}



func main( ){
	data,err := readFile( "input")
	if err != nil {
		fmt.Println("ERROR, no input file: ", err)
		os.Exit(1)
	}


}