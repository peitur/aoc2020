package main

import (
	"fmt"
	"os"
	"bufio"
//	"regexp"
//	"math"
//	"sort"
	"strings"
//	"strconv"
)

func parseGroupItems( group string ) map[byte]int {
	var keys map[byte]int = map[byte]int{}

	return keys
}


func parseGroups( data []string ) []string {

	var groups []string = []string{}
	pl := []string{}
	for i := range( data ){
		if len( data[i]) == 0 {
			groups = append( groups, strings.Join( pl, "" ) )
			pl = []string{}
		}else{
			pl = append( pl, strings.TrimSpace( data[i] ) )
		}
	}

	if len( pl ) > 0{
		groups = append( groups, strings.Join( pl, "" ) )
	}
	
	return groups

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

func main( ) {
	data,err := readFile( "input")
	if err != nil {
		fmt.Println("ERROR, no input file: ", err)
		os.Exit(1)
	}

	groups := parseGroups( data )
	for _, item := range( groups ){
		fmt.Println( item )
	}

}