package main

import (
	"fmt"
	p "gontext-progressbar/progressbar"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	//pp := p.NewCancelContext(p.Box7, 100)
	pp := p.NewTimeoutContext(p.Spin4, 100, time.Second*7)
	//pp := p.NewDeadlineContext(p.Emoji, 100, time.Now().Add(time.Second*7))

	cancel := pp.ShowMulti(" 5 second ", " deadline from now ")

	go func() {
		time.Sleep(time.Second * 5)
		fmt.Println("call cancel")
		cancel()
	}()

	wg.Wait()
}
