package account

import "time"

type Transaction struct {
	UID       string
	UserUID   string
	Amount    float32
	Reason    string
	CreatedAt time.Time
}
