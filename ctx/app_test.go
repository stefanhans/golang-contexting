package ctx_test

import (
	"fmt"
	. "github.com/stefanhans/golang-contexting/ctx"
	"testing"
)

var appDataTestTable_1 = []struct {
	appDataType    AppDataType
	strAppDataType string
}{
	{APP_DATA_TYPE_RZV, "APP_DATA_TYPE_RZV"},
	{AppDataType(255), "APP_DATA_TYPE_UNDEFINED"},
}

var appDataTestTable_2 = []struct {
	appDataArray []byte
}{
	{[]byte{0, 1, 2, 3}},
}

func TestApp(t *testing.T) {

	cip := CreateCip()

	for i, app := range appDataTestTable_1 {
		cip.SetAppData(app.appDataType, CIP_APP_ARRAY_RZV)
		appDataType, _ := cip.AppData()
		s := fmt.Sprintf("%s", appDataType)
		if s != app.strAppDataType {
			t.Errorf("%d: Value != Expected: %s != %s\n", i, s, app.strAppDataType)
		}
	}

	for i, app := range appDataTestTable_2 {

		var ada CipAppArray = CIP_APP_ARRAY_RZV
		ada[0] = byte(len(app.appDataArray))
		for i := 1; i <= len(app.appDataArray); i++ {
			ada[i] = app.appDataArray[i-1]
		}
		cip.SetAppData(APP_DATA_TYPE_RZV, ada)
		_, appDataArray := cip.AppData()
		s1 := fmt.Sprintf("%s", app.appDataArray)
		s2 := fmt.Sprintf("%s", appDataArray)
		if s1 != s2 {
			t.Errorf("%d: Value != Expected: %s != %s\n", i, s1, s2)
		}
	}
}
