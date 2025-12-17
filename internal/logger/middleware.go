package logger

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5/middleware"
	"go.uber.org/zap"
)

func LogRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// TODO: Отправьте лог в канал, чтобы не блокировать обработку запроса.
		// Запоминаем время начала обработки запроса
		start := time.Now()

		// Создаём обёртку над ResponseWriter для захвата статуса и размера ответа
		ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)

		// Вызываем следующий обработчик в цепочке
		next.ServeHTTP(ww, r)

		// Формируем лог после обработки запроса
		Log.Info(
			"[REQEST_INFO] >>> ",
			StringField("method", r.Method),
			StringField("path", r.RequestURI),
			IntField("status", ww.Status()),
			IntField("length", ww.BytesWritten()),
			DurationField("time", time.Since(start)),
		)
	})
}

func DurationField(s string, duration time.Duration) zap.Field {
	return zap.Duration(s, duration)
}

func StringField(s1, s2 string) zap.Field {
	return zap.String(s1, s2)
}

func IntField(s string, val int) zap.Field {
	return zap.Int(s, val)
}
