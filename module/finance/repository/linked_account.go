package repository

import "github.com/valyala/fasthttp"

type LinkedAccountRepository interface {
	FindOneByCode(*fasthttp.RequestCtx, string) (string, error)
}
