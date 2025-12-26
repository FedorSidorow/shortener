package middleware

import (
	"net/http"
	"strings"
	"time"

	"github.com/FedorSidorow/shortener/internal/gzip"
	"github.com/FedorSidorow/shortener/internal/logger"
	"github.com/go-chi/chi/v5/middleware"
)

func LogRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// TODO: Отправьте лог в канал, чтобы не блокировать обработку запроса.
		start := time.Now()
		ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)
		next.ServeHTTP(ww, r)

		logger.Log.Info(
			"[REQEST_INFO] >>> ",
			logger.StringField("method", r.Method),
			logger.StringField("path", r.RequestURI),
			logger.IntField("status", ww.Status()),
			logger.IntField("length", ww.BytesWritten()),
			logger.DurationField("time", time.Since(start)),
		)
	})
}

func GzipRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		var ow http.ResponseWriter
		acceptEncoding := req.Header.Get("Accept-Encoding")
		supportsGzip := strings.Contains(acceptEncoding, "gzip")

		if supportsGzip {
			cw := gzip.NewCompressWriter(res)
			ow = cw
			defer cw.Close()
		} else {
			ow = res
		}

		contentEncoding := req.Header.Get("Content-Encoding")
		if strings.Contains(contentEncoding, "gzip") {

			cr, err := gzip.NewCompressReader(req.Body)
			if err != nil {
				res.WriteHeader(http.StatusInternalServerError)
				return
			}
			req.Body = cr
			defer cr.Close()
		}

		next.ServeHTTP(ow, req)
	})
}
