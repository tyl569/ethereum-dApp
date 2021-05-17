package main

import (
	"github.com/ethereum/go-ethereum/accounts"
)
import "github.com/ethereum/go-ethereum/accounts/keystore"
import _ "github.com/ethereum/go-ethereum/common"
import "fmt"

func main() {
	ks := keystore.NewKeyStore("/path/to/keystore", keystore.StandardScryptN, keystore.StandardScryptP)
	am := accounts.NewManager(&accounts.Config{InsecureUnlockAllowed: false}, ks)
	fmt.Println(am)
}
