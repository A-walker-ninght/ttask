package timingwheel

import (
	"sync/atomic"
)

// 按照过期时间排序
type TaskList struct {
	head   *taskNode
	tail   *taskNode
	p      *taskNode // 当前指针
	length int64
}

func newTaskList() *TaskList {
	tasklist := &TaskList{
		head:   newTaskNode(nil),
		tail:   newTaskNode(nil),
		p:      nil,
		length: 0,
	}
	tasklist.head.next = tasklist.tail

	tasklist.tail.next = tasklist.head
	tasklist.head.prev = tasklist.tail
	return tasklist
}

func (l *TaskList) len() int64 {
	return l.length
}

func (l *TaskList) InsertTask(task *taskEntry) (bool, error) {

}

func (l *TaskList) FetchTask() *taskEntry {

}

type taskNode struct {
	prev *taskNode
	next *taskNode
	task *taskEntry
	lock uint32
}

func newTaskNode(task *taskEntry) *taskNode {
	return &taskNode{
		prev: nil,
		next: nil,
		task: task,
		lock: 1,
	}
}

func (n *taskNode) setTask(task *taskEntry) bool {
	if n.task != nil {
		return false
	}
	for {
		if atomic.AddUint32(&n.lock, -1) == 0 {
			n.task = task
			atomic.AddUint32(&n.lock, 1)
			return true
		}
	}
}

// 将任务取出
func (n *taskNode) fetchTask() *taskEntry {
	if task := n.task; task != nil {
		n.removeTask()
		return task
	}
	return nil
}

func (n *taskNode) removeTask() {
	for {
		if atomic.AddUint32(&n.lock, -1) == 0 {
			n.task = nil
			atomic.AddUint32(&n.lock, 1)
			return
		}
	}
}

func (n *taskNode) run() error {
	err := n.task.run()
	n.removeTask()
	return err
}

//func (n *taskNode) delaytime() (time.Duration, error) {
//	if n.task == nil {
//		return time
//	}
//	return n.task.getDeadline()
//}
