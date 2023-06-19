package postgres

import (
	"context"
	"fmt"
	"organization_service/genproto/organization_service"
	"organization_service/storage"

	"github.com/google/uuid"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type supplierRepo struct {
	db *pgxpool.Pool
}

func NewSupplierRepo(db *pgxpool.Pool) storage.SupplierRepoI {
	return &supplierRepo{
		db: db,
	}
}

func (b *supplierRepo) Create(ctx context.Context, req *organization_service.CreateSupplierRequest) (resp *organization_service.PrimaryKey, err error) {
	query := `INSERT INTO supplier(
				id, 
				first_name,
				last_name, 
				store_id,
				phone_number,
				active
				) VALUES (
					$1, 
					$2, 
					$3,
					$4,
					$5,
					$6
				)`

	uuid, err := uuid.NewRandom()
	if err != nil {
		return resp, err
	}

	_, err = b.db.Exec(ctx, query,
		uuid.String(),
		req.FirstName,
		req.LastName,
		req.StoreId,
		req.PhoneNumber,
		req.Active,
	)
	if err != nil {
		fmt.Println("hello")
		return resp, err
	}

	resp = &organization_service.PrimaryKey{
		Id: uuid.String(),
	}

	return
}

func (b *supplierRepo) Get(ctx context.Context, req *organization_service.PrimaryKey) (resp *organization_service.Supplier, err error) {
	resp = &organization_service.Supplier{}

	query := `SELECT 
		id, 
		first_name, 
		last_name,
		store_id,
		phone_number,
		active
	FROM supplier 
	WHERE id = $1`

	err = b.db.QueryRow(ctx, query, req.Id).Scan(
		&resp.Id,
		&resp.FirstName,
		&resp.LastName,
		&resp.StoreId,
		&resp.PhoneNumber,
		&resp.Active,
	)
	if err != nil {
		return resp, err
	}

	return
}

func (b *supplierRepo) GetList(ctx context.Context, req *organization_service.GetListSupplierRequest) (resp *organization_service.GetSuppliersListResponse, err error) {
	resp = &organization_service.GetSuppliersListResponse{}
	var (
		params      (map[string]interface{})
		filter      string
		order       string
		arrangement string
		offset      string
		limit       string
		q           string
	)
	params = map[string]interface{}{}

	query := `SELECT 
				id, 
				first_name, 
				last_name,
				store_id,
				phone_number,
				active
			FROM supplier `
	filter = " WHERE true"

	if len(req.GetFirstName()) > 0 {
		filter += " AND first_name ILIKE '%' || '" + req.GetFirstName() + "' || '%' "
	}

	if len(req.GetLastName()) > 0 {
		filter += " AND last_name ILIKE '%' || '" + req.GetLastName() + "' || '%' "
	}

	cQ := `SELECT count(1) FROM supplier` + filter

	err = b.db.QueryRow(ctx, cQ, pgx.NamedArgs(params)).Scan(
		&resp.Count,
	)
	if err != nil {
		return resp, err
	}

	q = query + filter + order + arrangement + offset + limit

	rows, err := b.db.Query(ctx, q, pgx.NamedArgs(params))
	if err != nil {
		return resp, err
	}
	defer rows.Close()

	for rows.Next() {
		result := &organization_service.Supplier{}

		err = rows.Scan(
			&result.Id,
			&result.FirstName,
			&result.LastName,
			&result.StoreId,
			&result.PhoneNumber,
			&result.Active,
		)
		if err != nil {
			return resp, err
		}

		resp.Suppliers = append(resp.Suppliers, result)
	}

	return
}

func (b *supplierRepo) Update(ctx context.Context, req *organization_service.UpdateSupplierRequest) (rowsAffected int, err error) {
	query := `UPDATE supplier SET
		first_name = @first_name,
		last_name = @last_name,
		store_id = @store_id,
		phone_number = @phone_number,
		active=@active
	WHERE
		id = @id`

	params := map[string]interface{}{
		"id":           req.Supplier.Id,
		"first_name":   req.Supplier.FirstName,
		"last_name":    req.Supplier.LastName,
		"phone_number": req.Supplier.PhoneNumber,
		"store_id":     req.Supplier.StoreId,
		"active":       req.Supplier.Active,
	}

	result, err := b.db.Exec(ctx, query, pgx.NamedArgs(params))
	if err != nil {
		return 0, err
	}

	rowsAffected = int(result.RowsAffected())

	return rowsAffected, err
}

func (b *supplierRepo) PatchUpdate(ctx context.Context, req *organization_service.PatchUpdateRequest) (rowsAffected int, err error) {
	columns := ""
	params := make(map[string]interface{})

	for _, v := range req.Params {
		params[v.Key] = v.Value
		columns += " " + v.Key + " = @" + v.Key
	}
	query := `
		UPDATE 
			supplier 
		SET` + columns + `
	WHERE 
		id = @id`

	params["id"] = req.Id

	result, err := b.db.Exec(ctx, query, pgx.NamedArgs(params))
	if err != nil {
		return 0, err
	}

	rowsAffected = int(result.RowsAffected())

	return rowsAffected, err
}

func (b *supplierRepo) Delete(ctx context.Context, req *organization_service.PrimaryKey) (rowsAffected int, err error) {
	query := `DELETE FROM supplier where id = $1`

	result, err := b.db.Exec(ctx, query, req.Id)
	if err != nil {
		return 0, err
	}

	rowsAffected = int(result.RowsAffected())

	return rowsAffected, err
}
