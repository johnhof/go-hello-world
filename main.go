package main

import (
    "fmt"
    "github.com/johnhof/go-hello-world/speaker"
)

func main()  {
    s, _ := speaker.New()
    fmt.Println("Starting Hello")
    go s.SayHello(5)
    fmt.Println("Starting World")
    go s.SayWorld(3)
    fmt.Println("Starting World")
    s.Say(".", 1)
}
