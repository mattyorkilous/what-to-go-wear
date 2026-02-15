package internal

import (
	"fmt"
	"os"
	"path/filepath"
)

func Initialize() {
	projectConfigDir := makeProjectConfigDir()

	configPath := filepath.Join(projectConfigDir, "config.yaml")

	defaultConfig := []byte("closet:\nofficeDays:")

	tryCreateFile(configPath, defaultConfig, true)

	statePath := filepath.Join(projectConfigDir, "state.yaml")

	tryCreateFile(
		statePath,
		[]byte("lastQueried:\ncurrentCasualIndex:\ncurrentOfficeIndex:"),
		false,
	)
}

func makeProjectConfigDir() string {
	configDir, err := os.UserConfigDir()
	if err != nil {
		fmt.Printf("Error getting user config directory: %v\n", err)
	}

	dir, err := os.Getwd()
	if err != nil {
		fmt.Printf("Error getting working directory: %v\n", err)
	}

	projectName := filepath.Base(dir)

	projectConfigDir := filepath.Join(configDir, projectName)

	os.MkdirAll(projectConfigDir, os.ModePerm)

	return projectConfigDir
}

func tryCreateFile(path string, defaultContent []byte, exitIfNotExists bool) {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		err = os.WriteFile(path, defaultContent, 0644)
		if err != nil {
			fmt.Printf("Error writing default file: %v\n", err)
		}

		if exitIfNotExists {
			fmt.Printf(
				`File does not exist. Created default as %s.
				Please complete the file.\n`,
				path,
			)

			os.Exit(0)
		}
	}
	if err != nil {
		fmt.Printf("Error checking file: %v\n", err)
	}
}
