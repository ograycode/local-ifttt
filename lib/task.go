package lib

import "log"
import "os/exec"
import "time"

type Task struct {
	Name          string
	IfThis        string
	ThenThat      string
	Sleep         int64
	AlwaysPerform bool
	LastSuccess   bool
}

func (t *Task) Run(done chan bool) {
	if t.ExecuteIfThis() {
		if t.AlwaysPerform || !t.LastSuccess {
			log.Println(t.Name, "if this successful, executing then that")
			t.ExecuteThenThat()
		}
		t.LastSuccess = true
	} else {
		t.LastSuccess = false
	}
	t.SleepNow()
	done <- true
}

func (t *Task) SleepNow() {
	time.Sleep(time.Duration(t.Sleep) * time.Second)
}

func (t *Task) ExecuteThenThat() bool {
	return t.execute(t.ThenThat)
}

func (t *Task) ExecuteIfThis() bool {
	return t.execute(t.IfThis)
}

func (t *Task) execute(command string) bool {
	cmd := exec.Command("bash", "-c", command)
	err := cmd.Run()
	if err != nil {
		return false
	}
	return true
}
