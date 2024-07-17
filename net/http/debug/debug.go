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
