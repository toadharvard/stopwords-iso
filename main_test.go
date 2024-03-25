package main

import (
	"os"
	"testing"
)

func TestNewStopwordsMapping(t *testing.T) {
	mapping, err := NewStopwordsMapping()
	if err != nil {
		t.Errorf("NewStopwordsMapping() error = %v, wantErr %v", err, false)
	}
	if len(mapping) == 0 {
		t.Errorf("NewStopwordsMapping() returned an empty mapping")
	}
	// Test if the mapping contains expected languages
	if _, exists := mapping["en"]; !exists {
		t.Errorf("NewStopwordsMapping() should contain 'en' key for English stopwords")
	}
	// More specific tests can be added here
}

func TestClearStringByLang(t *testing.T) {
	mapping, _ := NewStopwordsMapping() // Assuming that NewStopwordsMapping works correctly

	tests := []struct {
		name     string
		str      string
		language ISOCode639_1
		want     string
	}{
		{"English", "this is a special", "en", "special"},
		{"Non-Existent Language", "this should not change", "xx", "this should not change"},
		{"Case Insensitivity", "THIS is A Google", "EN", "Google"},
		{"No Stopwords", "uniqueword", "en", "uniqueword"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mapping.ClearStringByLang(tt.str, tt.language); got != tt.want {
				t.Errorf("StopwordsMapping.ClearStringByLang() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStopwordsMapping_ClearString(t *testing.T) {
	mapping, _ := NewStopwordsMapping()

	testStr := "А ИЛИ ОН it's ok vadim"
	want := "vadim"

	got := mapping.ClearString(testStr)
	if got != want {
		t.Errorf("StopwordsMapping.ClearString() got = %v, want %v", got, want)
	}
}

func TestNewStopwordsMapping_Error(t *testing.T) {
	// Temporarily rename the stopwords file to simulate an error
	os.Rename("stopwords-iso.json", "stopwords-iso_backup.json")
	defer os.Rename("stopwords-iso_backup.json", "stopwords-iso.json")

	_, err := NewStopwordsMapping()
	if err == nil {
		t.Errorf("Expected an error when stopwords-iso.json does not exist")
	}
}
