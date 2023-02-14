package repository

import "github.com/valyala/fasthttp"

type TaxRepository interface {
	GetByCoa(*fasthttp.RequestCtx, string) (int, error)
}
