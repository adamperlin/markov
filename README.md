## Markov
This is a simple markov chain algorithm implementation written in go. It is not particularly efficient, but does yield some interesting results.

## Usage
Take a look at this example:
```go

/**
*Reads from a file, input.txt, and generates a markov chain based on that output
*/

package main

import (
    "fmt"
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
		str := markov.NewChain(string(dat)).Build().Generate(200)
    fmt.Printf("%s\n", str)
	}()
	wg.Wait()
}
```
