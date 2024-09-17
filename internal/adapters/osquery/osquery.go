package osquery

import (
	"fmt"
	"os/exec"
)

type OsqueryAdapter struct{}

// CheckModifications implements core.FileCheckerPort.
func (o *OsqueryAdapter) CheckModifications(directory string) (string, error) {
	panic("unimplemented")
}

func NewOsqueryAdapter() *OsqueryAdapter {
	return &OsqueryAdapter{}
}

func (o *OsqueryAdapter) GetFileModifications(directory string) (string, error) {
	query := fmt.Sprintf("select path from osquery_file_events where parent_directory_name = '%s'", directory)
	cmd := exec.Command("osqueryi", "--json", query)
	output, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("failed to execute osquery: %w", err)
	}

	return string(output), nil
}
