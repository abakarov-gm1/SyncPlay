package main

import (
	"fmt"
)

type FirstDetail struct {
	Number string
}

type SecondDetail struct {
	Number string
}

type Composite struct {
	Number string
	First  FirstDetail
	Second SecondDetail
}

func result(firstCh chan *Composite, s []string) chan *Composite {
	resultCh := make(chan *Composite)

	go func() {
		for i, _ := range s {
			sec := SecondDetail{Number: fmt.Sprintf("%s-%s", s[i], "2")}
			c := <-firstCh
			c.Second = sec
			resultCh <- c
		}
		close(resultCh)
	}()

	return resultCh
}

func main() {
	firstCh := make(chan *Composite)
	//secondCh := make(chan *Composite)

	var n int
	var details string
	var s []string
	_, _ = fmt.Scan(&n)

	for range n {
		_, _ = fmt.Scan(&details)
		s = append(s, details)
	}

	go func() {
		for i, _ := range s {
			c := &Composite{Number: s[i]}
			f := FirstDetail{Number: fmt.Sprintf("%s-%s", s[i], "1")}
			c.First = f
			firstCh <- c
		}
	}()

	for value := range result(firstCh, s) {
		fmt.Println(value)
	}

}
