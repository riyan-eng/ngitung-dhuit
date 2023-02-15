package repository

import "github.com/valyala/fasthttp"

type COARepository interface {
	FindOneByCode(*fasthttp.RequestCtx, string) error
}
