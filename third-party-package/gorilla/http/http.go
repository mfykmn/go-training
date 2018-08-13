package http

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

type httpTime struct {
	time.Time
}

func (h *httpTime) String() string {
	return h.Format(time.RFC3339)
}

// New はドメインのサービスを追加したServerの構造体を返す関数。
func New(addr string, handlers *Handlers) *Server {
	return &Server{
		Server: http.Server{
			Addr: addr,
		},
		Handlers: *handlers,
	}
}

// Server はHTTPのサーバそのものを表す構造体。
type Server struct {
	http.Server
	Handlers
}

// Handlers は各種サービスを統括して管理する構造体。
type Handlers struct {
}

// ListenAndServe はServerを起動するメソッド。
func (s *Server) ListenAndServe() error {
	router := route(s)
	router = setMiddleware(router)

	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "HEAD"})
	credentialOk := handlers.AllowCredentials()
	s.Handler = handlers.CORS(originsOk, methodsOk, credentialOk)(router)
	return s.Server.ListenAndServe()
}

const (
	// Path
	usersPath = "/users"
)

// routeはエンドポイントと関数のルーティングをする関数。
func route(_ *Server) http.Handler {
	router := mux.NewRouter()

	// custom not found
	router.NotFoundHandler = http.HandlerFunc(customNotFound)

	// Health Check
	router.Path("/health").HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) { w.WriteHeader(http.StatusOK) },
	).Methods(http.MethodGet)

	v1 := router.PathPrefix("/v1").Handler(negroni.New(
		negroni.NewRecovery(),
		negroni.NewLogger(),
	)).Subrouter()

	// Rooting
	v1.Path(usersPath).HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"result":"ok"}`))
	}).Methods(http.MethodGet)
	return router
}

// customNotFound は404NotFoundのエラー内容をセットする関数。
func customNotFound(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte(`{"reason":"not found"}`))
}

// setMiddleware はMiddlewareをセットする関数。
func setMiddleware(handler http.Handler) http.Handler {
	n := negroni.New()
	n.UseHandler(handler)
	n.UseFunc(AuthMiddleware)
	return n
}

// AuthMiddleware は認証の前処理を行う関数。
func AuthMiddleware(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	// Authの処理
	next(rw, r)
}

func setStatusInternalServerResponse(w http.ResponseWriter, reason string) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(fmt.Sprintf(`{"reason":"%s"}`, reason)))
}
