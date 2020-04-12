package main

import (
	"testing"
)

func TestBetween(t *testing.T) {
	testData := []struct {
		value    string
		a        string
		b        string
		expected string
	}{
		{"Something (between) brackets", "(", ")", "between"},
		{"Something no ) between ( brackets", "(", ")", ""},
		{"Where there ( is no end", "(", ")", ""},
		{"Where there ) is no start", "(", ")", ""},
		{"Something *else# for a change", "*", "#", "else"},
		{"Something #between same# character", "#", "#", "between same"},
		{"Something (between) with multiple ) character", "(", ")", "between"},
		{"Something between a number of words", "between", "word", " a number of "},
	}

	for _, data := range testData {
		actual := Between(data.value, data.a, data.b)
		if actual != data.expected {
			t.Errorf("Between '%s' and '%s' from '%s' was wrong, got '%s', expected '%s'",
				data.a, data.b, data.value, actual, data.expected)
		}
	}
}

func TestBefore(t *testing.T) {
	testData := []struct {
		value    string
		str      string
		expected string
	}{
		{"Some string to check", "to", "Some string "},
		{"Some string that doesn't contain it", "else", ""},
		{"Some string that contains some of a word", "words", ""},
	}

	for _, data := range testData {
		actual := Before(data.value, data.str)
		if actual != data.expected {
			t.Errorf("Before '%s' within '%s' was wrong, got '%s', expected '%s'",
				data.str, data.value, actual, data.expected)
		}
	}
}

func TestAfter(t *testing.T) {
	testData := []struct {
		value    string
		str      string
		expected string
	}{
		{"Some string to check", "to", " check"},
		{"Some string that doesn't contain it", "else", ""},
	}

	for _, data := range testData {
		actual := After(data.value, data.str)
		if actual != data.expected {
			t.Errorf("After '%s' within '%s' was wrong, got '%s', expected '%s'",
				data.str, data.value, actual, data.expected)
		}
	}
}
