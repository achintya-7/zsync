package utils

import (
	"errors"
	"os"
	"strings"
)

const zshrcPreexecFunction = `
zsync_preexec() {
    local cmd="$1"
    local pwd="$PWD"

    # Use command substitution to execute curl silently and suppress output
    pid=$( curl -s -X POST -H "Content-Type: application/json" \
        -d '{"command":"'"$cmd"'", "cwd":"'"$pwd"'"}' \
        http://localhost:8080/log > /dev/null & echo $! )
}
`

func CheckAndFillZshrc() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "error in getting home directory", err
	}

	zshrcPath := homeDir + "/.zshrc"

	content, err := os.ReadFile(zshrcPath)
	if err != nil {
		return "error in reading zshrc file", err
	}

	if len(content) == 0 {
		return "zshrc file is empty", errors.New("zshrc file is empty")
	}

	contentStr := string(content)
	if strings.Contains(contentStr, "zsync_preexec()") {
		return "zshrc file already has zsync_preexec function", nil
	}

	contentStr += zshrcPreexecFunction

	err = os.WriteFile(zshrcPath, []byte(contentStr), 0644)
	if err != nil {
		return "error in writing zshrc file", err
	}

	return "zshrc file updated successfully", nil
}
