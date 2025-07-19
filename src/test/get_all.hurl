GET http://localhost:8080/tasks
HTTP 200

[Asserts]
jsonpath "$[0].id" == 1
jsonpath "$[0].task" == "Spin up service"
jsonpath "$[0].status" == "running"