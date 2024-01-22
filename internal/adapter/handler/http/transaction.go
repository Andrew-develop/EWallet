package http

import (
	"EWallet/internal/core/domain"
	"EWallet/internal/core/port"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"time"
)

type TransactionHandler struct {
	tsvc port.TransactionService
	wsvc port.WalletService
}

func NewTransactionHandler(tsvc port.TransactionService, wsvc port.WalletService) *TransactionHandler {
	return &TransactionHandler{
		tsvc,
		wsvc,
	}
}

type getTransactionUriRequest struct {
	From string `uri:"walletId" binding:"required,uuid"`
}

type getTransactionJsonRequest struct {
	To     string  `json:"to" binding:"required"`
	Amount float64 `json:"amount" binding:"required"`
}

func (th *TransactionHandler) CreateTransaction(ctx *gin.Context) {
	var uriReq getTransactionUriRequest
	var jsonReq getTransactionJsonRequest

	if err := ctx.ShouldBindJSON(&jsonReq); err != nil {
		validationError(ctx, err)
		return
	}

	if err := ctx.ShouldBindUri(&uriReq); err != nil {
		validationError(ctx, err)
		return
	}

	from, err := uuid.Parse(uriReq.From)
	if err != nil {
		handleError(ctx, domain.ErrBadRequest)
		return
	}

	wallet, err := th.wsvc.GetWallet(ctx, from)
	if err != nil {
		handleError(ctx, domain.ErrDataNotFound)
		return
	}

	if wallet.Balance < jsonReq.Amount {
		handleError(ctx, domain.ErrBadRequest)
		return
	}

	to, err := uuid.Parse(jsonReq.To)
	if err != nil {
		handleError(ctx, domain.ErrBadRequest)
		return
	}

	_, err = th.wsvc.GetWallet(ctx, to)
	if err != nil {
		handleError(ctx, domain.ErrBadRequest)
		return
	}

	transaction := domain.Transaction{
		Id:       uuid.New(),
		From:     from,
		To:       to,
		Amount:   jsonReq.Amount,
		DateTime: time.Now(),
	}

	_, err = th.tsvc.CreateTransaction(ctx, &transaction)
	if err != nil {
		handleError(ctx, err)
		return
	}

	rsp := newTransactionResponse(&transaction)

	ctx.JSON(http.StatusOK, rsp)
}

type listTransactionsRequest struct {
	WalletId string `uri:"walletId" binding:"required,uuid"`
}

func (th *TransactionHandler) ListTransactions(ctx *gin.Context) {
	var req listTransactionsRequest
	transactionsList := []transactionResponse{}

	if err := ctx.ShouldBindUri(&req); err != nil {
		validationError(ctx, err)
		return
	}

	id, err := uuid.Parse(req.WalletId)
	if err != nil {
		handleError(ctx, domain.ErrBadRequest)
		return
	}

	transactions, err := th.tsvc.ListTransactions(ctx, id)
	if err != nil {
		handleError(ctx, domain.ErrDataNotFound)
		return
	}

	for _, transaction := range transactions {
		transactionsList = append(transactionsList, newTransactionResponse(&transaction))
	}

	ctx.JSON(http.StatusOK, transactionsList)
}
