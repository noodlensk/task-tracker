// Package subscriber provides primitives to interact with the asyncapi
//
// Code generated by https://github.com/asyncapi/generator/ DO NOT EDIT.
package subscriber

type TaskEstimated struct {
	Id            string  `json:"id"`
	AssignedPrice float32 `json:"assignedPrice"`
	CompetedPrice float32 `json:"competedPrice"`
}
type UserCharged struct {
	UserUid string  `json:"user_uid"`
	Amount  float32 `json:"amount"`
	Reason  string  `json:"reason"`
}
type UserPayed struct {
	UserUid string  `json:"user_uid"`
	Amount  float32 `json:"amount"`
	Reason  string  `json:"reason"`
}
