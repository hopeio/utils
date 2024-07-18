package oauth

import (
	"context"

	"github.com/go-oauth2/oauth2/v4"
)

// ClientStore client information store
type ClientStore interface {
	GetByID(ctx context.Context, id string) (oauth2.ClientInfo, error)
	Set(cli oauth2.ClientInfo) (err error)
}
