package data

import "time"

type UsersQ interface {
	Insert(User) (int64, error)
	Select() ([]User, error)
	Get() (*User, error)
	Update() error

	FilterById(id int64) UsersQ
	FilterByEmail(email ...string) UsersQ

	UpdateName(name string) UsersQ
	UpdateEmail(email string) UsersQ
	UpdatePhone(phone string) UsersQ
	UpdateEmailVerified(isVerified bool) UsersQ
	UpdatePassword(password string) UsersQ
	UpdatePhotoUrl(url string) UsersQ
}

type User struct {
	Id                    int64     `db:"id" structs:"-" json:"-"`
	Name                  string    `db:"name" json:"name" structs:"name"`
	PhotoURL              *string   `db:"photo_url" json:"photo_url" structs:"photo_url"`
	PhoneNumber           *string   `db:"phone" json:"phone" structs:"phone"`
	Email                 string    `db:"email" json:"email" structs:"email"`
	Password              string    `db:"password" json:"password" structs:"password"`
	EmailVerified         bool      `db:"email_verified" json:"email_verified" structs:"email_verified"`
	Oauth2AccountProvider *string   `db:"oauth2_account_provider" json:"oauth2_account_provider" structs:"oauth2_account_provider"`
	CreatedAt             time.Time `db:"created_at" json:"created_at" structs:"created_at"`
}
