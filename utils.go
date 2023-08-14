package musttag

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
)

func getMainModule() (string, error) {
	args := []string{"go", "list", "-m", "-json"}

	data, err := exec.Command(args[0], args[1:]...).Output()
	if err != nil {
		return "", fmt.Errorf("running `%s`: %w", strings.Join(args, " "), err)
	}

	var module struct {
		Path      string `json:"Path"`
		Main      bool   `json:"Main"`
		Dir       string `json:"Dir"`
		GoMod     string `json:"GoMod"`
		GoVersion string `json:"GoVersion"`
	}

	cwd, _ := os.Getwd()
	decoder := json.NewDecoder(bytes.NewBuffer(data))

	for {
		if err := decoder.Decode(&module); err != nil {
			if errors.Is(err, io.EOF) {
				return "", fmt.Errorf("no main module in: %s", string(data))
			}

			return "", fmt.Errorf("decoding json: %w: %s", err, string(data))
		}

		if module.Main && strings.HasPrefix(cwd, module.Dir) {
			return module.Path, nil
		}
	}
}
