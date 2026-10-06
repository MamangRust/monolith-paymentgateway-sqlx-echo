package withdrawapimapper

type WithdrawResponseMapper interface {
	QueryMapper() WithdrawQueryResponseMapper
	CommandMapper() WithdrawCommandResponseMapper
	AmountStatsMapper() WithdrawStatsAmountResponseMapper
	StatusStatsMapper() WithdrawStatsStatusResponseMapper
}

type withdrawResponseMapper struct {
	queryMapper       WithdrawQueryResponseMapper
	commandMapper     WithdrawCommandResponseMapper
	amountStatsMapper WithdrawStatsAmountResponseMapper
	statusStatsMapper WithdrawStatsStatusResponseMapper
}

func NewWithdrawResponseMapper() WithdrawResponseMapper {
	return &withdrawResponseMapper{
		queryMapper:       NewWithdrawQueryResponseMapper(),
		commandMapper:     NewWithdrawCommandResponseMapper(),
		amountStatsMapper: NewWithdrawStatsAmountResponseMapper(),
		statusStatsMapper: NewWithdrawStatsStatusResponseMapper(),
	}
}

func (w *withdrawResponseMapper) QueryMapper() WithdrawQueryResponseMapper {
	return w.queryMapper
}

func (w *withdrawResponseMapper) CommandMapper() WithdrawCommandResponseMapper {
	return w.commandMapper
}

func (w *withdrawResponseMapper) AmountStatsMapper() WithdrawStatsAmountResponseMapper {
	return w.amountStatsMapper
}

func (w *withdrawResponseMapper) StatusStatsMapper() WithdrawStatsStatusResponseMapper {
	return w.statusStatsMapper
}
