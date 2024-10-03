package dynamicprogramming

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"slices"
	"strings"
	"unicode/utf8"

	"github.com/santakdalai90/golangdsa/src/util"
)

// Given a corpus of words and a word check if it's spelling is correct.
// If not, then suggest correct words sorted by likelihood

type IncorrectSpellingError struct{}

func (ise *IncorrectSpellingError) Error() string {
	return "Incorrect Spelling"
}

type SpellChecker struct {
	EditDistanceThreshold int
	LengthDiffThreshold   int
	MaxSuggestions        int
	dictionary            map[string]int
}

func NewSpellChecker(corpusPath string, editDistanceThreshold, lengthDiffThreshold, maxSuggestions int) (*SpellChecker, error) {
	sc := new(SpellChecker)

	sc.EditDistanceThreshold = editDistanceThreshold
	sc.LengthDiffThreshold = lengthDiffThreshold
	sc.MaxSuggestions = maxSuggestions

	err := sc.buildDictionary(corpusPath)

	return sc, err
}

// CheckSpelling checks if input s has correct spelling.
// If not, it returns IncorrectSpellingError and an slice of suggestions
func (sc *SpellChecker) CheckSpelling(s string) ([]string, error) {
	// check if the word exists in dictionary
	if _, ok := sc.dictionary[s]; ok {
		// return nil error if the work is found
		return []string{s}, nil
	}

	corrections := make([]string, 0)

	for word, _ := range sc.dictionary {
		if withinLengthThreshold(word, s, sc.LengthDiffThreshold) {
			if DamerauLevenshteinEditDistance(s, word) <= sc.EditDistanceThreshold {
				corrections = append(corrections, word)
			}
		}
	}

	slices.SortFunc(corrections, func(a, b string) int {
		return sc.dictionary[b] - sc.dictionary[a] // sort descending
	})

	if len(corrections) > sc.MaxSuggestions {
		corrections = corrections[:sc.MaxSuggestions]
	}

	return corrections, &IncorrectSpellingError{}
}

func (sc *SpellChecker) buildDictionary(corpusPath string) error {
	// initialize dictionary
	if sc.dictionary == nil {
		sc.dictionary = make(map[string]int)
	}
	// open file
	path, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Current Directory:", path)
	}
	corpusFile, err := os.Open(corpusPath)
	if err != nil {
		return fmt.Errorf("error opening corpus file. %w", err)
	}
	defer corpusFile.Close()

	// prepare scanner
	scanner := bufio.NewScanner(corpusFile)

	// prepare regex
	regex, err := regexp.Compile("[^a-zA-Z]+") // matches non alphabetic characters
	if err != nil {
		return fmt.Errorf("error compiling regular expression. %w", err)
	}

	for scanner.Scan() {
		// Read file line by line
		line := scanner.Text()
		// split words
		words := strings.Fields(line)
		// Update frequency
		for _, word := range words {
			cleanWord := regex.ReplaceAllString(word, "") // remove non alphabetic characters
			cleanWord = strings.ToLower(cleanWord)
			sc.dictionary[cleanWord]++
		}
	}

	// check scanner errors
	if err := scanner.Err(); err != nil {
		return fmt.Errorf("scanner error. %w", err)
	}

	return nil
}

func DamerauLevenshteinEditDistance(source, target string) int {
	//make unicode compatible
	s := []rune(source)
	t := []rune(target)

	// get lengths
	m := len(s)
	n := len(t)

	// create table
	E := util.Create2DSlice[int](m+1, n+1)

	// base case
	for i := 0; i <= m; i++ {
		E[i][0] = i
	}
	for j := 0; j <= n; j++ {
		E[0][j] = j
	}

	// fill table
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			cost := 0
			if s[i-1] != t[j-1] {
				cost = 1
			}

			E[i][j] = util.Min(1+E[i-1][j], 1+E[i][j-1], cost+E[i-1][j-1])

			// for transpose operation
			if (i > 1 && j > 1) && (source[i-1] == target[j-2] && source[i-2] == target[j-1]) {
				E[i][j] = util.Min(E[i][j], 1+E[i-2][j-2])
			}
		}
	}

	// solution
	return E[m][n]
}

func withinLengthThreshold(a, b string, threshold int) bool {
	// checks if a is withing length threshold of b
	return utf8.RuneCountInString(a) <= utf8.RuneCountInString(b)+threshold &&
		utf8.RuneCountInString(a) >= utf8.RuneCountInString(b)-threshold
}
