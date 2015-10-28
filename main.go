package main

import "encoding/json"
import "io/ioutil"
import "log"
import "os/exec"
import "time"

func main() {
	log.Println("Local-IFTTT Starting...")
	log.Println("Reading task config")
	configContents := readConfig()
	tasks := createTasks(configContents)
	schedule(tasks)
}

func schedule(tasks []Task) {
	allTasksDone := make(chan bool, len(tasks))
	for _, task := range tasks {
		log.Println(task.Name, "is starting...")
		go func(t Task, scheduleDone chan bool) {
			for true {
				taskDone := make(chan bool)
				go t.Run(taskDone)
				<-taskDone
				log.Println(t.Name, "is done, restarting...")
			}
			scheduleDone <- true
		}(task, allTasksDone)
	}
	for i := range allTasksDone {
		log.Panic("All tasks should never complete: ", i)
	}
}

func readConfig() []byte {
	contents, err := ioutil.ReadFile("config.json")
	if err != nil {
		log.Panic("Error reading config.json", err)
	}
	return contents
}

func createTasks(rawTasks []byte) []Task {
	var tasks []Task
	err := json.Unmarshal(rawTasks, &tasks)
	if err != nil {
		log.Panic("Unable to unmarshal tasks", err)
	}
	return tasks
}

type Task struct {
	Name          string
	IfThis        string
	ThenThat      string
	Sleep         int64
	AlwaysPerform bool
	LastSuccess   bool
}

func (t *Task) Run(done chan bool) {
	time.Sleep(time.Duration(t.Sleep) * time.Second)
	if t.ExecuteIfThis() {
		if t.AlwaysPerform || !t.LastSuccess {
			log.Println(t.Name, "if this successful, executing then that")
			t.ExecuteThenThat()
		}
		t.LastSuccess = true
	} else {
		t.LastSuccess = false
	}
	done <- true
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
