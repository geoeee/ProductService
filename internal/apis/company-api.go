package apis

import (
	"CompanyService/internal/resources"
	genModels "CompanyService/openapi/gen/companyservice/models"
	"CompanyService/openapi/gen/companyservice/server/operations"
	prodOps "CompanyService/openapi/gen/productservice/client/operations"
	"CompanyService/pkg/version"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/go-openapi/runtime/middleware"
	uuid "github.com/satori/go.uuid"
)

// CompanyAPI ...
type CompanyAPI struct {
}

var companies map[string]*genModels.Company

func init() {
	companies = make(map[string]*genModels.Company)
	c1 := &genModels.Company{
		CompanyID:   "86cf1699-00d3-494a-a49f-ec5230b0fadd",
		CompanyName: "c1",
	}
	companies[c1.CompanyID] = c1
	c2 := &genModels.Company{
		CompanyID:   "39266573-0f22-47fc-bac1-39ea8a351c72",
		CompanyName: "c2",
	}
	companies[c2.CompanyID] = c2
	c3 := &genModels.Company{
		CompanyID:   "3718a69f-e5f7-4395-9f61-1cc5680f5396",
		CompanyName: "c3",
	}
	companies[c3.CompanyID] = c3
}

func toList(cm map[string]*genModels.Company, req *http.Request) []*genModels.Company {
	var cl []*genModels.Company

	meta := struct {
		PodName string
		Version string
		Time    string
		Headers http.Header
	}{
		PodName: os.Getenv("HOSTNAME"),
		Version: version.Version,
		Time:    time.Now().String(),
		Headers: req.Header,
	}

	prodSvc := resources.NewProductServiceResource()
	for _, c := range cm {
		c.Meta = meta

		params := &prodOps.GetAPIV1ProductsParams{
			CompanyID: &c.CompanyID,
			Context:   req.Context(),
		}
		ps, err := prodSvc.GetProducts(params)
		if err != nil {
			fmt.Println("get product: ", err)
		} else {
			for _, p := range ps {
				c.Products = append(c.Products, p)
			}
		}
		cl = append(cl, c)
	}
	return cl
}

// Get ...
func (api *CompanyAPI) Get(params operations.GetAPIV1CompaniesCompanyIDParams) middleware.Responder {

	cID := params.CompanyID
	resp := companies[cID]
	meta := &struct {
		PodName string
		Version string
		Time    string
		Headers http.Header
	}{
		PodName: os.Getenv("HOSTNAME"),
		Version: version.Version,
		Time:    time.Now().String(),
		Headers: params.HTTPRequest.Header,
	}
	resp.Meta = meta
	prodSvc := resources.NewProductServiceResource()
	pParam := &prodOps.GetAPIV1ProductsParams{
		CompanyID: &cID,
		Context:   params.HTTPRequest.Context(),
	}
	ps, err := prodSvc.GetProducts(pParam)
	if err != nil {
		fmt.Println("get product: ", err)
	} else {
		for _, p := range ps {
			resp.Products = append(resp.Products, p)
		}
	}
	return operations.NewGetAPIV1CompaniesCompanyIDOK().WithPayload(resp)
}

// List ...
func (api *CompanyAPI) List(params operations.GetAPIV1CompaniesParams) middleware.Responder {

	resp := &genModels.PageCompanies{}
	resp.Limit = "0"
	resp.Offset = "0"
	resp.Count = fmt.Sprintf("%d", len(companies))
	resp.Elements = toList(companies, params.HTTPRequest)

	return operations.NewGetAPIV1CompaniesOK().WithPayload(resp)
}

// Create ...
func (api *CompanyAPI) Create(params operations.PostAPIV1CompaniesParams) middleware.Responder {
	resp := &genModels.Company{}

	resp.CompanyID = uuid.Must(uuid.NewV4()).String()
	comName := params.Body.CompanyName
	resp.CompanyName = *comName
	companies[resp.CompanyID] = resp
	meta := struct {
		PodName string
		Version string
		Time    string
		Headers http.Header
	}{
		PodName: os.Getenv("HOSTNAME"),
		Version: version.Version,
		Time:    time.Now().String(),
		Headers: params.HTTPRequest.Header,
	}
	resp.Meta = meta

	return operations.NewPostAPIV1CompaniesCreated().WithPayload(resp)
}
