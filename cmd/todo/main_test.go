package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"testing"
)

var (
	binName          = "todo"
	fileNameForTests = "todos.json"
)

func TestMain(m *testing.M) {
	// TestMain is where we do environment set ups of tear downs for this
	// suite of integration tests.
	err := os.Setenv("TODO_FILENAME", "test-todos.json")

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error setting test env variable")
		os.Exit(1)
	}

	if os.Getenv("TODO_FILENAME") != "" {
		fileNameForTests = os.Getenv("TODO_FILENAME")
	}

	fmt.Println("Building tool...")

	if runtime.GOOS == "windows" {
		binName += ".exe"
	}

	build := exec.Command("go", "build", "-o", binName)

	if err := build.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Cannot build tool %s: %s", binName, err)
		os.Exit(1)
	}

	result := m.Run() // Here, we run the tests.
	fmt.Println("Cleaning up...")
	err = os.Unsetenv("TODO_FILENAME")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error setting test env variable")
		os.Exit(1)
	}
	os.Remove(binName)
	os.Remove(fileNameForTests)
	os.Exit(result)
}

func TestTodoCli(t *testing.T) {
	task := "Vanquish foes"
	task2 := "Ride chocobo"
	dir, err := os.Getwd()
	cmdPath := filepath.Join(dir, binName)

	if err != nil {
		t.Fatal(err)
	}

	t.Run("AddNewTask", func(t *testing.T) {
		cmd := exec.Command(cmdPath, "-add", task)
		if err := cmd.Run(); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("AddNewTaskFromSTDIN", func(t *testing.T) {
		cmd := exec.Command(cmdPath, "-add")
		cmdStdIn, err := cmd.StdinPipe()
		if err != nil {
			t.Fatal(err)
		}
		io.WriteString(cmdStdIn, task2)
		cmdStdIn.Close()

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

		expected := fmt.Sprintf("  1: %s\n  2: %s\n", task, task2)
		if string(out) != expected {
			t.Errorf("Expected %q, got %q instead\n", expected, string(out))
		}
	})

	t.Run("CompleteTask", func(t *testing.T) {
		cmd := exec.Command(cmdPath, fmt.Sprintf("-complete=%d", 1))
		if err := cmd.Run(); err != nil {
			t.Fatal(err)
		}
	})
}
