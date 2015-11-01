package main

import "encoding/json"
import "io/ioutil"
import "log"
import "flag"

import "./lib"

var configLocation = flag.String("config", "config.json", "Sets the location and name of the config file")

func main() {
	flag.Parse()
	log.Println("Local-IFTTT Starting...")
	log.Println("Reading task config")
	configContents := readConfig()
	tasks := createTasks(configContents)
	schedule(tasks)
}

func schedule(tasks []lib.Task) {
	allTasksDone := make(chan bool, len(tasks))
	for _, task := range tasks {
		log.Println(task.Name, "is starting...")
		go func(t lib.Task, scheduleDone chan bool) {
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
	contents, err := ioutil.ReadFile(*configLocation)
	if err != nil {
		log.Panic("Error reading config.json: ", err)
	}
	return contents
}

func createTasks(rawTasks []byte) []lib.Task {
	var tasks []lib.Task
	err := json.Unmarshal(rawTasks, &tasks)
	if err != nil {
		log.Panic("Unable to unmarshal tasks", err)
	}
	return tasks
}
