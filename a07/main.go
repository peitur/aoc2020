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
	numof int64
	contains []*Bag
}


func  NewBag( ctype, color string, numof int64 ) *Bag {
	var b *Bag = new( Bag )
	b.color = color
	b.ctype = ctype
	b.numof = numof
	return b
}

func (b *Bag) AddBag( a * Bag ) {
	b.contains = append( b.contains, a )
}

func ( b *Bag ) Contains( ) []*Bag {
	return b.contains
}

func (b *Bag) IsTaget( ) bool {
	if b.ctype == "shiny" && b.color == "gold" {
		return true
	}
	return false
}

func (b *Bag) CanContainTarget( ) bool {
	if b.ctype == "bright" && b.color == "white" {
		return true
	}else if b.ctype == "muted" && b.color == "yellow" {
		return true
	}
	return false
}

func ( b *Bag ) Print( ) {
	fmt.Println( b )
	for _, x := range( b.contains ){
		x.Print()
	}
}

func ( b *Bag ) Name( ) string {
	return fmt.Sprintf("%s-%s", b.ctype, b.color )
}

func parseBagContent( ruleSubs string ) []*Bag {
	var bags []*Bag
	
	bp := regexp.MustCompile( `([0-9]+)\s+([a-z]+)\s+([a-z]+).*`)
	for _, rule := range( regexp.MustCompile(`,`).Split( ruleSubs , -1 ) ){
		m := bp.FindAllStringSubmatch( rule, -1 )
		if len( m ) > 0 && len( m[0] ) > 1 {
			var color string = m[0][3]
			var ctype string = m[0][2]
			numof, _ := str2int( m[0][1] )

			bags = append( bags, NewBag( ctype, color, numof ) )
		}
	}
	
	return bags
}

func checkTarget( b *Bag, mp map[string][]*Bag ) bool {
	var ret bool = false
	if b.IsTaget(){
		return true
	}else{
		for _, c := range( mp[ b.Name() ] ) {
			if checkTarget( c, mp ){
				return true
			}
		}
	}
	return ret
}

func parseBagRule( rule string ) *Bag {
	m := regexp.MustCompile( `^([a-z]+)\s+([a-z]+)\s+bags\s+contain\s+(.+)$` ).FindAllStringSubmatch( rule, -1 )
	var b * Bag = NewBag( m[0][1], m[0][2], 0 )
	for _, ab := range( parseBagContent( m[0][3]) ) {
		b.AddBag( ab )
	}
	return b
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
	var files []string = []string{"test", "input"}
//	var files []string = []string{"test"}

	for _, fname := range( files ) {
		data,err := readFile( fname )
		if err != nil {
			fmt.Println("ERROR, no input file: ", err)
			os.Exit(1)
		}

		var bmap map[string][]*Bag = make( map[string][]*Bag )
		var bags []*Bag = []*Bag{}
		var cancontain int = 0
		for _, rule := range( data ){
			x := parseBagRule( rule )
			bags = append( bags, x )
			ns := x.Name()
			if _, ok := bmap[ ns ]; !ok {
				bmap[ns] = []*Bag{}
			}
			for _, c := range( x.Contains() ){
				bmap[ns] = append( bmap[ns], c )
			}
			
		}
		for name := range( bmap ){
			fmt.Printf( "%s -> %d\n", name, len( bmap[name]) )
			for _, i := range( bmap[name] ) {
				fmt.Printf("\t%s\n", i.Name( ))
			}
		}

		for _, bag := range( bags ){
			fmt.Printf( "%s -> %b \n", bag.Name(), checkTarget( bag, bmap) )
			if checkTarget( bag, bmap){
				cancontain++
			}
		}

		fmt.Printf("Can contain: %d \n", cancontain - 1 )


	}
}