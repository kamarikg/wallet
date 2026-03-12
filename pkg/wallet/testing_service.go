package wallet

import (
	"reflect"
	"testing"
	"fmt"

	"github.com/kamarikg/wallet/pkg/types"
)

func TestService_Reject_succsess(t *testing.T)  {
	//создаём сервис
	s := &Service{}

	//регистрируем там пользователя
	phone := types.Phone("+992918654619")
	account, err := s.RegisterAccount(phone)
	if err != nil {
		t.Errorf("Reject(): can't register account, error = %v", err)
		return
	}

	//пополняем его счёт
	err = s.Deposit(account.ID, 10_000_00)
	if err !=nil {
		t.Errorf("Reject(): can't deposit account, error = %v", err)
		return
	}

	//осуществляем платёж на его счёт
	payment, err := s.Pay(account.ID, 1000_00, "auto")
	if err != nil {
		t.Errorf("Reject(): can't create payment, error = %v", err)
		return
	}

	//попробуем отменить платёж
	err = s.Reject(payment.ID)
	if err != nil {
		t.Errorf("Reject(): error = %v", err)
		return
	}

	// А как проверить статус платежа?
	// И баланс аккаунта?
}

func TestService_FindPaymentByID_success(t *testing.T) {
	//создаём сервис
	s := &Service{}

	// резистрируем там пользователя
	phone := types.Phone("+992918654619")
	account, err := s.RegisterAccount(phone)
	if err != nil {
		t.Errorf("FindPaymentByID(): can't register account, error = %v", err)
		return
	}

	// пополняем его счёт
	err = s.Deposit(account.ID, 10_000_00)
	if err != nil {
		t.Errorf("FindPaymentByID(): can't register account, error = %v", err)
		return
	}

	// осуществляем платёж на его счёт
	payment, err := s.Pay(account.ID, 1000_00, "auto")
	if err != nil {
		t.Errorf("FindPaymentByID(): can't create payment, error = %v", err)
		return
	}

	// пробуем найти платеж
	got, err := s.FindPaymentByID(payment.ID)
	if err != nil {
		t.Errorf("FindPaymentByID(): error = %v", err)
		return
	}

	// сравнимаем платежи
	if !reflect.DeepEqual(payment, got) {
		t.Errorf("FindPaymentByID(): wrong payment returned = %v", err)
		return
	}
}

type testService struct {
	*Service
}

func newTestService() *testService {
	return &testService{Service: &Service{}}
}

func (s *testService) addAccountWithBalance(phone types.Phone, balance types.Money) (*types.Account, error) {
	// ркгистрируем там пользователя
	account, err := s.RegisterAccount(phone)
	if err != nil {
		return nil, fmt.Errorf("can't register account, error = %v", err)
	}

	//пополняем его счёт
	err = s.Deposit(account.ID, balance)
	if err != nil {
		return nil, fmt.Errorf("can't deposit account, error = %v", err)
	}

	return account, nil
}