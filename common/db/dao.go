package db

import "context"

type Dao interface {
	FindOne(ctx context.Context, id interface{}, v interface{}) error
	Find(ctx context.Context, q QueryParams, v interface{}) error
	Insert(ctx context.Context, v interface{}) error
	Update(ctx context.Context, q QueryParams, v interface{}) error
	Upsert(ctx context.Context, q QueryParams, v interface{}) error
	Delete(ctx context.Context, q QueryParams) (int, error)
	DeleteCollection(ctx context.Context) error
	PipeAll(ctx context.Context, pipeline interface{}, result interface{}) error
	Count(ctx context.Context, q QueryParams) (int, error)
}

type QueryParams string // temp
