# About

This is a simplified and local program which was inspired by IFTTT. The concept is simple, it reads from config.json and executes the "if this" script. If the "if this" script returns a 0 status, then the "then this" script executes. This concept allows for a wide range of posibilities and complexities within the "if this" and "then this", anything from restarting processes to sending emails can be performed.

# Configuration

The program looks for `config.json` in the same location in which it is being executed. Options are below, and a sample `config.json` is available.

```json
[
  {
    "name": "Find File Task", //The name of the task
    "ifThis": "ls | grep x -q", //The script to run, returns 0 if thenThat should execute
    "thenThat": "touch found", //Executes if ifThis returns 0
    "sleep": 10, //The number of seconds between ifThis checks
    "alwaysPerform": false //thenThat will not execute if the last check returned 0 and the current one returns 0.
  }
]
```
