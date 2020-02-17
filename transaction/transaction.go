package transaction

import (
	"fmt"
)

type Transaction struct {
	From   string
	To     string
	Amount uint64
	//Timestamp time.Time
}

func (this *Transaction) Show() {
	fmt.Printf("From: %v ---%v---> %v\n", this.From, this.Amount, this.To)
}
