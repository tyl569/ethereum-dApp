package account_management

import "github.com/ethereum/go-ethereum/accounts/keystore"

func Delete(password string) {
	ks := keystore.NewKeyStore("/tmp/keystore", keystore.StandardScryptN, keystore.StandardScryptP)
	newAcc, _ := ks.NewAccount(password)
	_ = ks.Delete(newAcc, password)
}
