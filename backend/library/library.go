package library

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/viper"
)

func LoadLibrary(path string) (*Librarian, error) {
	librarian := &Librarian{
		Library:  make([]Story, 0),
		FilePath: path,
	}

	data, err := os.ReadFile(path)

	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(data, librarian); err != nil {
		return nil, fmt.Errorf("decode library: %w", err)
	}

	return librarian, nil

}

func (l *Librarian) SaveLibrary() error {
	data, err := json.MarshalIndent(l, "", "	")
	if err != nil {
		fmt.Printf("Error saving library: %v\n", err)
		return err
	}
	os.WriteFile(l.FilePath, data, 0644)
	return nil
}

func (l *Librarian) GetLibraryPath() string {
	if l.FilePath != "" {
		return l.FilePath
	}

	return viper.GetString("libraryPath")
}
