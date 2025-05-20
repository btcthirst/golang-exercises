package main

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"
)

// Middleware для логування
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		log.Printf("==> %s %s", r.Method, r.URL.String())
		for name, values := range r.Header {
			for _, v := range values {
				log.Printf("Header: %s=%s", name, v)
			}
		}

		// Обгортаємо ResponseWriter, щоб логувати статус
		lrw := &loggingResponseWriter{ResponseWriter: w, statusCode: 200}
		next.ServeHTTP(lrw, r)

		log.Printf("<== %d [%s]\n", lrw.statusCode, time.Since(start))
	})
}

// Обгортка для ResponseWriter для логування статусу
type loggingResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (lrw *loggingResponseWriter) WriteHeader(code int) {
	lrw.statusCode = code
	lrw.ResponseWriter.WriteHeader(code)
}

// Проксі до іншого сервера
func proxyHandler(target string) http.Handler {
	url, err := url.Parse(target)
	if err != nil {
		log.Fatalf("Invalid target URL: %v", err)
	}

	proxy := httputil.NewSingleHostReverseProxy(url)

	// Перевизначаємо Director, щоб зберегти шлях
	originalDirector := proxy.Director
	proxy.Director = func(req *http.Request) {
		originalDirector(req)
		// Можна змінити заголовки, якщо треба:
		req.Header.Set("X-Forwarded-Host", req.Host)
		req.Header.Set("X-Origin-Host", url.Host)
	}

	// Перевизначаємо ModifyResponse для логування тіла (опційно)
	proxy.ModifyResponse = func(resp *http.Response) error {
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		resp.Body.Close()
		resp.Body = io.NopCloser(bytes.NewReader(bodyBytes))
		log.Printf("Response Body: %s\n", string(bodyBytes))
		return nil
	}

	return proxy
}

func main() {
	targetServer := "http://httpbin.org" // Цільовий сервер для проксі

	mux := http.NewServeMux()
	mux.Handle("/", proxyHandler(targetServer))

	log.Println("Проксі-сервер запущено на :8080")
	err := http.ListenAndServe(":8080", loggingMiddleware(mux))
	if err != nil {
		log.Fatalf("Сервер завершив роботу: %v", err)
	}
}
