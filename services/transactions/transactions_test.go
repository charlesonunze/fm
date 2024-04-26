package transactions

import (
	"context"
	reqModel "fm/api/models/req"
	"fm/db"
	appModel "fm/models"
	"fm/services/tp"
	"log/slog"
	"reflect"
	"testing"
)

type mockTp struct{}

func (t *mockTp) CreateDeposit() (tp.Response, error) {
	return tp.Response{}, nil
}

type mockRepo struct{}

func (t *mockTp) CreateDeposit() (tp.Response, error) {
	return tp.Response{}, nil
}

func Test_paymentsService_Deposit(t *testing.T) {
	type fields struct {
		name      string
		repo      db.TransactionsRepo
		logger    *slog.Logger
		tpService tp.TpService
	}
	type args struct {
		ctx     context.Context
		deposit reqModel.CreateDeposit
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *appModel.Deposit
		wantErr bool
	}{
		{
			fields: fields{
				name:"",
				repo: db.NewTransactionsRepo(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &paymentsService{
				name:   tt.fields.name,
				repo:   tt.fields.repo,
				logger: tt.fields.logger,
			}
			got, err := s.Deposit(tt.args.ctx, tt.args.deposit)
			if (err != nil) != tt.wantErr {
				t.Errorf("paymentsService.Deposit() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("paymentsService.Deposit() = %v, want %v", got, tt.want)
			}
		})
	}
}
