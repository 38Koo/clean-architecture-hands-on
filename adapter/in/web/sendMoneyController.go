package adapterInWeb

import (
	"clean-architecture-hands-on/application/domain/model"
	portIn "clean-architecture-hands-on/application/port/in"
	"encoding/json"
	"net/http"
	"strconv"
)

type SendMoneyController struct {
	sendMoneyUseCase portIn.SendMoneyUseCase
}

func NewSendMoneyController(sendMoneyUseCase portIn.SendMoneyUseCase) *SendMoneyController {
	return &SendMoneyController{
		sendMoneyUseCase: sendMoneyUseCase,
	}
}

func (c *SendMoneyController) SendMoney(w http.ResponseWriter, r *http.Request) {
	sourceAccountIdStr := r.URL.Query().Get("sourceAccountId")
	targetAccountIdStr := r.URL.Query().Get("targetAccountId")
	amountStr := r.URL.Query().Get("amount")

	sourceAccountIdInt, err := strconv.ParseInt(sourceAccountIdStr, 10, 32)
	if err != nil {
		http.Error(w, "Invalid source account ID", http.StatusBadRequest)
		return
	}
	sourceAccountId := model.AccountId(sourceAccountIdInt)

	targetAccountIdInt, err := strconv.ParseInt(targetAccountIdStr, 10, 32)
	if err != nil {
		http.Error(w, "Invalid target account ID", http.StatusBadRequest)
		return
	}
	targetAccountId := model.AccountId(targetAccountIdInt)

	amount, err := strconv.ParseInt(amountStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid amount", http.StatusBadRequest)
		return
	}

	sendMoneyDto := portIn.SendMoneyDTO{
		SourceAccountId: sourceAccountId,
		TargetAccountId: targetAccountId,
		Money:           model.NewMoney(amount),
	}

	if ok := c.sendMoneyUseCase.SendMoney(sendMoneyDto); !ok {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Money sent successfully")

}
