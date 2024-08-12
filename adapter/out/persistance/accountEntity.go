package adapterOutPersistance

type AccountEntity struct {
	ID int32 `gorm:"primaryKey;autoIncrement"`
}

func (a AccountEntity) AccountTBL() string {
	return "accounts"
}
