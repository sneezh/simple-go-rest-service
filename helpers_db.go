package main

import (
	"github.com/whenspeakteam/pg/orm"
)

func getAllEntities(limit int, offset int, response interface{}) (err error) {
	err = db.Model(response).
		Apply(pagination(limit, offset)).
		Order("id ASC").
		Select()

	return err
}

func pagination(limit int, offset int) func(query *orm.Query) (*orm.Query, error) {
	filter := func(query *orm.Query) (*orm.Query, error) {
		if limit != 0 {
			limit = 10
		}
		query = query.Limit(limit)
		if offset != 0 {
			query = query.Offset(offset)
		}
		return query, nil
	}
	return filter
}
