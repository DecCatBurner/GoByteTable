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
                var init = bytes[op]
                SET(&bytes[op], val)
                fmt.Printf("Set %d from %d to %d \n", op, init, bytes[op])
                break
            case "ADD":
                ADD(&bytes[op], val)
                break
            case "SUB":
                SUB(&bytes[op], val)
                break
            case "MULT":
                MULT(&bytes[op], val)
                break
            case "DIV":
                DIV(&bytes[op], val)
                break
            case "MOD":
                MOD(&bytes[op], val)
                break
            // Byte by byte
            case "PTR":
                PTR(&bytes[op], &bytes[val])
                break
            case "MOV":
                MOV(&bytes[op], &bytes[val])
                break
            case "BADD":
                BADD(&bytes[op], &bytes[val])
                break
            case "BSUB":
                BSUB(&bytes[op], &bytes[val])
                break
            case "BMULT":
                BMULT(&bytes[op], &bytes[val])
                break
            case "BDIV":
                BDIV(&bytes[op], &bytes[val])
                break
            case "BMOD":
                BMOD(&bytes[op], &bytes[val])
                break
            // Byte Unary
            case "INC":
                INC(&bytes[op])
                break
            case "DNC":
                DNC(&bytes[op])
                break
            case "RAND":
                RAND(&bytes[op])
                break
            // Conversion
            case "HEX":
                HEX(&bytes[op])
                break
            case "BIT":
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