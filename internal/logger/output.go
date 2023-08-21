package logger

import (
	"fmt"
	"io/fs"
	"os"
	"path"
)

// outputTask is a line which must be written in the output file.
type outputTask struct {
	Filename string
	Line     string
}

// output if a buffer for lines which must be written in the output file.
var output = make(chan outputTask, 1024)

// starts file writer goroutine to prevent race.
func init() {
	go func() {
		for task := range output {
			// create logs dir
			if err := os.MkdirAll(path.Dir(task.Filename), fs.ModePerm); err != nil {
				fmt.Printf("[LOGGER] unable to use logs dir: %s: %v\n", task.Filename, err)
				continue
			}

			// check if file already exists
			if file, err := os.OpenFile(task.Filename, os.O_RDONLY, os.ModePerm); err != nil {
				if os.IsNotExist(err) {
					// create log file
					if file, err := os.OpenFile(task.Filename, os.O_CREATE|os.O_WRONLY, os.ModePerm); err != nil {
						fmt.Printf("[LOGGER] unable to create file: %s: %v\n", task.Filename, err)
						continue
					} else {
						file.Close()
					}
				}
			} else {
				file.Close()
			}

			// reopen file for appending
			if file, err := os.OpenFile(task.Filename, os.O_APPEND|os.O_WRONLY, os.ModePerm); err != nil {
				fmt.Printf("[LOGGER] unable to open file for appending: %s: %v\n", task.Filename, err)
				continue
			} else {
				if _, err := file.WriteString(task.Line + "\n"); err != nil {
					fmt.Printf("[LOGGER] unable to write data: %s: %v\n", task.Filename, err)
					continue
				}
				file.Close()
			}
		}
	}()
}

func outputWrite(filename string, line string) {
	output <- outputTask{
		Filename: filename,
		Line:     line,
	}
}
