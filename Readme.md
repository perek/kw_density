KW Density Tool
====================

Problem Description: 
------------------------
When implementing text search over a corpus (a collection of documents), systems often rely in part on the term frequency (TF) of words in a document.  
A TF score is computed for each word wi in each document djby computing the frequency of that word in that document.

Write a program that takes as input a set of documents (please use these sample documents) and a list of words, and returns the document with the highest TF score for each word and the TF score for that word in that document.  To break the document into words, you can strip out punctuation, split by whitespace, and convert everything to lowercase.  Please include documentation for running your program as well as the the output for the words “queequeg”, “whale”, and “sea”. Try to make it easy-to-use and efficiently implemented.

Building
====================
`make build`
Will output binary *kwd* into ./bin

Usage
====================
```$ ./kwd.sh -h
Usage of ./bin/kwd:
  -f string
        comma separated list containing text documents
  -w string
        comma separated list containing the words to check keyword density```

Sample:
```$ ./kwd.sh -f data/chapter1.txt,data/chapter2.txt,data/chapter3.txt,data/chapter4.txt,data/chapter5.txt -w queequeg,whale,sea
WORD: queequeg  FILE: data/chapter4.txt SCORE: 0.006634
WORD: whale     FILE: data/chapter1.txt SCORE: 0.001354
WORD: sea       FILE: data/chapter1.txt SCORE: 0.004515```
