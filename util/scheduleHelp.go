package util

type ScheduleFunc func(jobFun interface{}, params ...interface{}) error
