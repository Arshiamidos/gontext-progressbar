package progressbar

import (
	"fmt"
	"sync"
	"time"
)

var (
	Box1    = `⠋⠙⠹⠸⠼⠴⠦⠧⠇⠏`
	Box2    = `⠋⠙⠚⠞⠖⠦⠴⠲⠳⠓`
	Box3    = `⠄⠆⠇⠋⠙⠸⠰⠠⠰⠸⠙⠋⠇⠆`
	Box4    = `⠋⠙⠚⠒⠂⠂⠒⠲⠴⠦⠖⠒⠐⠐⠒⠓⠋`
	Box5    = `⠁⠉⠙⠚⠒⠂⠂⠒⠲⠴⠤⠄⠄⠤⠴⠲⠒⠂⠂⠒⠚⠙⠉⠁`
	Box6    = `⠈⠉⠋⠓⠒⠐⠐⠒⠖⠦⠤⠠⠠⠤⠦⠖⠒⠐⠐⠒⠓⠋⠉⠈`
	Box7    = `⠁⠁⠉⠙⠚⠒⠂⠂⠒⠲⠴⠤⠄⠄⠤⠠⠠⠤⠦⠖⠒⠐⠐⠒⠓⠋⠉⠈⠈`
	Spin1   = `|/-\`
	Spin2   = `◴◷◶◵`
	Spin3   = `◰◳◲◱`
	Spin4   = `◐◓◑◒`
	Spin5   = `▉▊▋▌▍▎▏▎▍▌▋▊▉`
	Spin6   = `▌▄▐▀`
	Spin7   = `╫╪`
	Spin8   = `■□▪▫`
	Spin9   = `←↑→↓`
	Emoji   = `😯😦😧😮😲😵😳😱😵😳😧😦😯😲`
	Default = Box1
)

type progressbar struct {
	prog  []rune
	delay int
}

func New(Type string, Delay int) *progressbar {
	if len(Type) > 1 {
		return &progressbar{
			delay: Delay,
			prog:  []rune(Type),
		}
	} else {
		return &progressbar{
			delay: Delay,
			prog:  []rune("😯😦😧😮😲😵😳😱😵😳😧😦😯😲"),
		}
	}
}
func (p *progressbar) Run() <-chan bool {
	t := time.Tick(time.Millisecond * time.Duration(p.delay))
	c := make(chan bool, 1)
	go func() {
		for {
			select {
			case <-t:
				p.prog = append(p.prog[1:], p.prog[0])
				c <- true
			}
		}
	}()
	return c
}
func (p *progressbar) Print(s string) string {
	//hide back 100000 clrline
	return fmt.Sprintf("%s%s", string(p.prog[0]), s)
}

func (p *progressbar) PrintLine(s string) string {
	//\r clrln
	return fmt.Sprintf("\r\033[2K%s\033[m%s", string(p.prog[0]), s)
}
func PrintMultiText(s ...string) string {

	_s := "\r\033[?25l\033[J"
	for i := 0; i < len(s); i++ {
		_s = fmt.Sprintf("%s%s\n", _s, s[i])
	}
	_s = fmt.Sprintf("%s\033[%dA", _s, len(s))
	return _s

}
func Race(c ...<-chan bool) chan bool {
	_c := make(chan bool, len(c))
	var wg sync.WaitGroup
	wg.Add(len(c))
	for _, v := range c {
		go func(vv <-chan bool) {
			if <-vv {
				_c <- true
				wg.Done()
			}
		}(v)

	}
	go func() {
		wg.Wait()
		close(_c)
	}()
	return _c
}

//
