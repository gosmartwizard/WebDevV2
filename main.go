package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/gosmartwizard/WebDevV2/controllers"
	"github.com/gosmartwizard/WebDevV2/templates"
	"github.com/gosmartwizard/WebDevV2/views"
)

func main() {

	r := chi.NewRouter()

	tpl, err := views.ParseFS(templates.FS, "layout-page.gohtml", "home-page.gohtml")
	views.Must(tpl, err)
	r.Get("/", controllers.StaticHandler(tpl))

	tpl, err = views.ParseFS(templates.FS, "layout-page.gohtml", "contact-page.gohtml")
	views.Must(tpl, err)
	r.Get("/contact", controllers.StaticHandler(tpl))

	tpl, err = views.ParseFS(templates.FS, "faq.gohtml")
	views.Must(tpl, err)
	r.Get("/faq", controllers.FAQ(tpl))

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not found", http.StatusNotFound)
	})

	fmt.Println("Starting the server on :3000...")

	http.ListenAndServe(":3000", r)
}
