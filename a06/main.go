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

func groupKeys( group map[byte]int ) []byte {
	var kList []byte = []byte{}
	for k := range( group ){
		kList = append( kList, k )
	}
	return kList
}

func parseGroupItems( group string ) map[byte]int {
	var keys map[byte]int = map[byte]int{}

	for _, b := range( []byte( group) ){
		if _, ok := keys[ b ]; !ok {
			keys[ b ] = 0
		}
		keys[b]++
	}

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
	var groupSumAns int = 0
	for _, item := range( groups ){
		groupMap := parseGroupItems( item )
		groupAns := len( groupKeys( groupMap ) )

		groupSumAns += groupAns
	}
	fmt.Printf("Sum Anyone: %d\n", groupSumAns )
}