package merchantapi

import (
	"github.com/mituan8/pay/internal/auth"
	"github.com/mituan8/pay/internal/bus"
	"github.com/mituan8/pay/internal/service/blockchain"
	"github.com/mituan8/pay/internal/service/merchant"
	"github.com/mituan8/pay/internal/service/payment"
	"github.com/mituan8/pay/internal/service/wallet"
	"github.com/rs/zerolog"
)

type BlockchainService interface {
	blockchain.Resolver
	blockchain.Convertor
}

type Handler struct {
	merchants  *merchant.Service
	tokens     *auth.TokenAuthManager
	payments   *payment.Service
	wallets    *wallet.Service
	blockchain BlockchainService
	publisher  bus.Publisher
	logger     *zerolog.Logger
}

func NewHandler(
	merchants *merchant.Service,
	tokens *auth.TokenAuthManager,
	payments *payment.Service,
	wallets *wallet.Service,
	blockchainService BlockchainService,
	publisher bus.Publisher,
	logger *zerolog.Logger,
) *Handler {
	log := logger.With().Str("channel", "dashboard_handler").Logger()

	return &Handler{
		merchants:  merchants,
		tokens:     tokens,
		payments:   payments,
		wallets:    wallets,
		blockchain: blockchainService,
		publisher:  publisher,
		logger:     &log,
	}
}

func (h *Handler) MerchantService() *merchant.Service {
	return h.merchants
}
