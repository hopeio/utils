/*
 * Copyright 2024 hopeio. All rights reserved.
 * Licensed under the MIT License that can be found in the LICENSE file.
 * @Created by jyb
 */

package loki

import (
	"github.com/grafana/loki-client-go/loki"
)

type Loki struct {
	Client loki.Client
}

func (lk *Loki) Write(b []byte) (n int, err error) {
	return
}

func (lk *Loki) Sync() error {
	return nil
}

func (lk *Loki) Close() error {
	lk.Client.Stop()
	return nil
}
