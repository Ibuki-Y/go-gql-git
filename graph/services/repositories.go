package services

import (
	"context"

	"github.com/Ibuki-Y/go-gql-git/graph/db"
	"github.com/Ibuki-Y/go-gql-git/graph/model"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type repositoryService struct {
	exec boil.ContextExecutor
}

func convertRepository(rep *db.Repository) *model.Repository {
	return &model.Repository{
		ID:        rep.ID,
		Owner:     &model.User{ID: rep.Owner},
		Name:      rep.Name,
		CreatedAt: rep.CreatedAt,
	}
}

func (r *repositoryService) GetRepositoryByID(ctx context.Context, id string) (*model.Repository, error) {
	rep, err := db.FindRepository(
		ctx, r.exec, id,
		db.RepositoryColumns.ID, db.RepositoryColumns.Name, db.RepositoryColumns.Owner, db.RepositoryColumns.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return convertRepository(rep), nil
}

func (r *repositoryService) GetRepositoryByFullName(ctx context.Context, owner, name string) (*model.Repository, error) {
	rep, err := db.Repositories(
		qm.Select(
			db.RepositoryColumns.ID,
			db.RepositoryColumns.Name,
			db.RepositoryColumns.Owner,
			db.RepositoryColumns.CreatedAt,
		),
		db.RepositoryWhere.Owner.EQ(owner),
		db.RepositoryWhere.Name.EQ(name),
	).One(ctx, r.exec)
	if err != nil {
		return nil, err
	}

	return convertRepository(rep), nil
}
