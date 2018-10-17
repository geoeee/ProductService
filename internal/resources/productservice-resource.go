package resources

import (
	"CompanyService/openapi/gen/productservice/client"
	prodOps "CompanyService/openapi/gen/productservice/client/operations"
	prodModels "CompanyService/openapi/gen/productservice/models"
	"fmt"
	"os"
)

var (
	// ProductServiceHost ...
	ProductServiceHost = client.DefaultHost
	// ProductServiceBasePath ...
	ProductServiceBasePath = client.DefaultBasePath
)

// ProductService ...
type ProductService struct {
	client *client.ProductService
}

// NewProductServiceResource ...
func NewProductServiceResource() IProductService {
	if h := os.Getenv("product_service_host"); h != "" {
		ProductServiceHost = h
	}

	if h := os.Getenv("product_service_basepath"); h != "" {
		ProductServiceBasePath = h
	}
	var resource ProductService
	fmt.Println("product host:", ProductServiceHost)
	fmt.Println("product basepath:", ProductServiceBasePath)
	cfg := client.DefaultTransportConfig().WithHost(ProductServiceHost).WithBasePath(ProductServiceBasePath)
	resource.client = client.NewHTTPClientWithConfig(nil, cfg)
	return &resource
}

// func (resource *ProductService) GetRuntimeClient() *httptransport.Runtime {
// 	return resource.client.Transport.(*httptransport.Runtime)
// }

// func (resource *ProductService) SetRuntimeClient(runtimeClient *httptransport.Runtime) {
// 	*resource.client.Transport.(*httptransport.Runtime) = *runtimeClient
// }

// GetProduct ...
// func (resource *ProductService) GetProduct(params *product.GetProductParams) (*product.GetProductOK, error) {
// 	return resource.client.Product.GetProduct(params)
// }

// GetProducts ...
func (resource *ProductService) GetProducts(params *prodOps.GetAPIV1ProductsParams) ([]*prodModels.Product, error) {
	ok, err := resource.client.Operations.GetAPIV1Products(params)
	if err != nil {
		return nil, err
	}
	return ok.Payload.Elements, nil
}
