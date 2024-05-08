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
	usersTable = "users"

	userId            = "id"
	userName          = "name"
	userPhotoUrl      = "photo_url"
	userPhoneNumber   = "phone"
	userPassword      = "password"
	userPhoneVerified = "phone_verified"
	userEmailVerified = "email_verified"
	userEmail         = "email"
)

type usersQ struct {
	database *sqlx.DB
	selector squirrel.SelectBuilder
	updater  squirrel.UpdateBuilder
}

func NewUsersQ(db *sqlx.DB) data.UsersQ {
	return &usersQ{
		database: db,
		selector: squirrel.Select(fmt.Sprintf("%s.*", usersTable)).PlaceholderFormat(squirrel.Dollar).From(usersTable),
		updater:  squirrel.Update(usersTable),
	}
}

func (q *usersQ) New() data.UsersQ {
	return NewUsersQ(q.database)
}

func (q *usersQ) Insert(user data.User) (int64, error) {
	stmt := squirrel.StatementBuilder.
		PlaceholderFormat(squirrel.Dollar).
		Insert(usersTable).
		Suffix("returning id").
		SetMap(structs.Map(&user))

	var id int64
	query, args, _ := stmt.ToSql()
	err := q.database.Get(&id, query, args...)

	return id, err
}

func (q *usersQ) Select() ([]data.User, error) {
	var res []data.User
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

func (q *usersQ) Get() (*data.User, error) {
	var res data.User
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

func (q *usersQ) Update() error {
	query, args, err := q.updater.PlaceholderFormat(squirrel.Dollar).ToSql()
	if err != nil {
		return errors.Wrap(err, "failed to build query")
	}
	_, err = q.database.Exec(query, args...)
	return err
}

func (q *usersQ) FilterById(id int64) data.UsersQ {
	q.selector = q.selector.Where(squirrel.Eq{userId: id})
	q.updater = q.updater.Where(squirrel.Eq{userId: id})

	return q
}
func (q *usersQ) FilterByEmail(email ...string) data.UsersQ {
	q.selector = q.selector.Where(squirrel.Eq{userEmail: email})
	q.updater = q.updater.Where(squirrel.Eq{userEmail: email})

	return q
}

func (q *usersQ) UpdateName(name string) data.UsersQ {
	q.updater = q.updater.Set(userName, name)

	return q
}

func (q *usersQ) UpdateEmail(email string) data.UsersQ {
	q.updater = q.updater.Set(userEmail, email)
	q.updater = q.updater.Set(userEmailVerified, false)

	return q
}

func (q *usersQ) UpdatePhone(phone string) data.UsersQ {
	q.updater = q.updater.Set(userPhoneNumber, phone)
	q.updater = q.updater.Set(userPhoneVerified, false)

	return q
}

func (q *usersQ) UpdateEmailVerified(isVerified bool) data.UsersQ {
	q.updater = q.updater.Set(userEmailVerified, isVerified)

	return q
}

func (q *usersQ) UpdatePassword(password string) data.UsersQ {
	q.updater = q.updater.Set(userPassword, password)

	return q
}

func (q *usersQ) UpdatePhotoUrl(url string) data.UsersQ {
	q.updater = q.updater.Set(userPhotoUrl, url)

	return q
}

func (q *usersQ) UpdatePhoneVerified(isVerified bool) data.UsersQ {
	q.updater = q.updater.Set(userPhoneVerified, isVerified)

	return q
}
