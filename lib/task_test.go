package lib

import (
	"os"
	"testing"
)

func buildTask() Task {
	return Task{"test", "touch x.test", "rm x.test", 1, true, false}
}

func TestSanity(t *testing.T) {
	x := 2 * 3
	if x != 6 {
		t.Error("Sanity failed")
	}
}

func TestBuildTask(t *testing.T) {
	task := buildTask()
	if task.Name != "test" {
		t.Error("Task name not test, it's:", task.Name)
	}
}

func TestExecuteIfThis(t *testing.T) {
	task := buildTask()
	success := task.ExecuteIfThis()
	_, err := os.Stat("x.test")
	if err != nil || !success {
		t.Error("ExecuteIfThis did not create x.test")
	} else {
		removeErr := os.Remove("x.test")
		if removeErr != nil {
			panic("failed to cleanup x.test")
		}
	}
}

func TestExecuteThenThat(t *testing.T) {
  task := buildTask()
  success := task.ExecuteIfThis()
  if !success {
    t.Error("Failed to create x.test for testing")
  } else {
    thenThatSuccess := task.ExecuteThenThat()
    _, err := os.Stat("x.test")
    if os.IsExist(err) || !thenThatSuccess {
      t.Error("ExecuteThenThat failed to clean up x.test")
    }
  }
}
