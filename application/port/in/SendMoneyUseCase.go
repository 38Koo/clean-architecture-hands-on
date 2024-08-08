package portIn

type SendMoneyUseCase interface {
	SendMoney(command SendMoneyCommand) bool
}