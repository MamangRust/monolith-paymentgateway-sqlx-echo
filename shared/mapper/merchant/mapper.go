package merchantapimapper

type MerchantResponseMapper interface {
	QueryMapper() MerchantQueryResponseMapper
	CommandMapper() MerchantCommandResponseMapper
	TransactionMapper() MerchantTransactionResponseMapper
	MethodStatsMapper() MerchantStatsMethodResponseMapper
	AmountStatsMapper() MerchantStatsAmountResponseMapper
	TotalAmountStatsMapper() MerchantStatsTotalAmountResponseMapper
}

type merchantResponseMapper struct {
	queryMapper       MerchantQueryResponseMapper
	commandMapper     MerchantCommandResponseMapper
	transactionMapper MerchantTransactionResponseMapper
	methodStatsMapper MerchantStatsMethodResponseMapper
	amountStatsMapper MerchantStatsAmountResponseMapper
	totalAmountMapper MerchantStatsTotalAmountResponseMapper
}

func NewMerchantResponseMapper() MerchantResponseMapper {
	return &merchantResponseMapper{
		queryMapper:       NewMerchantQueryResponseMapper(),
		commandMapper:     NewMerchantCommandResponseMapper(),
		transactionMapper: NewMerchantTransactionResponseMapper(),
		methodStatsMapper: NewMerchantStatsMethodResponseMapper(),
		amountStatsMapper: NewMerchantStatsAmountResponseMapper(),
		totalAmountMapper: NewMerchantStatsTotalAmountResponseMapper(),
	}
}

func (m *merchantResponseMapper) QueryMapper() MerchantQueryResponseMapper {
	return m.queryMapper
}

func (m *merchantResponseMapper) CommandMapper() MerchantCommandResponseMapper {
	return m.commandMapper
}

func (m *merchantResponseMapper) TransactionMapper() MerchantTransactionResponseMapper {
	return m.transactionMapper
}

func (m *merchantResponseMapper) MethodStatsMapper() MerchantStatsMethodResponseMapper {
	return m.methodStatsMapper
}

func (m *merchantResponseMapper) AmountStatsMapper() MerchantStatsAmountResponseMapper {
	return m.amountStatsMapper
}

func (m *merchantResponseMapper) TotalAmountStatsMapper() MerchantStatsTotalAmountResponseMapper {
	return m.totalAmountMapper
}
