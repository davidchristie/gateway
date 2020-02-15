package resolvers

import (
	"context"
	"os"

	"github.com/davidchristie/gateway/exec"
	"github.com/davidchristie/gateway/middleware"
	"github.com/davidchristie/gateway/model"
	identity "github.com/davidchristie/identity/client"
)

type mutationResolver struct{ *rootResolver }

type rootResolver struct {
	identity identity.Client
}

type queryResolver struct{ *rootResolver }

func NewRootResolver() exec.ResolverRoot {
	return &rootResolver{
		identity: identity.New(&identity.Options{
			Host: os.Getenv("IDENTITY_HOST"),
		}),
	}
}

func (r *rootResolver) Mutation() exec.MutationResolver {
	return &mutationResolver{r}
}

func (r *rootResolver) Query() exec.QueryResolver {
	return &queryResolver{r}
}

func (r *mutationResolver) Login(ctx context.Context, input model.LoginInput) (*model.LoginOutput, error) {
	accessToken, err := r.rootResolver.identity.Login(input.Email, input.Password)
	if err != nil {
		return nil, err
	}
	return &model.LoginOutput{
		AccessToken: *accessToken,
	}, nil
}

func (r *mutationResolver) Logout(ctx context.Context) (bool, error) {
	panic("not implemented")
}

func (r *mutationResolver) Signup(ctx context.Context, input model.SignupInput) (bool, error) {
	err := r.rootResolver.identity.Signup(input.Email, input.Password)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *queryResolver) User(ctx context.Context) (*model.User, error) {
	accessToken := middleware.AccessTokenForContext(ctx)
	if accessToken == nil {
		return nil, nil
	}
	user, err := r.rootResolver.identity.GetUser(*accessToken)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, nil
	}
	return &model.User{
		Email: user.Email(),
	}, nil
}
