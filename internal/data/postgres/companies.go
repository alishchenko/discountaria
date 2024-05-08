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
	companiesTable = "companies"

	companyId        = "id"
	companiesName    = "name"
	companiesLogoUrl = "logo_url"
)

type companiesQ struct {
	database *sqlx.DB
	selector squirrel.SelectBuilder
	updater  squirrel.UpdateBuilder
}

func NewCompaniesQ(db *sqlx.DB) data.CompaniesQ {
	return &companiesQ{
		database: db,
		selector: squirrel.Select(fmt.Sprintf("%s.*", companiesTable)).PlaceholderFormat(squirrel.Dollar).From(companiesTable),
		updater:  squirrel.Update(companiesTable),
	}
}

func (q *companiesQ) New() data.CompaniesQ {
	return NewCompaniesQ(q.database)
}

func (q *companiesQ) Insert(company data.Company) (int64, error) {
	stmt := squirrel.StatementBuilder.
		PlaceholderFormat(squirrel.Dollar).
		Insert(companiesTable).
		Suffix("returning id").
		SetMap(structs.Map(&company))

	var id int64
	query, args, _ := stmt.ToSql()
	err := q.database.Get(&id, query, args...)

	return id, err
}

func (q *companiesQ) Select() ([]data.Company, error) {
	var res []data.Company
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

func (q *companiesQ) Update() error {
	query, args, err := q.updater.PlaceholderFormat(squirrel.Dollar).ToSql()
	if err != nil {
		return errors.Wrap(err, "failed to build query")
	}
	_, err = q.database.Exec(query, args...)
	return err
}

func (q *companiesQ) Delete(id int64) error {
	query, args, err := squirrel.Delete(companiesTable).PlaceholderFormat(squirrel.Dollar).Where(
		squirrel.Eq{
			companyId: id,
		}).ToSql()

	if err != nil {
		return errors.Wrap(err, "failed to build query")
	}
	_, err = q.database.Exec(query, args...)
	return err
}

func (q *companiesQ) Get() (*data.Company, error) {
	var res data.Company
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

func (q *companiesQ) FilterById(id int64) data.CompaniesQ {
	q.selector = q.selector.Where(squirrel.Eq{companyId: id})
	q.updater = q.updater.Where(squirrel.Eq{companyId: id})

	return q
}
func (q *companiesQ) FilterByName(name string) data.CompaniesQ {
	q.selector = q.selector.Where(squirrel.Eq{companiesName: name})
	q.updater = q.updater.Where(squirrel.Eq{companiesName: name})

	return q
}

func (q *companiesQ) UpdateName(name string) data.CompaniesQ {
	q.updater = q.updater.Set(companiesName, name)

	return q
}

func (q *companiesQ) UpdateLogo(url string) data.CompaniesQ {
	q.updater = q.updater.Set(companiesLogoUrl, url)

	return q
}
