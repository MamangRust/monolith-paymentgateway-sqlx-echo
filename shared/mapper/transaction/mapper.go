package transactionapimapper

type TransactionResponseMapper interface {
	QueryMapper() TransactionQueryResponseMapper
	CommandMapper() TransactionCommandResponseMapper
	StatusStatsMapper() TransactionStatsStatusResponseMapper
	MethodStatsMapper() TransactionStatsMethodResponseMapper
	AmountStatsMapper() TransactionStatsAmountResponseMapper
}

type transactionResponseMapper struct {
	queryMapper   TransactionQueryResponseMapper
	commandMapper TransactionCommandResponseMapper
	statusMapper  TransactionStatsStatusResponseMapper
	methodMapper  TransactionStatsMethodResponseMapper
	amountMapper  TransactionStatsAmountResponseMapper
}

func NewTransactionResponseMapper() TransactionResponseMapper {
	return &transactionResponseMapper{
		queryMapper:   NewTransactionQueryResponseMapper(),
		commandMapper: NewTransactionCommandResponseMapper(),
		statusMapper:  NewTransactionStatsStatusResponseMapper(),
		methodMapper:  NewTransactionStatsMethodResponseMapper(),
		amountMapper:  NewTransactionStatsAmountResponseMapper(),
	}
}

func (t *transactionResponseMapper) QueryMapper() TransactionQueryResponseMapper {
	return t.queryMapper
}

func (t *transactionResponseMapper) CommandMapper() TransactionCommandResponseMapper {
	return t.commandMapper
}

func (t *transactionResponseMapper) StatusStatsMapper() TransactionStatsStatusResponseMapper {
	return t.statusMapper
}

func (t *transactionResponseMapper) MethodStatsMapper() TransactionStatsMethodResponseMapper {
	return t.methodMapper
}

func (t *transactionResponseMapper) AmountStatsMapper() TransactionStatsAmountResponseMapper {
	return t.amountMapper
}
