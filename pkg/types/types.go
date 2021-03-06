package types

//Money the minimal money unit
type Money int64

//Category пердставляет собой категорию, в которй был соверщен платеж(авто, аптеки, рестораны и т.д)
type Category string

//Status представляет собой статус платежа.
type Status string

//Предопределенные статусы платеже
const (
	StatusOk         Status = "OK"
	StatusFail       Status = "FAIL"
	StatusInProgress Status = "INPROGRESS"
)

//Payment представляет информация о платеже
type Payment struct {
	ID       int
	Amount   Money
	Category Category
	Status   Status
}
