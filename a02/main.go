package main

import (
	"fmt"
	"strconv"
	"os"
	"bufio"
	"regexp"
)


type DataContent struct {
	Min int64
	Max int64
	Char byte
	String string
}

func charCount( c byte, s string ) int {
	occ := 0

	bs := []byte( s )
	for x := range( bs ){
		if c == bs[x]{
			occ++
		}
	}
	return occ
}

func parseLine( line string ) *DataContent {
	re := regexp.MustCompile( `\s*([0-9]+)[-]+([0-9]+)\s*([a-zA-Z]+)\s*:\s*(\S+)$` )
	
	rx := re.FindAllStringSubmatch( line , -1 )[0]

	min, _ := strconv.ParseInt( rx[1], 10, 64 )
	max, _ := strconv.ParseInt( rx[2], 10, 64 )
	chr := []byte( rx[3] )
	str := rx[4]

	var d *DataContent = new( DataContent )
	d.Min = min
	d.Max = max
	d.Char = chr[0]
	d.String = str

	return d
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

	var nvalid int = 0
	for x := range( data ) {
		line := data[x]

		item := parseLine( line )
		cnum := int64( charCount( item.Char, item.String ) )

		if item.Min < cnum && item.Max <= cnum {
			fmt.Printf( "OK [%d:%d] > %s \n", len( item.String), cnum, line )
			nvalid++
		}else{
			fmt.Printf( "FF [%d:%d] > %s \n", len( item.String), cnum, line )
		} 

//		fmt.Println( line )
	}

	fmt.Println("Number of valid passwords: ", nvalid )
}