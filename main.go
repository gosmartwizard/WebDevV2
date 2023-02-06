package main

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"

	"github.com/go-chi/chi"
	"github.com/gosmartwizard/WebDevV2/controllers"
	"github.com/gosmartwizard/WebDevV2/views"
)

func executeTemplate(w http.ResponseWriter, filepath string) {

	tpl, err := views.Parse(filepath)
	if err != nil {
		log.Printf("parsing template: %v", err)
		http.Error(w, "There was an error parsing the template.", http.StatusInternalServerError)
		return
	}

	tpl.Execute(w, nil)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {

	tplPath := filepath.Join("templates", "home.gohtml")

	executeTemplate(w, tplPath)
}

func contactHandler(w http.ResponseWriter, r *http.Request) {

	tplPath := filepath.Join("templates", "contact.gohtml")

	executeTemplate(w, tplPath)
}

func faqHandler(w http.ResponseWriter, r *http.Request) {

	tplPath := filepath.Join("templates", "faq.gohtml")

	executeTemplate(w, tplPath)
}

func paramsHandler(w http.ResponseWriter, r *http.Request) {

	galleryID := chi.URLParam(r, "galleryID")

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	fmt.Fprint(w, "Got the galleryID : ", galleryID)
}

func main() {

	r := chi.NewRouter()
	/* r.Use(middleware.Logger)
	r.Get("/", homeHandler)
	r.Get("/contact", contactHandler)
	r.Get("/faq", faqHandler)

	r.With(middleware.Logger).Get("/galleries/{galleryID}", paramsHandler)
	*/

	tpl, err := views.Parse(filepath.Join("templates", "home.gohtml"))
	views.Must(tpl, err)
	r.Get("/", controllers.StaticHandler(tpl))

	tpl, err = views.Parse(filepath.Join("templates", "contact.gohtml"))
	views.Must(tpl, err)
	r.Get("/contact", controllers.StaticHandler(tpl))

	tpl, err = views.Parse(filepath.Join("templates", "faq.gohtml"))
	views.Must(tpl, err)
	r.Get("/faq", controllers.StaticHandler(tpl))

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not found", http.StatusNotFound)
	})

	fmt.Println("Starting the server on :3000...")

	http.ListenAndServe(":3000", r)
}
