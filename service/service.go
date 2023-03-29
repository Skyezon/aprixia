package service

import (
	"math/rand"
	"time"
	"unicode"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
//will return 6 character alphanumeric

func Generate6Alphanumeric() string {
    rand.Seed(time.Now().UnixNano())
    alphaFlag := false
    numericFlag := false
    b := make([]byte, 6)
    for i := range b {
        b[i] = charset[rand.Intn(len(charset))]
        if numericFlag == false && unicode.IsDigit(rune( b[i])) {
            numericFlag = true
        }
        if alphaFlag == false && unicode.IsLetter(rune(b[i])){
            alphaFlag = true
        }
    }
    if alphaFlag && numericFlag {
        return string(b)
    }else{
        return Generate6Alphanumeric()
    }
}
