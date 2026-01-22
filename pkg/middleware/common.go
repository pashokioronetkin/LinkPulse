package middleware

import "net/http"

type WrapperWriter struct { // WrapperWriter это структура, которая оборачивает http.ResponseWriter и добавляет поле StatusCode.
	http.ResponseWriter
	StatusCode int
}

func (w *WrapperWriter) WriteHeader(statusCode int) { // WriteHeader это метод, который устанавливает статус код ответа и вызывает WriteHeader у ResponseWriter.
	w.ResponseWriter.WriteHeader(statusCode)
	w.StatusCode = statusCode
}
