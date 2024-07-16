package main

import ("fmt"
"math/rand")

// Start Functions

// Conversion Functions

func HEX(by *byte){
    var h1, h2 int
    var h1s, h2s string
    h2 = int(*by) % 16
    h1 = (int(*by) - h2)/16
    switch h1{
        case 0:
            h1s = "0"
            break
        case 1:
            h1s = "1"
            break
        case 2:
            h1s = "2"
            break
        case 3:
            h1s = "3"
            break
        case 4:
            h1s = "4"
            break
        case 5:
            h1s = "5"
            break
        case 6:
            h1s = "6"
            break
        case 7:
            h1s = "7"
            break
        case 8:
            h1s = "8"
            break
        case 9:
            h1s = "9"
            break
        case 10:
            h1s = "A"
            break
        case 11:
            h1s = "B"
            break
        case 12:
            h1s = "C"
            break
        case 13:
            h1s = "D"
            break
        case 14:
            h1s = "E"
            break
        case 15:
            h1s = "F"
            break
        case 16:
            h1s = "G"
            break
    }
    switch h2{
        case 0:
            h2s = "0"
            break
        case 1:
            h2s = "1"
            break
        case 2:
            h2s = "2"
            break
        case 3:
            h2s = "3"
            break
        case 4:
            h2s = "4"
            break
        case 5:
            h2s = "5"
            break
        case 6:
            h2s = "6"
            break
        case 7:
            h2s = "7"
            break
        case 8:
            h2s = "8"
            break
        case 9:
            h2s = "9"
            break
        case 10:
            h2s = "A"
            break
        case 11:
            h2s = "B"
            break
        case 12:
            h2s = "C"
            break
        case 13:
            h2s = "D"
            break
        case 14:
            h2s = "E"
            break
        case 15:
            h2s = "F"
            break
        case 16:
            h2s = "G"
            break
    }
    fmt.Printf("0x" + h1s + h2s)
}

func BIT(by *byte){
    var fy byte = *by
    fys := [8]rune{}
    bitval := "0b"
    for i := 0; i < 8; i++{
        // Find if 1 or 0
        fys[i] = rune((fy%2)+48)
        fy /= 2
    }
    for i := 7; i >= 0; i--{
        bitval += string(fys[i])
    }
    fmt.Printf(bitval)
}

// Byte Urnary

func INC(by *byte){
    *by++
}

func DNC(by *byte){
    *by--
}

// Byte by value opperations

func SET(by *byte, n int){
    *by = byte(n)
}

func ADD(by *byte, n int){
    *by += byte(n)
}

func SUB(by *byte, n int){
    *by -= byte(n)
}

func MULT(by *byte, n int){
    *by *= byte(n)
}

func DIV(by *byte, n int){
    *by /= byte(n)
}

func MOD(by *byte, n int){
    *by %= byte(n)
}

func RAND(by *byte){
    *by = byte(rand.Intn(255))
}

// byte by byte opperations

func PTR(by, opby *byte){
    *by = *opby
}

func MOV(by, opby *byte){
    *opby = *by
    *by = 0
}

func BADD(by, opby *byte){
    *by += *opby
}

func BSUB(by, opby *byte){
    *by += *opby
}

func BMULT(by, opby *byte){
    *by *= *opby
}

func BDIV(by, opby *byte){
    *by /= *opby
}

func BMOD(by, opby *byte){
    *by %= *opby
}

func HELP(op int) {
	switch op{
	case 0:
		//Pages
		functs := [...]string{"Type as form: line(integer) opperation(All caps) input(integer) \n", 
		"Search page with: # ?", 
		"_____PAGES_____ \n", 
		"1 - Byte by Value. \n",
		"2 - Byte by Byte. \n",
		"3 - Byte Unary. \n", 
		"4 - Byte Conversions \n", 
		"5 - Program"}
		fmt.Println(functs)
		break
	case 1:
		//Byte by Value
        functs := [...]string{
            "_____OPPERATIONS_____ \n", 
            "SET - Sets the value of a byte \n", "ADD - Adds to the value of a byte \n", 
            "SUB - Subtracts a value from a byte \n", "MULT - Multiplies a byte by a value \n", 
            "DIV - Divides a byte by a value \n", 
            "MOD - Gets the remainder of a byte divided by a value"}
		fmt.Println(functs)
		break
	case 2:
		//Byte by Byte
        functs := [...]string{
            "_____OPPERATIONS_____ \n", 
            "PTR - Set a byte equal to another byte \n", 
            "MOV - Move the value of one byte to another and set the original to zero \n", 
            "BADD - Adds the value of one byte to another \n", 
            "BSUB - Subtracts the value of one byte from another \n", 
            "BMULT - Multipies the value of one byte with another \n", 
            "BDIV - Divides the value of one byte from another \n", 
            "BMOD - Gets the remainder of a byte divided by another byte \n"}
		fmt.Println(functs)
		break
	case 3:
		//Byte Unary
        functs := [...]string{
            "_____OPPERATIONS_____ \n", 
            "INC - Increment the value of a byte by 1", 
            "DNC - Decrement the value of a byte by 1", 
            "RAND - Randomizes the value of the byte \n"}
		fmt.Println(functs)
		break
	case 4:
		//Byte Conversions
        functs := [...]string{
            "_____OPPERATIONS_____ \n", 
            "HEX - Writes out the hexadecimal value of the byte \n", 
            "BIT - Writes out the binary value of the byte \n"}
		fmt.Println(functs)
		break
	case 5:
		//Program
        functs := [...]string{
            "_____OPPERATIONS_____ \n", 
            "PRINT - Print the table",
            "PRINTFT - Print the table from one line to another (Note a <= b)",
            "EXIT - Leaves the program \n", 
            "RESET - Resets the Program \n"}
		fmt.Println(functs)
		break
	default:
		fmt.Println("Page Not Found")
		break
	}
}
// End Functions

// Print Funct
func PrintByteLine(by *byte, ln int){
    fmt.Println(ln, by, *by)
}

func PrintByteTable(bys *[16]byte){
    // Writes out the byte table
    for i := 0; i < len(bys); i++{
        PrintByteLine(&bys[i], i)
    }
}

func PrintByteTableFromTo(bys *[16]byte, start, end int){
    // Writes out the byte table
    if (start > len(bys) || start < 0) { 
        fmt.Println("Value of start exceeds accepted range of list.") 
        return
    }
    if (end > len(bys) || end < start) { 
        fmt.Println("Value of end exceeds accepted range of list.")
        return
    }
    for i := start; i <= end; i++{
        PrintByteLine(&bys[i], i)
    }
}