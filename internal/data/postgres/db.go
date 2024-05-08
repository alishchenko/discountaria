package postgres

import (
	"fmt"
	"github.com/alishchenko/discountaria/internal/data"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type DB struct {
	db *sqlx.DB
}

func NewDB(storagePath string) (*DB, error) {
	db, err := sqlx.Open("postgres", storagePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}
	return &DB{
		db: db,
	}, nil
}

func (db *DB) Close() error {
	return db.db.Close()
}

func (db *DB) GetDB() *sqlx.DB {
	return db.db
}

func (db *DB) NewUsers() data.UsersQ {
	return NewUsersQ(db.db)
}

func (db *DB) NewOAuth2StatesQ() data.OAuth2StatesQ {
	return NewOAuth2StatesQ(db.db)
}

func (db *DB) NewCompanies() data.CompaniesQ {
	return NewCompaniesQ(db.db)
}

func (db *DB) NewOffers() data.OffersQ {
	return NewOffersQ(db.db)
}
