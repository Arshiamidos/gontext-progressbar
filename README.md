# gontext-progressbar
contextual progressbar with golang
```go
    import (
        ...
        p "gontext-progressbar/progressbar"
        ...
    )
    ...
    //senario one cancel context
    pp := p.NewCancelContext(p.Box7, 100)
    cancel := pp.Show(" yes it is progress bar ")

    go func() {
        time.Sleep(time.Second * 3)
        fmt.Println("call cancel")
        cancel()
    }()
```


```go
    //senario 2 multi line from one progress
    
    //pp := p.NewCancelContext(p.Box7, 100)
    //pp := p.NewTimeoutContex(p.Spin4, 100, time.Second*7)
    pp := p.NewDeadlineContext(p.Emoji, 100, time.Now().Add(time.Second*7))
	cancel := pp.ShowMulti(" it will be kill ", " after 10 second ")

	go func() {
		time.Sleep(time.Second * 10)
		fmt.Println("call cancel")
		cancel()
    }()
    
```




