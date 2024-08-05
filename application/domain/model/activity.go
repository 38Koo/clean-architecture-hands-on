package model

import "time"

type ActivityId int32

type Activity struct {
	Id              *ActivityId
	OwnerAccountId  AccountId
	SourceAccountId AccountId
	TargetAccountId AccountId
	Timestamp       time.Time
	Money           Money
}

func NewActivity(
	OwnerAccountId,
	SourceAccountId,
	TargetAccountId AccountId,
	Timestamp time.Time,
	Money Money) Activity {
	return Activity{
		Id:              nil,
		OwnerAccountId:  OwnerAccountId,
		SourceAccountId: SourceAccountId,
		TargetAccountId: TargetAccountId,
		Timestamp:       Timestamp,
		Money:           Money,
	}
}
