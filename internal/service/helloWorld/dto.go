package helloWorld

import "context"

type Query struct {
	name string
}

func (q Query) Valid(ctx context.Context) map[string]string {
	problems := make(map[string]string)
	if q.name == "" {
		problems["name"] = "name query is empty"
	}

	return problems
}
