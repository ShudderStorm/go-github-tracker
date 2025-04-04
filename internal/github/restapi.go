package github

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/ShudderStorm/go-github-tracker/pkg/oauth"
	"github.com/go-resty/resty/v2"
)

type apiurl string

const (
	apiUser      apiurl = "https://api.github.com/user"
	apiUserRepos apiurl = "https://api.github.com/user/repos"
)

const (
	acceptHeader string = "Accept"
	acceptType   string = "application/vnd.github+json"
)

type RestApi struct {
	access oauth.Access
	client resty.Client
}

func New(access oauth.Access) *RestApi {
	return &RestApi{
		access: access,
		client: *resty.New(),
	}
}

func (r *RestApi) GetUser() (User, error) {
	var user User

	resp, err := r.client.R().
		SetAuthToken(r.access.Token).
		SetHeader(acceptHeader, acceptType).
		Get(string(apiUser))

	if err != nil {
		return user, fmt.Errorf("%w: %w", RequestError, err)
	}

	err = handleStatusCode(resp.StatusCode())
	if err != nil {
		return user, err
	}

	err = json.Unmarshal(resp.Body(), &user)
	return user, err
}

func (r *RestApi) fetchRepos(params url.Values) ([]Repo, error) {
	var repos = make([]Repo, 0)

	resp, err := r.client.R().
		SetAuthToken(r.access.Token).
		SetHeader(acceptHeader, acceptType).
		SetQueryParamsFromValues(params).
		Get(string(apiUserRepos))

	if err != nil {
		return nil, fmt.Errorf("%w: %w", RequestError, err)
	}

	err = handleStatusCode(resp.StatusCode())
	if err != nil {
		return repos, err
	}

	err = json.Unmarshal(resp.Body(), &repos)
	return repos, err
}

func (r *RestApi) OwnedRepos() ([]Repo, error) {
	return r.fetchRepos(
		url.Values{
			"sort":        {"updated"},
			"direction":   {"desc"},
			"affiliation": {"owner"},
		},
	)
}

func (r *RestApi) CollaboratedRepos() ([]Repo, error) {
	return r.fetchRepos(
		url.Values{
			"sort":        {"updated"},
			"direction":   {"desc"},
			"affiliation": {"collaborator"},
		},
	)
}
