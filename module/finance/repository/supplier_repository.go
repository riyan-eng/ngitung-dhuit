package repository

import "github.com/valyala/fasthttp"

type SupplierRepository interface {
	FindOne(*fasthttp.RequestCtx, string) error
}
