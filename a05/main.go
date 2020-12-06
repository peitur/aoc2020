package main

import (
	"fmt"
	"os"
	"bufio"
	"regexp"
	"math"
	"sort"
//	"strings"
//	"strconv"
)

type PositionRaw struct {
	row []byte
	column []byte
}


const LENGTH=128
const WIDTH=8

func NewPositionRaw( r, c []byte ) *PositionRaw {
	var p *PositionRaw = new( PositionRaw )
	p.row = r
	p.column = c
	return p
}

func (p *PositionRaw ) Print( ) {
	r := p.CalcRow()
	c := p.CalcColumn()
	sid := p.SeatId()
	fmt.Printf( "R:%8s[%4d] C:%4s[%4d] Seat: %d\n", p.row,r, p.column,c, sid )
}

func (p *PositionRaw ) CalcRow( ) int {
	var pos, a, b int = 0, 0, int( math.Pow(2, float64( len( p.row ) ) ) ) - 1
	for _, c := range( p.row ){

		// F means to take the lower half
		// B means to take the upper half
		d := ( b - a ) / 2 + 1
		if c == 'F'{
			b = b - d
		} else if c == 'B' {
			a = a + d
		}
		pos = a
//		fmt.Printf( "Row: A:%d B:%d -> RowP:%d\n", a, b, pos )
	}

	return pos
}

func (p *PositionRaw ) CalcColumn( ) int {
	var pos, a, b int = 0, 0, int( math.Pow(2, float64( len( p.column ) ) ) ) - 1
	for _, c := range( p.column ){

		// R means to take the lower half
		// F means to take the upper half
		d := ( b - a ) / 2 + 1
		if c == 'L'{
			b = b - d
		} else if c == 'R' {
			a = a + d
		}
		pos = a
//		fmt.Printf( "Columnt: A:%d B:%d -> ColP:%d\n", a, b, pos )
	}

	return pos
}

func (p *PositionRaw ) SeatId( ) int {
	r := p.CalcRow()
	c := p.CalcColumn()
	sid := r * 8 + c
	return sid
}

func splitDescriptionRow( s string ) []byte {
	parts := regexp.MustCompile( `^([BF]+)([RL]+)$` ).FindAllStringSubmatch( s, -1 )

	if len( parts[0] ) != 3{
		return []byte{}
	}

	return []byte(parts[0][1])
}

func splitDescriptionColumn( s string ) []byte {
	parts := regexp.MustCompile( `^([BF]+)([RL]+)$` ).FindAllStringSubmatch( s, -1 )

	if len( parts[0] ) != 3 {
		return []byte{}
	}

	return []byte(parts[0][2])
}

func iterData( data []string ) []*PositionRaw {
	var ret []*PositionRaw 

	for _,item := range( data ){
		r := splitDescriptionRow( item )
		c := splitDescriptionColumn( item )
		ret = append( ret, NewPositionRaw( r, c )) 
	}

	return ret
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

	var topSid int = 0
	for _, p := range( iterData( data ) ){
		sid := p.SeatId()
		if sid > topSid{
			topSid = sid
		}
	}
	fmt.Printf("\nTop ID: %d\n", topSid )

	var slist []int = []int{}
	for _, p := range( iterData( data ) ){
		sid := p.SeatId()
		slist = append( slist, sid )
	}
	
	sort.Ints( slist )
	mySeat := 0
//	startSeat := slist[0]
//	lastSeat  := slist[ len( slist ) - 1 ] 
	for i := 1; i < len( slist ) - 1; i++ {
		pseat := slist[i-1]
		xseat := slist[i] 
		nseat := slist[i+1]
		if (xseat - pseat) > 1 || (nseat - xseat) > 1{
			fmt.Printf("P:%d X:%d N:%d \n", pseat, xseat, nseat)
			mySeat = xseat
		}
	}
	fmt.Printf("MySeat: %d\n", mySeat - 1 )
}