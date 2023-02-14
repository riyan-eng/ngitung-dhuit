package repository

import "github.com/valyala/fasthttp"

type TransactionRepository interface {
	Insert(*fasthttp.RequestCtx, string, float64) (string, error)
}
