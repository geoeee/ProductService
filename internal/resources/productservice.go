package resources

import (
	prodOps "CompanyService/openapi/gen/productservice/client/operations"
	prodModels "CompanyService/openapi/gen/productservice/models"
)

// IProductService ...
type IProductService interface {
	GetProducts(params *prodOps.GetAPIV1ProductsParams) ([]*prodModels.Product, error)
}
