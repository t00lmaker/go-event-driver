package models

const InspectionContextKey = "inspection"

type Inspection struct {
	ID        string `json:"id"`
	BuildID   string `json:"build_id"`
	ProjectID string `json:"project_id"`
	Branch    string `json:"branch"`
	Commit    string `json:"commit"`
	Status    string `json:"status"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func (i *Inspection) ContextKey() string {
	return InspectionContextKey
}

func (i *Inspection) ContextValue() map[string]string {
	return map[string]string{
		"id":         i.ID,
		"build_id":   i.BuildID,
		"project_id": i.ProjectID,
		"branch":     i.Branch,
		"commit":     i.Commit,
		"status":     i.Status,
		"created_at": i.CreatedAt,
		"updated_at": i.UpdatedAt,
	}
}
func (i *Inspection) String() string {
	return "Inspection{" +
		"ID: " + i.ID +
		", BuildID: " + i.BuildID +
		", ProjectID: " + i.ProjectID +
		", Branch: " + i.Branch +
		", Commit: " + i.Commit +
		", Status: " + i.Status +
		", CreatedAt: " + i.CreatedAt +
		", UpdatedAt: " + i.UpdatedAt +
		"}"
}
