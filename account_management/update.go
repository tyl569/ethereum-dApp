package account_management

import (
	"github.com/ethereum/go-ethereum/accounts/keystore"
)

func Update(oldPassword string, newPassword string) {
	ks := keystore.NewKeyStore("/tmp/keystore", keystore.StandardScryptN, keystore.StandardScryptP)
	newAcc, _ := ks.NewAccount(oldPassword)
	_ := ks.Update(newAcc, oldPassword, newPassword)
}
