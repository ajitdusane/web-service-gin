package main

import (
	"os"

	"path/filepath"

	"encoding/json"

	"github.com/case-framework/case-backend/pkg/study/types"

	"log/slog"

	"github.com/ajitdusane/web-service-gin/app_config"
)

func WriteParticipantToFile(participant types.Participant) {

	jsonParticipant, err := json.Marshal(participant)

	if err != nil {
		slog.Error("Participant could not be converted to json", slog.String("error", err.Error()))
		return
	}

	if err := os.MkdirAll(app_config.AppConfig.Data.Dir, os.ModePerm); err != nil {
		slog.Error("could not create "+app_config.AppConfig.Data.Dir, slog.String("error", err.Error()))
		return
	}

	filePath := filepath.Join(app_config.AppConfig.Data.Dir, participant.ParticipantID+".json")

	if err := os.WriteFile(filePath, jsonParticipant, 0644); err != nil {
		slog.Error("Error writing json to file", slog.String("error", err.Error()))
	}
}

// Reads participants stored in the directory (path in config.yaml) as json files
func ReadParticipants() {

	jsonFiles, err := os.ReadDir(app_config.AppConfig.Data.Dir)
	slog.Info("read participants from "+app_config.AppConfig.Data.Dir, slog.Int("length", len(jsonFiles)))

	if err != nil {
		slog.Error("could not read " + app_config.AppConfig.Data.Dir)
	}

	for _, file := range jsonFiles {

		if filepath.Ext(file.Name()) != ".json" {
			continue
		}

		var participant types.Participant

		if json.Unmarshal(ReadParticipantAsJson(file.Name()), &participant) != nil {
			slog.Error("error while converting " + file.Name() + " content to json")
		}

		Participants = append(Participants, participant)
	}
}

// Read participant.json file
func ReadParticipantAsJson(filename string) []byte {

	filePath := filepath.Join(app_config.AppConfig.Data.Dir, filename)

	content, err := os.ReadFile(filePath)
	if err != nil {
		slog.Error("could not read " + filePath)
	}

	return content
}

func ParticipantExists(filename string) bool {

	_, err := os.Stat(filepath.Join(app_config.AppConfig.Data.Dir, filename))
	return !os.IsNotExist(err)
}
