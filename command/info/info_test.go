package info

import (
	"strings"
	"testing"

	"github.com/hashicorp/consul/agent"
	"github.com/mitchellh/cli"
)

func TestInfoCommand_noTabs(t *testing.T) {
	if strings.ContainsRune(New(nil).Help(), '\t') {
		t.Fatal("usage has tabs")
	}
}

func TestInfoCommand(t *testing.T) {
	a1 := agent.NewTestAgent(t.Name(), ``)
	defer a1.Shutdown()

	ui := cli.NewMockUi()
	cmd := New(ui)
	args := []string{"-http-addr=" + a1.HTTPAddr()}

	code := cmd.Run(args)
	if code != 0 {
		t.Fatalf("bad: %d. %#v", code, ui.ErrorWriter.String())
	}

	if !strings.Contains(ui.OutputWriter.String(), "agent") {
		t.Fatalf("bad: %#v", ui.OutputWriter.String())
	}
}
