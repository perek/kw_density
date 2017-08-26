package main

import (
	"bufio"
	"os"
	"strings"
)

//Indexer takes a filename and calculates word counts and densities, should be created via NewIndexer
type Indexer struct {
	filename string
	counts   map[string]int
	wcount   int
}

//NewIndexer returns a new instance of Indexer and intilizes the word index
func NewIndexer(filename string) *Indexer {
	ret := &Indexer{filename: filename, counts: map[string]int{}}
	ret.index()
	return ret
}

//index builds the word index
func (i *Indexer) index() {
	file, err := os.Open(i.filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() { // internally, it advances token based on sperator
		lowerCase := strings.ToLower(scanner.Text())
		words := strings.Split(lowerCase, " ")
		for _, word := range words {
			word = strings.Trim(word, "`~!@#$%^&*()_-+={}[]:;\"'\\|<>,.?/ “”")
			if word == "" {
				continue
			}
			i.wcount++
			i.counts[word]++
		}
	}
}

//Count takes a string and returns the count of occurences in the given file
func (i *Indexer) Count(word string) int {
	return i.counts[word]
}

//Total returns the total amount of words in the index
func (i *Indexer) Total() int {
	return i.wcount
}

//Filename returns the path to file indexed
func (i *Indexer) Filename() string {
	return i.filename
}

//Density takes a word string and returns the density of that word in the file
func (i *Indexer) Density(word string) float64 {
	if i.wcount == 0 {
		return 0
	}
	return float64(i.counts[word]) / float64(i.wcount)
}
