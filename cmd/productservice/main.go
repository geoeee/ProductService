package main

import (
	"ProductService/internal/apis"
	"ProductService/openapi/gen/productservice/server"
	"ProductService/openapi/gen/productservice/server/operations"
	"flag"
	"log"

	"github.com/go-openapi/loads"
)

var portFlag = flag.Int("port", 8080, "Port to run this service on")

func main() {
	// load embedded swagger file
	swaggerSpec, err := loads.Analyzed(server.SwaggerJSON, "")
	if err != nil {
		log.Fatalln(err)
	}

	// create new service API
	api := operations.NewProductServiceAPI(swaggerSpec)
	s := server.NewServer(api)
	defer func() {
		_ = s.Shutdown()
	}()

	// parse flags
	flag.Parse()
	s.Port = *portFlag
	productAPI := apis.ProductAPI{}

	api.GetAPIV1ProductsHandler = operations.GetAPIV1ProductsHandlerFunc(productAPI.List)

	api.GetAPIV1ProductsProductIDHandler = operations.GetAPIV1ProductsProductIDHandlerFunc(productAPI.Get)

	api.PostAPIV1ProductsHandler = operations.PostAPIV1ProductsHandlerFunc(productAPI.Create)

	// serve API
	if err := s.Serve(); err != nil {
		log.Fatalln(err)
	}
}
