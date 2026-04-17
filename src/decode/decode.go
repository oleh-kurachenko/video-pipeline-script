package decode

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"

	"video-pipeline-script/common"
)

type DecodingError struct {
	message string
}

func (e *DecodingError) Error() string {
	return e.message
}

var inputFilenamePattern = regexp.MustCompile(`^\d+\.mp4$`)

func makeTranscodingTasks(dirs []string,
	outputPrefix string) (tasks []common.TranscodingTask,
	err error) {
	for _, dir := range dirs {
		fileInfo, err := os.Stat(dir)

		if os.IsNotExist(err) {
			return nil,
				&DecodingError{fmt.Sprintf("%s does not exist", dir)}
		}
		if !fileInfo.Mode().IsDir() {
			return nil,
				&DecodingError{fmt.Sprintf("%s is not a directory", dir)}
		}
		if err != nil {
			return nil, err
		}

		files, err := os.ReadDir(dir)
		if err != nil {
			return nil, &DecodingError{fmt.Sprintf(
				"directory %s could not be read: %s", dir, err.Error())}
		}

		for _, file := range files {
			if file.IsDir() || !inputFilenamePattern.MatchString(file.Name()) {
				continue
			}

			task := common.TranscodingTask{
				InputFilename:  filepath.Join(dir, file.Name()),
				OutputFilename: filepath.Join(dir, outputPrefix+file.Name()),
				Options:        []string{},
			}
			tasks = append(tasks, task)
		}
	}

	return tasks, nil
}

func Decode(args []string) error {
	_, err := makeTranscodingTasks(args, "dnxhd_")
	if err != nil {
		return err
	}

	return nil
}
