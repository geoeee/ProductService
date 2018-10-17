package main

import (
	"CompanyService/internal/apis"
	"CompanyService/openapi/gen/companyservice/server"
	"CompanyService/openapi/gen/companyservice/server/operations"
	"flag"
	"log"

	"github.com/go-openapi/loads"
)

var portFlag = flag.Int("port", 8081, "Port to run this service on")

func main() {
	// load embedded swagger file
	swaggerSpec, err := loads.Analyzed(server.SwaggerJSON, "")
	if err != nil {
		log.Fatalln(err)
	}

	// create new service API
	api := operations.NewCompanyServiceAPI(swaggerSpec)
	s := server.NewServer(api)
	defer func() {
		_ = s.Shutdown()
	}()

	// parse flags
	flag.Parse()
	s.Port = *portFlag
	comAPI := apis.CompanyAPI{}

	api.GetAPIV1CompaniesHandler = operations.GetAPIV1CompaniesHandlerFunc(comAPI.List)
	api.GetAPIV1CompaniesCompanyIDHandler = operations.GetAPIV1CompaniesCompanyIDHandlerFunc(comAPI.Get)
	api.PostAPIV1CompaniesHandler = operations.PostAPIV1CompaniesHandlerFunc(comAPI.Create)
	// serve API
	if err := s.Serve(); err != nil {
		log.Fatalln(err)
	}
}
