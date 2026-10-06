package transferapimapper

type TransferResponseMapper interface {
	QueryMapper() TransferQueryResponseMapper
	CommandMapper() TransferCommandResponseMapper
	StatusStatsMapper() TransferStatsStatusResponseMapper
	AmountStatsMapper() TransferStatsAmountResponseMapper
}

type transferResponseMapper struct {
	queryMapper   TransferQueryResponseMapper
	commandMapper TransferCommandResponseMapper
	statusStats   TransferStatsStatusResponseMapper
	amountStats   TransferStatsAmountResponseMapper
}

func NewTransferResponseMapper() TransferResponseMapper {
	return &transferResponseMapper{
		queryMapper:   NewTransferQueryResponseMapper(),
		commandMapper: NewTransferCommandResponseMapper(),
		statusStats:   NewTransferStatsStatusResponseMapper(),
		amountStats:   NewTransferStatsAmountResponseMapper(),
	}
}

func (t *transferResponseMapper) QueryMapper() TransferQueryResponseMapper {
	return t.queryMapper
}

func (t *transferResponseMapper) CommandMapper() TransferCommandResponseMapper {
	return t.commandMapper
}

func (t *transferResponseMapper) StatusStatsMapper() TransferStatsStatusResponseMapper {
	return t.statusStats
}

func (t *transferResponseMapper) AmountStatsMapper() TransferStatsAmountResponseMapper {
	return t.amountStats
}
