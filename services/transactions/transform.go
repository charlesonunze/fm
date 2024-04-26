package transactions

import (
	reqModel "fm/api/models/req"
	appModel "fm/models"
)

func toDepositAppModel(payment reqModel.CreateDeposit) appModel.Deposit {
	return appModel.Deposit{}
}
