package command

import (
	"context"
	"testing"
)

func Test_cmd(t *testing.T) {

	sess := NewSession()
	sess.ShowLog = true

	sess.SetDir("/tmp")

	a, b := sess.Run(context.Background(), "ls -l && du -h", false)

	//
	t.Log("ok...", a, b)

}

func Test_cmd1(t *testing.T) {

	sess := NewSession()
	sess.ShowLog = true

	a, b := sess.Run(context.Background(), "./a.sh", false)

	//
	t.Log("ok...", a, b)

}

func Test_cmd2(t *testing.T) {

	sess := NewSession()
	sess.ShowLog = true

	a, b := sess.Run(context.Background(), "echo 123444 > 1.log", true)

	//
	t.Log("ok...", a, b)

}
