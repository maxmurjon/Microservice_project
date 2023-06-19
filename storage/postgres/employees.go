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

type employeesRepo struct {
	db *pgxpool.Pool
}

func NewEmployeesRepo(db *pgxpool.Pool) storage.EmployeesRepoI {
	return &employeesRepo{
		db: db,
	}
}

func (b *employeesRepo) Create(ctx context.Context, req *organization_service.CreateEmployeesRequest) (resp *organization_service.PrimaryKey, err error) {
	query := `INSERT INTO employees(
				id, 
				first_name, 
				last_name,
				store_id,
				phone_number,
				login,
				password_hash,
				role
				) VALUES (
					$1, 
					$2, 
					$3,
					$4,
					$5,
					$6,
					$7,
					$8
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
		req.Login,
		req.PasswordHash,
		req.Role,
	)

	if err != nil {
		return resp, err
	}

	resp = &organization_service.PrimaryKey{
		Id: uuid.String(),
	}

	return
}

func (b *employeesRepo) Get(ctx context.Context, req *organization_service.PrimaryKey) (resp *organization_service.Employees, err error) {
	resp = &organization_service.Employees{}

	query := `SELECT 
		id, 
		first_name, 
		last_name,
		store_id,
		phone_number,
		login,
		password_hash,
		role
	FROM employees 
	WHERE id = $1`

	err = b.db.QueryRow(ctx, query, req.Id).Scan(
		&resp.Id,
		&resp.FirstName,
		&resp.LastName,
		&resp.StoreId,
		&resp.PhoneNumber,
		&resp.Login,
		&resp.PasswordHash,
		&resp.Role,
	)
	if err != nil {
		return resp, err
	}

	return
}

func (b *employeesRepo) GetList(ctx context.Context, req *organization_service.GetEmployeesListRequest) (resp *organization_service.GetEmployeessListResponse, err error) {
	resp = &organization_service.GetEmployeessListResponse{}
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
				login,
				password_hash,
				role
			FROM employees `
	filter = " WHERE true"

	if len(req.GetFirstName()) > 0 {
		filter += " AND first_name ILIKE '%' || '" + req.GetFirstName() + "' || '%' "
	}

	if len(req.GetLastName()) > 0 {
		filter += " AND last_name ILIKE '%' || '" + req.GetLastName() + "' || '%' "
	}

	if len(req.GetPhoneNumber()) > 0 {
		filter += " AND phone_number ILIKE '%' || '" + req.GetPhoneNumber() + "' || '%' "
	}

	cQ := `SELECT count(1) FROM employees` + filter

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
		result := &organization_service.Employees{}

		err = rows.Scan(
			&result.Id,
			&result.FirstName,
			&result.LastName,
			&result.StoreId,
			&result.PhoneNumber,
			&result.Login,
			&result.PasswordHash,
			&result.Role,
		)
		if err != nil {
			return resp, err
		}

		resp.Employeess = append(resp.Employeess, result)
	}

	return
}

func (b *employeesRepo) Update(ctx context.Context, req *organization_service.UpdateEmployeesRequest) (rowsAffected int, err error) {
	query := `UPDATE employees SET
		first_name = @first_name,
		last_name = @last_name,
		store_id = @store_id,
		phone_number = @phone_number,
		login = @login,
		password_hash = @password_hash,
		role = @role
	WHERE
		id = @id`

	params := map[string]interface{}{
		"id":            req.Employees.Id,
		"first_name":    req.Employees.FirstName,
		"last_name":     req.Employees.LastName,
		"store_id":      req.Employees.StoreId,
		"phone_number":  req.Employees.PhoneNumber,
		"login":         req.Employees.Login,
		"password_hash": req.Employees.PasswordHash,
		"role":          req.Employees.Role,
	}

	result, err := b.db.Exec(ctx, query, pgx.NamedArgs(params))
	if err != nil {
		return 0, err
	}

	rowsAffected = int(result.RowsAffected())

	return rowsAffected, err
}

func (b *employeesRepo) PatchUpdate(ctx context.Context, req *organization_service.PatchUpdateRequest) (rowsAffected int, err error) {
	columns := ""
	params := make(map[string]interface{})

	for _, v := range req.Params {
		params[v.Key] = v.Value
		columns += " " + v.Key + " = @" + v.Key
	}
	query := `
		UPDATE 
			employees 
		SET` + columns + `
	WHERE 
		id = @id`

	params["id"] = req.Id
	result, err := b.db.Exec(ctx, query, pgx.NamedArgs(params))
	if err != nil {
		fmt.Println("esngisr")
		return 0, err
	}

	rowsAffected = int(result.RowsAffected())

	return rowsAffected, err
}

func (b *employeesRepo) Delete(ctx context.Context, req *organization_service.PrimaryKey) (rowsAffected int, err error) {
	query := `DELETE FROM employees where id = $1`

	result, err := b.db.Exec(ctx, query, req.Id)
	if err != nil {
		return 0, err
	}

	rowsAffected = int(result.RowsAffected())

	return rowsAffected, err
}
