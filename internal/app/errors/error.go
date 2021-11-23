package errors

import (
	"errors"
	"fmt"
	"strings"
)

type (
	err struct {
		tags   map[string]interface{}
		err    error
		stacks []error
	}
)

func (e *err) Error() string {
	return e.err.Error()
}

func NewError(e error) *err {
	return &err{
		err:    e,
		tags:   make(map[string]interface{}),
		stacks: make([]error, 0),
	}
}

func Is(e, target error) bool {
	if e == nil {
		return false
	}

	current := e.(*err)
	return errors.Is(current.err, target)
}

func (e *err) AddCtx(key, value string) *err {
	e.tags[key] = value
	return e
}

func (e *err) AddStack(err error) *err {
	e.stacks = append(e.stacks, err)
	return e
}

func Format(e error) string {
	if e == nil {
		return ""
	}

	current := e.(*err)

	var tags string
	for k, v := range current.tags {
		tags += fmt.Sprintf("[%s: %v]", k, v)
	}

	traces := make([]string, len(current.stacks))
	for i, e := range current.stacks {
		traces[i] = e.Error()
	}

	var trace string
	if len(traces) > 0 {
		trace = fmt.Sprintf("[stacks: %s]", strings.Join(traces, " - "))
	}

	return fmt.Sprintf("[msg: %v][err: %v]%s%v",
		current.Error(), current.Error(), trace, tags)
}
