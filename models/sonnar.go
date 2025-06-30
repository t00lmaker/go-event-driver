package models

const SonnarContextKey = "sonnar"

type Sonnar struct {
	ID        string `json:"id"`
	ProjectID string `json:"project_id"`
	Branch    string `json:"branch"`
	Commit    string `json:"commit"`
	Status    string `json:"status"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func (s *Sonnar) ContextKey() string {
	return BuildContextKey
}

func (s *Sonnar) ContextValue() map[string]string {
	return map[string]string{
		"id":         s.ID,
		"project_id": s.ProjectID,
		"branch":     s.Branch,
		"commit":     s.Commit,
		"status":     s.Status,
		"created_at": s.CreatedAt,
		"updated_at": s.UpdatedAt,
	}
}
func (s *Sonnar) String() string {
	return "Sonnar{" +
		"ID: " + s.ID +
		", ProjectID: " + s.ProjectID +
		", Branch: " + s.Branch +
		", Commit: " + s.Commit +
		", Status: " + s.Status +
		", CreatedAt: " + s.CreatedAt +
		", UpdatedAt: " + s.UpdatedAt +
		"}"
}
