package responses

import (
	"github.com/alishchenko/discountaria/internal/data"
	"net/http"
	"time"
)

type CompanyResponse struct {
	Id          int64     `db:"id" structs:"-" json:"id"`
	Name        string    `db:"name" json:"name" structs:"name"`
	Description *string   `db:"description" json:"description" structs:"description"`
	Category    string    `db:"category" json:"category" structs:"category"`
	LogoURL     *string   `db:"logo_url" json:"logo_url" structs:"logo_url"`
	URL         *string   `db:"url" json:"url" structs:"url"`
	UserId      int64     `db:"user_id" json:"user_id" structs:"user_id"`
	CreatedAt   time.Time `db:"created_at" json:"created_at" structs:"created_at"`
}

type CompanyListResponse struct {
	Data  []CompanyResponse
	Links Links
}

func ToGetCompanyResponse(company data.Company) CompanyResponse {
	return CompanyResponse{
		Id:          company.Id,
		Name:        company.Name,
		Description: company.Description,
		Category:    company.Category,
		LogoURL:     company.LogoURL,
		URL:         company.URL,
		UserId:      company.UserId,
		CreatedAt:   company.CreatedAt,
	}
}

func ToListCompanyResponse(r *http.Request, companies []data.Company) CompanyListResponse {
	var resp CompanyListResponse
	for _, company := range companies {
		resp.Data = append(resp.Data, ToGetCompanyResponse(company))
	}
	resp.Links.Next = "/companies?" + r.URL.RawQuery

	return resp
}
