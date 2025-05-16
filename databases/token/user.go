package dbtoken

import (
	"fmt"
	"slices"
	"sync"
)

var mutex sync.Mutex
var userTokenSlice []string // This 'database' is simply a slice of token strings

// Return token is token is found and not expired
func CheckUserTokenExists(tokenStr string) bool {
	return slices.Contains(userTokenSlice, tokenStr)
}

func AddToken(tokenStr string) {
	defer mutex.Unlock()

	mutex.Lock()
	if CheckUserTokenExists(tokenStr) {
		return
	}

	userTokenSlice = append(userTokenSlice, tokenStr)
}

// This is mainly used for signing out
func DeleteToken(tokenStr string) {
	if !CheckUserTokenExists(tokenStr) {
		return
	}

	userTokenSlice = slices.DeleteFunc(userTokenSlice, func(token string) bool {
		return token == tokenStr
	})
}

func PrintAmountOfTokens() {
	fmt.Println("# of user tokens: ", len(userTokenSlice))
}
