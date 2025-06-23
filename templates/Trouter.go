package templates

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/ptdrpg/chi/init/scrypt/handler"
)

func WritteRouter(directory string) {
	r := fmt.Sprintf(`
package router

import (
	"net/http"
	"%s/controller"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

type Router struct {
	R *chi.Mux
	C *controller.Controller
}

func NewRouter(c *controller.Controller) *Router {
	r := chi.NewRouter()

	// Middleware CORS
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	}))

	return &Router{
		R: r,
		C: c,
	}
}

func (r *Router) RegisterRouter() {
	// r.R.Group(func(public chi.Router) {
	// 	public.Route("/api/v1/login", func(login chi.Router) {
	// 		login.Post("/", r.C.Login)
	// 	})
	// })

	// r.R.Group(func(private chi.Router) {
	// 	// private.Use(lib.JWTMiddleware)
	// 	private.Route("/api/v1", func(v1 chi.Router) {
	// 		v1.Route("/users", func(users chi.Router) {
	// 			users.Get("/", r.C.GetAllUsers)
	// 			users.Post("/newVisitor", r.C.CreateUser)
	// 			users.Put("/reabilite/{id}", r.C.ActualiseUser)
	// 			users.Delete("/delete/{id}", r.C.DeleteUser)
	// 		})
	// 		v1.Route("/movies", func(movies chi.Router) {
	// 			movies.Get("/list", r.C.MoviesList)
	// 			movies.Get("/stream", r.C.Stream)
	// 		})
	// 	})
	// })
}

func (r *Router) Handler() http.Handler {
	return r.R
}
	
	`, directory)

	routerPath := filepath.Join(directory, "router", "router.go")
	router, err := os.Create(routerPath)
	if err != nil {
		handler.ErrorHandler(err)
	}
	router.Write([]byte(r))
	router.Close()
}
