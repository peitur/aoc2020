package main

import (
	"fmt"
	"os"
	"bufio"
	"regexp"
	"strings"
	"strconv"
)

type Passport struct {
	byr string 
	iyr string
	eyr string
	hgt string
	hcl string
	ecl string
	pid string
	cid string
}


func NewPassport() *Passport {
	return new( Passport )
}

func (p *Passport ) validByr( ) bool { 
	b, err := strconv.ParseInt( p.byr, 10, 64 )	
	if err != nil { return false }
	if ! regexp.MustCompile(`^[0-9]{4}$`).MatchString( p.byr ) { return false  }
	if b < 1920 || b > 2002 { return false }

	return true 
}

func (p *Passport ) validIyr( ) bool { 
	b, err := strconv.ParseInt( p.iyr, 10, 64 )
	if err != nil { return false }
	if ! regexp.MustCompile(`^[0-9]{4}$`).MatchString( p.iyr ) { return false  }
	if b < 2010 || b > 2020 { return false }
	return true 
}

func (p *Passport ) validEyr( ) bool {
	b, err := strconv.ParseInt( p.eyr, 10, 64 )	
	if err != nil { return false }
	if ! regexp.MustCompile(`^[0-9]{4}$`).MatchString( p.eyr ) { return false  }
	if b < 2020 || b > 2030 { return false }
	return true
}

func (p *Passport ) validHgt( ) bool {
	if p.hgt == "" { return false }

	parts := regexp.MustCompile(`^([0-9]+)(cm|in)$`).FindAllStringSubmatch( p.hgt, -1 )
	if len( parts ) > 0 {

		if len( parts[0] ) == 3 {
			b, err := strconv.ParseInt( string( parts[0][1]), 10, 64 )	
			if err != nil { return false }

			if string( parts[0][2] ) == "cm" {
				if b < 150 || b > 193 { return false }

			} else if string( parts[0][2] ) == "in" {
				if b < 59 || b > 76 { return false }
			}
		}
	}else{ 
		return false 
	}

	return true 
}

func (p *Passport ) validEcl( ) bool { 
	if regexp.MustCompile(`^(amb|blu|brn|gry|grn|hzl|oth)$`).MatchString( p.ecl ){ return true }
	return false
}

func (p *Passport ) validHcl( ) bool { 
	if regexp.MustCompile(`^#[0-9a-f]{6}$`).MatchString( p.hcl ){ return true }
	return false
}

func (p *Passport ) validPid( ) bool { 
	if regexp.MustCompile(`^[0-9]{9}$`).MatchString( p.pid ){ return true }


	return false
}

func (p *Passport ) validCid( ) bool { 
	return true
}

func (p *Passport ) validPassportIsh( ) bool {
	if p.byr == "" { return false } 
	if p.iyr == "" { return false } 
	if p.eyr == "" { return false } 
	if p.hgt == "" { return false }
	if p.hcl == "" { return false } 
	if p.ecl == "" { return false }
	if p.pid == "" { return false }
//	if p.cid == "" { return false }
	return true
}

func (p *Passport ) validPassportStrict( ) bool {
	if ! p.validByr() { return false }
	if ! p.validIyr() { return false }
	if ! p.validEyr() { return false }
	if ! p.validHgt() { return false }
	if ! p.validHcl() { return false }
	if ! p.validEcl() { return false }
	if ! p.validPid() { return false }
	if ! p.validCid() { return false }
	return true
}


func (p *Passport) printPassport() {

	if 	p.validPassportIsh() { 
		fmt.Printf("%-6s", "OK" )
	}else{
		fmt.Printf("%-6s", "BAD" )
	}

	if 	p.validPassportStrict() { 
		fmt.Printf("%-6s", "OK" )
	}else{
		fmt.Printf("%-6s", "BAD" )
	}

	fmt.Printf(" >> byr: %-6s", p.byr )
	fmt.Printf(" >> iyr: %-6s", p.iyr )
	fmt.Printf(" >> eyr: %-6s", p.eyr )
	fmt.Printf(" >> hgt: %-8s", p.hgt )
	fmt.Printf(" >> hcl: %-8s", p.hcl )
	fmt.Printf(" >> ecl: %-10s", p.ecl )
	fmt.Printf(" >> pid: %-12s", p.pid )
	fmt.Printf(" >> cid: %-6s", p.cid )
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

			if ld[0] == "byr" { r.byr = strings.TrimSpace( ld[1] ) }
			if ld[0] == "iyr" { r.iyr = strings.TrimSpace( ld[1] ) }
			if ld[0] == "eyr" { r.eyr = strings.TrimSpace( ld[1] ) }
			if ld[0] == "hgt" { r.hgt = strings.TrimSpace( ld[1] ) }
			if ld[0] == "hcl" { r.hcl = strings.TrimSpace( ld[1] ) }
			if ld[0] == "ecl" { r.ecl = strings.TrimSpace( ld[1] ) }
			if ld[0] == "pid" { r.pid = strings.TrimSpace( ld[1] ) }
			if ld[0] == "cid" { r.cid = strings.TrimSpace( ld[1] ) }

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
			pl = append( pl, strings.TrimSpace( data[i] ) )
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

	var validPassportsIsh int = 0
	var validPassportsStrict int = 0
	var passports []*Passport = parsePassports( data )

	fmt.Printf("Found %d passorts in document\n", len( passports ))
	for i := range( passports ){
		p := passports[i]
		if p.validPassportIsh(){
			validPassportsIsh++
		}
	}

	for i := range( passports ){
		p := passports[i]
		if p.validPassportStrict(){
			validPassportsStrict++
		}
	}

	for i := range( passports ){
		passports[i].printPassport()
	}


	fmt.Printf("Number of valid(ish) passports : %d\n", validPassportsIsh )
	fmt.Printf("Number of valid strict passports : %d\n", validPassportsStrict )

}