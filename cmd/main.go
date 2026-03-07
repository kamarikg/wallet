package main

import (
    "fmt"

    "github.com/kamarikg/wallet/pkg/wallet"
    "github.com/kamarikg/wallet/pkg/types"
)

func main() {
    svc := &wallet.Service{}
    account, err := svc.RegisterAccount(types.Phone("+992918654619"))
    if err != nil {
        fmt.Println("Ошибка регистрации:", err)
        return
    }

    err = svc.Deposit(account.ID, 10)
    if err != nil {
        fmt.Println("Ошибка депозита:", err)
        return
    }

    fmt.Println("Баланс аккаунта:", account.Balance) // 10
}
