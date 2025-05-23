package internal

import (
	"errors"
	"fmt"
	"github.com/zalando/go-keyring"
	"regexp"
	"strings"
)

const service = "r2p-2fa"

func IsValidName(name string) bool {
	regex := `^[a-zA-Z0-9_-]+$`
	return regexp.MustCompile(regex).MatchString(name)
}

func Has(name string) bool {
	_, err := keyring.Get(service, name)
	return err == nil
}

func Set(name, code string) error {
	if Has(name) {
		return fmt.Errorf("name %s already exists, consider removing it with `2fa remove %s`", name, name)
	}
	if !IsValidName(name) {
		return fmt.Errorf("name %s is invalid, it call only contain alphanumeric characters, underscores and dashes", name)
	}
	err := keyring.Set(service, name, code)
	if err != nil {
		return fmt.Errorf("failed to set keyring entry: %w", err)
	}
	return appendIndex(name)
}

func Get(name string) (string, error) {
	if !Has(name) {
		return "", fmt.Errorf("name %s does not exist, you can add it with `2fa add %s <code>`", name, name)
	}
	return keyring.Get(service, name)
}

func Remove(name string) error {
	err := keyring.Delete(service, name)
	if err != nil && !errors.Is(err, keyring.ErrNotFound) {
		return fmt.Errorf("failed to delete keyring entry: %w", err)
	}
	return removeIndex(name)
}

func List() ([]string, error) {
	return getIndex()
}

func getIndex() ([]string, error) {
	index, err := keyring.Get(service, "__index__")
	if err != nil && !errors.Is(err, keyring.ErrNotFound) {
		return nil, err
	}

	return strings.Split(strings.Trim(string(index), ";"), ";"), nil
}

func saveIndex(index []string) error {
	if len(index) == 0 {
		return keyring.Delete(service, "__index__")
	}
	return keyring.Set(service, "__index__", strings.Join(index, ";"))
}

func appendIndex(name string) error {
	index, err := getIndex()
	if err != nil {
		return err
	}
	index = append(index, name)
	return saveIndex(index)
}

func removeIndex(name string) error {
	index, err := getIndex()
	if err != nil {
		return err
	}
	for i, n := range index {
		if n == name {
			index = append(index[:i], index[i+1:]...)
			break
		}
	}
	return saveIndex(index)
}
