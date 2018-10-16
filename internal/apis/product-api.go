package apis

import (
	genModels "ProductService/openapi/gen/productservice/models"
	"ProductService/openapi/gen/productservice/server/operations"
	"ProductService/pkg/version"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/go-openapi/runtime/middleware"
	uuid "github.com/satori/go.uuid"
)

// ProductAPI ...
type ProductAPI struct {
}

var products map[string]*genModels.Product

func init() {
	products = make(map[string]*genModels.Product)
	p1 := &genModels.Product{
		ProductID:   "31577e70-9d4c-441a-9210-b3e5471af3ff",
		CompanyID:   "1",
		ProductName: "p1",
	}
	products[p1.ProductID] = p1
	p2 := &genModels.Product{
		ProductID:   "190e8697-15cd-4590-a304-99303d7a7cbf",
		CompanyID:   "2",
		ProductName: "p2",
	}
	products[p2.ProductID] = p2
	p3 := &genModels.Product{
		ProductID:   "b5863fd4-8eea-4c94-bdb4-ce8ed3fbcdc7",
		CompanyID:   "3",
		ProductName: "p3",
	}
	products[p3.ProductID] = p3

}
func toList(pm map[string]*genModels.Product, comID *string, req *http.Request) []*genModels.Product {
	var pl []*genModels.Product

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

	if comID != nil {
		for _, p := range pm {
			if *comID == p.CompanyID {
				p.Meta = meta
				pl = append(pl, p)
			}
		}
		return pl
	}
	for _, p := range pm {
		p.Meta = meta
		pl = append(pl, p)
	}
	return pl
}

// Get ...
func (api *ProductAPI) Get(params operations.GetAPIV1ProductsProductIDParams) middleware.Responder {

	pID := params.ProductID
	resp := products[pID]
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
	return operations.NewGetAPIV1ProductsProductIDOK().WithPayload(resp)
}

// List ...
func (api *ProductAPI) List(params operations.GetAPIV1ProductsParams) middleware.Responder {

	comID := params.CompanyID
	resp := &genModels.PageProducts{}
	resp.Limit = "0"
	resp.Offset = "0"
	resp.Count = fmt.Sprintf("%d", len(products))
	resp.Elements = toList(products, comID, params.HTTPRequest)

	return operations.NewGetAPIV1ProductsOK().WithPayload(resp)
}

// Create ...
func (api *ProductAPI) Create(params operations.PostAPIV1ProductsParams) middleware.Responder {

	resp := &genModels.Product{}

	resp.ProductID = uuid.Must(uuid.NewV4()).String()
	resp.ProductName = params.Body.ProductName
	resp.CompanyID = params.Body.CompanyID
	products[resp.ProductID] = resp
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
	return operations.NewPostAPIV1ProductsCreated().WithPayload(resp)
}
