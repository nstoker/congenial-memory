package kudo

import (
	"strconv"

	"github.com/nstoker/congenial-memory/pkg/core"
)

// GitHubRepo does what you might think
type GitHubRepo struct {
	RepoID      int64  `json:"id"`
	RepoURL     string `json:"html_url"`
	RepoName    string `json:"full_name"`
	Language    string `json:"language"`
	Description string `json:"description"`
	Notes       string `json:"notes"`
}

// Service another does what it says on the tin
type Service struct {
	userID string
	repo   core.Repository
}

// GetKudos get some kudos
func (s Service) GetKudos() ([]*core.Kudo, error) {
	return s.repo.FindAll(map[string]interface{}{"userId": s.userID})
}

// CreateKudoFor yes
func (s Service) CreateKudoFor(gitHubRepo GitHubRepo) (*core.Kudo, error) {
	kudo := s.gitHubRepoToKudo(githubRepo)
	err := s.repo.Create(kudo)
	if err != nil {
		return nil, err
	}

	return kudo, nil
}

// UpdateKudoWith updates kudo
func (s Service) UpdateKudoWith(gitHubRepo GitHubRepo) (*core.Kudo, error) {
	kudo := s.githubRepoToKudo(gitHubRepo)
	err := s.repo.Create(kudo)
	if err != nil {
		return nil, err
	}

	return kudo, nil
}

// RemoveKudo 'cos we're sick of it
func (s Service) RemoveKudo(githubRepo GitHubRepo) (*core.Kudo, error) {
	kudo := s.githubRepoToKudo(githubRepo)
	err := s.repo.Delete(kudo)
	if err != nil {
		return nil, err
	}

	return kudo, nil
}

func (s Service) githubRepoToKudo(githubRepo GitHubRepo) *core.Kudo {
	return &core.Kudo{
		UserID:      s.userID,
		RepoID:      strconv.Itoa(int(githubRepo.RepoID)),
		RepoName:    githubRepo.RepoName,
		RepoURL:     githubRepo.RepoURL,
		Language:    githubRepo.Language,
		Description: githubRepo.Description,
		Notes:       githubRepo.Notes,
	}
}

// NewService creates a new service. The descriptions are so edgy
func NewService(repo core.Repository, userID string) Service {
	return Service{
		repo:   repo,
		userID: userID,
	}
}
