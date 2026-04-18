package decode

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"video-pipeline-script/common"
)

type DecodingError struct {
	message string
}

func (e *DecodingError) Error() string {
	return e.message
}

var inputFilenamePattern = regexp.MustCompile(`^(\d+)\.mp4$`)

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
			fileName := strings.ToLower(file.Name())

			if file.IsDir() || !inputFilenamePattern.MatchString(fileName) {
				continue
			}

			newFileName := outputPrefix + inputFilenamePattern.
				ReplaceAllString(fileName, "$1.mov")

			task := common.TranscodingTask{
				InputFilename:  filepath.Join(dir, file.Name()),
				OutputFilename: filepath.Join(dir, newFileName),
			}
			tasks = append(tasks, task)
		}
	}

	return tasks, nil
}

func Decode(args []string) error {
	tasks, err := makeTranscodingTasks(args, "dnxhd_")
	if err != nil {
		return err
	}

	for i := range tasks {
		tasks[i].InputOptions = map[string]string{
			"loglevel": "warning",
			"hwaccel":  "cuda",
		}
		tasks[i].OutputOptions = map[string]string{
			"loglevel":  "warning",
			"c:v":       "dnxhd",
			"profile:v": "dnxhr_hqx",
			"pix_fmt":   "yuv422p10le",
			"c:a":       "pcm_s16le",
		}
	}

	return common.TranscodeTasks(tasks)
}
