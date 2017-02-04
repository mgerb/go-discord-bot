package webserver

import (
	"../config"
	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
	"io"
	"log"
	"os"
)

func logger(next fasthttp.RequestHandler) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		logger := ctx.Logger()
		logger.Printf(ctx.RemoteAddr().String())
		next(ctx)
	}
}

func applyMiddleware(handler fasthttp.RequestHandler) fasthttp.RequestHandler {
	newHandler := logger(handler)

	return newHandler
}

func registerRoutes(router *fasthttprouter.Router) {

	router.PUT("/upload", fileUpload)

	router.ServeFiles("/static/*filepath", "./static")

	router.NotFound = func(ctx *fasthttp.RequestCtx) {
		fasthttp.ServeFile(ctx, "./index.html")
	}
}

func Start() {
	router := fasthttprouter.New()

	registerRoutes(router)

	// apply our middleware
	handlers := applyMiddleware(router.Handler)

	// start web server
	log.Fatal(fasthttp.ListenAndServe(config.Config.ServerAddr, handlers))
}

func fileUpload(ctx *fasthttp.RequestCtx) {
	password := ctx.FormValue("password")

	if string(password) != config.Config.UploadPassword {
		ctx.Error("Invalid password.", 400)
		return
	}

	file, err := ctx.FormFile("file")
	if err != nil {
		ctx.Error("Error reading file.", 400)
		return
	}

	src, err := file.Open()
	if err != nil {
		ctx.Error("Error opening file.", 400)
		return
	}

	defer src.Close()

	// create uploads folder if it does not exist
	if _, err := os.Stat(config.Config.SoundsPath); os.IsNotExist(err) {
		os.Mkdir(config.Config.SoundsPath, os.ModePerm)
	}

	// check if file already exists
	if _, err := os.Stat(config.Config.SoundsPath + file.Filename); err == nil {
		ctx.Error("File already exists.", 400)
		return
	}

	dst, err := os.Create(config.Config.SoundsPath + file.Filename)
	if err != nil {
		ctx.Error("Error creating file.", 400)
		return
	}

	defer dst.Close()

	if _, err = io.Copy(dst, src); err != nil {
		ctx.Error("Error writing file.", 400)
		return
	}

	ctx.Success("application/json", []byte("Success!"))
}
