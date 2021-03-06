package handlers

/**
 * Before filterライフサイクルに値する。
 * ハンドラー実行の直前に実行される。
 */

import (
	"net/http"
)

func setHeader(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Credentials", "false")
	w.Header().Set("Access-Control-Allow-Header", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")

	if err := r.ParseForm(); err != nil {
		panic(err)
	}
}
