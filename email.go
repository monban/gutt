package main

type email struct {
	from    string
	subject string
	body    string
}

func (e email) FilterValue() string {
	return e.subject
}

func (e email) Title() string {
	return e.subject
}

func (e email) Description() string {
	return e.from
}
