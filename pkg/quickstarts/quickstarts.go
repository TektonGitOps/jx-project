package quickstarts

import (
	"strings"
)

func (q *Quickstart) SurveyName() string {
	if q.Owner == JenkinsXQuickstartsOwner {
		return q.Name
	}
	// TODO maybe make a nicer string?
	return q.ID
}

func (f *QuickstartFilter) Matches(q *Quickstart) bool {
	if strings.Contains(q.ID, "WIP-") {
		return false
	}
	text := f.Text
	if text != "" && !strings.Contains(q.ID, text) {
		return false
	}
	owner := strings.ToLower(f.Owner)
	if owner != "" && strings.EqualFold(q.Owner, owner) {
		return false
	}
	language := strings.ToLower(f.Language)
	if language != "" && strings.EqualFold(q.Language, language) {
		return false
	}
	framework := strings.ToLower(f.Framework)
	if framework != "" && strings.EqualFold(q.Framework, framework) {
		return false
	}
	if !f.AllowML && strings.HasPrefix(q.Name, "ML-") {
		return false
	}
	return true
}

// GetGitServer returns the git server to use
func (q *Quickstart) GetGitServer() string {
	if q.GitServer == "" {
		q.GitServer = "https://github.com"
	}
	return q.GitServer
}

// GetGitKind returns the kind of git provider to use
func (q *Quickstart) GetGitKind() string {
	if q.GitKind == "" {
		q.GitKind = "github"
	}
	return q.GitKind
}
