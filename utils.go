package musttag

import (
	"fmt"
	"os/exec"
	"strings"
)

// mainModulePackages returns a set of packages that belong to the main module.
func mainModulePackages() (map[string]struct{}, error) {
	// https://pkg.go.dev/cmd/go#hdr-Package_lists_and_patterns
	// > When using modules, "all" expands to all packages in the main module
	// > and their dependencies, including dependencies needed by tests of any of those.

	// NOTE(junk1tm): the command may run out of file descriptors if go version <= 1.18,
	// especially on macOS, which has the default soft limit set to 256 (ulimit -nS).
	// Since go1.19 the limit is automatically increased to the maximum allowed value;
	// see https://github.com/golang/go/issues/46279 for details.
	cmd := [...]string{"go", "list", "-f={{if and (not .Standard) .Module.Main}}{{.ImportPath}}{{end}}", "all"}

	out, err := exec.Command(cmd[0], cmd[1:]...).Output()
	if err != nil {
		return nil, fmt.Errorf("running go list: %w", err)
	}

	list := strings.TrimSpace(string(out))
	m := make(map[string]struct{}, len(list))
	for _, pkg := range strings.Split(list, "\n") {
		m[pkg] = struct{}{}
	}

	return m, nil
}
