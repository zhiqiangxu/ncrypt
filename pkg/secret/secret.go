package secret

import (
	"crypto/sha256"
	"fmt"
	"syscall"

	"golang.org/x/crypto/ssh/terminal"
)

// Secret for deciding secret
type Secret struct {
	s []byte
}

// New is ctor for Secret
func New() *Secret {
	return &Secret{}
}

// GetSecret returns the decided secret
func (s *Secret) GetSecret() []byte {
	return s.s
}

// DecideSecret will prompt for parties and finally generate a secret
func (s *Secret) DecideSecret(parties []string) (err error) {

	h := sha256.New()

	var bytePassword []byte
	for _, party := range parties {
		fmt.Printf("Enter Pass for: %s\n", party)
		bytePassword, err = terminal.ReadPassword(int(syscall.Stdin))
		if err != nil {
			return
		}

		_, err = h.Write(bytePassword)
		if err != nil {
			return
		}
	}

	s.s = h.Sum(nil)

	return
}
