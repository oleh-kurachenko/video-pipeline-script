package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
)

type decodingError struct {
	message string
}

func (e *decodingError) Error() string {
	return e.message
}

var inputFilenamePattern = regexp.MustCompile(`^\d+\.mp4$`)

func makeTranscodingTasks(dirs []string,
	outputPrefix string) (tasks []transcodingTask,
	err error) {
	for _, dir := range dirs {
		fileInfo, err := os.Stat(dir)

		if os.IsNotExist(err) {
			return nil,
				&decodingError{fmt.Sprintf("%s does not exist", dir)}
		}
		if !fileInfo.Mode().IsDir() {
			return nil,
				&decodingError{fmt.Sprintf("%s is not a directory", dir)}
		}
		if err != nil {
			return nil, err
		}

		files, err := os.ReadDir(dir)
		if err != nil {
			return nil, &decodingError{fmt.Sprintf(
				"directory %s could not be read: %s", dir, err.Error())}
		}

		for _, file := range files {
			if file.IsDir() || !inputFilenamePattern.MatchString(file.Name()) {
				continue
			}

			task := transcodingTask{
				inputFilename:  filepath.Join(dir, file.Name()),
				outputFilename: filepath.Join(dir, outputPrefix+file.Name()),
				options:        []string{},
			}
			tasks = append(tasks, task)
		}
	}

	return tasks, nil
}

func decode(args []string) error {
	_, err := makeTranscodingTasks(args, "dnxhd_")
	if err != nil {
		return err
	}

	return nil
}
