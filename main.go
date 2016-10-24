package main

import (
  "markov"
  "sync"
  "io/ioutil"
)

func main() {
	dat, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		markov.NewChain(string(dat)).Build().Generate(200)
	}
	wg.Wait()
}
