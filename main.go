package main
 
import "fmt"
 
func main() {
    var command string
    for {
        fmt.Print("enter your command:")
		fmt.Scan(&command)
        switch command {
		case "h":
            fmt.Println("Help!")
        case "p":
            fmt.Println("Print something!")
        case "q":
			fmt.Println("Quit!")
            return
        default:
            fmt.Println("Command is undefined!")
        }
 
    }
}