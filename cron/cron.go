package cron

import (
	"fmt"

	"github.com/robfig/cron"
)

func SetupCron(frequencyInMinutes int) {
	cronInstance := cron.New()

	cronInstance.AddFunc(fmt.Sprintf("@every %dm", frequencyInMinutes), syncCommands())

	cronInstance.Start()
}
