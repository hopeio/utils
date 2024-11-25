/*
 * Copyright 2024 hopeio. All rights reserved.
 * Licensed under the MIT License that can be found in the LICENSE file.
 * @Created by jyb
 */

package debug

import (
	_ "expvar"
	"net/http"
	_ "net/http/pprof"
	"runtime/debug"
)

func init() {
	http.Handle("/debug/stack", http.HandlerFunc(Stack))
}

func Handler() http.Handler {
	return http.DefaultServeMux
}

func Stack(w http.ResponseWriter, r *http.Request) {
	w.Write(debug.Stack())
}
