// by @knottedcat https://github.com/knottedcat/passwordgen

package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

const (
	consonants = "bcdfghjkmnpqrstvwxz" // 19 no letters that look similar
	vowels     = "aeiouy"              // 7
)

func randInt(max int) int {
	n, err := rand.Int(rand.Reader, big.NewInt(int64(max)))
	if err != nil {
		panic(err)
	}
	return int(n.Int64())
}
func randChar(charset string) byte {
	length := len(charset)
	index := randInt(length)
	return charset[index]
}
func createSegment() []byte {
	// first creates a empty list with 6 spaces
	segment := make([]byte, 6)

	// now add cvccvc for the segment
	segment[0] = randChar(consonants)
	segment[1] = randChar(vowels)
	segment[2] = randChar(consonants)
	segment[3] = randChar(consonants)
	segment[4] = randChar(vowels)
	segment[5] = randChar(consonants)

	return segment
}
func main() {
	seg1 := createSegment()
	seg2 := createSegment()
	seg3 := createSegment()

	// put in list
	segmentList := [][]byte{seg1, seg2, seg3}

	// pick random segment and spot
	randomNumSegment := randInt(3)
	randomNumPos := randInt(6)

	// pick random number
	randomDigit := byte('0' + randInt(10))

	// swap the number into it
	segmentList[randomNumSegment][randomNumPos] = randomDigit

	// now add capital

	randomCapPos := randInt(6)
	randomCapSegment := randInt(3)

	// this checks for collisions between the number and capitol and rerolls the positon; segment is fine
	for randomCapSegment == randomNumSegment && randomCapPos == randomNumPos {
		randomCapPos = randInt(6)
	}

	// subtract 32 to make the ascii captal
	segmentList[randomCapSegment][randomCapPos] = segmentList[randomCapSegment][randomCapPos] - 32

	// finally print out the end pw
	fmt.Printf("%s-%s-%s\n", string(segmentList[0]), string(segmentList[1]), string(segmentList[2]))
}
