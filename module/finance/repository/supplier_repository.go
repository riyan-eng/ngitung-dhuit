package repository

import "github.com/valyala/fasthttp"

type SupplierRepository interface {
	FindOneByCode(*fasthttp.RequestCtx, string) error
}
