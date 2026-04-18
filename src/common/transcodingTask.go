package common

import (
	"fmt"
	"log/slog"

	"github.com/u2takey/ffmpeg-go"
)

type TranscodingTask struct {
	InputFilename  string
	OutputFilename string
	InputOptions   map[string]string
	OutputOptions  map[string]string
}

func makeKwArgs(args map[string]string) ffmpeg_go.KwArgs {
	kwargs := ffmpeg_go.KwArgs{}
	for key, value := range args {
		kwargs[key] = value
	}

	return kwargs
}

func (task TranscodingTask) Transcode() error {
	return ffmpeg_go.
		Input(task.InputFilename, makeKwArgs(task.InputOptions)).
		Output(task.OutputFilename, makeKwArgs(task.OutputOptions)).
		OverWriteOutput().
		Silent(true).
		Run()
}

func TranscodeTasks(tasks []TranscodingTask) error {
	for i, task := range tasks {
		slog.Info(fmt.Sprintf("transcoding task %d/%d: %s -> %s", i+1,
			len(tasks), task.InputFilename, task.OutputFilename))
		err := task.Transcode()
		if err != nil {
			return err
		}
		slog.Info(fmt.Sprintf("done transcoding %d/%d: %s -> %s", i+1,
			len(tasks), task.InputFilename, task.OutputFilename))
	}

	return nil
}
