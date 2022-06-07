package formula1_test

import (
	"os"
	"testing"

	f1E "github.com/TtMyth123/kit/ExcelFormula/engine"
	"github.com/tealeg/xlsx"
)

var xlFile *xlsx.File

func TestMain(m *testing.M) {
	if _, ferr := os.Stat("../testdocs/hcp-pricer.xlsx"); os.IsNotExist(ferr) {
		os.Exit(0)
		return
	}

	var err error
	xlFile, err = xlsx.OpenFile("../testdocs/hcp-pricer.xlsx")
	if err != nil {
		os.Exit(1)
	}

	code := m.Run()
	os.Exit(code)
}

func TestOne(t *testing.T) {
	var err error
	engine := f1E.NewEngine(xlFile)
	inputs := map[string]string{
		"Pricer!B4":  "18",
		"Pricer!B16": "Plan 1",
	}

	outputs := map[string]f1E.OutParam{}

	err = engine.Execute(inputs, &outputs)
	if err != nil {
		t.Errorf("Unexpected error %v", err)
		return
	}
	//if strings.HasPrefix((*outputs)["Pricer!B17"], "函数未实现") {
	//	t.Errorf("Expected: Numeric value\tActual: %s", (*outputs)["Pricer!B17"])
	//}
	t.Error("aaa", outputs)

	//if strings.HasPrefix(outputs["Pricer!B17"], "函数未实现") {
	//	t.Errorf("Expected: Numeric value\tActual: %s", (*outputs)["Pricer!B17"])
	//}
}
