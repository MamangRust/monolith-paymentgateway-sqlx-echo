package roleapimapper

type RoleResponseMapper interface {
	QueryMapper() RoleQueryResponseMapper
	CommandMapper() RoleCommandResponseMapper
}

type roleResponseMapper struct {
	queryMapper   RoleQueryResponseMapper
	commandMapper RoleCommandResponseMapper
}

func NewRoleResponseMapper() RoleResponseMapper {
	return &roleResponseMapper{
		queryMapper:   NewRoleQueryResponseMapper(),
		commandMapper: NewRoleCommandResponseMapper(),
	}
}

func (r *roleResponseMapper) QueryMapper() RoleQueryResponseMapper {
	return r.queryMapper
}

func (r *roleResponseMapper) CommandMapper() RoleCommandResponseMapper {
	return r.commandMapper
}
