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

	for n := range( data ){
		for i := 0; i < 64; i++ {
			fmt.Printf( "%c", charPositionInLine( []byte( data[n] ), i ) )
		}
		fmt.Println()
	}
}