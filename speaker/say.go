package speaker

import (
    "fmt"
    "time"
)

func Say (str string, s int) {
    for true {
        time.Sleep(time.Duration(s) * time.Second)
        fmt.Print(str)
    }
}
