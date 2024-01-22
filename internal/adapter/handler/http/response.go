package http

import (
	"EWallet/internal/core/domain"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"net/http"
	"time"
)

type walletResponse struct {
	Id      uuid.UUID
	Balance float64
}

func newWalletResponse(wallet *domain.Wallet) walletResponse {
	return walletResponse{
		Id:      wallet.Id,
		Balance: wallet.Balance,
	}
}

type transactionResponse struct {
	Time   time.Time
	From   uuid.UUID
	To     uuid.UUID
	Amount float64
}

func newTransactionResponse(transaction *domain.Transaction) transactionResponse {
	return transactionResponse{
		Time:   transaction.DateTime,
		From:   transaction.From,
		To:     transaction.To,
		Amount: transaction.Amount,
	}
}

var errorStatusMap = map[error]int{
	domain.ErrBadRequest:   http.StatusBadRequest,
	domain.ErrDataNotFound: http.StatusNotFound,
}

func validationError(ctx *gin.Context, err error) {
	errMsgs := parseError(err)
	errRsp := newErrorResponse(errMsgs)
	ctx.JSON(http.StatusBadRequest, errRsp)
}

func handleError(ctx *gin.Context, err error) {
	statusCode, ok := errorStatusMap[err]
	if !ok {
		statusCode = http.StatusInternalServerError
	}

	errMsg := parseError(err)
	errRsp := newErrorResponse(errMsg)
	ctx.JSON(statusCode, errRsp)
}

func parseError(err error) []string {
	var errMsgs []string

	if errors.As(err, &validator.ValidationErrors{}) {
		for _, err := range err.(validator.ValidationErrors) {
			errMsgs = append(errMsgs, err.Error())
		}
	} else {
		errMsgs = append(errMsgs, err.Error())
	}

	return errMsgs
}

type errorResponse struct {
	Success  bool     `json:"success"`
	Messages []string `json:"messages"`
}

func newErrorResponse(errMsgs []string) errorResponse {
	return errorResponse{
		Success:  false,
		Messages: errMsgs,
	}
}
