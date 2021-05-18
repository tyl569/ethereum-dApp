package account_management

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/keystore"
)

func Create() {
	ks := keystore.NewKeyStore("/tmp/keystore", keystore.StandardScryptN, keystore.StandardScryptP)
	newAcc, _ := ks.NewAccount("Tyl569")
	fmt.Println(newAcc)
}
