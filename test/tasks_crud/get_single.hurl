GET http://localhost:8080/tasks/3

HTTP 200

[Asserts]
jsonpath "$.id" == "3"
jsonpath "$.task" == "Deploy update"
jsonpath "$.status" == "completed"