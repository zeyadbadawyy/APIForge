package storage

import (
	"encoding/json"
	"os"
)

const FileName = "data/tasks.json"

func SaveTasks() error {

	data, err :=
		json.MarshalIndent(
			Tasks,
			"",
			"  ",
		)

	if err != nil {
		return err
	}

	return os.WriteFile(
		FileName,
		data,
		0644,
	)
}

func LoadTasks() error {

	data, err :=
		os.ReadFile(FileName)

	if err != nil {
		return err
	}

	return json.Unmarshal(
		data,
		&Tasks,
	)
}
