package middleware

import (
	"fmt"
	"net/http"
	"runtime/debug"

	"github.com/gorilla/mux"
	"github.com/zhuiye8/Lyss-chat-server/internal/util"
	"github.com/zhuiye8/Lyss-chat-server/pkg/logger"
)

// Recover 中间件处�?panic
func Recover(log *logger.Logger) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				if err := recover(); err != nil {
					// 记录 panic 堆栈
					stack := debug.Stack()
					log.Errorf("PANIC: %v\n%s", err, stack)

					// 返回 500 错误
					util.InternalServerError(w, "服务器内部错�?)
				}
			}()

			next.ServeHTTP(w, r)
		})
	}
}

