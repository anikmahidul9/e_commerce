package middlewares

import "net/http"

type Middleware func(http.Handler) http.Handler

type Manager struct {
	globalMiddlewares []Middleware
}

func NewManager() *Manager {
	return &Manager{
		globalMiddlewares: []Middleware{},
	}
}

func (m *Manager) Use(mw ...Middleware) {
	m.globalMiddlewares = append(m.globalMiddlewares, mw...)
}

func (mnger *Manager) With(handler http.Handler, middlewares ...Middleware) http.Handler {
	h := handler
	for _, middleware := range middlewares {
		h = middleware(h)
	}
	return h
}

func (m *Manager) WrapMux(handler http.Handler) http.Handler {

	h := handler //h = logger(hudai(corsWithPreflight()))
	for _, middleware := range m.globalMiddlewares {
		h = middleware(h)
	}

	return h
}