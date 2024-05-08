package data

import "time"

type NoncesQ interface {
	Insert(Nonce) (int64, error)
	Select() ([]Nonce, error)
	Get() (*Nonce, error)

	FilterById(id int64) NoncesQ
}

type Nonce struct {
	Id         int64      `db:"id" structs:"-" json:"-"`
	Identifier string     `db:"identifier" structs:"identifier"`
	Nonce      string     `db:"nonce" structs:"nonce"`
	ExpiredAt  *time.Time `db:"expired_at" structs:"expired_at"`
}
