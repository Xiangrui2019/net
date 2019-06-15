package web

import "net/http"

type Router struct {
	mux     *http.ServeMux
	pattern string
}

func NewRouter(mux *http.ServeMux) *Router {
	return &Router{mux: mux}
}

func (r *Router) join(pattern string) string {
	return r.pattern + pattern
}

func (r *Router) Group(pattern string) *Router {
	return &Router{mux: r.mux, pattern: r.join(pattern)}
}


func (r *Router) Handle(method, pattern string, handlers ...Handler) {
	r.mux.HandleFunc(r.join(pattern), func(w http.ResponseWriter, r *http.Request) {
		handler(method, w, r, handlers)
	})
}

func (r *Router) HandleFunc(method, pattern string, handlers ...HandleFunc) {
	r.mux.HandleFunc(r.join(pattern), func (w http.ResponseWriter, r *http.Request) {
		handleFunc(method, w, r, handlers)
	})
}

func (r *Router) GET(pattern string, handlers ...Handler) {
	r.mux.HandleFunc(r.join(pattern), func (w http.ResponseWriter, r *http.Request) {
		handler("GET", w, r, handlers)
	})
}

func (r *Router) POST(pattern string, handlers ...Handler) {
	r.mux.HandleFunc(r.join(pattern), func (w http.ResponseWriter, r *http.Request) {
		handler("POST", w, r, handlers)
	})
}

func (r *Router) PUT(pattern string, handlers ...Handler) {
	r.mux.HandleFunc(r.join(pattern), func (w http.ResponseWriter, r *http.Request) {
		handler("PUT", w, r, handlers)
	})
}

func (r *Router) DELETE(pattern string, handlers ...Handler) {
	r.mux.HandleFunc(r.join(pattern), func (w http.ResponseWriter, r *http.Request) {
		handler("DELETE", w, r, handlers)
	})
}

func (r *Router) GETServiceMethod(pattern string, handlers ...HandleFunc) {
	r.mux.HandleFunc(r.join(pattern), func(w http.ResponseWriter, r *http.Request) {
		handleFunc("GET", w, r, handlers)
	})
}

func (r *Router) POSTServiceMethod(pattern string, handlers ...HandleFunc) {
	r.mux.HandleFunc(r.join(pattern), func(w http.ResponseWriter, r *http.Request) {
		handleFunc("POST", w, r, handlers)
	})
}

func (r *Router) PUTServiceMethod(pattern string, handlers ...HandleFunc) {
	r.mux.HandleFunc(r.join(pattern), func (w http.ResponseWriter, r *http.Request) {
		handleFunc("PUT", w, r, handlers)
	})
}

func handler(method string, w http.ResponseWriter, r *http.Request, handlers []Handler) {
	if r.Method != method {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	c := new(Context)
	c.Request = r
	c.Response = w
	for _, h := range handlers {
		h.ServeHTTP(c)
		if c.done {
			break
		}
	}
}

func handleFunc(method string, w http.ResponseWriter, r *http.Request, handlers []HandleFunc) {
	if r.Method != method {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	c := new(Context)
	c.Request = r
	c.Response = w
	for _, h := range handlers {
		h(c)
		if c.done {
			break
		}
	}
}
