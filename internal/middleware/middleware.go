package middleware

import (
	"bytes"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/FedorSidorow/shortener/config"
	"github.com/FedorSidorow/shortener/internal/auth"
	"github.com/FedorSidorow/shortener/internal/gzip"
	"github.com/FedorSidorow/shortener/internal/logger"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/google/uuid"
)

// LogRequest — middleware-логер для входящих HTTP-запросов.
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

// GzipRequest - определяет запрос на сжатие и при необходимости возвращает сжатый ответ.
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

// AuthCookieMiddleware(next http.Handler) http.Handler — middleware-для входящих HTTP-запросов.
// Выдаёт пользователю симметрично подписанную куку, содержащую уникальный идентификатор пользователя,
// если такой куки не существует или она не проходит проверку подлинности.
func AuthCookieMiddleware(next http.Handler, options *config.Options) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var (
			UserID      uuid.UUID
			tokenString string

			ctx        = r.Context()
			newCookie  = http.Cookie{Name: auth.NameCookie}
			token, err = r.Cookie(auth.NameCookie)
		)

		if err == nil {
			tokenString = token.Value
			UserID = auth.GetUserID(options, tokenString)

			if UserID == uuid.Nil {
				http.Error(w, "", http.StatusUnauthorized)
				return
			}

		} else {
			UserID = uuid.New()

			tokenString, err = auth.BuildJWTString(options, UserID)

			if err != nil {
				logger.Log.Error(err.Error())
				http.Error(w, "", http.StatusInternalServerError)
				return
			}

			newCookie.Value = tokenString
			http.SetCookie(w, &newCookie)
		}

		ctx = auth.WithUserID(ctx, UserID)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}

// AuditMiddleware - создает аудит запись.
func AuditMiddleware(next http.HandlerFunc, action string, pub *Publisher) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var originalURL string
		if action == "shorten" {
			body, err := io.ReadAll(r.Body)
			if err != nil {
				logger.Log.Error("failed to read request body for audit", logger.ErrorField(err))
			} else {
				originalURL = string(body)
				r.Body = io.NopCloser(bytes.NewReader(body))
			}
		}

		ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)
		var locationURL string

		next(ww, r)

		status := ww.Status()
		if status >= 400 {
			return
		}

		if action == "follow" {
			loc := ww.Header().Get("Location")
			if loc != "" {
				locationURL = loc
			}
		}

		url := originalURL
		if action == "follow" && locationURL != "" {
			url = locationURL
		}

		userID, _ := auth.UserIDFrom(r.Context())
		userIDStr := ""
		if userID != uuid.Nil {
			userIDStr = userID.String()
		}

		event := auditEvent{
			TS:     time.Now().Unix(),
			Action: action,
			UserID: userIDStr,
			URL:    url,
		}

		pub.Notify(event)
	})
}
