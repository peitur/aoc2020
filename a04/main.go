package main

import (
	"fmt"
	"strconv"
	"os"
	"bufio"
	"regexp"
)

type Passport struct {
	byr int64 
	iyr int64
	eyr int64
	hgt string
	hcl string
	ecl string
	pid int64
	cid int64
}


func NewPassport() *Passport {
	return new( Passport )
}

func (p *Passport ) validPassportIsh( ) bool {
	if p.byr == 0 { return false }
	if p.iyr == 0 { return false }
	if p.eyr == 0 { return false }
	if p.hgt == "" { return false }
	if p.hcl == "" { return false }
	if p.ecl == "" { return false }
	if p.pid == 0 { return false }
//	if p.cid == 0 { return false }
	return true
}

func (p *Passport) printPassport() {
	
	if 	p.validPassportIsh() { 
		fmt.Printf("%-6s", "OK" )
	}else{
		fmt.Printf("%-6s", "BAD" )
	}

	fmt.Printf(" >> byr: %-6d", p.byr )
	fmt.Printf(" >> iyr: %-6d", p.iyr )
	fmt.Printf(" >> eyr: %-6d", p.eyr )
	fmt.Printf(" >> hgt: %-8s", p.hgt )
	fmt.Printf(" >> hcl: %-8s", p.hcl )
	fmt.Printf(" >> ecl: %-10s", p.ecl )
	fmt.Printf(" >> pid: %-12d", p.pid )
	fmt.Printf(" >> cid: %-6d", p.cid )
	fmt.Printf("\n")
}

func parsePassport( data []string ) *Passport {
	var r *Passport = NewPassport()

	rx := regexp.MustCompile(`(\s+|\n)`)
	fx := regexp.MustCompile(`(:)`)

	for i := range( data ){
		xl := rx.Split( data[i], -1 )
		for x := range( xl ){
			ld := fx.Split( xl[x] , -1 )

			if ld[0] == "byr" { r.byr, _ = strconv.ParseInt( ld[1], 10, 64 ) }
			if ld[0] == "iyr" { r.iyr, _ = strconv.ParseInt( ld[1], 10, 64 ) }
			if ld[0] == "eyr" { r.eyr, _ = strconv.ParseInt( ld[1], 10, 64 ) }
			if ld[0] == "hgt" { r.hgt = ld[1] }
			if ld[0] == "hcl" { r.hcl = ld[1] }
			if ld[0] == "ecl" { r.ecl = ld[1] }
			if ld[0] == "pid" { r.pid, _ = strconv.ParseInt( ld[1], 10, 64 ) }
			if ld[0] == "cid" { r.cid, _ = strconv.ParseInt( ld[1], 10, 64 ) }

		}
	}

	return r
}

func parsePassports( data []string ) []*Passport {

	passports := []*Passport{}
	pl := []string{}
	for i := range( data ){
		if len( data[i]) == 0 {
			passports = append( passports, parsePassport( pl ))
			pl = []string{}
		}else{
			pl = append( pl, data[i] )
		}
	}

	if len( pl ) > 0{
		passports = append( passports, parsePassport( pl ))
	}
	
	return passports
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

	var validPassports int = 0
	var passports []*Passport = parsePassports( data )

	fmt.Printf("Found %d passorts in document\n", len( passports ))
	for i := range( passports ){
		p := passports[i]
		p.printPassport()

		if p.validPassportIsh(){
			validPassports++
		}
	}
	fmt.Printf("Number of valid(ish) passports : %d\n", validPassports )
}