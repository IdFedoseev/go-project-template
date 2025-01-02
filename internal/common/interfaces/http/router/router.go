package router

import (
	"net/http"
	middleware2 "proj/internal/common/interfaces/http/middleware"
	"strings"
)

// Group представляет группу маршрутов с общим middleware
type Group struct {
	mux        *http.ServeMux
	prefix     string
	middleware func(http.HandlerFunc) http.HandlerFunc
}

// HandleFunc регистрирует обработчик с учетом middleware
func (g *Group) HandleFunc(pattern string, handler http.HandlerFunc) {
	method, path, _ := strings.Cut(pattern, " ")
	g.mux.HandleFunc(method+" "+g.prefix+path, g.middleware(handler))
}

type RouterRegistrar interface {
	RegisterRoutes(group *Group)
}

func NewRouter(registrars ...RouterRegistrar) http.Handler {
	mux := http.NewServeMux()

	api := &Group{
		mux:    mux,
		prefix: "/api",
		middleware: middleware2.Chain(
			middleware2.Trace(),
			middleware2.RequestLogger(),
			middleware2.Auth(),
		),
	}

	for _, r := range registrars {
		r.RegisterRoutes(api)
	}

	return mux
}
