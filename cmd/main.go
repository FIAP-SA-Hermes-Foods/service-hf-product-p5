package main

import (
	"context"
	"log"
	"net/http"
	"os"
	l "service-hf-product-p5/external/logger"
	productrpc "service-hf-product-p5/internal/adapters/rpc"
	"service-hf-product-p5/internal/core/application"
	httpH "service-hf-product-p5/internal/handler/http"

	"github.com/marcos-dev88/genv"
)

func init() {
	if err := genv.New(); err != nil {
		l.Errorf("", "error set envs %v", " | ", err)
	}
}

func main() {

	router := http.NewServeMux()

	ctx := context.Background()

	productRPC := productrpc.NewProductRPC(ctx, os.Getenv("HOST_PRODUCT"), os.Getenv("PORT_PRODUCT"))

	productWorkerRPC := productrpc.NewProductWorkerRPC(ctx, os.Getenv("HOST_PRODUCT"), os.Getenv("PORT_PRODUCT"))

	app := application.NewApplication(ctx, productRPC, productWorkerRPC)

	h := httpH.NewHandler(app)

	router.Handle("/hermes_foods/product/", http.StripPrefix("/", httpH.Middleware(h.HandlerProduct)))
	router.Handle("/hermes_foods/product", http.StripPrefix("/", httpH.Middleware(h.HandlerProduct)))

	log.Fatal(http.ListenAndServe(":"+os.Getenv("API_HTTP_PORT"), router))
}
