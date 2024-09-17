package osquery

import (
	"file-modification-tracker/internal/core"
	"fmt"
	"os/exec"
	"time"
	"encoding/json"
)

type OsqueryAdapter struct{}

type osqueryFileEvent struct {
    Path string `json:"path"`
    Time string `json:"time"`
}

func (o *OsqueryAdapter) CheckModifications(directory string) (string, error) {
	panic("unimplemented")
}

func NewOsqueryAdapter() *OsqueryAdapter {
	return &OsqueryAdapter{}
}

func (o *OsqueryAdapter) GetFileModifications(directory string) ([]core.FileModification, error) {
	query := fmt.Sprintf("select path, time from osquery_file_events where parent_directory_name = '%s'", directory)
	cmd := exec.Command("osqueryi", "--json", query)
	output, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("failed to execute osquery: %w", err)
	}
	var events []osqueryFileEvent
	if err := json.Unmarshal(output, &events); err != nil {
		return nil, fmt.Errorf("failed to unmarshal osquery output: %w", err)
	}
	var modifications []core.FileModification

	return modifications, nil
}

type FileModification struct {
	Filename     string
	LastModified time.Time
}

type MockOsqueryAdapter struct {
	MockFileModifications []core.FileModification
}

func NewMockOsqueryAdapter() *MockOsqueryAdapter {
	return &MockOsqueryAdapter{}
}

func (m *MockOsqueryAdapter) GetFileModifications(directory string) ([]core.FileModification, error) {
	return m.MockFileModifications, nil
}
