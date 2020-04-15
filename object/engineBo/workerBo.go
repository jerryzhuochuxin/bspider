package engineBo

type dealwith func(worker WorkerBo)

type WorkerBo struct {
	Url    string
	Method dealwith
	Data   interface{}
	Queue  *QueueBo
}
