package time

import (
	"context"
	"golang.org/x/time/rate"
	"log"
	"testing"
	"time"
)

func Test_RateLimiter(t *testing.T) {
	/**
	func NewLimiter(r Limit, b int) *Limiter
	type Limiter struct {
		mu     sync.Mutex
		limit  Limit
		burst  int
		tokens float64
		// last is the last time the limiter's tokens field was updated
		last time.Time
		// lastEvent is the latest time of a rate-limited event (past or future)
		lastEvent time.Time
	}
	limit每秒传输的个数，burst桶内总个数
	*/
	l := rate.NewLimiter(2, 1)
	log.Println(l.Limit(), l.Burst())

	for i := 0; i < 100; i++ {
		//阻塞等待直到取到第一个token
		log.Println("before Wait")
		log.Println("reserve Delay", l.Reserve().Delay(), l.Tokens())
		c, _ := context.WithTimeout(context.Background(), time.Second*2)
		if err := l.Wait(c); err != nil {
			log.Println("limiter wait err: " + err.Error())
		}
		log.Println("after Wait")

		//返回需要等待多久才有新的token，这样就可以等待指定时间执行任务
		r := l.Reserve()
		log.Println("reserve Delay", r.Delay(), l.Tokens())
		//time.Sleep(r.Delay())
		//判断当时是否可以取到token
		a := l.Allow()
		log.Println("Allow: ", a)
	}
}
