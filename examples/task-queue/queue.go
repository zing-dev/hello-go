package main

import (
	"errors"
	"fmt"
	"log"
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

type job struct {
	name     string        // 任务名
	at       time.Time     // 任务执行时间
	duration time.Duration // 最小的延时
	action   func() error  // 执行函数
}

type jobs []job

type queue struct {
	jobs     jobs           // 任务队列
	jobChAdd chan job       // 添加的任务
	jobChPop chan job       // 可以消费的任务
	stop     chan struct{}  // 关闭队列标志
	once     sync.Once      // 单例执行
	wg       sync.WaitGroup // 等待协程完成
	mu       sync.Mutex     // 锁
}

// do 任务监控
func (q *queue) do() {
	q.wg.Add(1)
	go func() {
		defer q.wg.Done()
		for {
			select {
			case <-q.stop:
				close(q.stop)
				close(q.jobChAdd)
				close(q.jobChPop)
				return
			case job, ok := <-q.jobChAdd:
				if !ok {
					return
				}
				// 将来的执行时间
				job.at = time.Now().Add(job.duration)
				q.mu.Lock()
				q.jobs = append(q.jobs, job)
				log.Printf("add job %s run at %s size %d\n", job.name, job.at.Format("2006/01/02 15:04:05"), len(q.jobs))
				// 对队列里面的任务根据最小的时延排序,最先执行的任务在队首
				sort.Slice(q.jobs, func(i, j int) bool {
					return q.jobs[i].at.Before(q.jobs[j].at)
				})
				q.mu.Unlock()
			default:
				q.mu.Lock()
				if len(q.jobs) > 0 {
					sub := q.jobs[0].at.Sub(time.Now())
					if sub > time.Duration(0) {
						time.Sleep(sub)
					}
					select {
					case q.jobChPop <- q.jobs[0]:
						if len(q.jobs) == 1 {
							q.jobs = nil
						} else {
							q.jobs = q.jobs[1:]
						}
						//default: // 消费队列已满,丢掉当前的任务,或去掉default分支阻塞消费
						//	log.Printf("break %s", q.jobs[0].name)
					}
				}
				q.mu.Unlock()
				time.Sleep(time.Millisecond)
			}
		}
	}()
}

// add 添加任务
// 1. 当队列里面的任务满了,会丢弃添加的任务
// 2. 如果不丢弃的,去掉default分支,会阻塞当前任务添加,当队列里面的任务有被消费时候添加
func (q *queue) add(job job, duration time.Duration) {
	job.duration = duration
	select {
	case q.jobChAdd <- job:
		//default:
	}
}

// pop 任务出队
func (q *queue) pop(handler func() error) {
	job, ok := <-q.jobChPop
	if !ok {
		panic("sorry close channel")
	}
	if err := handler(); err != nil {
		log.Printf("pop %s error \n", job.name)
		q.add(job, job.duration*2) // 执行失败,最小的延时*2
		return
	}
	log.Printf("pop %s success\n", job.name)

}

// 关闭队列
func (q *queue) close() {
	q.once.Do(func() {
		q.stop <- struct{}{}
		q.wg.Wait()
		log.Println("close...")
	})
}

func main() {
	q := queue{
		jobChAdd: make(chan job, 10),
		jobChPop: make(chan job, 10),
		stop:     make(chan struct{}),
	}
	q.do()
	for i := 1; i < 110; i++ {
		go q.add(job{name: fmt.Sprintf("%d", i), action: func() error { return nil }}, time.Second*time.Duration(i%4+1))
		if i%10 == 0 {
			go q.pop(func() error { return errors.New("pop err") })
		} else {
			go q.pop(func() error { return nil })
		}
	}
	for i := 0; i < 50; i++ {
		go q.pop(func() error { return nil })
	}

	time.AfterFunc(time.Second*30, func() {
		q.close()
		q.pop(func() error { return nil })
	})
	time.Sleep(time.Second * 100)
}
