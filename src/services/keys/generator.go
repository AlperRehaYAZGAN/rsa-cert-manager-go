// This is the simple go file for encrypt and decrypt text with symmetric and asymmetric keys.

package keys

import "crypto/rand"

func GenerateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}

	return b, nil
}
