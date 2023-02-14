package repository

import "github.com/valyala/fasthttp"

type COARepository interface {
	GetByCode(*fasthttp.RequestCtx, string) error
}
