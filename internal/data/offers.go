package data

import (
	"github.com/alishchenko/discountaria/internal/server/requests"
	"time"
)

type OffersQ interface {
	Insert(Offer) (int64, error)
	InsertUsers(int64, ...User) error
	Select() ([]Offer, error)
	Get() (*Offer, error)
	Delete(id int64) error

	PageParams(params requests.PaginationParams) OffersQ
	FilterByCompanyName(name string) OffersQ
	FilterById(id int64) OffersQ
}

type Offer struct {
	Id          int64      `db:"id" structs:"id" json:"-"`
	CompanyId   int64      `db:"company_id" json:"company_id" structs:"company_id"`
	CompanyName string     `db:"company_name" json:"company_name" structs:"company_name"`
	Sale        int64      `db:"sale" json:"sale" structs:"sale"`
	IsPersonal  bool       `db:"is_personal" json:"is_personal" structs:"is_personal"`
	CreatedAt   time.Time  `db:"created_at" json:"created_at" structs:"created_at"`
	ExpiredAt   *time.Time `db:"expired_at" json:"expired_at" structs:"expired_at"`
}

type UsersOffers struct {
	Id      int64 `db:"id" structs:"-" json:"-"`
	OfferId int64 `db:"company_id" structs:"company_id"`
	UserId  int64 `db:"users" structs:"users"`
}
