package adapterOutPersistance

type AccountEntity struct {
	ID uint `gorm:"primaryKey;autoIncrement"`
}

func (a AccountEntity) AccountTBL() string {
	return "accounts"
}
