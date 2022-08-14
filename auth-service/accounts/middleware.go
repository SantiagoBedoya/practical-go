package accounts

import (
	"context"
	"time"

	"github.com/go-kit/kit/log"
)

// LogginMiddleware define data struct for loggin middleware
type LogginMiddleware struct {
	Logger log.Logger
	Next   AuthService
}

// SignIn implements AuthService's method and make log
func (mw LogginMiddleware) SignIn(ctx context.Context, email string, password string) (output string, err error) {
	defer func(begin time.Time) {
		_ = mw.Logger.Log(
			"method", "sign-in",
			"email", email,
			"output", output,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())
	output, err = mw.Next.SignIn(ctx, email, password)
	return
}

// SignUp implements AuthService's method and make log
func (mw LogginMiddleware) SignUp(ctx context.Context, email string, password string) (output string, err error) {
	defer func(begin time.Time) {
		_ = mw.Logger.Log(
			"method", "sign-up",
			"email", email,
			"output", output,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())
	output, err = mw.Next.SignUp(ctx, email, password)
	return
}
