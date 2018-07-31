package ddevexpect

import (
	"fmt"
	"os"
	"os/exec"
	"testing"

	expect "github.com/Netflix/go-expect"
	"github.com/hinshun/vt10x"
	"github.com/stretchr/testify/require"
)

func TestDdevConfig(t *testing.T) {
	err := os.RemoveAll("test")
	require.NoError(t, err)

	err = os.Mkdir("test", os.ModePerm)
	require.NoError(t, err)

	err = os.Chdir("test")
	require.NoError(t, err)

	logFileName := fmt.Sprintf("%s.log", t.Name())
	logFile, err := os.OpenFile(logFileName, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	require.NoError(t, err)
	defer logFile.Close()

	c, state, err := vt10x.NewVT10XConsole(expect.WithStdout(logFile))
	require.NoError(t, err)

	donec := make(chan struct{})
	go func() {
		defer close(donec)
		c.ExpectString("Project name")
		c.SendLine("")
		c.ExpectString("Docroot Location")
		c.SendLine("")
		c.ExpectString("Project Type")
		c.SendLine("")
		c.ExpectEOF()
	}()

	cmd := exec.Command("ddev", "config")
	cmd.Stdin = c.Tty()
	cmd.Stdout = c.Tty()
	cmd.Stderr = c.Tty()

	err = cmd.Run()
	require.NoError(t, err)

	c.Tty().Close()
	<-donec

	t.Log(expect.StripTrailingEmptyLines(state.String()))
}
