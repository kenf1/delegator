GET http://localhost:8080/auth/uncreate/{{token}}
Content-Type: application/json

HTTP 200

[Asserts]
jsonpath "$.claims.Id"           exists
jsonpath "$.claims.Email"        exists
jsonpath "$.claims.Roles"        exists
jsonpath "$.claims.Permissions"  exists
jsonpath "$.claims.Org_id"       exists