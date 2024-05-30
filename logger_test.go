package skogul_test

import (
	"testing"

	"github.com/telenornms/skogul"
)

func TestLogPrint(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	x := skogul.NewLog()
	x.LogBus.AddLogWriter("test", func(msg string) {
		if msg == "" {
			t.Error("Empty message in test")
		}
		t.Logf("MESSAGE PASSED TO TEST %s", msg)
	})

	skogul.Print("Testing print")
}

func TestLogPrintf(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	x := skogul.NewLog()
	x.LogBus.AddLogWriter("test", func(msg string) {
		if msg != "" {
			t.Logf("MESSAGE PASSED TO TEST %s", msg)
		}
	})

	skogul.Printf("Testing printf")
}

func TestLogPrintln(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	x := skogul.NewLog()
	x.LogBus.AddLogWriter("test", func(msg string) {
		if msg != "" {
			t.Logf("MESSAGE PASSED TO TEST %s", msg)
		}
	})

	skogul.Println("Testing println")
}

func TestLogPrintFatal(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	x := skogul.NewLog()
	x.LogBus.AddLogWriter("test", func(msg string) {
		if msg != "" {
			t.Logf("MESSAGE PASSED TO TEST %s", msg)
		}
	})

	skogul.Fatal("Testing fatal")
}

func TestLogPrintFatalf(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	x := skogul.NewLog()
	x.LogBus.AddLogWriter("test", func(msg string) {
		if msg != "" {
			t.Logf("MESSAGE PASSED TO TEST %s", msg)
		}
	})

	skogul.Fatalf("Testing fatalf %s", "testing arg")
}

func TestLogPrintFatalln(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	x := skogul.NewLog()
	x.LogBus.AddLogWriter("test", func(msg string) {
		if msg != "" {
			t.Logf("MESSAGE PASSED TO TEST %s", msg)
		}
	})

	skogul.Fatalln("Testing fatalln")
}
