package repository

import "github.com/valyala/fasthttp"

type TaxRepository interface {
	FindOneByCoa(*fasthttp.RequestCtx, string) (int, error)
}
