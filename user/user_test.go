package user

import (
	"context"
	"github.com/riku179/regisys/models"
	"testing"
)

func TestNewAndFromContext(t *testing.T) {
	ctx := context.Background()
	u := &models.User{ID: 114514}
	ctx = NewContext(ctx, u)
	user, ok := FromContext(ctx)
	if !ok {
		t.Error("failed to cast User")
	}
	if user.ID != 114514 {
		t.Error("User don't match with already set")
	}
}
