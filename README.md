## Markov
This is a simple markov chain algorithm implementation written in go. It is not particularly efficient, but does yield some interesting results. 

## Usage
If you do want to use this, it should be relatively easy. Run:

`git clone https://github.com/adamperlin/markov.git`
If you haven't already, you'll need to install (go)[https://golang.org]
Next, simply run `go build markov.go`, and you should have an executable. 

To generate a chain based on specified file, run: 
`$ ./markov -f [file]`
The default file is `markov.txt`

To generate a specific number of words, run:
`$ ./markov -n [number of words]`
Default is 200



