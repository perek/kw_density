package main

import (
	"flag"
	"fmt"
	"strings"
)

func main() {
	filesFlag := flag.String("f", "", "comma separated list containing text documents")
	wordsFlag := flag.String("w", "", "comma separated list containing the words to check keyword density")
	//verboseFlag := flag.Bool("v", false, "verbose output")

	flag.Parse()
	files := strings.Split(*filesFlag, ",")
	words := strings.Split(*wordsFlag, ",")

	corpus := NewCorpus(files)

	for _, word := range words {
		file, score := corpus.MaxDensity(word)
		fmt.Printf("WORD: %s	FILE: %s	SCORE: %f\n", word, file, score)
	}
}
