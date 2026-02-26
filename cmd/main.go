package cmd

import (
	"fmt"

	"github.com/kamarikg/wallet/pkg/wallet"
)

func main() {
	svc := &wallet.Service{}
	account, err := svc.RegisterAccount("+992918654619")
	if err != nil {
		fmt.Println(err)
		fmt.Println(account)
		return
	}
}
