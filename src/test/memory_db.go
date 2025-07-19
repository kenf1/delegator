package test

import "github.com/kenf1/delegator/src/models"

var Tasks = []models.TaskDBRow{
	{Id: "1", Task: "Spin up service", Status: "running"},
	{Id: "2", Task: "Connect mainframe", Status: "queued"},
	{Id: "3", Task: "Deploy update", Status: "completed"},
	{Id: "4", Task: "Backup database", Status: "failed"},
	{Id: "5", Task: "Monitor logs", Status: "running"},
	{Id: "10", Task: "Run tests", Status: "queued"},
	{Id: "20", Task: "Update documentation", Status: "completed"},
	{Id: "30", Task: "Restart application", Status: "failed"},
	{Id: "40", Task: "Optimize performance", Status: "running"},
	{Id: "50", Task: "Schedule maintenance", Status: "queued"},
}
