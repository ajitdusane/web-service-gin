package main

import (
	"os"

	"path/filepath"

	"encoding/json"

	"github.com/case-framework/case-backend/pkg/study/types"

	"log/slog"

	"github.com/ajitdusane/web-service-gin/app_config"
)

// serializes a Participant and saves it to a JSON file.
func WriteParticipantToFile(participant types.Participant) error {

	jsonParticipant, err := json.Marshal(participant)

	if err != nil {
		return logError("Failed to marshal participant to JSON", err)
	}

	// Ensure the directory exists
	if err := createDirectory(app_config.AppConfig.Data.Dir); err != nil {
		return err
	}

	// Write JSON to file
	filePath := filepath.Join(app_config.AppConfig.Data.Dir, participant.ParticipantID+".json")
	if err := os.WriteFile(filePath, jsonParticipant, 0644); err != nil {
		return logError("Failed to write participant JSON to file", err)
	}

	return nil
}

// creates the directory if it does not exist
func createDirectory(dir string) error {
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return logError("Failed to create directory: "+dir, err)
	}
	return nil
}

func logError(message string, err error) error {
	slog.Error(message, slog.String("error", err.Error()))
	return err
}

// Reads participants stored in the directory (path in config.yaml) as json files
func ReadAllParticipants() error {

	jsonFiles, err := os.ReadDir(app_config.AppConfig.Data.Dir)
	if err != nil {
		return logError("Failed to read directory: "+app_config.AppConfig.Data.Dir, err)
	}

	slog.Info("Reading participants", slog.Int("file_count", len(jsonFiles)))

	for _, file := range jsonFiles {

		if filepath.Ext(file.Name()) != ".json" {
			continue
		}

		participant, err := ReadParticipantFromFile(file.Name())
		if err != nil {
			slog.Error("Failed to unmarshal participant from file: "+file.Name(), slog.String("error", err.Error()))
			continue
		}

		Participants = append(Participants, participant)
	}

	return nil
}

// readParticipantFromFile reads a single participant JSON file and unmarshals it into a Participant struct.
func ReadParticipantFromFile(filename string) (types.Participant, error) {
	var participant types.Participant

	content, err := ReadFileAsBytes(filename)
	if err != nil {
		return participant, err
	}

	if err := json.Unmarshal(content, &participant); err != nil {
		return participant, err
	}

	return participant, nil
}

// reads the content of a file and returns it as a byte slice.
func ReadFileAsBytes(filename string) ([]byte, error) {
	filePath := filepath.Join(app_config.AppConfig.Data.Dir, filename)

	content, err := os.ReadFile(filePath)
	if err != nil {
		return nil, logError("Failed to read file: "+filePath, err)
	}

	return content, nil
}

func ParticipantExists(filename string) bool {

	_, err := os.Stat(filepath.Join(app_config.AppConfig.Data.Dir, filename))
	return !os.IsNotExist(err)
}
