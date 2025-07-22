package models

type UserInfo struct {
	Id          string
	Email       string
	Roles       []string
	Permissions []string
	Org_id      int
}
