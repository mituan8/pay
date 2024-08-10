package test

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgtype"
	"github.com/mituan8/pay/internal/auth"
	"github.com/mituan8/pay/internal/db/repository"
	kmsapi "github.com/mituan8/pay/internal/kms/api"
	"github.com/mituan8/pay/internal/lock"
	"github.com/mituan8/pay/internal/log"
	"github.com/mituan8/pay/internal/money"
	tatumprovider "github.com/mituan8/pay/internal/provider/tatum"
	"github.com/mituan8/pay/internal/provider/trongrid"
	httpServer "github.com/mituan8/pay/internal/server/http"
	"github.com/mituan8/pay/internal/server/http/merchantapi"
	merchantauth "github.com/mituan8/pay/internal/server/http/merchantapi/auth"
	"github.com/mituan8/pay/internal/server/http/middleware"
	"github.com/mituan8/pay/internal/server/http/paymentapi"
	"github.com/mituan8/pay/internal/server/http/webhook"
	"github.com/mituan8/pay/internal/service/blockchain"
	"github.com/mituan8/pay/internal/service/merchant"
	"github.com/mituan8/pay/internal/service/payment"
	"github.com/mituan8/pay/internal/service/processing"
	"github.com/mituan8/pay/internal/service/registry"
	"github.com/mituan8/pay/internal/service/transaction"
	"github.com/mituan8/pay/internal/service/user"
	"github.com/mituan8/pay/internal/service/wallet"
	"github.com/mituan8/pay/internal/test/fakes"
	"github.com/mituan8/pay/internal/util"
	kmsmock "github.com/mituan8/pay/pkg/api-kms/v1/mock"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/require"
)

type IntegrationTest struct {
	*Client

	Context context.Context

	Database   *Database
	Repository *repository.Queries
	Storage    *repository.Store

	Services *Services
	KMS      *KMS

	Providers *Providers

	Logger *zerolog.Logger

	Fakes *fakes.Fakes
	Clear *Clear
	Must  *Must

	server *httpServer.Server
}

type Services struct {
	AuthTokenManager *auth.TokenAuthManager
	Merchants        *merchant.Service
	Users            *user.Service
	Wallet           *wallet.Service
	Payment          *payment.Service
	Transaction      *transaction.Service
	Blockchain       *blockchain.Service
	Processing       *processing.Service
	Registry         *registry.Service
	Locker           *lock.Locker
	JobLogger        *log.JobLogger
}

type Providers struct {
	KMS          *kmsmock.ClientService
	Tatum        *tatumprovider.Provider
	TatumMock    *TatumMock
	Trongrid     *trongrid.Provider
	TrongridMock *fakes.Trongrid
}

var processingConfig = processing.Config{
	WebhookBasePath:         "http://localhost/webhok",
	PaymentFrontendBasePath: "https://pay.aefbay.com",
	PaymentFrontendSubPath:  "/",
	DefaultServiceFee:       0.015, // 1.5%
}

func NewIntegrationTest(t *testing.T) *IntegrationTest {
	ctx := context.Background()
	logger := zerolog.Nop()

	// DB
	db := NewDB(ctx)
	t.Cleanup(func() { db.TearDown() })

	repo := repository.New(db.Conn())
	storage := repository.NewStore(db.Conn())

	kv := registry.New(repo, &logger)

	// Providers
	tatumProvider, tatumMock := NewTatum(kv, &logger)
	trongridProvider, trongridMock := fakes.NewTrongrid(&logger)
	kmsWalletsClient := &kmsmock.ClientService{}

	providers := &Providers{
		KMS:          kmsWalletsClient,
		Tatum:        tatumProvider,
		Trongrid:     trongridProvider,
		TatumMock:    tatumMock,
		TrongridMock: trongridMock,
	}

	// Services
	currencies := blockchain.NewCurrencies()
	if err := blockchain.DefaultSetup(currencies); err != nil {
		panic("unable to setup blockchain service:" + err.Error())
	}

	blockchainService := blockchain.New(
		currencies,
		blockchain.Providers{
			Tatum:    tatumProvider,
			Trongrid: trongridProvider,
		},
		false,
		&logger,
	)

	globalFaker := fakes.New(t, blockchainService)

	locker := lock.New(storage)

	authTokenManager := auth.NewTokenAuth(repo, &logger)
	merchantsService := merchant.New(repo, blockchainService, &logger)
	usersService := user.New(storage, globalFaker.Bus, kv, &logger)
	walletsService := wallet.New(kmsWalletsClient, globalFaker.ConvertorProxy, storage, &logger)
	transactionsService := transaction.New(storage, globalFaker.CurrencyResolver, walletsService, &logger)

	paymentsService := payment.New(
		repo,
		processingConfig.PaymentFrontendBasePath,
		transactionsService,
		merchantsService,
		walletsService,
		globalFaker,
		globalFaker,
		&logger,
	)

	processingService := processing.New(
		processingConfig,
		walletsService,
		merchantsService,
		paymentsService,
		transactionsService,
		globalFaker,
		tatumProvider,
		globalFaker.Bus,
		locker,
		&logger,
	)

	jobLogger := log.NewJobLogger(storage)

	googleConfig := auth.GoogleConfig{ClientID: "1", ClientSecret: "2", RedirectCallback: "3"}
	googleAuthService := auth.NewGoogleOAuth(googleConfig, &logger)

	// HTTP Handlers
	merchantAPIHandler := merchantapi.NewHandler(
		merchantsService,
		authTokenManager,
		paymentsService,
		walletsService,
		globalFaker,
		globalFaker.Bus,
		&logger,
	)

	dashboardAuthHandler := merchantauth.NewHandler(googleAuthService, usersService, nil, &logger)

	paymentAPIHandler := paymentapi.New(
		paymentsService,
		merchantsService,
		blockchainService,
		processingService,
		&logger,
	)

	webhookHandler := webhook.New(processingService, &logger)

	// KMS
	kms := setupKMS(t, trongridProvider, &logger)

	// Server
	webConfig := httpServer.Config{
		Address: "0.0.0.0",
		Port:    "8888",
		Session: middleware.SessionConfig{
			FilesystemPath: setupTmpDir(t, "sessions"),
			Secret:         "secret",
			CookieMaxAge:   60,
		},
	}

	srv := httpServer.New(
		webConfig,
		false,
		httpServer.WithLogger(&logger),
		httpServer.WithDashboardAPI(
			webConfig,
			merchantAPIHandler,
			dashboardAuthHandler,
			authTokenManager,
			usersService,
			true,
			true,
		),
		httpServer.WithMerchantAPI(merchantAPIHandler, authTokenManager),
		httpServer.WithPaymentAPI(paymentAPIHandler, webConfig),
		httpServer.WithWebhookAPI(webhookHandler),
		kmsapi.SetupRoutes(kmsapi.New(kms.Service, &logger)),
	)

	tc := &IntegrationTest{
		Context:    ctx,
		Database:   db,
		Repository: repo,
		Storage:    storage,
		Fakes:      globalFaker,
		Logger:     &logger,
		Providers:  providers,
		Services: &Services{
			AuthTokenManager: authTokenManager,
			Merchants:        merchantsService,
			Users:            usersService,
			Wallet:           walletsService,
			Payment:          paymentsService,
			Processing:       processingService,
			Transaction:      transactionsService,
			Blockchain:       blockchainService,
			Registry:         kv,
			Locker:           locker,
			JobLogger:        jobLogger,
		},
		KMS:    kms,
		server: srv,
		Client: &Client{handler: srv.Echo().ServeHTTP},
	}

	tc.Clear = &Clear{tc}
	tc.Must = &Must{tc}

	return tc
}

func (i *IntegrationTest) CreateSamplePayment(t *testing.T, merchantID int64, opts ...payment.CreateOpt) *payment.Payment {
	return i.CreatePayment(t, merchantID, money.USD, 12.34, opts...)
}

func (i *IntegrationTest) CreatePayment(
	t *testing.T,
	merchantID int64,
	fiat money.FiatCurrency,
	amount float64,
	opts ...payment.CreateOpt,
) *payment.Payment {
	price, err := money.FiatFromFloat64(fiat, amount)
	require.NoError(t, err)

	props := payment.CreatePaymentProps{
		MerchantOrderUUID: uuid.New(),
		Money:             price,
	}

	p, err := i.Services.Payment.CreatePayment(i.Context, merchantID, props, opts...)

	require.NoError(t, err)

	return p
}

func (i *IntegrationTest) CreateRawPayment(
	t *testing.T,
	merchantID int64,
	opts ...func(*repository.CreatePaymentParams),
) repository.Payment {
	price, err := money.USD.MakeAmount("1000")
	require.NoError(t, err)

	create := repository.CreatePaymentParams{
		PublicID:          uuid.New(),
		CreatedAt:         time.Now(),
		UpdatedAt:         time.Now(),
		Type:              payment.TypePayment.String(),
		Status:            payment.StatusPending.String(),
		MerchantID:        merchantID,
		MerchantOrderUuid: uuid.New(),
		MerchantOrderID:   sql.NullString{},
		ExpiresAt:         sql.NullTime{},
		Price:             repository.MoneyToNumeric(price),
		Decimals:          int32(price.Decimals()),
		Currency:          price.Ticker(),
		Description:       sql.NullString{},
		RedirectUrl:       "",
		Metadata:          pgtype.JSONB{Status: pgtype.Null},
		IsTest:            false,
	}

	for _, opt := range opts {
		opt(&create)
	}

	pt, err := i.Repository.CreatePayment(i.Context, create)
	require.NoError(t, err)

	return pt
}

//nolint:errcheck
func (i *IntegrationTest) TearDown() {
	_ = i.server.Shutdown(context.Background())
	i.Database.TearDown()
}

func setupTmpDir(t *testing.T, dir string) string {
	fullPath := fmt.Sprintf("%s%s/%s", os.TempDir(), util.Strings.Random(6), dir)

	require.NoError(t, os.MkdirAll(fullPath, os.ModePerm))
	t.Cleanup(func() { require.NoError(t, os.RemoveAll(fullPath)) })

	return fullPath
}
