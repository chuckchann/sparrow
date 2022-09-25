package util

import (
	"math/rand"
	"strconv"
	"time"
	"unsafe"
)

const (
	allCharacters  = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	allIdxBits     = 6                    // 62=111110b  6 bits to represent a letter index
	allIdxMask     = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	allIdxSections = 63 / letterIdxBits

	letterCharacters  = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	letterIdxBits     = 6                    // 52=110100b  6 bits to represent a letter index
	letterIdxMask     = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxSections = 63 / letterIdxBits

	upLettersCharacters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	upLetterIdxBits     = 5 // 26=11010b  5 bits to represent a letter index
	upLetterIdxMask     = 1<<letterIdxBits - 1
	upLetterIdxSections = 63 / letterIdxBits

	lowLettersCharacters = "abcdefghijklmnopqrstuvwxyz"
	lowLetterIdxBits     = 5 // 26=11010b  5 bits to represent a letter index
	lowLetterIdxMask     = 1<<letterIdxBits - 1
	lowLetterIdxSections = 63 / letterIdxBits

	numLettersCharacters = "1234567890"
	numLetterIdxBits     = 4 // 1010=11010b  5 bits to represent a letter index
	numLetterIdxMask     = 1<<letterIdxBits - 1
	numLetterIdxSections = 63 / letterIdxBits
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func randomBaseString(n, idxBits, idxMask, sections int64, characters string) string {
	b := make([]byte, n)
	for i, cache, remain := n-1, rand.Int63(), sections; i >= 0; {
		if remain == 0 {
			cache, remain = rand.Int63(), sections
		}
		if idx := int(cache & idxMask); idx < len(characters) {
			b[i] = characters[idx]
			i--
		}
		remain--
		cache >>= idxBits
	}

	return *(*string)(unsafe.Pointer(&b))
}

func RandomString(n int64) string {
	return randomBaseString(n, allIdxBits, allIdxMask, allIdxSections, allCharacters)
}

func RandomLetterString(n int64) string {
	return randomBaseString(n, letterIdxBits, letterIdxMask, letterIdxSections, letterCharacters)
}

func RandomUpLetterString(n int64) string {
	return randomBaseString(n, upLetterIdxBits, upLetterIdxMask, upLetterIdxSections, upLettersCharacters)
}

func RandomLowLetterString(n int64) string {
	return randomBaseString(n, lowLetterIdxBits, lowLetterIdxMask, lowLetterIdxSections, lowLettersCharacters)
}

func RandomNumString(n int64) string {
	return randomBaseString(n, numLetterIdxBits, numLetterIdxMask, numLetterIdxSections, numLettersCharacters)
}

func RandomTimeString() string {
	return strconv.FormatInt(time.Now().UnixNano(), 10)
}

func RandomIntNum(n int) int {
	return rand.Intn(n)
}
