PATCH http://localhost:8080/tasks/patch
Content-Type: application/json

{
  "id": "5",
  "task": "Patched task name"
}

HTTP 200

[Asserts]
jsonpath "$.id" == "5"
jsonpath "$.task" == "Patched task name"
jsonpath "$.status" == "running"