package main

// func (b *bookRepo) GetList(ctx context.Context, req *book_service.GetBooksListRequest) (resp *book_service.GetBooksListResponse, err error) {
// 	resp = &book_service.GetBooksListResponse{}
// 	var (
// 			params      (map[string]interface{})
// 			filter      string
// 			order       string
// 			arrangement string
// 			offset      string
// 			limit       string
// 			q           string
// 	)

// 	params = map[string]interface{}{}

// 	query := `select 
// 							id, 
// 							name, 
// 							number_of_pages,
// 							created_at,
// 							updated_at
// 					from books`
// 	filter = " WHERE true"
// 	order = " ORDER BY created_at"
// 	arrangement = " DESC"
// 	offset = " OFFSET 0"
// 	limit = " LIMIT 10"

// 	if req.Page > 0 {
// 			req.Page = (req.Page - 1) * req.Limit
// 			params["offset"] = req.Page
// 			offset = " OFFSET @offset"
// 	}

// 	if req.Limit > 0 {
// 			params["limit"] = req.Limit
// 			limit = " LIMIT @limit"
// 	}

// 	cQ := `SELECT count(1) FROM books` + filter

// 	err = b.db.QueryRow(ctx, cQ, pgx.NamedArgs(params)).Scan(
// 			&resp.Count,
// 	)

// 	if err != nil {
// 			return resp, err
// 	}

// 	q = query + filter + order + arrangement + offset + limit

// 	rows, err := b.db.Query(ctx, q, pgx.NamedArgs(params))
// 	if err != nil {
// 			return resp, err
// 	}

// defer rows.Close()

// 	for rows.Next() {
// 			book := &book_service.Book{}
// 			result := &Book{}

// 			err = rows.Scan(
// 					&result.Id,
// 					&result.Name,
// 					&result.NumberOfPages,
// 					&result.CreatedAt,
// 					&result.UpdatedAt,
// 			)

// 			if err != nil {
// 					return resp, err
// 			}

// 			if result.CreatedAt.Valid {
// 					book.CreatedAt = result.CreatedAt.String
// 			}

// 			if result.UpdatedAt.Valid {
// 					book.UpdatedAt = result.UpdatedAt.String
// 			}

// 			book.Id = result.Id
// 			book.Name = result.Name
// 			book.NumberOfPages = result.NumberOfPages

// 			resp.Books = append(resp.Books, book)
// 	}

// 	return
// }