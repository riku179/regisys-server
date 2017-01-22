package user

import (
	"github.com/riku179/regisys/models"
	"golang.org/x/net/context"
)

var userKey = "user"

func NewContext(ctx context.Context, u *models.User) context.Context {
	return context.WithValue(ctx, userKey, u)
}

func FromContext(ctx context.Context) (*models.User, bool) {
	u, ok := ctx.Value(userKey).(*models.User)
	return u, ok
}
