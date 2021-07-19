package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/akrylysov/pogreb"
	"github.com/valyala/fasthttp"
)

var db *pogreb.DB

func main() {

	var err error
	db, err = pogreb.Open("pogreb.test", nil)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer db.Close()

	requestHandler := func(ctx *fasthttp.RequestCtx) {
		switch string(ctx.Path()) {
		case "/register":
			registerHandler(ctx)
		case "/list":
			listHandler(ctx)
		case "/":
			homeHandler(ctx)
		default:
			ctx.Error("Unsupported path", fasthttp.StatusNotFound)
		}
	}

	fasthttp.ListenAndServe(":80", requestHandler)
}

func registerHandler(ctx *fasthttp.RequestCtx) {
	switch string(ctx.Method()) {
	case "GET":
		fmt.Fprintf(ctx, "Hello, Register!\n\n")
	case "POST":
		k := strconv.FormatUint(ctx.ID(), 10)

		err := db.Put([]byte(k), ctx.FormValue("data"))
		if err != nil {
			ctx.Error("Error can't put data to db", fasthttp.StatusInternalServerError)
		}
	default:
		fmt.Fprintf(ctx, "Request method is %q\n", ctx.Method())
	}

}

func homeHandler(ctx *fasthttp.RequestCtx) {
	fmt.Fprintf(ctx, "Hello, HOME!\n\n")
}

func listHandler(ctx *fasthttp.RequestCtx) {
	it := db.Items()
	for {
		key, val, err := it.Next()
		if err == pogreb.ErrIterationDone {
			break
		}
		if err != nil {
			ctx.Error("Error can't get data from db", fasthttp.StatusInternalServerError)

		}
		fmt.Printf("[key]=%s  [data]=%s\n", key, val)

	}
}
