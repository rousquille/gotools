package gotools

import (
	"log"
	"regexp"
	"strings"
)

// check if string contains element from elementsToCheck
func checkFromList(str string, elementsToCheck []string) bool {
	for _, toCheck := range elementsToCheck {
		if strings.Contains(str, toCheck) {
			return true
		}
	}
	return false
}

// remove all elements not contains from slice of string
func removeElementsNotContainsFromSlices(sliceToAnalyse []string, elementsToKeep []string) []string {
	for i := 0; i < len(sliceToAnalyse); i++ {
		file := sliceToAnalyse[i]
		for _, toKeep := range elementsToKeep {
			if !strings.Contains(file, toKeep) {
				sliceToAnalyse = append(sliceToAnalyse[:i], sliceToAnalyse[i+1:]...)
				i--
				break
			}
		}
	}
	return sliceToAnalyse
}

// remove whose contains several string elements from slice of string
func removeContainsElementsFromSlices(sliceToAnalyse []string, elementsToRemove []string) []string {
	for i := 0; i < len(sliceToAnalyse); i++ {
		file := sliceToAnalyse[i]
		for _, excluded := range elementsToRemove {
			if strings.Contains(file, excluded) {
				sliceToAnalyse = append(sliceToAnalyse[:i], sliceToAnalyse[i+1:]...)
				i--
				break
			}
		}
	}
	return sliceToAnalyse
}

// remove regex string from slice of string
func removeWithMatchingRegex(s []string, regex string) []string {
	for i, v := range s {
		matched, err := regexp.MatchString(regex, v)
		if err != nil {
			log.Fatalln(err)
		}
		if matched {
			return append(s[:i], s[i+1:]...)
		}
	}
	return s
}

// remove specific string from slice of string
func removeString(s []string, r string) []string {
	for i, v := range s {
		if v == r {
			return append(s[:i], s[i+1:]...)
		}
	}
	return s
}
