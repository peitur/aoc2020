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

type Ansvers struct {
	ansvers []string
	size int
}

func NewAnsvers( astrings []string ) *Ansvers {
	var a *Ansvers = new( Ansvers )
	a.ansvers = astrings
	a.size = len( astrings )
	return a
}

func (a *Ansvers ) InOneAnsves( ) string {
	return strings.Join( a.ansvers, "" )
}

func (a *Ansvers ) Size( ) int {
	return a.size
}

func (a *Ansvers) AnsveredByAll( ) int {
	var items map[byte]int = parseGroupItems( a.InOneAnsves() )
	var sumSeen int = 0
	for item := range( items ){
		if items[item] == a.size {
			sumSeen++
		}
	}
	return sumSeen
}

/*
For each group, get map for 
Number of ansvers eq number of people
*/
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


func parseGroups( data []string ) []*Ansvers {

	var groups []*Ansvers
	pl := []string{}
	for i := range( data ){
		if len( data[i]) == 0 {
			groups = append( groups, NewAnsvers( pl ) )
			pl = []string{}
		}else{
			pl = append( pl, strings.TrimSpace( data[i] ) )
		}
	}

	if len( pl ) > 0{
		groups = append( groups, NewAnsvers( pl ) )
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
	var peopleSumAns int = 0

	for _, item := range( groups ){
		groupMap := parseGroupItems( item.InOneAnsves() )
		groupAns := len( groupKeys( groupMap ) )
		seenByAll := item.AnsveredByAll()

		groupSumAns += groupAns
		peopleSumAns += seenByAll
	}
	fmt.Printf("Sum Anyone: %d\n", groupSumAns )
	fmt.Printf("Sum SeenBy All in group: %d\n", peopleSumAns )

}