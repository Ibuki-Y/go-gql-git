package services

import (
	"context"
	"log"

	"github.com/Ibuki-Y/go-gql-git/graph/db"
	"github.com/Ibuki-Y/go-gql-git/graph/model"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type issueService struct {
	exec boil.ContextExecutor
}

func convertIssue(issue *db.Issue) *model.Issue {
	issueURL, err := model.UnmarshalURI(issue.URL)
	if err != nil {
		log.Println("invalid URI: ", issue.URL)
	}

	return &model.Issue{
		ID:         issue.ID,
		URL:        issueURL,
		Title:      issue.Title,
		Closed:     (issue.Closed == 1),
		Number:     int(issue.Number),
		Author:     &model.User{ID: issue.Author},
		Repository: &model.Repository{ID: issue.Repository},
	}
}

func (i *issueService) GetIssueByID(ctx context.Context, id string) (*model.Issue, error) {
	issue, err := db.FindIssue(
		ctx, i.exec, id,
		db.IssueColumns.ID,
		db.IssueColumns.URL,
		db.IssueColumns.Title,
		db.IssueColumns.Closed,
		db.IssueColumns.Number,
		db.IssueColumns.Author,
		db.IssueColumns.Repository,
	)
	if err != nil {
		return nil, err
	}

	return convertIssue(issue), nil
}

func (i *issueService) GetIssueByRepoAndNumber(ctx context.Context, repID string, number int) (*model.Issue, error) {
	issue, err := db.Issues(
		qm.Select(
			db.IssueColumns.ID,
			db.IssueColumns.URL,
			db.IssueColumns.Title,
			db.IssueColumns.Closed,
			db.IssueColumns.Number,
			db.IssueColumns.Author,
			db.IssueColumns.Repository,
		),
		db.IssueWhere.Repository.EQ(repID),
		db.IssueWhere.Number.EQ(int64(number)),
	).One(ctx, i.exec)
	if err != nil {
		return nil, err
	}

	return convertIssue(issue), nil
}
