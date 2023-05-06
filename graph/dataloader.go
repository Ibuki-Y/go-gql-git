package graph

import (
	"context"
	"errors"

	"github.com/Ibuki-Y/go-gql-git/graph/model"
	"github.com/Ibuki-Y/go-gql-git/graph/services"
	"github.com/graph-gophers/dataloader/v7"
)

type Loaders struct {
	UserLoader dataloader.Interface[string, *model.User]
	RepoLoader dataloader.Interface[string, *model.Repository]
}

func NewLoaders(Srv services.Services) *Loaders {
	userBatcher := &userBatcher{Srv: Srv}

	return &Loaders{
		UserLoader: dataloader.NewBatchedLoader(userBatcher.BatchGetUsers),
	}
}

type userBatcher struct {
	Srv services.Services
}

func (u *userBatcher) BatchGetUsers(ctx context.Context, IDs []string) []*dataloader.Result[*model.User] {
	results := make([]*dataloader.Result[*model.User], len(IDs))
	for i := range results {
		results[i] = &dataloader.Result[*model.User]{
			Error: errors.New("not found"),
		}
	}

	// 検索条件であるIDが引数でもらったIDsスライスの何番目のインデックスに格納されていたのか検索できるようにmap化する
	indexs := make(map[string]int, len(IDs))
	for i, ID := range IDs {
		indexs[ID] = i
	}

	users, err := u.Srv.ListUsersByID(ctx, IDs)

	for _, user := range users {
		var res *dataloader.Result[*model.User]
		if err != nil {
			res = &dataloader.Result[*model.User]{
				Error: err,
			}
		} else {
			res = &dataloader.Result[*model.User]{
				Data: user,
			}
		}
		results[indexs[user.ID]] = res
	}

	return results
}
