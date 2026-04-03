package microsoft

import (
	"container/list"
	"context"
	"sync"
	"time"
)

//设计并实现一个线程安全的泛型异步队列 AsyncQueue<T>，支持可选的 容量上限 和 元素 TTL（过期时间）。
// 并提供一个使用示例：1 个生产者线程、多个消费者线程并发工作

// Java
// public interface AsyncQueue<T> {
// 入队：如果队列已满，则阻塞直到有空间（或超时/中断）
//     void put(T item, long ttlMillis) throws InterruptedException;
// 出队：如果队列为空，则阻塞直到有元素（或超时/中断）
//     T take() throws InterruptedException;
// 出队（带超时）：在 timeoutMillis 内未取到元素返回 null
//     T take(long timeoutMillis) throws InterruptedException;
// 当前队列中“未过期元素”的估计数量（允许近似，但不能明显错误）
//     int size();
// }

// ttlMillis <= 0 表示永不过期
// 元素过期后不能被 take() 返回

type Node[T any] struct {
	val       T
	expiredAt time.Time
}

type AsyncQueue[T any] struct {
	cap  int
	list *list.List
	mu   sync.Mutex
	cond *sync.Cond
}

func NewAsyncQueue[T any](cap int) *AsyncQueue[T] {
	q := &AsyncQueue[T]{
		cap:  cap,
		list: list.New(),
	}
	q.cond = sync.NewCond(&q.mu)
	return q
}

func (q *AsyncQueue[T]) Put(val T, ttl time.Duration) error {
	q.mu.Lock()
	defer q.mu.Unlock()

	for {
		for q.list.Len() >= q.cap {
			q.clean()
			if q.list.Len() >= q.cap {
				q.cond.Wait()
			}
		}
		var expiredAt time.Time
		if ttl > 0 {
			expiredAt = time.Now().Add(ttl)
		}
		n := &Node[T]{
			val:       val,
			expiredAt: expiredAt,
		}
		q.list.PushBack(n)
		q.cond.Signal()
		return nil
	}
}

func (q *AsyncQueue[T]) TakeWithTimeout(ctx context.Context) (T, error) {
	q.mu.Lock()
	defer q.mu.Unlock()

	for {
		for q.list.Len() == 0 {
			select {
			case <-ctx.Done():
				var zero T
				return zero, ctx.Err()
			default:
				q.cond.Wait()
			}
		}

		elem := q.list.Front()
		n := elem.Value.(*Node[T])
		if !n.expiredAt.IsZero() && time.Now().After(n.expiredAt) {
			q.list.Remove(elem)
			continue
		}

		// take success
		q.list.Remove(elem)
		return n.val, nil
	}
}

func (q *AsyncQueue[T]) Take() (T, error) {
	for {
		for q.list.Len() == 0 {
			q.cond.Wait()
		}
		elem := q.list.Front()
		n := elem.Value.(*Node[T])
		if time.Now().After(n.expiredAt) {
			q.list.Remove(elem)
			continue
		}

		// take success
		q.list.Remove(elem)
		return n.val, nil
	}
}

func (q *AsyncQueue[T]) Size() int {
	q.mu.Lock()
	defer q.mu.Unlock()
	q.clean()
	return q.list.Len()
}

func (q *AsyncQueue[T]) clean() {
	q.mu.Lock()
	defer q.mu.Unlock()
	now := time.Now()
	for e := q.list.Front(); e != nil; {
		next := e.Next()
		n := e.Value.(*Node[T])
		if !n.expiredAt.IsZero() && now.After(n.expiredAt) {
			q.list.Remove(e)
		}
		e = next
	}
}
