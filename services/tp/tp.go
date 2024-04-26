package tp

import "time"

type TpService interface {
	CreateDeposit() (Response, error)
}

type tpService struct{}

type Req struct {
	Amount    string
	Reference string
	CreatedAt time.Time
}

type Response struct {
	Status    string
	Amount    string
	Reference string
}

func New() TpService {
	return &tpService{}
}

func (t *tpService) CreateDeposit() (Response, error) {
	return Response{}, nil
}
