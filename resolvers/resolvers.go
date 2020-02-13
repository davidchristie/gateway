package resolvers

import (
	"context"

	"github.com/davidchristie/gateway/exec"
	"github.com/davidchristie/gateway/model"
)

type mutationResolver struct{ *rootResolver }

type rootResolver struct{}

type queryResolver struct{ *rootResolver }

func NewRootResolver() exec.ResolverRoot {
	return &rootResolver{}
}

func (r *rootResolver) Mutation() exec.MutationResolver {
	return &mutationResolver{r}
}

func (r *rootResolver) Query() exec.QueryResolver {
	return &queryResolver{r}
}

func (r *mutationResolver) Login(ctx context.Context, input model.LoginInput) (*model.LoginOutput, error) {
	panic("not implemented")
}
func (r *mutationResolver) Logout(ctx context.Context) (bool, error) {
	panic("not implemented")
}
func (r *mutationResolver) Signup(ctx context.Context, input model.SignupInput) (bool, error) {
	panic("not implemented")
}

func (r *queryResolver) User(ctx context.Context) (*model.User, error) {
	panic("not implemented")
}
