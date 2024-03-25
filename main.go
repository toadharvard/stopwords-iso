package main

import (
	"encoding/json"
	"os"
	"regexp"
	"slices"
	"strings"
)

type StopwordsMapping map[string][]string
type ISOCode639_1 = string

var wordSegmenter = regexp.MustCompile(`[\pL\p{Mc}\p{Mn}-_']+`)

func standardizeSpaces(str string) string {
	return strings.Join(strings.Fields(str), " ")
}

// NewStopwordsMapping initializes a new StopwordsMapping from a JSON file.
//
// Returns:
// - StopwordsMapping: a map containing language to stopwords mapping.
// - error: an error object if an error occurred while reading or unmarshaling the JSON file.
func NewStopwordsMapping() (StopwordsMapping, error) {
	jsonFile, err := os.ReadFile("stopwords-iso.json")
	if err != nil {
		return *new(StopwordsMapping), err
	}

	mapping := make(map[string][]string)
	json.Unmarshal(jsonFile, &mapping)
	return mapping, nil
}

// ClearStringByLang clears the given string by removing all stopwords in the specified language.
//
// Parameters:
// - str: the string to be cleared.
// - language: the language of the stopwords to be removed in ISO 639-1 format.
//
// Return:
// - string: the cleared string.
func (m *StopwordsMapping) ClearStringByLang(str string, language ISOCode639_1) string {
	language = strings.ToLower(language)

	str = standardizeSpaces(str)

	words := wordSegmenter.FindAllString(str, -1)

	filtered := []string{}
	for _, word := range words {
		if !slices.Contains((*m)[language], strings.ToLower(word)) {
			filtered = append(filtered, word)
		}
	}
	return strings.Join(filtered, " ")
}

// ClearString clears the given string by removing stopwords for all languages.
//
// Parameters:
// - str: the string to be cleared.
//
// Returns:
// - string: the cleared string.
func (m *StopwordsMapping) ClearString(str string) string {
	for language := range *m {
		str = m.ClearStringByLang(str, language)
	}
	return str
}