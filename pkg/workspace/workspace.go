package workspace

import (
	"strings"
)

func GetDir(workspaceID string) string {
	provider, path, _ := strings.Cut(workspaceID, "://")
	if provider == "directory" {
		return path
	}
	panic("workspaceID must start with directory://")
}
