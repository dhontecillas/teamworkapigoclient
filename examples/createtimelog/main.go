package main

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/dhontecillas/teamworkapigoclient/pkg/config"
	"github.com/dhontecillas/teamworkapigoclient/pkg/openapi/pmv1"
	"github.com/dhontecillas/teamworkapigoclient/pkg/openapi/projv3"
	"github.com/dhontecillas/teamworkapigoclient/pkg/twapi"
)

const explanation string = `

This example requires the following env vars to be set:

export TWAPICLIENT_APIKEY="twp_putHereTheApiKeyForYourUser"
export TWAPICLIENT_HOST="yoursubdomainname.teamwork.com"

`

func main() {
	cnf, err := config.LoadConf()
	if err != nil {
		fmt.Printf(explanation)
		fmt.Printf("\n\nError: %s\n", err.Error())
	}

	// create a new set of api client for teamwork projects
	cs := twapi.NewClientSet(cnf.Host, cnf.APIKey, context.Background())

	projID, err := createProject(cs)
	if err != nil {
		panic(err.Error())
	}
	/*
		_ = deleteBudget(cs, 3)
		_ = createRepeatingBudget(cs, 9, time.Date(2018, time.June, 1, 0, 0, 0, 0, time.UTC))
	*/
	_, err = createTask(cs, projID, "Task AUTO TEST", "This *is* an __autotest__")
	if err != nil {
		panic(err.Error())
	}

	tasklistID, err := createTaskList(cs, projID, "Tasklist AUTO TEST",
		"This *is* an __autotest__")
	if err != nil {
		panic(err.Error())
	}

	tid, err := createTaskInList(cs, tasklistID,
		fmt.Sprintf("AUTO TK %d-%d", tasklistID, 0),
		fmt.Sprintf("AUTO generated task **%d-%d**", tasklistID, 0))
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("Created task %d", tid)

	tid, err = createTimeLog(cs, tid, time.Now().AddDate(0, 0, -2))
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("Created timelog %d\n", tid)
}

func createProject(cs *twapi.ClientSet) (int64, error) {
	tm := time.Now().Format("20060102_150405")
	name := fmt.Sprintf("AUTO %s", tm)
	desc := fmt.Sprintf("This is an autogenerated project %s", tm)
	resp, _, err := cs.Pm.ProjectsApi.POSTProjectsJson(cs.PmCtx).Body(
		pmv1.InlineObject50{
			Project: pmv1.ProjectsJsonProject{
				Name:        name,
				Description: &desc,
			},
		}).Execute()

	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		return 0, err
	}
	projID, err := strconv.ParseInt(*resp.Id, 10, 64)
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		return 0, err
	}
	fmt.Printf("Created: %s (Id: %d)\n", name, projID)
	return projID, nil
}

func createTaskList(cs *twapi.ClientSet, projectID int64, name string, description string) (int64, error) {
	strProjectID := strconv.FormatInt(projectID, 10)
	resp, _, err := cs.Pm.TaskListsApi.POSTProjectsIdTasklistsJson(cs.PmCtx,
		strProjectID).Body(pmv1.InlineObject73{
		TodoList: &pmv1.ProjectsIdTasklistsJsonTodoList{
			Name:        &name,
			Description: &description,
		},
	}).Execute()
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		return 0, err
	}
	fmt.Printf("Success: %v\n", resp)
	var rID int64
	rID, err = strconv.ParseInt(*resp.TASKLISTID, 10, 64)
	if err != nil {
		return 0, err
	}
	return rID, nil
}

func createTaskInList(cs *twapi.ClientSet, tasklistID int64, name string, description string) (int64, error) {
	strTasklistID := strconv.FormatInt(tasklistID, 10)
	var repFreq = "norepeat"
	resp, _, err := cs.Pm.TasksApi.POSTTasklistsIdTasksJson(cs.PmCtx, strTasklistID).
		Body(pmv1.InlineObject91{
			TodoItem: pmv1.TasklistsIdTasksJsonTodoItem{
				Content:     name,
				Description: &description,
				RepeatOptions: &pmv1.TasklistsIdTasksJsonTodoItemRepeatOptions{
					RepeatsFreq: &repFreq,
				},
			},
		}).Execute()
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		return 0, err
	}
	fmt.Printf("Success: %v\n", resp)
	var rID int64
	rID, err = strconv.ParseInt(*resp.Id, 10, 64)
	if err != nil {
		return 0, err
	}
	return rID, nil
}

func createTask(cs *twapi.ClientSet, projectID int64, name string, description string) (int64, error) {
	strProjId := strconv.FormatInt(projectID, 10)
	resp, _, err := cs.Pm.TasksApi.POSTProjectsProjIdTasksJson(cs.PmCtx, strProjId).
		Body(pmv1.InlineObject77{
			TodoItem: &pmv1.ProjectsProjIdTasksJsonTodoItem{
				Content:     name,
				Description: &description,
			},
		}).Execute()
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		return 0, err
	}
	fmt.Printf("Created task: %v\n", resp)
	var rID int64
	rID, err = strconv.ParseInt(*resp.Id, 10, 64)
	if err != nil {
		return 0, err
	}
	return rID, nil
}

// we should create from 4 to 10 timelogs per day per user
func createTimeLog(cs *twapi.ClientSet, taskID int64, from time.Time) (int64, error) {
	desc := "Why do not have an id ?"
	hours := "1"
	isBillable := "true"
	minutes := "20"
	personId := "2"

	strTaskID := strconv.FormatInt(taskID, 10)
	resp, _, err := cs.Pm.TimeTrackingApi.POSTTasksIdTimeEntriesJson(cs.PmCtx, strTaskID).
		Body(pmv1.InlineObject101{
			TimeEntry: &pmv1.ProjectsIdTimeEntriesJsonTimeEntry{
				Date:        from.Format(twapi.V1_DATEFORMATSTR),
				Description: &desc,
				Hours:       hours,
				Isbillable:  &isBillable,
				Minutes:     minutes,
				PersonId:    &personId,
				Time:        from.Format(twapi.V1_TIMEFORMATSTR),
			},
		}).Execute()
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		return 0, err
	}
	fmt.Printf("Success: %v\n", resp)
	/*
		var rID int64
		rID, err = strconv.ParseInt(resp.Id, 10, 64)
		if err != nil {
			return 0, err
		}
		return rID, nil
	*/
	return int64(*resp.Id), nil
}

func createRepeatingBudget(cs *twapi.ClientSet, projectID int64, from time.Time) error {
	repeatUnit := "WEEK"
	isRepeating := true
	status := "ACTIVE"
	budgetType := "TIME"
	capacity := int32(60 * 8 * 5) // 40 weekly hours
	projID := int32(projectID)
	start := from.Format(twapi.TIMEFORMATSTR)
	end := from.AddDate(2, 0, 0).Format(twapi.TIMEFORMATSTR)
	all := "ALL"
	repeatPeriod := int32(1)
	resp, _, err := cs.V3.BudgetsApi.POSTProjectsApiV3BudgetsJson(cs.V3Ctx).BudgetRequest(
		projv3.BudgetRequest{
			Budget: &projv3.BudgetProjectBudget{
				ProjectId:     &projID,
				Capacity:      &capacity,
				StartDateTime: &start,
				EndDateTime:   &end,
				RepeatPeriod:  &repeatPeriod,
				RepeatUnit:    &repeatUnit,
				IsRepeating:   &isRepeating,
				Type:          &budgetType,
				Status:        &status,
				TimelogType:   &all,
			},
		}).Execute()
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		return err
	}
	fmt.Printf("Success: %v\n", resp)
	return nil
}

func deleteBudget(cs *twapi.ClientSet, budgetId int32) error {
	resp, err := cs.V3.BudgetsApi.DELETEProjectsApiV3BudgetsidJson(cs.V3Ctx, int32(budgetId)).Execute()
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		return err
	}
	fmt.Printf("Success: %v\n", resp)
	return nil
}