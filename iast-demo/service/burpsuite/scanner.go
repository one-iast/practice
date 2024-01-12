package burpsuite

import (
	"fmt"
	"github.com/go-co-op/gocron"
	"iast-demo/util"
	"log"
	"time"
)

type Runner struct {
}

func (r *Runner) Run() {
	fmt.Println("iast-burpsuite start")
	s := gocron.NewScheduler(time.Local)
	_, err := s.Every(5).Second().Do(util.PacketFlow)
	_, err = s.Every(30).Second().Do(SubmitTask)
	_, err = s.Every(10).Second().Do(GetTaskResult)
	_, err = s.Every(10).Second().Do(ShowTaskResult, "all")
	if err != nil {
		log.Fatal(err)
	}
	s.StartAsync()
}
