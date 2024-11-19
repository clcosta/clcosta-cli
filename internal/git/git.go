package git

import (
	"errors"
	"fmt"
	"os/exec"
)

// Configure sets the Git user configuration
func Configure(user string) error {
	var email, name string

	switch user {
	case "personal":
		email = "personal@example.com"
		name = "PersonalUser"
	case "work":
		email = "work@example.com"
		name = "WorkUser"
	default:
		return errors.New("invalid user type. Valid options: personal, work")
	}

	_, err := exec.Command("git", "config", "--local", "user.email", email).Output()
	if err != nil {
		return fmt.Errorf("failed to set Git email: %w", err)
	}

	_, err = exec.Command("git", "config", "--local", "user.name", name).Output()
	if err != nil {
		return fmt.Errorf("failed to set Git name: %w", err)
	}

	fmt.Printf("Git user configured: %s <%s>\n", name, email)
	return nil
}
