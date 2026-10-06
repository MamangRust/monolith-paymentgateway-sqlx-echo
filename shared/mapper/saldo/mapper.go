package saldoapimapper

type saldoResponseMapper struct {
	queryMapper   SaldoQueryResponseMapper
	commandMapper SaldoCommandResponseMapper
	balanceStats  SaldoStatsBalanceResponseMapper
	totalStats    SaldoStatsTotalResponseMapper
}

type SaldoResponseMapper interface {
	QueryMapper() SaldoQueryResponseMapper
	CommandMapper() SaldoCommandResponseMapper
	BalanceStatsMapper() SaldoStatsBalanceResponseMapper
	TotalStatsMapper() SaldoStatsTotalResponseMapper
}

func NewSaldoResponseMapper() SaldoResponseMapper {
	return &saldoResponseMapper{
		queryMapper:   NewSaldoQueryResponseMapper(),
		commandMapper: NewSaldoCommandResponseMapper(),
		balanceStats:  NewSaldoStatsBalanceResponseMapper(),
		totalStats:    NewSaldoStatsTotalResponseMapper(),
	}
}

func (s *saldoResponseMapper) QueryMapper() SaldoQueryResponseMapper {
	return s.queryMapper
}

func (s *saldoResponseMapper) CommandMapper() SaldoCommandResponseMapper {
	return s.commandMapper
}

func (s *saldoResponseMapper) BalanceStatsMapper() SaldoStatsBalanceResponseMapper {
	return s.balanceStats
}

func (s *saldoResponseMapper) TotalStatsMapper() SaldoStatsTotalResponseMapper {
	return s.totalStats
}
