package postgres

import (
	"database/sql"
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/alishchenko/discountaria/internal/data"
	"github.com/alishchenko/discountaria/internal/server/requests"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"strings"
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
		selector: squirrel.Select(fmt.Sprintf("%s.*, companies.name as company_name", offersTable)).PlaceholderFormat(squirrel.Dollar).
			From(offersTable).
			Join("companies ON companies.id = offers.company_id").
			Join(fmt.Sprintf("%s ON %s.%s = %s.%s",
				usersOffersTable, offersTable, offerId,
				usersOffersTable, usersOffersOfferId)).
			GroupBy("offers.id", "companies.id"),
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
		Columns("company_id", "sale", "is_personal", "created_at", "expired_at").
		Values(offer.CompanyId, offer.Sale, offer.IsPersonal, offer.CreatedAt, offer.ExpiredAt)

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
	q.selector = q.selector.Join("JOIN users ON users.id = users_offers.user_id;")

	return q
}

func (q *offersQ) JoinCompanies() data.OffersQ {
	q.selector = q.selector.Join("JOIN cities ON companies.id = offers.company_id;")
	return q
}

func (q *offersQ) FilterByCompanyName(name string) data.OffersQ {
	q.selector = q.selector.Where(squirrel.Like{`LOWER(companies.name)`: "%" + strings.ToLower(name) + "%"})
	q.updater = q.updater.Where(squirrel.Like{`LOWER(companies.name)`: "%" + strings.ToLower(name) + "%"})

	return q
}

func (q *offersQ) PageParams(params requests.PaginationParams) data.OffersQ {
	if params.Order == "" {
		params.Order = "desc"
	}
	if params.Limit == 0 {
		params.Limit = 15
	}
	if params.Number == 0 {
		params.Number = 1
	}
	if params.Sort == "" {
		params.Sort = "offers.id"
	}
	q.selector = q.selector.Limit(params.Limit).Offset((params.Number - 1) * params.Limit).OrderBy(fmt.Sprintf("%s %s", params.Sort, params.Order))
	return q
}
