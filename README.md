# rapi
Basic RestAPI in Go without Framework

# enpoints
POST    http://127.0.0.1:8080/tasks     Create task
Payload
{
    "name": "test"
    "description": "test"
}

GET     http://127.0.0.1:8080/tasks/id  Read task

GET     http://127.0.0.1:8080/tasks     Read tasks

PUT     http://127.0.0.1:8080/tasks/id  Update task
{
    "name": "test"
    "description": "test"
}

DELETE  http://127.0.0.1:8080/tasks/id  Delete task