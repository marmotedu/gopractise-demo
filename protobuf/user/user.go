package user

import (
	"context"

	pb "github.com/marmotedu/api/proto/apiserver/v1"

	"github.com/marmotedu/iam/internal/apiserver/store"
)

type User struct {
}

func (c *User) GetUser(ctx context.Context, r *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	if r.Username != nil {
		return store.Client().Users().GetUserByName(r.Class, r.Username)
	}

	return store.Client().Users().GetUserByID(r.Class, r.UserId)
}
