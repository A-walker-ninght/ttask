package timingwheel

import (
	"fmt"
	"time"
)

type task func() error // 任务

type taskEntry struct {
	name    string
	t       task
	delayms time.Duration // 超时时间
}

func (t *taskEntry) getName() string {
	return t.name
}

func (t *taskEntry) run() error {
	defer func() {
		if err := recover(); err != nil {
			//log.Fatalf(),打印日志
		}
	}()
	if t.t == nil {
		panic(fmt.Sprintf("the task: %s is nil", t.name))
	}
	return t.t()
}

func (t *taskEntry) getDeadline() time.Duration {
	return t.delayms
}
