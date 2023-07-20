/*
 * Copyright Bytedance Ltd. and/or its affiliates.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package pool

import (
	"github.com/bytedance/arishem/typedef"
	"github.com/bytedance/gopkg/util/logger"
	"runtime"
	"runtime/debug"
	"sync/atomic"
	"time"
)

// taskHolder is the pool task define which hold parameter
type taskHolder struct {
	fn typedef.Task
	// param is the parameter that task can hold and pass it when task is scheduled
	param interface{}
}

// WorkPool is a work pool that bind to some idl goroutines and put task to run.
type WorkPool struct {
	size  int32
	tasks chan *taskHolder

	// maxIdle is the number of the max idle workers in the pool.
	maxIdle int32

	maxIdleTime time.Duration
}

// NewWorkPool construct a new work pool
func NewWorkPool(maxIdle int, maxIdleTime time.Duration) *WorkPool {
	if maxIdle <= 0 {
		maxIdle = runtime.NumCPU() * 2
	}
	if maxIdleTime <= 0 {
		maxIdleTime = time.Second * 30
	}
	return &WorkPool{
		tasks:   make(chan *taskHolder),
		maxIdle: int32(maxIdle),
	}
}

// Submit a task and run
func (wp *WorkPool) Submit(t typedef.Task, p interface{}) {
	if t == nil {
		return
	}
	wp.submitInternal(&taskHolder{
		fn:    t,
		param: p,
	})
}

// SubmitWait submit the task and wait it util finish
func (wp *WorkPool) SubmitWait(t typedef.Task, p interface{}) {
	if t == nil {
		return
	}
	dc := make(chan struct{})
	wp.submitInternal(&taskHolder{
		fn: func(param interface{}) {
			t(p)
			close(dc)
		},
		param: p,
	})
	<-dc
}

func (wp *WorkPool) submitInternal(t *taskHolder) {
	select {
	case wp.tasks <- t:
		// reuse goroutine
		return
	default:
	}

	atomic.AddInt32(&wp.size, 1)
	go func() {
		defer func() {
			if r := recover(); r != nil {
				logger.Errorf("panic in work pool, stack=>%s", r, debug.Stack())
			}
			atomic.AddInt32(&wp.size, -1)
		}()
		// remove debug log
		//logger.Info("goroutine created!")

		t.fn(t.param)

		if atomic.LoadInt32(&wp.size) > wp.maxIdle {
			return
		}

		// waiting for new task
		idleTimer := time.NewTimer(wp.maxIdleTime)
		for {
			select {
			case t = <-wp.tasks:
				if t != nil {
					t.fn(t.param)
				}
			case <-idleTimer.C:
				// waiting timeout, exits
				return
			}

			// case when it does a task, stop the time timer, and receive the signal left, and start a new loop
			if !idleTimer.Stop() {
				<-idleTimer.C
			}
			// continue waiting
			idleTimer.Reset(wp.maxIdleTime)
		}
	}()
}
