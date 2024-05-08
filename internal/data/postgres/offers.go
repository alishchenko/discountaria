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
	offersTable        = "offers"
	usersOffersTable   = "users_offers"
	usersOffersOfferId = "offer_id"
	usersOffersUserId  = "user_id"

	offerId            = "id"
	offerName          = "name"
	offerPhotoUrl      = "photo_url"
	offerPhoneNumber   = "phone"
	offerPassword      = "password"
	offerPhoneVerified = "phone_verified"
	offerEmailVerified = "email_verified"
	offerEmail         = "email"
)

type offersQ struct {
	database *sqlx.DB
	selector squirrel.SelectBuilder
	updater  squirrel.UpdateBuilder
}

func NewOffersQ(db *sqlx.DB) data.OffersQ {
	return &offersQ{
		database: db,
		selector: squirrel.Select(fmt.Sprintf("%s.*", offersTable)).PlaceholderFormat(squirrel.Dollar).
			From(offersTable).
			Join(fmt.Sprintf("%s ON %s.%s = %s.%s",
				usersOffersTable, offersTable, offerId,
				usersOffersTable, usersOffersOfferId)).
			GroupBy(offerId),
		updater: squirrel.Update(offersTable),
	}
}

func (q *offersQ) New() data.OffersQ {
	return NewOffersQ(q.database)
}

func (q *offersQ) Insert(offer data.Offer) (int64, error) {
	stmt := squirrel.StatementBuilder.
		PlaceholderFormat(squirrel.Dollar).
		Insert(offersTable).
		Suffix("returning id").
		SetMap(structs.Map(&offer))

	var id int64
	query, args, _ := stmt.ToSql()
	err := q.database.Get(&id, query, args...)
	stmt = squirrel.StatementBuilder.
		PlaceholderFormat(squirrel.Dollar).Insert(usersOffersTable)
	return id, err
}

func (q *offersQ) InsertUsers(offerId int64, users ...data.User) error {
	statement := squirrel.
		Insert(usersOffersTable).
		PlaceholderFormat(squirrel.Dollar).
		Columns(usersOffersOfferId, usersOffersUserId)
	for _, user := range users {
		statement = statement.Values(offerId, user.Id)
	}
	query, args, err := statement.ToSql()
	if err != nil {
		return errors.Wrap(err, "failed to build query")
	}
	_, err = q.database.Exec(query, args...)
	return err
}

func (q *offersQ) Select() ([]data.Offer, error) {
	var res []data.Offer
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

func (q *offersQ) Get() (*data.Offer, error) {
	var res data.Offer
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

func (q *offersQ) Update() error {
	query, args, err := q.updater.PlaceholderFormat(squirrel.Dollar).ToSql()
	if err != nil {
		return errors.Wrap(err, "failed to build query")
	}
	_, err = q.database.Exec(query, args...)
	return err
}

func (q *offersQ) Delete(id int64) error {
	query, args, err := squirrel.Delete(offersTable).PlaceholderFormat(squirrel.Dollar).Where(
		squirrel.Eq{
			offerId: id,
		}).ToSql()

	if err != nil {
		return errors.Wrap(err, "failed to build query")
	}
	_, err = q.database.Exec(query, args...)
	return err
}

func (q *offersQ) FilterById(id int64) data.OffersQ {
	q.selector = q.selector.Where(squirrel.Eq{offerId: id})
	q.updater = q.updater.Where(squirrel.Eq{offerId: id})

	return q
}
func (q *offersQ) FilterByEmail(email string) data.OffersQ {
	q.selector = q.selector.Where(squirrel.Eq{offerEmail: email})
	q.updater = q.updater.Where(squirrel.Eq{offerEmail: email})

	return q
}

func (q *offersQ) UpdateName(name string) data.OffersQ {
	q.updater = q.updater.Set(offerName, name)

	return q
}

func (q *offersQ) UpdateEmail(email string) data.OffersQ {
	q.updater = q.updater.Set(offerEmail, email)
	q.updater = q.updater.Set(offerEmailVerified, false)

	return q
}

func (q *offersQ) UpdatePhone(phone string) data.OffersQ {
	q.updater = q.updater.Set(offerPhoneNumber, phone)
	q.updater = q.updater.Set(offerPhoneVerified, false)

	return q
}

func (q *offersQ) UpdateEmailVerified(isVerified bool) data.OffersQ {
	q.updater = q.updater.Set(offerEmailVerified, isVerified)

	return q
}

func (q *offersQ) UpdatePassword(password string) data.OffersQ {
	q.updater = q.updater.Set(offerPassword, password)

	return q
}

func (q *offersQ) UpdatePhotoUrl(url string) data.OffersQ {
	q.updater = q.updater.Set(offerPhotoUrl, url)

	return q
}

func (q *offersQ) UpdatePhoneVerified(isVerified bool) data.OffersQ {
	q.updater = q.updater.Set(offerPhoneVerified, isVerified)

	return q
}

func (q *offersQ) JoinUsers() data.OffersQ {
	q.selector = q.selector.Join("")

	return q
}
