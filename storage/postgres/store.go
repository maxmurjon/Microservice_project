package postgres

import (
	"context"
	"organization_service/genproto/organization_service"
	"organization_service/storage"

	"github.com/google/uuid"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type storeRepo struct {
	db *pgxpool.Pool
}

func NewStoreRepo(db *pgxpool.Pool) storage.StoreRepoI {
	return &storeRepo{
		db: db,
	}
}

func (b *storeRepo) Create(ctx context.Context, req *organization_service.CreateStoreRequest) (resp *organization_service.PrimaryKey, err error) {
	query := `INSERT INTO store(
				id, 
				branch_id, 
				name
				) VALUES (
					$1, 
					$2, 
					$3
				)`

	uuid, err := uuid.NewRandom()
	if err != nil {
		return resp, err
	}

	_, err = b.db.Exec(ctx, query,
		uuid.String(),
		req.BranchId,
		req.Name,
	)

	if err != nil {
		return resp, err

	}

	resp = &organization_service.PrimaryKey{
		Id: uuid.String(),
	}

	return
}

func (b *storeRepo) Get(ctx context.Context, req *organization_service.PrimaryKey) (resp *organization_service.Store, err error) {
	resp = &organization_service.Store{}

	query := `SELECT 
		id, 
		branch_id, 
		name
	FROM store 
	WHERE id = $1`

	err = b.db.QueryRow(ctx, query, req.Id).Scan(
		&resp.Id,
		&resp.BranchId,
		&resp.Name,
	)
	if err != nil {
		return resp, err
	}

	return
}

func (b *storeRepo) GetList(ctx context.Context, req *organization_service.GetListStoreRequest) (resp *organization_service.GetStoresListResponse, err error) {
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

	// query := `select
	// 			id,
	// 			name,
	// 			number_of_pages,
	// 			created_at,
	// 			updated_at
	// 		from books`
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

func (b *storeRepo) Update(ctx context.Context, req *organization_service.UpdateStoreRequest) (rowsAffected int, err error) {
	query := `UPDATE store SET
		branch_id = @branch_id,
		name = @name
	WHERE
		id = @id`

	params := map[string]interface{}{
		"id":          req.Store.Id,
		"branch_id": req.Store.BranchId,
		"name":req.Store.Name,
	}

	result, err := b.db.Exec(ctx, query, pgx.NamedArgs(params))
	if err != nil {
		return 0, err
	}

	rowsAffected = int(result.RowsAffected())

	return rowsAffected, err
}

func (b *storeRepo) PatchUpdate(ctx context.Context, req *organization_service.PatchUpdateRequest) (rowsAffected int, err error) {
	columns := ""
	params := make(map[string]interface{})

	for _, v := range req.Params {
		params[v.Key] = v.Value
		columns += " " + v.Key + " = @" + v.Key
	}
	query := `
		UPDATE 
			store 
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

func (b *storeRepo) Delete(ctx context.Context, req *organization_service.PrimaryKey) (rowsAffected int, err error) {
	query := `DELETE FROM store where id = $1`

	result, err := b.db.Exec(ctx, query, req.Id)
	if err != nil {
		return 0, err
	}

	rowsAffected = int(result.RowsAffected())

	return rowsAffected, err
}
