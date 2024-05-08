package server

import (
	"github.com/alishchenko/discountaria/internal/server/handlers"
	"github.com/alishchenko/discountaria/internal/server/helpers"
	"github.com/alishchenko/discountaria/internal/server/middleware"
	"github.com/go-chi/chi/v5"
)

func (s *service) router() chi.Router {
	r := chi.NewRouter()

	r.Use(middleware.CtxMiddleWare(
		helpers.CtxLog(s.log),
		helpers.CtxDB(s.db),
		helpers.CtxTokens(s.cfg.Tokens),
		helpers.SetOAuth2StateConfig(s.cfg.OAuth2StateConfig),
		helpers.SetOAuth2FacebookConfig(s.cfg.OAuth2FacebookConfig.ToFacebookOauth2()),
		helpers.SetOAuth2GoogleConfig(s.cfg.OAuth2GoogleConfig.ToGoogleOauth2()),
		helpers.SetOAuth2LinkedinConfig(s.cfg.OAuth2LinkedinConfig.ToLinkedinOauth2()),
	))
	r.Route("/users", func(r chi.Router) {
		r.Post("/register", handlers.RegisterUser)
		r.Post("/login", handlers.LoginUser)
		r.Get("/", handlers.ListUsers)
		r.Route("/{id}", func(r chi.Router) {
			r.Get("/", handlers.GetUser)
			r.Patch("/", handlers.UpdateUser)
			r.Post("/companies", handlers.CreateCompany)
		})
	})
	r.Route("/oauth2", func(r chi.Router) {
		r.Route("/facebook", func(r chi.Router) {
			r.Get("/", handlers.OAuth2Facebook)
			r.Post("/callback", handlers.OAuth2FacebookCallback)
		})
		r.Route("/google", func(r chi.Router) {
			r.Get("/", handlers.OAuth2Google)
			r.Post("/callback", handlers.OAuth2GoogleCallback)
		})
		r.Route("/linkedin", func(r chi.Router) {
			r.Get("/", handlers.OAuth2Linkedin)
			r.Post("/callback", handlers.OAuth2LinkedinCallback)
		})
	})
	r.Route("/companies", func(r chi.Router) {
		r.Get("/", handlers.ListCompanies)
		r.Route("/{id}", func(r chi.Router) {
			r.Get("/", handlers.GetCompany)
			r.Patch("/", handlers.UpdateCompany)
			r.Delete("/", handlers.DeleteCompany)
			r.Post("/offers", handlers.CreateOffer)
		})
	})
	r.Route("/photo", func(r chi.Router) {
		r.Post("/", handlers.UploadPhoto)
	})
	r.Route("/offers", func(r chi.Router) {
		r.Get("/", handlers.ListOffers)
		r.Route("/{id}", func(r chi.Router) {
			r.Get("/", handlers.GetOffer)
			r.Delete("/", handlers.DeleteOffer)
		})
	})

	return r
}
