package adapterOutPersistance

import "time"

type ActivityEntity struct {
	ID              uint      `gorm:"primaryKey;autoIncrement"`
	LocalDataTime   time.Time `gorm:"not null"`
	OwnerAccountID  int64     `gorm:"not null"`
	SourceAccountID int64     `gorm:"not null"`
	TargetAccountID int64     `gorm:"not null"`
	Amount          int64     `gorm:"not null"`
}

func (a ActivityEntity) ActivityTBL() string {
	return "account_activities"
}
