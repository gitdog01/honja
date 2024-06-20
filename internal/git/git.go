package git

import (
	"fmt"
	"os/exec"
)

func CloneRepo(repoURL, destDir string) error {
	cmd := exec.Command("git", "clone", repoURL, destDir)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to clone repo: %v, output: %s", err, string(out))
	}
	fmt.Printf("Repository cloned: %s\n", string(out))
	return nil
}
