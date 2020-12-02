package main

import (
	"fmt"
	"strconv"
	"os"
	"bufio"
)

const COMPLE=2020


func isCompleted( s int64 ) bool {
	if s == COMPLE{
		return true
	}
	return false
}

func str2int( s string ) (int64, error) {
	u, err := strconv.ParseInt( s, 10, 64)
	if err != nil {
		return 0, err
	}
	return u, nil
}

func readFile( filename string ) ([]int64, error) {
	var ret []int64 = []int64{}

	fd, err := os.Open( filename )

    if err != nil {
		return  nil, err
	}

    defer fd.Close()

    scanner := bufio.NewScanner(fd)
    for scanner.Scan() {
		v, err := str2int( scanner.Text() )
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

func asum( ) int64 {

	data, err := readFile("input" )
	if err != nil {
		fmt.Println("ERROR: ", err )
		os.Exit(1)
	}
	
	for x := range( data ){
		a := data[x]

		for y := range( data ){
			b := data[y]
			sum := a + b

			if isCompleted( int64( sum ) ){
				prod := a * b
				fmt.Printf( "%d + %d = %d -> %d\n", a, b, sum, prod )
				return prod
			}
		}


	}
	return 0
}

func bsum() int64 {

	data, err := readFile("input" )
	if err != nil {
		fmt.Println("ERROR: ", err )
		os.Exit(1)
	}
	
	for x := range( data ){
		a := data[x]

		for y := range( data ){
			b := data[y]

			for z := range( data ){
				c := data[z]
				sum := a + b + c
	
				if isCompleted( int64( sum ) ){
					prod := a * b * c 
					fmt.Printf( "%d + %d + %d = %d -> %d\n", a, b, c, sum, prod )
					return prod
				}
			}
		}


	}
	return 0	
}

func main( ){
	asum()
	bsum()
}