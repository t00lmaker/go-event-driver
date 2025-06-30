package models

const BuildContextKey = "build"

type Build struct {
	ID        string `json:"id"`
	ProjectID string `json:"project_id"`
	Branch    string `json:"branch"`
	Commit    string `json:"commit"`
	Status    string `json:"status"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func (b *Build) ContextKey() string {
	return BuildContextKey
}

func (b *Build) ContextValue() map[string]string {
	return map[string]string{
		"id":         b.ID,
		"project_id": b.ProjectID,
		"branch":     b.Branch,
		"commit":     b.Commit,
		"status":     b.Status,
		"created_at": b.CreatedAt,
		"updated_at": b.UpdatedAt,
	}
}

func (b *Build) String() string {
	return "Build{" +
		"ID: " + b.ID +
		", ProjectID: " + b.ProjectID +
		", Branch: " + b.Branch +
		", Commit: " + b.Commit +
		", Status: " + b.Status +
		", CreatedAt: " + b.CreatedAt +
		", UpdatedAt: " + b.UpdatedAt +
		"}"
}
