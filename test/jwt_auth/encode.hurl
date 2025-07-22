POST http://localhost:8080/auth/create
Content-Type: application/json

{
    "Id": "12345",
    "Email": "alice@example.com",
    "Roles": ["admin", "editor"],
    "Permissions": ["read", "write"],
    "Org_id": 1001
}

HTTP 200

[Asserts]
header "Content-Type" contains "application/json"
jsonpath "$.token" exists
jsonpath "$.token" startsWith "eyJ"