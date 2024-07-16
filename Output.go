package main
import ("fmt"
"math/rand"
"time")
//import "GitHub/GoByteTable/Functions.go"
//C:\Program Files\Go\src\

func main() {
    reset:
    fmt.Println("Help: 0 ?")
    // Sets the values of the bytes randomly
    bytes := [16]byte{}
    rand.Seed(int64(time.Now().Nanosecond()))
    for i := 0; i < len(bytes); i++{
        bytes[i] = byte(rand.Intn(255))
    }
    PrintByteTable(&bytes)
    // Start
    for (true) {
        fmt.Printf("\n")
        // Determines the opperation to use on the bytes
        fmt.Printf("What do you want to do?\n")
        var val, op int
        var cmd string
        fmt.Scanln(&op, &cmd, &val)
        // Opperates
        //op = line
        switch cmd{
            // Program cmd
            case "?":
                HELP(op)
                break
            case "HELP":
                HELP(op)
                break
            case "":
                fmt.Println(op)
                break
            case "PRINT":
                PrintByteTable(&bytes)
                break
            case "PRINTFT":
                PrintByteTableFromTo(&bytes, op, val)
                break
            case "EXIT":
                goto exit
                break
            case "RESET":
                goto reset
                break
            // Byte by value
            case "SET":
                //var init = bytes[op]
                PrintByteLine(&bytes[op], op)
                fmt.Print("-->")
                SET(&bytes[op], val)
                PrintByteLine(&bytes[op], op)
                //fmt.Printf("Set %d from %d to %d \n", op, init, bytes[op])
                break
            case "ADD":
                PrintByteLine(&bytes[op], op)
                fmt.Print("-->")
                ADD(&bytes[op], val)
                PrintByteLine(&bytes[op], op)
                break
            case "SUB":
                PrintByteLine(&bytes[op], op)
                fmt.Print("-->")
                SUB(&bytes[op], val)
                PrintByteLine(&bytes[op], op)
                break
            case "MULT":
                PrintByteLine(&bytes[op], op)
                fmt.Print("-->")
                MULT(&bytes[op], val)
                PrintByteLine(&bytes[op], op)
                break
            case "DIV":
                PrintByteLine(&bytes[op], op)
                fmt.Print("-->")
                DIV(&bytes[op], val)
                PrintByteLine(&bytes[op], op)
                break
            case "MOD":
                PrintByteLine(&bytes[op], op)
                fmt.Print("-->")
                MOD(&bytes[op], val)
                PrintByteLine(&bytes[op], op)
                break
            // Byte by byte
            case "PTR":
                PrintByteLine(&bytes[op], op)
                PrintByteLine(&bytes[val], val)
                fmt.Print("-->")
                PTR(&bytes[op], &bytes[val])
                PrintByteLine(&bytes[op], op)
                PrintByteLine(&bytes[val], val)
                break
            case "MOV":
                PrintByteLine(&bytes[op], op)
                PrintByteLine(&bytes[val], val)
                fmt.Print("-->")
                MOV(&bytes[op], &bytes[val])
                PrintByteLine(&bytes[op], op)
                PrintByteLine(&bytes[val], val)
                break
            case "BADD":
                PrintByteLine(&bytes[op], op)
                PrintByteLine(&bytes[val], val)
                fmt.Print("-->")
                BADD(&bytes[op], &bytes[val])
                PrintByteLine(&bytes[op], op)
                PrintByteLine(&bytes[val], val)
                break
            case "BSUB":
                PrintByteLine(&bytes[op], op)
                PrintByteLine(&bytes[val], val)
                fmt.Print("-->")
                BSUB(&bytes[op], &bytes[val])
                PrintByteLine(&bytes[op], op)
                PrintByteLine(&bytes[val], val)
                break
            case "BMULT":
                PrintByteLine(&bytes[op], op)
                PrintByteLine(&bytes[val], val)
                fmt.Print("-->")
                BMULT(&bytes[op], &bytes[val])
                PrintByteLine(&bytes[op], op)
                PrintByteLine(&bytes[val], val)
                break
            case "BDIV":
                PrintByteLine(&bytes[op], op)
                PrintByteLine(&bytes[val], val)
                fmt.Print("-->")
                BDIV(&bytes[op], &bytes[val])
                PrintByteLine(&bytes[op], op)
                PrintByteLine(&bytes[val], val)
                break
            case "BMOD":
                PrintByteLine(&bytes[op], op)
                PrintByteLine(&bytes[val], val)
                fmt.Print("-->")
                BMOD(&bytes[op], &bytes[val])
                PrintByteLine(&bytes[op], op)
                PrintByteLine(&bytes[val], val)
                break
            // Byte Unary
            case "INC":
                PrintByteLine(&bytes[op], op)
                fmt.Print("-->")
                INC(&bytes[op])
                PrintByteLine(&bytes[op], op)
                break
            case "DNC":
                PrintByteLine(&bytes[op], op)
                fmt.Print("-->")
                DNC(&bytes[op])
                PrintByteLine(&bytes[op], op)
                break
            case "SFT":
                PrintByteLine(&bytes[op], op)
                fmt.Print("-->")
                SHIFT(&bytes[op], val)
                PrintByteLine(&bytes[op], op)
            case "RAND":
                PrintByteLine(&bytes[op], op)
                fmt.Print("-->")
                RAND(&bytes[op])
                PrintByteLine(&bytes[op], op)
                break
            // Conversion
            case "HEX":
                PrintByteLine(&bytes[op], op)
                fmt.Print("-->")
                HEX(&bytes[op])
                break
            case "BIT":
                PrintByteLine(&bytes[op], op)
                fmt.Print("-->")
                BIT(&bytes[op]);
                break
            // Variables
            default:
                fmt.Printf("Not an accepted opperation, type 0 ? for more info. \n")
                break
        }
        fmt.Printf("\n")
    }
    exit:
    fmt.Printf("Bye")
}