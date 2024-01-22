package http

import (
	"EWallet/internal/core/domain"
	"EWallet/internal/core/port"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

type WalletHandler struct {
	svc port.WalletService
}

func NewWalletHandler(svc port.WalletService) *WalletHandler {
	return &WalletHandler{
		svc,
	}
}

func (wh *WalletHandler) Create(ctx *gin.Context) {
	wallet := domain.Wallet{
		Id:      uuid.New(),
		Balance: 100.0,
	}

	_, err := wh.svc.Create(ctx, &wallet)
	if err != nil {
		handleError(ctx, domain.ErrBadRequest)
		return
	}

	rsp := newWalletResponse(&wallet)

	ctx.JSON(http.StatusOK, rsp)
}

type getWalletRequest struct {
	Id string `uri:"walletId" binding:"required,uuid"`
}

func (wh *WalletHandler) GetWallet(ctx *gin.Context) {
	var req getWalletRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		validationError(ctx, err)
		return
	}

	id, err := uuid.Parse(req.Id)
	if err != nil {
		handleError(ctx, domain.ErrBadRequest)
		return
	}

	wallet, err := wh.svc.GetWallet(ctx, id)
	if err != nil {
		handleError(ctx, domain.ErrDataNotFound)
		return
	}

	rsp := newWalletResponse(wallet)

	ctx.JSON(http.StatusOK, rsp)
}
