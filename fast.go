package fast

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/a-h/templ"
	"github.com/julienschmidt/httprouter"
)

func defaultErrorHandler(err error, c *Context) error {
	slog.Error("error", "err", err)
	return nil
}

type ErrorHandler func(error, *Context) error

type Context struct {
	response http.ResponseWriter
	request  *http.Request
	ctx      context.Context
}

func (c *Context) Render(comp templ.Component) error {
	return comp.Render(c.ctx, c.response)
}

type Handler func(c *Context) error

type Fast struct {
	ErrorHandler ErrorHandler
	router       *httprouter.Router
}

func New() *Fast {
	return &Fast{
		ErrorHandler: defaultErrorHandler,
		router:       httprouter.New(),
	}
}

func (f *Fast) Start(port string) error {
	return http.ListenAndServe(port, f.router)
}

func (f *Fast) Get(path string, h Handler, plug ...Handler) {
	f.router.GET(path, f.makeHTTPRouterHandler(h))
}

func (f *Fast) makeHTTPRouterHandler(h Handler) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		ctx := &Context{
			response: w,
			request:  r,
			ctx:      context.Background(),
		}
		if err := h(ctx); err != nil {
			f.ErrorHandler(err, ctx)
		}
	}
}
