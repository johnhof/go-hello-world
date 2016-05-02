package main

import (
    "fmt"
    "github.com/johnhof/go-hello-world/speaker"
)

func main()  {
    fmt.Println("Starting Hello")
    go speaker.SayHello(5)
    fmt.Println("Starting World")
    go speaker.SayWorld(3)
    fmt.Println("Starting World")
    speaker.Say(".", 1)
}
