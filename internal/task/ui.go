/*
Copyright © 2024 Alec Carpenter
*/
package task

import (
	"io"
	"slices"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
)

func drawTasks(w io.Writer, t *Tasks) {
	tbl := table.NewWriter()
	tbl.SetOutputMirror(w)
	tbl.AppendHeader(table.Row{"Timestamp", "ID", "Task", "Group"})
	for _, taskObject := range t.Task {
		tbl.AppendRow([]any{taskObject.Timestamp, taskObject.ID, taskObject.Item, taskObject.Group})
	}
	styledTable := styleTable(tbl)
	styledTable.Render()
}

func drawTasksGroup(w io.Writer, t *Tasks, name string) {
	tbl := table.NewWriter()
	tbl.SetOutputMirror(w)
	tbl.AppendHeader(table.Row{"Timestamp", "ID", "Task"})
	var taskIDs []int
	for _, group := range t.Groups {
		if name == group.Name {
			taskIDs = group.TaskIDs
		}
	}
	for _, taskObject := range t.Task {
		if slices.Contains(taskIDs, taskObject.ID) {
			tbl.AppendRow([]any{taskObject.Timestamp, taskObject.ID, taskObject.Item})
		}
	}
	styledTable := styleTable(tbl)
	styledTable.Render()
}

func drawListGroups(w io.Writer, t *Tasks) {
	tbl := table.NewWriter()
	tbl.SetOutputMirror(w)
	tbl.AppendHeader(table.Row{"Group", "# of Tasks"})
	for _, groupObject := range t.Groups {
		numberOfTasks := len(groupObject.TaskIDs)
		tbl.AppendRow([]any{groupObject.Name, numberOfTasks})
	}
	styledTable := styleTable(tbl)
	styledTable.Render()
}

func styleTable(tbl table.Writer) table.Writer {
	tbl.SetStyle(table.StyleBold)
	tbl.Style().Color.Header = text.Colors{text.BgBlack, text.FgRed}
	tbl.Style().Color.Row = text.Colors{text.BgBlack, text.FgYellow}
	tbl.Style().Color.RowAlternate = text.Colors{text.BgHiBlack, text.FgWhite}
	tbl.Style().Color.Separator = text.Colors{text.BgHiBlack, text.FgGreen}
	tbl.Style().Color.Border = text.Colors{text.BgHiBlack, text.FgGreen}
	return tbl
}
