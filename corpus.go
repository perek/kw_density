package main

import (
	"strings"
)

//Corpus is used to query multiple file indices
type Corpus struct {
	indices []*Indexer
}

//NewCorpus returns a new instance of Corpus and initilizes the indexes
func NewCorpus(files []string) *Corpus {
	corpus := Corpus{indices: []*Indexer{}}
	for _, file := range files {
		index := NewIndexer(file)
		corpus.indices = append(corpus.indices, index)
	}
	return &corpus
}

//MaxDensity returns the filename and score for the highest density score in the corpus
func (c *Corpus) MaxDensity(word string) (string, float64) {
	word = strings.ToLower(word)
	maxDensity := 0.0
	file := "none"
	for _, index := range c.indices {
		if index.Density(word) > maxDensity {
			maxDensity = index.Density(word)
			file = index.Filename()
		}
	}

	return file, maxDensity
}
