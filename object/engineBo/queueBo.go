package engineBo

import (
	"github.com/sirupsen/logrus"
)

type QueueBo chan WorkerBo

func (r *QueueBo) AddWork(w WorkerBo) {
	*r <- w
	logrus.Debug("add worker", w)
}

func (r *QueueBo) RunNextWork() {
	w := <-*r

	defer func() {
		if info := recover(); info != nil {
			logrus.Error("catch: ", info)
		} else {
			logrus.Debug("remove and run worker ", w)
		}
	}()

	w.Method(w)
}
