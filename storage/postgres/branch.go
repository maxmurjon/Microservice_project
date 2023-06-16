package postgres

import (
	"context"
	"organization_service/genproto/organization_service"
	"organization_service/storage"

	"github.com/google/uuid"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type branchRepo struct {
	db *pgxpool.Pool
}

func NewBranchRepo(db *pgxpool.Pool) storage.BranchRepoI {
	return &branchRepo{
		db: db,
	}
}

func (b *branchRepo) Create(ctx context.Context, req *organization_service.CreateBranchRequest) (resp *organization_service.PrimaryKey, err error) {
	query := `INSERT INTO branch(
				id, 
				branch_code, 
				name,
				address,
				phone_number
				) VALUES (
					$1, 
					$2, 
					$3,
					$4,
					$5
				)`

	uuid, err := uuid.NewRandom()
	if err != nil {
		return resp, err
	}

	_, err = b.db.Exec(ctx, query,
		uuid.String(),
		req.BranchCode,
		req.Name,
		req.Address,
		req.PhoneNumber,
	)

	if err != nil {
		return resp, err
	}

	resp = &organization_service.PrimaryKey{
		Id: uuid.String(),
	}

	return
}

func (b *branchRepo) Get(ctx context.Context, req *organization_service.PrimaryKey) (resp *organization_service.Branch, err error) {
	resp = &organization_service.Branch{}

	query := `SELECT 
		id, 
		branch_code, 
		name,
		address,
		phone_number
	FROM branch 
	WHERE id = $1`

	err = b.db.QueryRow(ctx, query, req.Id).Scan(
		&resp.Id,
		&resp.BranchCode,
		&resp.Name,
		&resp.Address,
		&resp.PhoneNumber)
	if err != nil {
		return resp, err
	}

	return
}

func (b *branchRepo) GetList(ctx context.Context, req *organization_service.GetListBranchRequest) (resp *organization_service.GetBranchsListResponse, err error) {
	// resp = &organization_service.GetBranchsListResponse{}
	// var (
	// 	params      (map[string]interface{})
	// 	filter      string
	// 	order       string
	// 	arrangement string
	// 	offset      string
	// 	limit       string
	// 	q           string
	// )

	// params = map[string]interface{}{}

	// query := `SELECT 
	// 			id, 
	// 			branch_code, 
	// 			name,
	// 			address,
	// 			phone_number
	// 		FROM branch `
	// filter = " WHERE true"
	// order = " ORDER BY created_at"
	// arrangement = " DESC"
	// offset = " OFFSET 0"
	// limit = " LIMIT 10"

	// if req.Page > 0 {
	// 	req.Page = (req.Page - 1) * req.Limit
	// 	params["offset"] = req.Page
	// 	offset = " OFFSET @offset"
	// }

	// if req.Limit > 0 {
	// 	params["limit"] = req.Limit
	// 	limit = " LIMIT @limit"
	// }

	// cQ := `SELECT count(1) FROM books` + filter

	// err = b.db.QueryRow(ctx, cQ, pgx.NamedArgs(params)).Scan(
	// 	&resp.Count,
	// )

	// if err != nil {
	// 	return resp, err
	// }

	// q = query + filter + order + arrangement + offset + limit

	// rows, err := b.db.Query(ctx, q, pgx.NamedArgs(params))
	// if err != nil {
	// 	return resp, err
	// }
	// defer rows.Close()

	// for rows.Next() {
	// 	book := &book_service.Book{}
	// 	result := &Book{}

	// 	err = rows.Scan(
	// 		&result.Id,
	// 		&result.Name,
	// 		&result.NumberOfPages,
	// 		&result.CreatedAt,
	// 		&result.UpdatedAt,
	// 	)

	// 	if err != nil {
	// 		return resp, err
	// 	}

	// 	if result.CreatedAt.Valid {
	// 		book.CreatedAt = result.CreatedAt.String
	// 	}

	// 	if result.UpdatedAt.Valid {
	// 		book.UpdatedAt = result.UpdatedAt.String
	// 	}

	// 	book.Id = result.Id
	// 	book.Name = result.Name
	// 	book.NumberOfPages = result.NumberOfPages

	// 	resp.Books = append(resp.Books, book)
	// }

	return
}

func (b *branchRepo) Update(ctx context.Context, req *organization_service.UpdateBranchRequest) (rowsAffected int, err error) {
	query := `UPDATE branch SET
		branch_code = @branch_code,
		name = @name,
		address = @address,
		phone_number = @phone_number
	WHERE
		id = @id`

	params := map[string]interface{}{
		"id":           req.Branch.Id,
		"branch_code":  req.Branch.BranchCode,
		"name":         req.Branch.Name,
		"address":      req.Branch.Address,
		"phone_number": req.Branch.PhoneNumber,
	}

	result, err := b.db.Exec(ctx, query, pgx.NamedArgs(params))
	if err != nil {
		return 0, err
	}

	rowsAffected = int(result.RowsAffected())

	return rowsAffected, err
}

func (b *branchRepo) PatchUpdate(ctx context.Context, req *organization_service.PatchUpdateRequest) (rowsAffected int, err error) {
	columns := ""
	params := make(map[string]interface{})

	for _, v := range req.Params {
		params[v.Key] = v.Value
		columns += " " + v.Key + " = @" + v.Key
	}
	query := `
		UPDATE 
			branch 
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

func (b *branchRepo) Delete(ctx context.Context, req *organization_service.PrimaryKey) (rowsAffected int, err error) {
	query := `DELETE FROM branch where id = $1`

	result, err := b.db.Exec(ctx, query, req.Id)
	if err != nil {
		return 0, err
	}

	rowsAffected = int(result.RowsAffected())

	return rowsAffected, err
}
