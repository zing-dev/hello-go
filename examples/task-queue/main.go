package main

import (
	"container/list"
	"context"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"sort"
	"sync"
	"time"
)

//1. 队列存储调度系统中各种操作任务
//2. 支持add,pop,close接口，会有多个并发的调用方,pop接口如果没有任务可以pop的话会阻塞
//3. add函数包含一个额外参数指定任务能被pop的最早延时，默认为0(马上能pop)
//4. pop函数支持传入一个处理函数，如果处理函数失败，自动把任务重新add到队列中，重新add时的延时和任务重试的次数有关系，重试次数越多，等待越长
//5. pop时按照任务能被pop的最早时间顺序，没有到任务指定的时间不能pop
//6. 队列需要支持关闭，关闭后pop操作直接报错

type Handler func() error

type Task struct {
	Name     string
	MinPopAt time.Duration //最小的调用时延
	RunAt    time.Time
	Handler  Handler // 处理函数
}

type Queue struct {
	ctx    context.Context
	cancel context.CancelFunc
	wg     *sync.WaitGroup
	Locker *sync.Mutex
	Tasks  []*Task
	Tasks2 *list.List
	Task   chan *Task
}

func NewQueue(size uint) *Queue {
	ctx, cancel := context.WithCancel(context.Background())
	q := &Queue{
		ctx:    ctx,
		cancel: cancel,
		wg:     new(sync.WaitGroup),
		Locker: new(sync.Mutex),
		Task:   make(chan *Task, size),
		Tasks2: list.New(),
	}

	q.wg.Add(1)
	go func() {
		defer q.wg.Done()
		for {
			select {
			case <-ctx.Done():
				return
			default:
				q.Locker.Lock()
				if len(q.Tasks) > 0 {
					time.Sleep(q.Tasks[0].RunAt.Sub(time.Now()))
					q.Task <- q.Tasks[0]
					if len(q.Tasks) > 1 {
						q.Tasks = q.Tasks[1:]
					}else {
						q.Tasks = nil
					}
					//for i, task := range q.Tasks {
					//	if task.RunAt.Before(time.Now()) {
					//		q.Task <- task
					//		log.Println("i ", i, len(q.Tasks))
					//		q.Tasks = append(q.Tasks[:i], q.Tasks[i+1:]...)
					//	}
					//}
				}
				q.Locker.Unlock()
			}
		}
	}()
	return q
}

// Add 添加任务
func (q *Queue) Add(task *Task, minPopAt time.Duration) {
	task.RunAt = time.Now().Add(minPopAt)
	task.MinPopAt = minPopAt
	q.Locker.Lock()
	//if len(q.Tasks) == 0 {
	//	q.Tasks = append(q.Tasks, task)
	//}
	//for i, v := range q.Tasks {
	//if v.RunAt.After(task.RunAt) {
	//	q.Tasks = append(q.Tasks[:i], task)
	//	q.Tasks = append(q.Tasks, q.Tasks[i+1:]...)
	//	break
	//}
	//}
	// 1 2
	log.Printf("add task %s, pop at %s", task.Name, task.RunAt.Format("2006/01/02 15:04:05"))
	//if q.Tasks2.Len() == 0 {
	//	q.Tasks2.PushBack(task)
	//} else {
	//	for e := q.Tasks2.Front(); e != nil; e = e.Next() {
	//		if e.Value.(*Task).RunAt.Before() {}
	//}
	//}
	q.Tasks = append(q.Tasks, task)
	sort.Slice(q.Tasks, func(i, j int) bool {
		return q.Tasks[i].RunAt.Before(q.Tasks[j].RunAt)
	})
	q.Locker.Unlock()
}

// Pop 执行任务
func (q *Queue) Pop(handler Handler) {
	task, ok := <-q.Task
	if !ok {
		panic("channel already close")
	}
	log.Printf("pop task %s", task.Name)
	if err := handler(); err != nil {
		log.Printf("pop task %s err reAdd", task.Name)
		q.Add(task, task.MinPopAt*2)
		return
	}
}

// Close 关闭队列
func (q *Queue) Close() {
	q.cancel()
	close(q.Task)
	q.wg.Wait()
	log.Println("close queue...")
}

func main1() {
	queue := NewQueue(10)
	for i := 0; i < 10; i++ {
		go queue.Add(&Task{Name: fmt.Sprintf("%d", i+1), Handler: func() error { return nil }}, time.Second*time.Duration(rand.Intn(5)+2))
	}

	for i := 0; i < 8; i++ {
		if i == 6 {
			go queue.Pop(func() error { return errors.New("pop err") })
		}
		go queue.Pop(func() error { return nil })
	}
	time.AfterFunc(time.Second*10, func() {
		queue.Close()
	})
	go queue.Pop(func() error { return nil })
	time.Sleep(time.Second * 10)
}
