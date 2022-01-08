package publish

import (
	"Telegraph/internal/handler/api"
	"context"
)

type Publish interface {
	Store(ctx context.Context, key string) error
	All(ctx context.Context, key string) ([]api.Publish, error)
}
