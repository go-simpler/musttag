package musttag

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os/exec"
	"strings"
)

type modInfo struct {
	Path      string `json:"Path"`
	Dir       string `json:"Dir"`
	GoMod     string `json:"GoMod"`
	GoVersion string `json:"GoVersion"`
	Main      bool   `json:"Main"`
}

func getMainModule() (string, error) {
	args := []string{"go", "list", "-m", "-json"}

	raw, err := exec.Command(args[0], args[1:]...).Output()
	if err != nil {
		return "", fmt.Errorf("running `%s`: %w", strings.Join(args, " "), err)
	}

	var v modInfo
	err = json.NewDecoder(bytes.NewBuffer(raw)).Decode(&v)
	if err != nil {
		return "", fmt.Errorf("unmarshaling error: %w: %s", err, string(raw))
	}

	return v.Path, nil
}
