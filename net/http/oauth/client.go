/*
 * Copyright 2024 hopeio. All rights reserved.
 * Licensed under the MIT License that can be found in the LICENSE file.
 * @Created by jyb
 */

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
