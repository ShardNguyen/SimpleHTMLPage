package dbtoken

import (
	"fmt"
	"slices"
	"sync"
)

var mutex sync.Mutex
var userTokenMap []string // Map has key being the user's token and value being the username

// Return token is token is found and not expired
func CheckUserTokenExists(tokenStr string) bool {
	return slices.Contains(userTokenMap, tokenStr)
}

func AddToken(tokenStr string) {
	defer mutex.Unlock()

	mutex.Lock()
	if CheckUserTokenExists(tokenStr) {
		return
	}

	userTokenMap = append(userTokenMap, tokenStr)
}

// This is mainly used for signing out
func DeleteToken(tokenStr string) {
	if !CheckUserTokenExists(tokenStr) {
		return
	}

	userTokenMap = slices.DeleteFunc(userTokenMap, func(token string) bool {
		return token == tokenStr
	})
}

func PrintAmountOfTokens() {
	fmt.Println("# of user tokens: ", len(userTokenMap))
}
