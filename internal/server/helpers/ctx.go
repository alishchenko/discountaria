package helpers

import (
	"context"
	"github.com/alishchenko/discountaria/internal/config"
	"github.com/alishchenko/discountaria/internal/data/postgres"
	"golang.org/x/oauth2"
	"log/slog"
	"net/http"
)

type ctxKey int

const (
	logCtxKey                  ctxKey = iota
	dbCtxKey                   ctxKey = iota
	tokensCtxKey               ctxKey = iota
	oAuth2FacebookConfigCtxKey ctxKey = iota
	oAuth2StateConfigCtxKey    ctxKey = iota
	oAuth2GoogleConfigCtxKey   ctxKey = iota
	oAuth2LinkedinConfigCtxKey ctxKey = iota
	mimeTypesCtxKey            ctxKey = iota
	awsCfgKey                  ctxKey = iota
	oSignatureConfigCtxKey     ctxKey = iota
)

func CtxLog(entry *slog.Logger) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, logCtxKey, entry)
	}
}

func Log(r *http.Request) *slog.Logger {
	return r.Context().Value(logCtxKey).(*slog.Logger)
}

func CtxDB(entry *postgres.DB) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, dbCtxKey, entry)
	}
}

func DB(r *http.Request) *postgres.DB {
	return r.Context().Value(dbCtxKey).(*postgres.DB)
}

func CtxTokens(entry config.Tokens) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, tokensCtxKey, entry)
	}
}

func Tokens(r *http.Request) config.Tokens {
	return r.Context().Value(tokensCtxKey).(config.Tokens)
}

func SetOAuth2FacebookConfig(cfg *oauth2.Config) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, oAuth2FacebookConfigCtxKey, cfg)
	}
}

func OAuth2FacebookConfig(r *http.Request) *oauth2.Config {
	return r.Context().Value(oAuth2FacebookConfigCtxKey).(*oauth2.Config)
}

func SetOAuth2StateConfig(cfg config.OAuth2StateConfig) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, oAuth2StateConfigCtxKey, cfg)
	}
}

func OAuth2StateConfig(r *http.Request) config.OAuth2StateConfig {
	return r.Context().Value(oAuth2StateConfigCtxKey).(config.OAuth2StateConfig)
}

func SetOAuth2GoogleConfig(cfg *oauth2.Config) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, oAuth2GoogleConfigCtxKey, cfg)
	}
}

func OAuth2GoogleConfig(r *http.Request) *oauth2.Config {
	return r.Context().Value(oAuth2GoogleConfigCtxKey).(*oauth2.Config)
}

func SetOAuth2LinkedinConfig(cfg *oauth2.Config) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, oAuth2LinkedinConfigCtxKey, cfg)
	}
}

func OAuth2LinkedinConfig(r *http.Request) *oauth2.Config {
	return r.Context().Value(oAuth2LinkedinConfigCtxKey).(*oauth2.Config)
}

func SetSignatureConfig(cfg config.Signature) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, oSignatureConfigCtxKey, cfg)
	}
}

func SignatureConfig(r *http.Request) config.Signature {
	return r.Context().Value(oSignatureConfigCtxKey).(config.Signature)
}

func CtxMimeTypes(entry *config.MimeTypes) func(ctx context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, mimeTypesCtxKey, entry)
	}
}

func CtxAwsConfig(entry *config.AWSConfig) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, awsCfgKey, entry)
	}
}

func MimeTypes(r *http.Request) *config.MimeTypes {
	return r.Context().Value(mimeTypesCtxKey).(*config.MimeTypes)
}

func AwsConfig(r *http.Request) *config.AWSConfig {
	return r.Context().Value(awsCfgKey).(*config.AWSConfig)
}
