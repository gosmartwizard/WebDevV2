package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/gosmartwizard/WebDevV2/controllers"
	"github.com/gosmartwizard/WebDevV2/models"
	"github.com/gosmartwizard/WebDevV2/templates"
	"github.com/gosmartwizard/WebDevV2/views"
)

func main() {

	// Setup a database connection
	cfg := models.DefaultPostgresConfig()
	db, err := models.Open(cfg)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Setup our model services
	userService := models.UserService{
		DB: db,
	}

	// Setup our controllers
	usersC := controllers.Users{
		UserService: &userService,
	}

	r := chi.NewRouter()

	tpl, err := views.ParseFS(templates.FS, "home.gohtml", "tailwind.gohtml")
	views.Must(tpl, err)
	r.Get("/", controllers.StaticHandler(tpl))

	tpl, err = views.ParseFS(templates.FS, "contact.gohtml", "tailwind.gohtml")
	views.Must(tpl, err)
	r.Get("/contact", controllers.StaticHandler(tpl))

	tpl, err = views.ParseFS(templates.FS, "faq.gohtml", "tailwind.gohtml")
	views.Must(tpl, err)
	r.Get("/faq", controllers.FAQ(tpl))

	tpl, err = views.ParseFS(templates.FS, "signup.gohtml", "tailwind.gohtml")
	views.Must(tpl, err)
	usersC.Templates.New = tpl
	r.Get("/signup", usersC.New)
	r.Post("/signup", usersC.Create)

	tpl, err = views.ParseFS(templates.FS, "signin.gohtml", "tailwind.gohtml")
	views.Must(tpl, err)
	usersC.Templates.SignIn = tpl
	r.Get("/signin", usersC.SignIn)
	r.Post("/signin", usersC.ProcessSignIn)

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not found", http.StatusNotFound)
	})

	fmt.Println("Starting the server on :3000...")

	http.ListenAndServe(":3000", r)
}
