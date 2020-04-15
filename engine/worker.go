package engine

type dealwith func(worker Worker)

type Worker struct {
	Url    string
	Method dealwith
	Data   interface{}
	Queue  *Queue
}
