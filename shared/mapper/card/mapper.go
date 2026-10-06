package cardapimapper

type CardResponseMapper interface {
	QueryMapper() CardQueryResponseMapper
	CommandMapper() CardCommandResponseMapper
	AmountStatsMapper() CardStatsAmountResponseMapper
	BalanceStatsMapper() CardStatsBalanceResponseMapper
	DashboardMapper() CardDashboardResponseMapper
}

type cardResponseMapper struct {
	queryMapper        CardQueryResponseMapper
	commandMapper      CardCommandResponseMapper
	amountStatsMapper  CardStatsAmountResponseMapper
	balanceStatsMapper CardStatsBalanceResponseMapper
	dashboardMapper    CardDashboardResponseMapper
}

func NewCardResponseMapper() CardResponseMapper {
	return &cardResponseMapper{
		queryMapper:        NewCardQueryResponseMapper(),
		commandMapper:      NewCardCommandResponseMapper(),
		amountStatsMapper:  NewCardStatsAmountResponseMapper(),
		balanceStatsMapper: NewCardStatsBalanceResponseMapper(),
		dashboardMapper:    NewCardDashboardResponseMapper(),
	}
}

func (c *cardResponseMapper) QueryMapper() CardQueryResponseMapper {
	return c.queryMapper
}

func (c *cardResponseMapper) CommandMapper() CardCommandResponseMapper {
	return c.commandMapper
}

func (c *cardResponseMapper) AmountStatsMapper() CardStatsAmountResponseMapper {
	return c.amountStatsMapper
}

func (c *cardResponseMapper) BalanceStatsMapper() CardStatsBalanceResponseMapper {
	return c.balanceStatsMapper
}

func (c *cardResponseMapper) DashboardMapper() CardDashboardResponseMapper {
	return c.dashboardMapper
}
