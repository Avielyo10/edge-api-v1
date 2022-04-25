package image

import (
	"encoding/json"
	netURL "net/url"
	"strings"

	"github.com/Avielyo10/edge-api/internal/common/models"
)

// Repo is a repository of an image.
type Repo struct {
	name string
	url  string
}

// Repos is a list of repos.
type Repos struct {
	repos []*Repo
}

// NewRepo returns a new repo.
func NewRepo(name, url string) *Repo {
	if !isValidRepo(name, url) {
		return nil
	}
	return &Repo{
		name: name,
		url:  url,
	}
}

// NewRepos returns a new list of repos.
func NewRepos(repos ...*Repo) Repos {
	if len(repos) == 0 {
		return Repos{}
	}
	r := Repos{repos: make([]*Repo, 0, len(repos))}
	for _, repo := range repos {
		r.Add(repo)
	}
	return r
}

// isValidRepo returns true if the repo is valid.
func isValidRepo(name, url string) bool {
	_, err := netURL.ParseRequestURI(url)
	return err == nil && name != "" && strings.TrimSpace(name) != ""
}

// IsZero returns true if the repo is empty.
func (r Repo) IsZero() bool {
	return r == Repo{}
}

// Add adds a repo to the list.
func (r *Repos) Add(repos ...*Repo) {
	for _, repo := range repos {
		if repo != nil {
			r.repos = append(r.repos, repo)
		}
	}
}

// String returns the string representation of the repo.
func (r Repo) String() string {
	return `{"name":"` + r.name + `","url":"` + r.url + `"}`
}

// StringArray returns a string array of the repos.
func (r Repos) StringArray() []string {
	var repos []string
	for _, repo := range r.repos {
		repos = append(repos, repo.String())
	}
	return repos
}

// MarshalJSON creates a custom JSON marshaller.
func (r Repo) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	}{
		Name: r.name,
		URL:  r.url,
	})
}

// UnmarshalJSON creates a custom JSON unmarshaller.
func (r *Repo) UnmarshalJSON(b []byte) error {
	var tmp struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	}
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	r.name = tmp.Name
	r.url = tmp.URL
	return nil
}

// MarshalJSON creates a custom JSON marshaller.
func (r Repos) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.repos)
}

// UnmarshalJSON creates a custom JSON unmarshaller.
func (r *Repos) UnmarshalJSON(b []byte) error {
	if string(b) == `[]` {
		return nil
	}
	var repos []Repo
	if err := json.Unmarshal(b, &repos); err != nil {
		return err
	}
	for _, repo := range repos {
		r.Add(NewRepo(repo.name, repo.url))
	}
	return nil
}

// ReposFromArray creates a new list of repos from a map.
func ReposFromArray(repos []interface{}) ([]*Repo, error) {
	// validate the repos by marshalling & unmarshalling
	jrepo, _ := json.Marshal(repos) // nolint: errcheck // ignore error, since this is always valid
	var tmp Repos
	if err := json.Unmarshal(jrepo, &tmp); err != nil {
		return nil, err
	}
	return tmp.repos, nil
}

// MarshalGorm marshals the repos to a gorm model.
func (r Repos) MarshalGorm(account string) []models.Repo {
	var repos []models.Repo
	for _, repo := range r.repos {
		repos = append(repos, models.Repo{
			Account: account,
			Name:    repo.name,
			URL:     repo.url,
		})
	}
	return repos
}
