package main
 
import "fmt"
 
func main() {
    var command string
    for {
        fmt.Print("请输入你的指令:")
		fmt.Scan(&command)
        switch command {
            case "p":
                fmt.Println("Print something!")
            case "q":
                fmt.Println("Quit!")
                return
            default:
                fmt.Println("Command undefined!")
        }
 
    }
}