package types

import (
	"fmt"

	// "github.com/kamarikg/wallet/pkg/wallet"
)

// Финансовые типы
type Money int64

type PaymentCategory string
type PaymentStatus string

const (
	PaymentStatusOK         PaymentStatus = "OK"
	PaymentStatusFail       PaymentStatus = "FAIL"
	PaymentStatusInProgress PaymentStatus = "INPROGRESS"
)

type Payment struct {
	ID        string
	AccountID int64
	Amount    Money
	Category  PaymentCategory
	Status    PaymentStatus
}

type Phone string

type Account struct {
	ID      int64
	Phone   Phone
	Balance Money
}

// Ошибки
type Error string

func (e Error) Error() string {
	return string(e)
}

// Коммуникации
type Messenger interface {
	Send(message string) bool
	Receive() (message string, ok bool)
}

type Telegram struct{}

func (t *Telegram) Send(message string) bool {
	return true
}

func (t *Telegram) Receive() (message string, ok bool) {
	return "", true
}

func New(text string) error {
	return &errorString{text}
}

type errorString struct {
	s string
}

func (e *errorString) Error() string {
	return e.s
}

// UUID — 16 байт
type UUID [16]byte

// Преобразование UUID в строку
func (uuid UUID) String() string {
	var buf [36]byte
	encodeHex(buf[:], uuid)
	return string(buf[:])
}

// Реализация encodeHex
func encodeHex(b []byte, uuid UUID) {
	hex := "0123456789abcdef"
	for i, v := range uuid {
		b[i*2] = hex[v>>4]
		b[i*2+1] = hex[v&0x0f]
	}
	// Добавляем дефисы в нужные позиции
	copy(b[32:], []byte("-"))
	// Для простоты можно использовать готовые пакеты, например github.com/google/uuid
}

// Тип common для тестов
type common struct {
	failed bool
}

func (c *common) log(msg string) {
	fmt.Println("LOG:", msg)
}

func (c *common) Fail() {
	c.failed = true
}

func (c *common) Error(args ...interface{}) {
	c.log(fmt.Sprintln(args...))
	c.Fail()
}

// Структура T
type T struct {
	common
	isParallel bool
	context    *testContext
}

type testContext struct {
	name string
}
