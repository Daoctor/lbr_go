package fetchword

import (
	"github.com/hermanschaaf/enchant"
	"github.com/reiver/go-porterstemmer"
	"regexp"
	"sort"
	"strings"
)

// Word struct word
type Word struct {
	word  string
	count int
}

//GetValidWord get one valid word
func GetValidWord(word string) string {
	var wordPattern *regexp.Regexp
	wordPattern, _ = regexp.Compile(`\w{3,15}`)
	enchant, err := enchant.NewEnchant()
	enchant.LoadDict("en_GB")
	if err != nil {
		panic("Enchant error: " + err.Error())
	}
	noResult := ""
	m := wordPattern.MatchString(word)
	stem := porterstemmer.StemString(word)
	rO := enchant.Check(word)
	rA := enchant.Check(stem)
	if !(m && (rO || rA)) {
		return noResult
	}
	if rO && rA && len(stem) < 3 {
		return word
	}
	if rO && rA && len(stem) > 3 {
		return stem
	}
	if rO {
		return word
	}
	if len(stem) < 3 {
		return noResult
	}
	return stem
}

// GetWords get all valid word from list
func GetWords(wordList []string) []Word {
	wordDict := make(map[string]int)
	// generate word freq map
	for _, word := range wordList {
		word = strings.TrimSpace(word)
		word = GetValidWord(word)
		if len(word) == 0 {
			continue
		}
		_, exsits := wordDict[word]
		if exsits {
			wordDict[word]++
		} else {
			wordDict[word] = 1
		}
	}
	// sort word by freq
	var sortedWords = []Word{}
	for k := range wordDict {
		w := Word{k, wordDict[k]}
		sortedWords = append(sortedWords, w)
	}
	sort.Slice(sortedWords, func(i, j int) bool {
		return sortedWords[i].count > sortedWords[j].count
	})
	return sortedWords
}

