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
        fmt.Println(err)
        return
    }

    err = svc.Deposit(account.ID, 10)
    if err != nil {
        switch err {
        case wallet.ErrAmountMustBePositive:
            fmt.Println("Сумма должна бать положительной")
        case wallet.ErrAccountNotFound:
            fmt.Println("Аккаут пользователя не найден")
        }
        return
    }

    fmt.Println(account.Balance) // 10
}
