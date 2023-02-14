package repository

import "github.com/valyala/fasthttp"

type LinkedAccountRepository interface {
	GetByCode(*fasthttp.RequestCtx, string) (string, error)
}
