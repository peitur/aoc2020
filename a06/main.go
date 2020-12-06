package main

import (
	"fmt"
	"os"
	"bufio"
//	"regexp"
//	"math"
//	"sort"
//	"strings"
//	"strconv"
)



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

func main( ) {
	data,err := readFile( "input")
	if err != nil {
		fmt.Println("ERROR, no input file: ", err)
		os.Exit(1)
	}

}