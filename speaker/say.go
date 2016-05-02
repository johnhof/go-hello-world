package speaker

func (s Speaker) Say (str string, sec int) {
    for true {
        s.Wait(sec)
        s.Log(str)
    }
}
