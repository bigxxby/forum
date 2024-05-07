package crypto

import "github.com/gofrs/uuid"

func GenerateUUID() (string, error) {
	id, err := uuid.DefaultGenerator.NewV4()
	if err != nil {
		return "", err
	}
	return id.String(), nil
}
