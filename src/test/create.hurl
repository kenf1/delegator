POST http://localhost:8080/tasks
Content-Type: application/json

{
    "task": "Enter mainframe",
    "status": "queued"
}

HTTP 201

[Asserts]
jsonpath "$.id" != null
jsonpath "$.task" == "Enter mainframe"
jsonpath "$.status" == "queued"