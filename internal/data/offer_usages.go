package data

import "time"

type OfferUsagesQ interface {
	Insert(OfferUsages) (int64, error)
	Select() ([]OfferUsages, error)
	Get() (*OfferUsages, error)

	FilterById(id int64) OfferUsagesQ
}

type OfferUsages struct {
	Id      int64     `db:"id" structs:"-" json:"-"`
	OfferId int64     `db:"company_id" structs:"company_id"`
	UserId  int64     `db:"users" structs:"users"`
	UsedAt  time.Time `db:"used_at" structs:"used_at"`
}
