package main

import ("fmt"
        "os"
        "strings"
        "bufio"
)

func main() {
    reader := bufio.NewScanner(os.Stdin)

    for { 
        fmt.Print("Pokedex > ")
        success := reader.Scan()
        
        if !success {
            break
        }
        text := reader.Text()
        lowered := strings.ToLower(text)
        words := strings.Fields(lowered)

        if len(words) == 0 {
            continue 
        }

        fmt.Printf("Your command was: %s\n", words[0])
    }

}

