package speaker

func (s Speaker) Say (str string, sec int) {
    for {
        s.Wait(sec)
        s.Log(str)
    }
}
