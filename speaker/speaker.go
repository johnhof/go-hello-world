package speaker

import (
    "fmt"
    "time"
)

type Speaker struct {}

func New() (Speaker, error) {
    return Speaker{}, nil
}

func (s Speaker) Log(str string) {
    fmt.Print(str)
}

func (s Speaker) Wait(sec int) {
    time.Sleep(time.Duration(sec) * time.Second)
}
