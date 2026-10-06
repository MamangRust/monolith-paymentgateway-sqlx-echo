package topupapimapper

type TopupResponseMapper interface {
	QueryMapper() TopupQueryResponseMapper
	CommandMapper() TopupCommandResponseMapper
	StatusStatsMapper() TopupStatsStatusResponseMapper
	AmountStatsMapper() TopupStatsAmountResponseMapper
	MethodStatsMapper() TopupStatsMethodResponseMapper
}

type topupResponseMapper struct {
	queryMapper   TopupQueryResponseMapper
	commandMapper TopupCommandResponseMapper
	statusStats   TopupStatsStatusResponseMapper
	amountStats   TopupStatsAmountResponseMapper
	methodStats   TopupStatsMethodResponseMapper
}

func NewTopupResponseMapper() TopupResponseMapper {
	return &topupResponseMapper{
		queryMapper:   NewTopupQueryResponseMapper(),
		commandMapper: NewTopupCommandResponseMapper(),
		statusStats:   NewTopupStatsStatusResponseMapper(),
		amountStats:   NewTopupStatsAmountResponseMapper(),
		methodStats:   NewTopupStatsMethodResponseMapper(),
	}
}

func (t *topupResponseMapper) QueryMapper() TopupQueryResponseMapper {
	return t.queryMapper
}

func (t *topupResponseMapper) CommandMapper() TopupCommandResponseMapper {
	return t.commandMapper
}

func (t *topupResponseMapper) StatusStatsMapper() TopupStatsStatusResponseMapper {
	return t.statusStats
}

func (t *topupResponseMapper) AmountStatsMapper() TopupStatsAmountResponseMapper {
	return t.amountStats
}

func (t *topupResponseMapper) MethodStatsMapper() TopupStatsMethodResponseMapper {
	return t.methodStats
}
