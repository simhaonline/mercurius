package sql

import (
	"context"
	"fmt"
)

type sqlCtxKey string

const (
	dbCtx sqlCtxKey = "db"
)

// WithContext creates a new context containing the database transaction, which is usually request-scoped value
// crossing at least different repositories, or even controllers for a specific use case (domain driven design).
// Considering https://tip.golang.org/pkg/context/ and balancing the need of creating all repositories, controllers
// and use cases for each request, just to satisfy an orthogonal constraints, seems not worth it. Neither from
// the complexity nor from the performance aspects. At the end, sql repositories expect their database or transaction
// instances in the context.
func WithContext(ctx context.Context, db DBTX) context.Context {
	return context.WithValue(ctx, dbCtx, db)
}

// FromContext is the counterpart of WithContext and returns a DBTX instance from the request-scoped context.
func FromContext(ctx context.Context) (DBTX, error) {
	if v := ctx.Value(dbCtx); v != nil {
		return v.(DBTX), nil
	}
	return nil, fmt.Errorf("context without DBTX, prepare with sql.WithContext")
}
