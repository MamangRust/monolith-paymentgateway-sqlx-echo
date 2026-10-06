package repository

import (
	"github.com/jmoiron/sqlx"

	pbuser "github.com/MamangRust/monolith-payment-gateway-pb/user"
	merchantstatsrepository "github.com/MamangRust/monolith-payment-gateway-merchant/repository/stats"
	merchantstatsapikeyrepository "github.com/MamangRust/monolith-payment-gateway-merchant/repository/statsbyapikey"
	merchantstatsmerchantrepository "github.com/MamangRust/monolith-payment-gateway-merchant/repository/statsbymerchant"
	"github.com/MamangRust/monolith-payment-gateway-pkg/adapter"
	useradapter "github.com/MamangRust/monolith-payment-gateway-pkg/adapter/user"
)

type GuardOptions struct {
	User []adapter.GuardOption
}

type Repositories interface {
	MerchantQueryRepository
	MerchantCommandRepository
	MerchantDocumentQueryRepository
	MerchantDocumentCommandRepository
	MerchantTransactionRepository
	merchantstatsrepository.MerchantStatsRepository
	merchantstatsapikeyrepository.MerchantStatsByApiKeyRepository
	merchantstatsmerchantrepository.MerchantStatsByMerchantRepository
	UserRepository
}

type repositories struct {
	MerchantQueryRepository
	MerchantCommandRepository
	MerchantDocumentQueryRepository
	MerchantDocumentCommandRepository
	MerchantTransactionRepository
	merchantstatsrepository.MerchantStatsRepository
	merchantstatsapikeyrepository.MerchantStatsByApiKeyRepository
	merchantstatsmerchantrepository.MerchantStatsByMerchantRepository
	UserRepository
}

func NewRepositories(db *sqlx.DB, userQueryClient pbuser.UserQueryServiceClient, guards ...GuardOptions) Repositories {
	var g GuardOptions

	if len(guards) > 0 {
		g = guards[0]
	}

	return &repositories{
		NewMerchantQueryRepository(db),
		NewMerchantCommandRepository(db),
		NewMerchantDocumentQueryRepository(db),
		NewMerchantDocumentCommandRepository(db),
		NewMerchantTransactionRepository(db),
		merchantstatsrepository.NewMerchantStatsRepository(db),
		merchantstatsapikeyrepository.NewMerchantStatsByApiKeyRepository(db),
		merchantstatsmerchantrepository.NewMerchantStatsByMerchantRepository(db),
		useradapter.NewQueryAdapter(userQueryClient, g.User...),
	}
}
