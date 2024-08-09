package portIn

type SendMoneyUseCase interface {
	SendMoney(command SendMoneyDTO) bool
}