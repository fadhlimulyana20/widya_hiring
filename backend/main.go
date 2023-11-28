package main

import (
	"gitlab.com/project-quiz/cmd"
	_ "gitlab.com/project-quiz/utils/env"
	_ "gitlab.com/project-quiz/utils/log"
)

func main() {
	cmd.Execute()
}
