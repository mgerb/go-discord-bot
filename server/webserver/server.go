package webserver

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/crypto/acme/autocert"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/mgerb/go-discord-bot/server/config"
	"github.com/mgerb/go-discord-bot/server/webserver/handlers"
)

func getRouter() *chi.Mux {

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.DefaultCompress)

	if !config.Flags.Prod {
		r.Use(middleware.Logger)
	}

	workDir, _ := os.Getwd()

	FileServer(r, "/static", http.Dir(filepath.Join(workDir, "./dist/static")))
	FileServer(r, "/public/sounds", http.Dir(filepath.Join(workDir, "./sounds")))
	FileServer(r, "/public/youtube", http.Dir(filepath.Join(workDir, "./youtube")))

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./dist/index.html")
	})

	// configure end points
	r.Get("/soundlist", handlers.SoundList)
	r.Put("/upload", handlers.FileUpload)
	r.Get("/ytdownloader", handlers.Downloader)

	return r
}

// Start -
func Start() {
	router := getRouter()

	if config.Flags.TLS {

		// start server on port 80 to redirect
		go http.ListenAndServe(":80", http.HandlerFunc(redirect))

		// start TLS server
		log.Fatal(http.Serve(autocert.NewListener(), router))

	} else {

		// start basic server
		http.ListenAndServe(config.Config.ServerAddr, router)
	}
}

// FileServer conveniently sets up a http.FileServer handler to serve
// static files from a http.FileSystem.
func FileServer(r chi.Router, path string, root http.FileSystem) {
	if strings.ContainsAny(path, "{}*") {
		panic("FileServer does not permit URL parameters.")
	}

	fs := http.StripPrefix(path, http.FileServer(root))

	if path != "/" && path[len(path)-1] != '/' {
		r.Get(path, http.RedirectHandler(path+"/", 301).ServeHTTP)
		path += "/"
	}
	path += "*"

	r.Get(path, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fs.ServeHTTP(w, r)
	}))
}

// redirect to https
func redirect(w http.ResponseWriter, req *http.Request) {
	// remove/add not default ports from req.Host
	target := "https://" + req.Host + req.URL.Path
	if len(req.URL.RawQuery) > 0 {
		target += "?" + req.URL.RawQuery
	}

	http.Redirect(w, req, target, http.StatusTemporaryRedirect)
}
