package main

import "time"

type email struct {
	from    string
	subject string
	body    string
	time    time.Time
}

func (e email) FilterValue() string {
	return e.subject
}

func (e email) Title() string {
	return e.subject
}

func (e email) Description() string {
	return e.time.Local().Format(time.DateTime)
}
