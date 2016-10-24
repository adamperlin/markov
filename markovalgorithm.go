//generate a markov chain
package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
	"sync"
	"flag"
)

type (
	Prefix struct { //prefix tuple struct
		Prefix1 string
		Prefix2 string
	}
	Suffix struct {
		Suffixes []string
	}
	Markov struct {
		pairs map[*Prefix]*Suffix //map pointers to prefix and suffix
	}
)

func NewChain() *Markov {
	return &Markov{
		pairs: make(map[*Prefix]*Suffix),
	}
}
func (s *Suffix) Add(val string) {
	s.Suffixes = append(s.Suffixes, val)
}
func (p *Prefix) Compare(p2 *Prefix) bool {
	return p.Prefix1 == p2.Prefix1 && p.Prefix2 == p2.Prefix2 //compare two prefixes for equality
}

func (m *Markov) FindPrefixMatches(ref *Prefix) *Suffix { //necessary to have a method to find a prefix which matches another
	for k, v := range m.pairs { //iterate through the map
		if k.Compare(ref) {
			return v
		}
	}
	return nil
}
func (m *Markov) Build(text []string) {
	var prefix *Prefix //current prefix
	var suffix *Suffix //current suffix
	for i := range text {
		if i+2 > len(text)-1 {break}
			p1 := text[i]     //first prefix string
			p2 := text[i+1]   //second prefix string
			suff := text[i+2] //suffix string
			prefix = &Prefix{ //create a new prefix pair
				Prefix1: p1,
				Prefix2: p2,
			}
			suffix = &Suffix{ //create a new suffix with current suffix string
				Suffixes: []string{suff},
			}
			if s := m.FindPrefixMatches(prefix); s != nil { //check to see if the current prefix exists
				//if current prefix exists, get the corresponding suffix to that instance, and add its values to the current suffix
				for _, str := range s.Suffixes {
					suffix.Add(str)
				}
			}
			m.pairs[prefix] = suffix //add new prefix/suffix pair to the hash table
		}
}


func (m *Markov) PrintMapString() { //debug method to make sure the map was built correctly
	for k, v := range m.pairs {
		suffs := strings.Join(v.Suffixes, " ")
		fmt.Printf("%s: %s \n", string(k.Prefix1)+" "+string(k.Prefix2), suffs)
	}
}

func (m *Markov) Generate(length int, firstWord string, secondWord string) { //print out n words in a markov chain based on a starting prefix pair
	var pref1 string = firstWord
	var pref2 = secondWord
	fmt.Print(firstWord + " " + secondWord)
	for i := 0; i < length; i++ {
		suff := m.GetNextSuffix(pref1, pref2)
		fmt.Print(suff + " ")
		pref1, pref2 = pref2, suff
	}
}
func (m *Markov) GetNextSuffix(w1 string, w2 string) string {
	rand.Seed(time.Now().UnixNano()) //seed the random number generator
	for k, v := range m.pairs {
		if k.Prefix1 == w1 && k.Prefix2 == w2 {
			return v.Suffixes[rand.Intn(len(v.Suffixes))] //randomly select a suffix
		}
	}
	return ""
}

func main() {
	var inputFile = flag.String("f", "markov.txt", "usage: markov -f <filename")
	flag.Parse()
	dat, err := ioutil.ReadFile(*inputFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	c := NewChain()
	arr := strings.Split(string(dat), " ")
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func(){
		defer wg.Done()
		c.Build(arr)
	}()
	wg.Wait()
	//c.PrintMapString()
	//fmt.Print(strings.Join(arr[0:2], " ") + " ")
	wg.Add(1)
	go func(){
		defer wg.Done()
		c.Generate(200, arr[0], arr[1])
	}()
	wg.Wait()
}
