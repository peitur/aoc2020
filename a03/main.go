package main

import (
	"fmt"
	"strconv"
	"os"
	"bufio"

)



func isTree( c byte ) bool {
	if c == '#'{
		return true
	}
	return false
}

func isFree( c byte ) bool {
	if c == '.'{
		return true
	}
	return false
}

func charPositionInLine( line []byte, pos int ) byte {
	return line[ pos % len( line ) ]
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

	skips := [5][2]int{ {1,1},{3,1},{5,1},{7,1},{1,2} }

	for i := 0; i < len(  skips ); i++ {
		v := skips[i]
		var hitTree int = 0
		var hpos int = 0
	
		for vpos := 0; vpos < len( data ); {
			line := []byte( data[vpos] )

			c := charPositionInLine( line, hpos)
			if isTree( c ){
				hitTree++
			}

			hpos += v[0]
			vpos += v[1]
		}

		fmt.Printf("\nHit %d trees\n", hitTree )
	}
}