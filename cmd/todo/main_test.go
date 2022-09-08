package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"testing"
)

var (
	binName  = "todo"
	fileName = "todos.json"
)

func TestMain(m *testing.M) {
	fmt.Println("Building tool...")

	if runtime.GOOS == "windows" {
		binName += ".exe"
	}

	build := exec.Command("go", "build", "-o", binName)

	if err := build.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Cannot build tool %s: %s", binName, err)
		os.Exit(1)
	}

	result := m.Run()
	fmt.Println("Cleaning up...")
	os.Remove(binName)
	os.Remove(fileName)
	os.Exit(result)
}

func TestTodoCli(t *testing.T) {
	task := "Vanquish foes"
	dir, err := os.Getwd()
	cmdPath := filepath.Join(dir, binName)

	if err != nil {
		t.Fatal(err)
	}

	t.Run("AddNewTask", func(t *testing.T) {
		cmd := exec.Command(cmdPath, fmt.Sprintf("-add=%s", task))

		if err := cmd.Run(); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("ListTasks", func(t *testing.T) {
		cmd := exec.Command(cmdPath, "-list")
		out, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatal(err)
		}

		expected := "Vanquish foes\n"
		if string(out) != expected {
			t.Errorf("Expected %q, got %q instead\n", expected, string(out))
		}
	})
}
