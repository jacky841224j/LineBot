package route

// import (
// 	"net/http"

// 	"github.com/gorilla/mux"
// )

// type Route struct {
// 	Method     string
// 	Pattern    string
// 	Handle     http.HandlerFunc
// 	Middleware mux.MiddlewareFunc
// }

// var routes []Route

// func init() {
// 	register("POST", "/api/todo", controller.AddTodo, nil)
// }

// func NewRouter() *mux.Router {
// 	r := mux.NewRouter()
// 	for _, route := range routes {
// 		r.Methods(route.Method).Path(route.Pattern).Handler(route.Handle)

// 		if route.Middleware != nil {
// 			r.Use(route.Middleware)
// 		}
// 	}

// 	return r
// }

// func register(method, pattern string, handler http.HandlerFunc, middleware mux.MiddlewareFunc) {
// 	routes = append(routes, Route{method, pattern, handler, middleware})
// }
