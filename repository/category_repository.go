package repository

import (
	"context"
	"database/sql"
	"rochiekop/golang-restful-api/model/domain"
)

//Repository (pattern) is data access layer to domain (category)

//First create construct using Interface (Best Practice)

type CategoryRepository interface {
	Save(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category
	Update(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category
	Delete(ctx context.Context, tx *sql.Tx, category domain.Category)
	FindById(ctx context.Context, tx *sql.Tx, categoryId int) (domain.Category, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Category
}

//Implementation using struct
