package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"
	"sync"
)

type wordFreq struct {
	word  string
	count int
}

func main() {
	var wg sync.WaitGroup
	var mu sync.Mutex

	files := []string{"sample.txt"}

	wordCount := make(map[string]int)

	re := regexp.MustCompile(`[a-zA-Z]+`)

	for _, file := range files {
		wg.Add(1)

		go func(fname string) {
			defer wg.Done()
			f, err := os.Open(fname)
			if err != nil {
				fmt.Errorf("error opening in file %s", err)
				return
			}
			defer f.Close()
			scanner := bufio.NewScanner(f)
			for scanner.Scan() {
				line := strings.ToLower(scanner.Text())
				words := re.FindAllString(line, -1)

				for _, word := range words {
					mu.Lock()
					wordCount[word]++
					mu.Unlock()
				}
			}
			if err := scanner.Err(); err != nil {
				fmt.Errorf("error while reading the file %s", err)
			}
		}(file)
	}

	wg.Wait()

	sortedFreq := sortByMap(wordCount)

	for _, wf := range sortedFreq {
		fmt.Printf("%s  %d\n", wf.word, wf.count)
	}
}

func sortByMap(wordCount map[string]int) []wordFreq {
	var freqs []wordFreq
	for k, v := range wordCount {
		freqs = append(freqs, wordFreq{k, v})
	}
	sort.Slice(freqs, func(i, j int) bool {
		if freqs[i].count == freqs[j].count {
			return freqs[i].word < freqs[j].word
		}
		return freqs[i].count > freqs[j].count
	})
	return freqs
}
