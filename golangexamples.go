package golangexamples

import (
	"github.com/ehteshamz/greetings"
)

// takes input a slice and returns its contents in a string, separated by a '-'
func ConcatSlice(sliceToConcat []byte) string {
	if len(sliceToConcat) == 0 { // check if input has values
		return ""
	}
	retStr := ""                        // the string to be returned
	for _, val := range sliceToConcat { // loop over all the elements in the slice
		retStr += string(val) + "-" // append each character to retStr
	}
	return retStr[:len(retStr)-1] // remove last '-' and return
}

// takes input a slice and returns an encrypted slice, using Ceaser Cypher
func Encrypt(sliceToEncrypt []byte, ceaserCount int) {
	if ceaserCount > 0 { // checking for a positive ceaserCount
		for ind, _ := range sliceToEncrypt { // update each element in the slice
			sliceToEncrypt[ind] = byte(int(sliceToEncrypt[ind]) + ceaserCount)
		}
	}
}

// imports greetings package and calls the appropriate function, returning a string
func EZGreetings(name string) string {
	return greetings.PrintGreetings(name)
}
