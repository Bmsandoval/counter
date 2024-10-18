package pkg

import "os/user"

// HomeDir returns the user's home directory
func HomeDir() (string, error) {
	usr, err := user.Current()
	if err != nil {
		return "", err
	}
	return usr.HomeDir, nil
}
