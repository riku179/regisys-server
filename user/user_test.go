package user

import (
	"context"
	"testing"

	"github.com/riku179/regisys/models"
)

func TestNewAndFromContext(t *testing.T) {
	ctx := context.Background()
	u := &models.User{ID: 114514}
	ctx = NewContext(ctx, u)
	user := FromContext(ctx)
	if user.ID != 114514 {
		t.Error("User don't match with already set")
	}
}
