package repository

import "github.com/valyala/fasthttp"

type SubsidiaryLedgerRepository interface {
	InsertOnePayable(*fasthttp.RequestCtx, string, string, float64) error
}
