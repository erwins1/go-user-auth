package middleware

type Middleware struct {
	ByPassAuthEndpoint map[string]bool
}

func Init(m Middleware) *Middleware {
	return &Middleware{
		ByPassAuthEndpoint: m.ByPassAuthEndpoint,
	}
}
