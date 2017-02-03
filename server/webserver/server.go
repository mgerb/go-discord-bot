package webserver

import (
	"fmt"
	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
	"io"
	"log"
	"os"
	"os/exec"
	"strings"
	"sync"
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
	//newHandler = fasthttp.CompressHandler(newHandler)

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
	log.Fatal(fasthttp.ListenAndServe("0.0.0.0:8080", handlers))
}

func fileUpload(ctx *fasthttp.RequestCtx) {

	file, err := ctx.FormFile("file")
	if err != nil {
		ctx.SetStatusCode(500)
		fmt.Fprint(ctx, "Error reading file.")
		return
	}

	src, err := file.Open()
	if err != nil {
		ctx.SetStatusCode(500)
		fmt.Fprint(ctx, "Error opening file.")
		return
	}

	defer src.Close()

	// create uploads folder if it does not exist
	if _, err := os.Stat("./uploads"); os.IsNotExist(err) {
		os.Mkdir("./uploads", os.ModePerm)
	}

	// check if file already exists
	if _, err := os.Stat("./uploads/" + file.Filename); err == nil {
		ctx.SetStatusCode(403)
		fmt.Fprint(ctx, "File already exists.")
		return
	}

	dst, err := os.Create("./uploads/" + file.Filename)
	if err != nil {
		ctx.SetStatusCode(500)
		fmt.Fprint(ctx, "Error writing file.")
		return
	}

	defer dst.Close()

	if _, err = io.Copy(dst, src); err != nil {
		ctx.SetStatusCode(500)
		fmt.Fprint(ctx, "Error writing file.")
		return
	}

	err = convertFile(file.Filename)
	if err != nil {
		ctx.SetStatusCode(500)
		fmt.Fprint(ctx, err.Error())
		return
	}

	fmt.Fprint(ctx, "Success")
}

func convertFile(fileName string) error {
	trimmedName := strings.Split(fileName, ".")[0]
	outputPath := "./sounds/"
	inputPath := "./uploads/"

	fmt.Println(inputPath + fileName)
	fmt.Println(outputPath + trimmedName + ".dca")

	command := "./dca-rs " + "--raw " + "--i " + inputPath + fileName + " > " + outputPath + trimmedName + ".dca"

	wg := new(sync.WaitGroup)
	commands := []string{command}
	for _, str := range commands {
		wg.Add(1)
		go exe_cmd(str, wg)
	}
	wg.Wait()

	/*
		cmd := exec.Command("./dca-rs", "--raw", "--i", inputPath+fileName+" > "+outputPath+trimmedName+".dca")

		_, err := cmd.Output()

		if err != nil {
			fmt.Println(err.Error())
		}

		err = cmd.Wait()
		if err == nil {
			fmt.Println(err.Error())
		}
	*/

	/*
		outFile, err := os.Create(outputPath + trimmedName + ".dca")
		if err != nil {
			return err
		}
		defer outFile.Close()

		stdoutPipe, err := cmd.StdoutPipe()
		if err != nil {
			return err
		}

		writer := bufio.NewWriter(outFile)
		defer writer.Flush()

		fmt.Println(stdoutPipe)
		go io.Copy(writer, stdoutPipe)
	*/

	fmt.Println("working")
	return nil
}

func exe_cmd(cmd string, wg *sync.WaitGroup) {
	fmt.Println(cmd)
	parts := strings.Fields(cmd)
	out, err := exec.Command(parts[0], parts[1]).Output()
	if err != nil {
		fmt.Println("error occured")
		fmt.Printf("%s", err)
	}
	fmt.Printf("%s", out)
	wg.Done()
}
