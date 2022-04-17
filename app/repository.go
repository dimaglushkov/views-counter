package app

import (
	"context"
	"fmt"
)

type Repository interface {
	Visit(ctx context.Context, url string) (int64, error)
}

type UnknownUrlError struct {
	Url string
}

func (e UnknownUrlError) Error() string {
	return fmt.Sprintf("unknown url: %s", e.Url)
}
