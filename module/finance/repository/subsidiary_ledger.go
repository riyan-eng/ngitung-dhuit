package repository

import "github.com/valyala/fasthttp"

type SubsidiaryLedgerRepository interface {
	InsertPayable(*fasthttp.RequestCtx, string, string, float64) error
}
