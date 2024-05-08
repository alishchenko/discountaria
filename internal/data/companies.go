package data

import "time"

type CompaniesQ interface {
	Insert(Company) (int64, error)
	Update() error
	Select() ([]Company, error)
	Get() (*Company, error)
	Delete(id int64) error

	FilterById(id int64) CompaniesQ
	FilterByName(name string) CompaniesQ

	UpdateName(name string) CompaniesQ
	UpdateLogo(url string) CompaniesQ
}

type Company struct {
	Id        int64     `db:"id" structs:"-" json:"-"`
	Name      string    `db:"name" json:"name" structs:"name"`
	LogoURL   *string   `db:"logo_url" json:"logo_url" structs:"logo_url"`
	UserId    int64     `db:"user_id" json:"user_id" structs:"user_id"`
	CreatedAt time.Time `db:"created_at" json:"created_at" structs:"created_at"`
}
