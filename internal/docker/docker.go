package docker

import (
	"fmt"
	"os/exec"
)

func RunDockerContainer(image string) error {
	cmd := exec.Command("docker", "run", "-d", image)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to run docker container: %v, output: %s", err, string(out))
	}
	fmt.Printf("Container started: %s\n", string(out))
	return nil
}
