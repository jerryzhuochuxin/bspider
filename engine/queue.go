package engine

type Queue chan Worker

func (r *Queue) AddWork(w Worker) {
	*r <- w
}

func (r *Queue) GetWork() Worker {
	return <-*r
}
