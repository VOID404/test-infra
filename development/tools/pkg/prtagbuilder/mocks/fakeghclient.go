// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	context "context"
	"net/http"

	github "github.com/google/go-github/v48/github"
	"github.com/kyma-project/test-infra/development/tools/pkg/prtagbuilder"
	mock "github.com/stretchr/testify/mock"
)

func NewFakeGitHubClient(httpClient *http.Client) *prtagbuilder.GitHubClient {
	return &prtagbuilder.GitHubClient{
		Repositories: &GithubRepoService{},
		PullRequests: &GithubPullRequestsService{},
		// any other services your app uses
	}
}

// githubPullRequestsService is an autogenerated mock type for the githubPullRequestsService type
type GithubPullRequestsService struct {
	mock.Mock
}

// Get provides a mock function with given fields: ctx, owner, repo, number
func (_m GithubPullRequestsService) Get(ctx context.Context, owner string, repo string, number int) (*github.PullRequest, *github.Response, error) {
	ret := _m.Called(ctx, owner, repo, number)

	var r0 *github.PullRequest
	if rf, ok := ret.Get(0).(func(context.Context, string, string, int) *github.PullRequest); ok {
		r0 = rf(ctx, owner, repo, number)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*github.PullRequest)
		}
	}

	var r1 *github.Response
	if rf, ok := ret.Get(1).(func(context.Context, string, string, int) *github.Response); ok {
		r1 = rf(ctx, owner, repo, number)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*github.Response)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, string, string, int) error); ok {
		r2 = rf(ctx, owner, repo, number)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// githubRepoService is an autogenerated mock type for the githubRepoService type
type GithubRepoService struct {
	mock.Mock
}

// GetBranch provides a mock function with given fields: ctx, owner, repo, branch
func (_m GithubRepoService) GetBranch(ctx context.Context, owner string, repo string, branch string, followRedirects bool) (*github.Branch, *github.Response, error) {
	ret := _m.Called(ctx, owner, repo, branch)

	var r0 *github.Branch
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string) *github.Branch); ok {
		r0 = rf(ctx, owner, repo, branch)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*github.Branch)
		}
	}

	var r1 *github.Response
	if rf, ok := ret.Get(1).(func(context.Context, string, string, string) *github.Response); ok {
		r1 = rf(ctx, owner, repo, branch)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*github.Response)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, string, string, string) error); ok {
		r2 = rf(ctx, owner, repo, branch)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetCommit provides a mock function with given fields: ctx, owner, repo, sha
func (_m GithubRepoService) GetCommit(ctx context.Context, owner string, repo string, sha string, opts *github.ListOptions) (*github.RepositoryCommit, *github.Response, error) {
	ret := _m.Called(ctx, owner, repo, sha)

	var r0 *github.RepositoryCommit
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string) *github.RepositoryCommit); ok {
		r0 = rf(ctx, owner, repo, sha)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*github.RepositoryCommit)
		}
	}

	var r1 *github.Response
	if rf, ok := ret.Get(1).(func(context.Context, string, string, string) *github.Response); ok {
		r1 = rf(ctx, owner, repo, sha)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*github.Response)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, string, string, string) error); ok {
		r2 = rf(ctx, owner, repo, sha)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}
