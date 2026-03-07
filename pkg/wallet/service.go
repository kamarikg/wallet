package wallet

import (
    "errors"

    "github.com/kamarikg/wallet/pkg/types"
)

type Service struct {
    nextAccountID int64
    accounts      []*types.Account
    payments      []*types.Payment
}

// Метод регистрации аккаунта
func (s *Service) RegisterAccount(phone types.Phone) (*types.Account, error) {
    for _, account := range s.accounts {
        if account.Phone == phone {
            return nil, errors.New("phone already registered")
        }
    }

    s.nextAccountID++
    account := &types.Account{
        ID:      s.nextAccountID,
        Phone:   phone,
        Balance: 0,
    }
    s.accounts = append(s.accounts, account)

    return account, nil
}

// Метод пополнения баланса
func (s *Service) Deposit(accountID int64, amount types.Money) error {
    if amount <= 0 {
        return errors.New("amount must be greater than 0")
    }

    var account *types.Account
    for _, acc := range s.accounts {
        if acc.ID == accountID {
            account = acc
            break
        }
    }

    if account == nil {
        return errors.New("account not found")
    }

    // Зачисление средств пока не рассматриваем как платёж
    account.Balance += amount
    return nil
}
