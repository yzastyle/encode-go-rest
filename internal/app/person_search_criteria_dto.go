package app

type PersonSearchCriteriaDTO struct {
	Limit     uint64 `query:"limit"`
	Offset    uint64 `query:"offset"`
	Email     string `query:"email"`
	Phone     string `query:"phone"`
	FirstName string `query:"first-name"`
	LastName  string `query:"last-name"`
}
