package user

import (
	"backend/src/governance/entitiy/user"
	"context"

	"github.com/niko-labs/libs-go/helper/paginator"
)

func (r *RepositoryUser) GetUserWithFilter(ctx context.Context, pagination paginator.Pagination) ([]*user.User, *int, error) {
	db := r.GetDB()

	SQL := `
	SELECT 
		id, name, email, avatar, created_at, updated_at, deleted_at, 
		COUNT(*) OVER() AS total 
		FROM users 
			WHERE 
				deleted_at IS NULL 
			ORDER BY 
				name DESC
		LIMIT $1
		OFFSET $2;
	`

	rows, _ := db.Query(ctx, SQL, pagination.Limit, pagination.Offset())

	defer rows.Close()
	err := rows.Err()
	if err != nil {
		return nil, nil, err
	}

	users := []*user.User{}
	var total *int

	for rows.Next() {
		user := &user.User{}
		err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Avatar, &user.CreatedAt, &user.UpdatedAt, &user.DeletedAt, &total)
		if err != nil {
			return nil, nil, err
		}
		users = append(users, user)
	}
	return users, total, nil
}