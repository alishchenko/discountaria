package postgres

import (
	"database/sql"
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/alishchenko/discountaria/internal/data"
	"github.com/fatih/structs"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

const (
	offerUsagesTable = "offer_sages"

	offerUsagesId            = "id"
	offerUsagesName          = "name"
	offerUsagesPhotoUrl      = "photo_url"
	offerUsagesPhoneNumber   = "phone"
	offerUsagesPassword      = "password"
	offerUsagesPhoneVerified = "phone_verified"
	offerUsagesEmailVerified = "email_verified"
	offerUsagesEmail         = "email"
)

type offerUsagesQ struct {
	database *sqlx.DB
	selector squirrel.SelectBuilder
	updater  squirrel.UpdateBuilder
}

func NewOfferUsagesQ(db *sqlx.DB) data.OfferUsagesQ {
	return &offerUsagesQ{
		database: db,
		selector: squirrel.Select(fmt.Sprintf("%s.*", offerUsagesTable)).PlaceholderFormat(squirrel.Dollar).From(offerUsagesTable),
		updater:  squirrel.Update(offerUsagesTable),
	}
}

func (q *offerUsagesQ) New() data.OfferUsagesQ {
	return NewOfferUsagesQ(q.database)
}

func (q *offerUsagesQ) Insert(offerUsages data.OfferUsages) (int64, error) {
	stmt := squirrel.StatementBuilder.
		PlaceholderFormat(squirrel.Dollar).
		Insert(offerUsagesTable).
		Suffix("returning id").
		SetMap(structs.Map(&offerUsages))

	var id int64
	query, args, _ := stmt.ToSql()
	err := q.database.Get(&id, query, args...)

	return id, err
}

func (q *offerUsagesQ) Select() ([]data.OfferUsages, error) {
	var res []data.OfferUsages
	query, args, err := q.selector.ToSql()
	if err != nil {
		return nil, errors.Wrap(err, "failed to build query")
	}

	err = q.database.Select(&res, query, args...)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return res, nil
}

func (q *offerUsagesQ) Get() (*data.OfferUsages, error) {
	var res data.OfferUsages
	query, args, err := q.selector.ToSql()
	if err != nil {
		return nil, errors.Wrap(err, "failed to build query")
	}

	err = q.database.Get(&res, query, args...)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return &res, nil
}

func (q *offerUsagesQ) FilterById(id int64) data.OfferUsagesQ {
	q.selector = q.selector.Where(squirrel.Eq{offerUsagesId: id})
	q.updater = q.updater.Where(squirrel.Eq{offerUsagesId: id})

	return q
}
func (q *offerUsagesQ) FilterByEmail(email string) data.OfferUsagesQ {
	q.selector = q.selector.Where(squirrel.Eq{offerUsagesEmail: email})
	q.updater = q.updater.Where(squirrel.Eq{offerUsagesEmail: email})

	return q
}
