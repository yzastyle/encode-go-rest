package postgre

import (
	"github.com/gocraft/dbr/v2"
)

type QueryBuilder struct {
	s *dbr.SelectStmt
}

func (q *QueryBuilder) CreateQuery(s *dbr.SelectStmt) *QueryBuilder {
	return &QueryBuilder{s: s}
}

func (q *QueryBuilder) HasFirstName(firstName string) *QueryBuilder {
	if firstName != "" {
		q.s.Where("first_name = ?", firstName)
	}
	return q
}

func (q *QueryBuilder) HasLastName(lastName string) *QueryBuilder {
	if lastName != "" {
		q.s.Where("last_name = ?", lastName)
	}
	return q
}

func (q *QueryBuilder) HasEmail(email string) *QueryBuilder {
	if email != "" {
		q.s.Where("email = ?", email)
	}
	return q
}

func (q *QueryBuilder) HasPhone(phone string) *QueryBuilder {
	if phone != "" {
		q.s.Where("phone = ?", phone)
	}
	return q
}

func (q *QueryBuilder) WithLimit(l uint64) *QueryBuilder {
	if l != 0 {
		q.s.Limit(l)
	}
	return q
}

func (q *QueryBuilder) WithOffest(o uint64) *QueryBuilder {
	if o != 0 {
		q.s.Offset(o)
	}
	return q
}

func (q *QueryBuilder) Build() *dbr.SelectStmt {
	return q.s
}
