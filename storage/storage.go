package storage

import (
	"context"
	"organization_service/genproto/organization_service"
)

type StorageI interface {
	CloseDB()
	Branch() BranchRepoI
	Store() StoreRepoI
	Employees() EmployeesRepoI
	Supplier() SupplierRepoI
	Category() CategoryRepoI
	Product() ProductRepoI
}

type BranchRepoI interface {
	Create(ctx context.Context, req *organization_service.CreateBranchRequest) (resp *organization_service.PrimaryKey, err error)
	Get(ctx context.Context, req *organization_service.PrimaryKey) (resp *organization_service.Branch, err error)
	GetList(ctx context.Context, req *organization_service.GetListBranchRequest) (resp *organization_service.GetBranchsListResponse, err error)
	Update(ctx context.Context, req *organization_service.UpdateBranchRequest) (rowsAffected int, err error)
	PatchUpdate(ctx context.Context, req *organization_service.PatchUpdateRequest) (rowsAffected int, err error)
	Delete(ctx context.Context, req *organization_service.PrimaryKey) (rowsAffected int, err error)
}

type StoreRepoI interface {
	Create(ctx context.Context, req *organization_service.CreateStoreRequest) (resp *organization_service.PrimaryKey, err error)
	Get(ctx context.Context, req *organization_service.PrimaryKey) (resp *organization_service.Store, err error)
	GetList(ctx context.Context, req *organization_service.GetListStoreRequest) (resp *organization_service.GetStoresListResponse, err error)
	Update(ctx context.Context, req *organization_service.UpdateStoreRequest) (rowsAffected int, err error)
	PatchUpdate(ctx context.Context, req *organization_service.PatchUpdateRequest) (rowsAffected int, err error)
	Delete(ctx context.Context, req *organization_service.PrimaryKey) (rowsAffected int, err error)
}

type EmployeesRepoI interface {
	Create(ctx context.Context, req *organization_service.CreateEmployeesRequest) (resp *organization_service.PrimaryKey, err error)
	Get(ctx context.Context, req *organization_service.PrimaryKey) (resp *organization_service.Employees, err error)
	GetList(ctx context.Context, req *organization_service.GetEmployeesListRequest) (resp *organization_service.GetEmployeessListResponse, err error)
	Update(ctx context.Context, req *organization_service.UpdateEmployeesRequest) (rowsAffected int, err error)
	PatchUpdate(ctx context.Context, req *organization_service.PatchUpdateRequest) (rowsAffected int, err error)
	Delete(ctx context.Context, req *organization_service.PrimaryKey) (rowsAffected int, err error)
}

type SupplierRepoI interface {
	Create(ctx context.Context, req *organization_service.CreateSupplierRequest) (resp *organization_service.PrimaryKey, err error)
	Get(ctx context.Context, req *organization_service.PrimaryKey) (resp *organization_service.Supplier, err error)
	GetList(ctx context.Context, req *organization_service.GetListSupplierRequest) (resp *organization_service.GetSuppliersListResponse, err error)
	Update(ctx context.Context, req *organization_service.UpdateSupplierRequest) (rowsAffected int, err error)
	PatchUpdate(ctx context.Context, req *organization_service.PatchUpdateRequest) (rowsAffected int, err error)
	Delete(ctx context.Context, req *organization_service.PrimaryKey) (rowsAffected int, err error)
}

type CategoryRepoI interface {
	Create(ctx context.Context, req *organization_service.CreateCategoryRequest) (resp *organization_service.PrimaryKey, err error)
	Get(ctx context.Context, req *organization_service.PrimaryKey) (resp *organization_service.Category, err error)
	GetList(ctx context.Context, req *organization_service.GetCategorysListRequest) (resp *organization_service.GetCategorysListResponse, err error)
	Update(ctx context.Context, req *organization_service.UpdateCategoryRequest) (rowsAffected int, err error)
	PatchUpdate(ctx context.Context, req *organization_service.PatchUpdateRequest) (rowsAffected int, err error)
	Delete(ctx context.Context, req *organization_service.PrimaryKey) (rowsAffected int, err error)
}

type ProductRepoI interface {
	Create(ctx context.Context, req *organization_service.CreateProductRequest) (resp *organization_service.PrimaryKey, err error)
	Get(ctx context.Context, req *organization_service.PrimaryKey) (resp *organization_service.Product, err error)
	GetList(ctx context.Context, req *organization_service.GetListProductRequest) (resp *organization_service.GetProductsListResponse, err error)
	Update(ctx context.Context, req *organization_service.UpdateProductRequest) (rowsAffected int, err error)
	PatchUpdate(ctx context.Context, req *organization_service.PatchUpdateRequest) (rowsAffected int, err error)
	Delete(ctx context.Context, req *organization_service.PrimaryKey) (rowsAffected int, err error)
}
