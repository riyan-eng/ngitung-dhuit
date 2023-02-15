package repository

import "github.com/valyala/fasthttp"

type TransactionRepository interface {
	InsertOne(*fasthttp.RequestCtx, string, float64) (string, error)
}
