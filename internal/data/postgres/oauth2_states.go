package postgres

import (
	"database/sql"
	"github.com/alishchenko/discountaria/internal/data"
	"github.com/jmoiron/sqlx"

	sq "github.com/Masterminds/squirrel"
	"github.com/pkg/errors"
)

const oAuth2StatesTable = "oauth2_states"

var ErrNoSuchState = errors.New("no such state")

type oAuth2StatesQ struct {
	db            *sqlx.DB
	selectBuilder sq.SelectBuilder
	deleteBuilder sq.DeleteBuilder
}

func NewOAuth2StatesQ(db *sqlx.DB) data.OAuth2StatesQ {
	return &oAuth2StatesQ{
		db:            db,
		selectBuilder: sq.Select("*").From(oAuth2StatesTable).PlaceholderFormat(sq.Dollar),
		deleteBuilder: sq.Delete(oAuth2StatesTable).PlaceholderFormat(sq.Dollar),
	}
}

func (q *oAuth2StatesQ) New() data.OAuth2StatesQ {
	return NewOAuth2StatesQ(q.db)
}

func (q *oAuth2StatesQ) Create(oAuth2State data.OAuth2State) (*data.OAuth2State, error) {
	clauses := map[string]interface{}{
		"state":      oAuth2State.State,
		"valid_till": oAuth2State.ValidTill,
	}
	result := data.OAuth2State{}
	stmt := sq.Insert(oAuth2StatesTable).
		PlaceholderFormat(sq.Dollar).SetMap(clauses).Suffix("RETURNING *")
	query, args, _ := stmt.ToSql()
	err := q.db.Get(&result, query, args...)
	return &result, errors.Wrap(err, "failed to create oauth2 state")
}

func (q *oAuth2StatesQ) Get() (*data.OAuth2State, error) {
	result := new(data.OAuth2State)
	query, args, _ := q.selectBuilder.ToSql()
	err := q.db.Get(result, query, args...)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}

	return result, errors.Wrap(err, "failed to get oauth2 state")
}

func (q *oAuth2StatesQ) Delete() error {
	query, args, err := q.deleteBuilder.ToSql()
	if err != nil {
		return errors.Wrap(err, "failed to build query")
	}
	_, err = q.db.Exec(query, args...)
	return err
}

func (q *oAuth2StatesQ) FilterByID(id int64) data.OAuth2StatesQ {
	q.selectBuilder = q.selectBuilder.Where(sq.Eq{"id": id})
	q.deleteBuilder = q.deleteBuilder.Where(sq.Eq{"id": id})

	return q
}
func (q *oAuth2StatesQ) FilterByState(state string) data.OAuth2StatesQ {
	q.selectBuilder = q.selectBuilder.Where(sq.Eq{"state": state})
	q.deleteBuilder = q.deleteBuilder.Where(sq.Eq{"state": state})

	return q
}
