package function

import (
	"context"
	"fmt"
	"sync/atomic"

	"github.com/baidu/openedge/logger"
	"github.com/baidu/openedge/module/function/runtime"
	"github.com/jolestar/go-commons-pool"
	"github.com/juju/errors"
	"github.com/sirupsen/logrus"
)

// Function function
type Function struct {
	cfg   FunctionConfig
	pool  *pool.ObjectPool
	man   *Manager
	log   *logrus.Entry
	count int64
}

func newFunction(m *Manager, c FunctionConfig) *Function {
	pc := pool.NewDefaultPoolConfig()
	pc.MinIdle = c.Instance.Min
	pc.MaxIdle = c.Instance.Max
	pc.MaxTotal = c.Instance.Max
	pc.MinEvictableIdleTime = c.Instance.IdleTime
	pc.TimeBetweenEvictionRuns = c.Instance.IdleTime / 2
	f := &Function{
		cfg: c,
		man: m,
		log: logger.WithFields("function", c.Name),
	}
	f.pool = pool.NewObjectPool(context.Background(), newFuncionFactory(f), pc)
	return f
}

func (f *Function) newFunclet() (*funclet, error) {
	id := fmt.Sprintf("%s-%s-%d", f.man.cfg.Name, f.cfg.Name, atomic.AddInt64(&f.count, 1))
	fl := &funclet{
		id:  id,
		cfg: f.cfg,
		man: f.man,
		log: f.log.WithField("id", id),
	}
	err := fl.start()
	if err != nil {
		f.log.WithError(err).Errorf("Failed to create function instance")
		fl.Close()
		return nil, errors.Trace(err)
	}
	return fl, nil
}

// Invoke call funtion to handle message and return result message
func (f *Function) Invoke(msg *runtime.Message) (*runtime.Message, error) {
	item, err := f.pool.BorrowObject(context.Background())
	if err != nil {
		return nil, errors.Trace(err)
	}
	fl := item.(*funclet)
	res, err := fl.handle(msg)
	if err != nil {
		f.log.WithError(err).Error("Failed to talk with function instance")
		err1 := f.pool.InvalidateObject(context.Background(), item)
		if err1 != nil {
			fl.Close()
			f.log.WithError(err).Error("Failed to invalidate function instance")
		}
		return nil, errors.Trace(err)
	}
	f.pool.ReturnObject(context.Background(), item)
	return res, nil

}

// Stop all function instances
// The function instance will be stopped in three cases:
// 1. function instance returns a system error
// 2. function instance is not invoked for a period of time [TODO]
// 3. function manager is closed
func (f *Function) close() {
	f.pool.Close(context.Background())
	f.log.Debug("Function closed")
}
