package types

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
