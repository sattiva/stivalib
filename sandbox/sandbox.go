package sandbox

import (
	"context"
	"errors"
	"os/exec"
	"time"
)

type Result struct {
	Stdout   []byte
	Stderr   []byte
	ExitCode int
	TimedOut bool
}

func ExecBounded(parentContext context.Context, timeout time.Duration, name string, args ...string) (*Result, error) {
	ctx, cancel := context.WithTimeout(parentContext, timeout)
	defer cancel()

	cmd := exec.CommandContext(ctx, name, args...)
	var stdout, stderr []byte
	var err error

	stdout, err = cmd.Output()
	if err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			stderr = exitErr.Stderr
			res := &Result{
				Stdout:   stdout,
				Stderr:   stderr,
				ExitCode: exitErr.ExitCode(),
				TimedOut: ctx.Err() == context.DeadlineExceeded,
			}
			return res, nil
		}
		if ctx.Err() == context.DeadlineExceeded {
			return &Result{TimedOut: true, ExitCode: -1}, errors.New("command execution timed out")
		}
		return nil, err
	}

	return &Result{
		Stdout:   stdout,
		Stderr:   stderr,
		ExitCode: 0,
		TimedOut: false,
	}, nil
}
