package engine

import (
	"encoding/json"
	"fmt"
	"math"
	"os"
	"testing"

	f1Formula "github.com/TtMyth123/kit/ExcelFormula/formula"
	"github.com/tealeg/xlsx"
)

var EPSILON = math.Pow(10, -9)
var xlFile *xlsx.File

func TestMain(m *testing.M) {
	// setup
	var err error
	xlFile, err = xlsx.OpenFile("../testdocs/formula1-x1.xlsx")
	if err != nil {
		os.Exit(1)
	}

	code := m.Run()
	os.Exit(code)
}

func TestOutParamMarshalJSON(t *testing.T) {
	outParam := NewOutParam("string")

	outParam.Value = `10s`
	if serialized, err := json.Marshal(outParam); err != nil {
		t.Error(err)
	} else if string(serialized) != `"10s"` {
		t.Errorf("Expected: 10s\tActual: %s", serialized)
	}

	outParam.Value = []string{"10s", "20s"}
	if serialized, err := json.Marshal(outParam); err != nil {
		t.Error(err)
	} else if string(serialized) != `["10s","20s"]` {
		t.Errorf("Expected: 10s\tActual: %s", serialized)
	}

	outParam.Value = [][]string{{"10s", "20s"}, {"11s", "22s"}}
	if serialized, err := json.Marshal(outParam); err != nil {
		t.Error(err)
	} else if string(serialized) != `[["10s","20s"],["11s","22s"]]` {
		t.Errorf("Expected: 10s\tActual: %s", serialized)
	}
}

func TestRangeToSlice(t *testing.T) {
	var cellRange Range
	var slice []Cell
	var ok bool

	cellRange = Range{
		rowCount: 2,
		colCount: 1,
	}
	cellRange.cells = make([]Cell, 2)

	slice, ok = cellRange.ToSlice()
	if ok != true {
		t.Errorf("Unexpected error")
	}
	if result := len(slice); result != 2 {
		t.Errorf("Expected: 2\tActual: %v", result)
	}

	cellRange = Range{
		rowCount: 1,
		colCount: 2,
	}
	cellRange.cells = make([]Cell, 2)

	slice, ok = cellRange.ToSlice()
	if ok != true {
		t.Errorf("Unexpected error")
	}
	if result := len(slice); result != 2 {
		t.Errorf("Expected: 2\tActual: %v", result)
	}

	cellRange = Range{
		rowCount: 2,
		colCount: 2,
	}
	cellRange.cells = make([]Cell, 4)

	slice, ok = cellRange.ToSlice()
	if ok == true {
		t.Errorf("Expected error")
	}
}

func TestRangeTo2DSlice(t *testing.T) {
	var cellRange Range
	var slice [][]Cell

	cellRange = Range{
		rowCount: 2,
		colCount: 2,
	}
	cellRange.cells = make([]Cell, 4)

	slice, ok := cellRange.To2DSlice()
	if ok != true {
		t.Errorf("Unexpected error")
	}
	if result := len(slice); result != cellRange.rowCount {
		t.Errorf("Expected: %d\tActual: %v", cellRange.rowCount, result)
	}
	if result := len(slice[0]); result != cellRange.colCount {
		t.Errorf("Expected: %d\tActual: %v", cellRange.colCount, result)
	}
	if result := len(slice[1]); result != 2 {
		t.Errorf("Expected: %d\tActual: %v", cellRange.colCount, cellRange.colCount)
	}
}

func TestNumberLiteral(t *testing.T) {
	engine := NewEngine(xlFile)
	formula := f1Formula.NewFormula(`=10.1`)

	var result interface{}
	result, _ = engine.EvalFormula(formula)
	if result != 10.1 {
		t.Errorf("Expected: 10.1\tActual: %v", result)
	}
}

func TestInfixOperationsOf2Literal(t *testing.T) {
	var engine *Engine
	var formula *f1Formula.Formula
	var result interface{}

	engine = NewEngine(xlFile)
	formula = f1Formula.NewFormula(`=1.1 + 2.2`)
	result, _ = engine.EvalFormula(formula)
	if r, ok := result.(float64); !ok || (r-3.3) > EPSILON {
		t.Errorf("Expected: 3.3\tActual: %v", result)
	}

	engine = NewEngine(xlFile)
	formula = f1Formula.NewFormula(`=10 * 2`)
	result, _ = engine.EvalFormula(formula)
	if r, ok := result.(float64); !ok || (r-20) > EPSILON {
		t.Errorf("Expected: 20\tActual: %v", result)
	}

	engine = NewEngine(xlFile)
	formula = f1Formula.NewFormula(`=10 / 2`)
	result, _ = engine.EvalFormula(formula)
	if r, ok := result.(float64); !ok || (r-5) > EPSILON {
		t.Errorf("Expected: 5\tActual: %v", result)
	}
}

func TestAdditionOf3Literals(t *testing.T) {
	engine := NewEngine(xlFile)
	formula := f1Formula.NewFormula(`=1.1 + 2.2 + 10`)

	var result interface{}
	result, _ = engine.EvalFormula(formula)
	if r, ok := result.(float64); !ok || (r-13.3) > EPSILON {
		t.Errorf("Expected: 13.3\tActual: %v", result)
	}
}

func TestArithOfLiterals(t *testing.T) {
	var engine *Engine
	var formula *f1Formula.Formula
	var result interface{}
	engine = NewEngine(xlFile)
	formula = f1Formula.NewFormula(`=20 - 29 + 10`)
	result, _ = engine.EvalFormula(formula)
	if r, ok := result.(float64); !ok || (r-1) > EPSILON {
		t.Errorf("Expected: 0\tActual: %v", result)
	}

	engine = NewEngine(xlFile)
	formula = f1Formula.NewFormula(`=10 + 3 * 2`)
	result, _ = engine.EvalFormula(formula)
	if r, ok := result.(float64); !ok || (r-16) > EPSILON {
		t.Errorf("Expected: 16\tActual: %v", result)
	}

	engine = NewEngine(xlFile)
	formula = f1Formula.NewFormula(`=10 + 3 / 2`)
	result, _ = engine.EvalFormula(formula)
	if r, ok := result.(float64); !ok || (r-11.5) > EPSILON {
		t.Errorf("Expected: 11.5\tActual: %v", result)
	}

	engine = NewEngine(xlFile)
	formula = f1Formula.NewFormula(`=10 + 20 + 30 - 1`)
	result, _ = engine.EvalFormula(formula)
	if r, ok := result.(float64); !ok || (r-59) > EPSILON {
		t.Errorf("Expected: 0\tActual: %v", result)
	}

	engine = NewEngine(xlFile)
	formula = f1Formula.NewFormula(`=10 - 1 + 20 + 30`)
	result, _ = engine.EvalFormula(formula)
	if r, ok := result.(float64); !ok || (r-59) > EPSILON {
		t.Errorf("Expected: 0\tActual: %v", result)
	}

	engine = NewEngine(xlFile)
	formula = f1Formula.NewFormula(`=10 + 20 + 30 - 1 - 2`)
	result, _ = engine.EvalFormula(formula)
	if r, ok := result.(float64); !ok || (r-57) > EPSILON {
		t.Errorf("Expected: 0\tActual: %v", result)
	}

	engine = NewEngine(xlFile)
	formula = f1Formula.NewFormula(`=10 - 1 - 2 + 20 + 30`)
	result, _ = engine.EvalFormula(formula)
	if r, ok := result.(float64); !ok || (r-57) > EPSILON {
		t.Errorf("Expected: 0\tActual: %v", result)
	}

	engine = NewEngine(xlFile)
	formula = f1Formula.NewFormula(`=2 * (5 - 1)`)
	result, _ = engine.EvalFormula(formula)
	if r, ok := result.(float64); !ok || math.Abs(r-8) > EPSILON {
		t.Errorf("Expected: 8\tActual: %v", result)
	}

	engine = NewEngine(xlFile)
	formula = f1Formula.NewFormula(`=(5 - 1) * 2`)
	result, _ = engine.EvalFormula(formula)
	if r, ok := result.(float64); !ok || math.Abs(r-8) > EPSILON {
		t.Errorf("Expected: 8\tActual: %v", result)
	}

	engine = NewEngine(xlFile)
	formula = f1Formula.NewFormula(`=2 / (5 - 1)`)
	result, _ = engine.EvalFormula(formula)
	if r, ok := result.(float64); !ok || math.Abs(r-0.5) > EPSILON {
		t.Errorf("Expected: 0.5\tActual: %v", result)
	}

	engine = NewEngine(xlFile)
	formula = f1Formula.NewFormula(`=(5 - 1) / 2`)
	result, _ = engine.EvalFormula(formula)
	if r, ok := result.(float64); !ok || math.Abs(r-2) > EPSILON {
		t.Errorf("Expected: 2\tActual: %v", result)
	}
}

func TestLogicalOperators(t *testing.T) {
	var engine *Engine
	var formula *f1Formula.Formula
	var result interface{}

	engine = NewEngine(xlFile)
	formula = f1Formula.NewFormula(`=5 > 1`)
	result, _ = engine.EvalFormula(formula)
	if r, ok := result.(bool); !ok || r != true {
		t.Errorf("Expected: true\tActual: %v", result)
	}

	engine = NewEngine(xlFile)
	formula = f1Formula.NewFormula(`=5 = 5`)
	result, _ = engine.EvalFormula(formula)
	if r, ok := result.(bool); !ok || r != true {
		t.Errorf("Expected: true\tActual: %v", result)
	}

	engine = NewEngine(xlFile)
	formula = f1Formula.NewFormula(`="hello" = "hello"`)
	result, _ = engine.EvalFormula(formula)
	if r, ok := result.(bool); !ok || r != true {
		t.Errorf("Expected: true\tActual: %v", result)
	}

	engine = NewEngine(xlFile)
	formula = f1Formula.NewFormula(`=5 < 1`)
	result, _ = engine.EvalFormula(formula)
	if r, ok := result.(bool); !ok || r != false {
		t.Errorf("Expected: false\tActual: %v", result)
	}

	engine = NewEngine(xlFile)
	formula = f1Formula.NewFormula(`=5 <= 5`)
	result, _ = engine.EvalFormula(formula)
	if r, ok := result.(bool); !ok || r != true {
		t.Errorf("Expected: true\tActual: %v", result)
	}

	engine = NewEngine(xlFile)
	formula = f1Formula.NewFormula(`=5 >= 1`)
	result, _ = engine.EvalFormula(formula)
	if r, ok := result.(bool); !ok || r != true {
		t.Errorf("Expected: true\tActual: %v", result)
	}
}

func TestSimpleCellRef(t *testing.T) {
	var engine *Engine
	var formula *f1Formula.Formula
	var result interface{}

	engine = NewEngine(xlFile)
	formula = f1Formula.NewFormula(`=B2`)
	result, _ = engine.EvalFormula(formula)
	if r, ok := result.(float64); !ok || (r-10) > EPSILON {
		t.Errorf("Expected: 10\tActual: %v", result)
	}

	engine = NewEngine(xlFile)
	formula = f1Formula.NewFormula(`=Discounts!E2`)
	result, _ = engine.EvalFormula(formula)
	if r, ok := result.(string); !ok || r != "Cheap" {
		t.Errorf("Expected: 10\tActual: %v", result)
	}
}

func TestArithWithCellRef(t *testing.T) {
	var engine *Engine
	var formula *f1Formula.Formula
	var result interface{}

	engine = NewEngine(xlFile)
	formula = f1Formula.NewFormula(`=B2 + 1`) // 10 + 1
	result, _ = engine.EvalFormula(formula)
	if r, ok := result.(float64); !ok || (r-11) > EPSILON {
		t.Errorf("Expected: 11\tActual: %v", result)
	}

	engine = NewEngine(xlFile)
	formula = f1Formula.NewFormula(`=B2 + C2 + D2`) // 10 + 11 + 13
	result, _ = engine.EvalFormula(formula)
	if r, ok := result.(float64); !ok || (r-34) > EPSILON {
		t.Errorf("Expected: 34\tActual: %v", result)
	}
}

func TestIndirectCellRef(t *testing.T) {
	var engine *Engine
	var formula *f1Formula.Formula
	var result interface{}

	engine = NewEngine(xlFile)
	formula = f1Formula.NewFormula(`=Input!B3`) // B3=Discounts!E2
	result, _ = engine.EvalFormula(formula)
	if r, ok := result.(string); !ok || r != "Cheap" {
		t.Errorf("Expected: Cheap\tActual: %v", result)
	}
}

func TestFunOfLiteral(t *testing.T) {
	var engine *Engine
	var formula *f1Formula.Formula
	var result interface{}

	engine = NewEngine(xlFile)
	formula = f1Formula.NewFormula(`=FLOOR(10.1)`)
	result, _ = engine.EvalFormula(formula)
	if (result.(float64) - 10) > EPSILON {
		t.Errorf("Expected: 10\tActual: %v", result)
	}

	engine = NewEngine(xlFile)
	formula = f1Formula.NewFormula(`=SUM(2.1)`)
	result, _ = engine.EvalFormula(formula)
	if (result.(float64) - 2.1) > EPSILON {
		t.Errorf("Expected: 2.1\tActual: %v", result)
	}

	engine = NewEngine(xlFile)
	formula = f1Formula.NewFormula(`=SUM(10, 20)`)
	result, _ = engine.EvalFormula(formula)
	if (result.(float64) - 30) > EPSILON {
		t.Errorf("Expected: 30\tActual: %v", result)
	}

	engine = NewEngine(xlFile)
	formula = f1Formula.NewFormula(`=POWER(2, 3)`)
	result, _ = engine.EvalFormula(formula)
	if (result.(float64) - 8) > EPSILON {
		t.Errorf("Expected: 8\tActual: %v", result)
	}
}

func TestSumOfRefs(t *testing.T) {
	var engine *Engine
	var formula *f1Formula.Formula
	var result interface{}

	engine = NewEngine(xlFile)
	formula = f1Formula.NewFormula(`=SUM(Input!B2)`)
	result, _ = engine.EvalFormula(formula)
	if math.Abs(result.(float64)-10) > EPSILON {
		t.Errorf("Expected: 10\tActual: %v", result)
	}

	engine = NewEngine(xlFile)
	formula = f1Formula.NewFormula(`=SUM(Input!B2, Input!C2)`)
	result, _ = engine.EvalFormula(formula)
	if math.Abs(result.(float64)-21) > EPSILON {
		t.Errorf("Expected: 21\tActual: %v", result)
	}

	engine = NewEngine(xlFile)
	formula = f1Formula.NewFormula(`=SUM(Input!B2:D2)`)
	result, _ = engine.EvalFormula(formula)
	if math.Abs(result.(float64)-34) > EPSILON {
		t.Errorf("Expected: 34\tActual: %v", result)
	}

	engine = NewEngine(xlFile)
	formula = f1Formula.NewFormula(`=SUM(Discounts!A2:B6)`)
	result, _ = engine.EvalFormula(formula)
	if math.Abs(result.(float64)-31.9) > EPSILON {
		t.Errorf("Expected: 31.9\tActual: %v", result)
	}
}

func TestSimpleIf(t *testing.T) {
	var engine *Engine
	var formula *f1Formula.Formula
	var result interface{}

	engine = NewEngine(xlFile)
	formula = f1Formula.NewFormula(`=IF(TRUE(), 2)`)
	result, _ = engine.EvalFormula(formula)
	if math.Abs(result.(float64)-2) > EPSILON {
		t.Errorf("Expected: 2\tActual: %v", result)
	}

	engine = NewEngine(xlFile)
	formula = f1Formula.NewFormula(`=IF(FALSE(), 2)`)
	result, _ = engine.EvalFormula(formula)
	if result.(bool) != false {
		t.Errorf("Expected: false\tActual: %v", result)
	}

	engine = NewEngine(xlFile)
	formula = f1Formula.NewFormula(`=IF(TRUE(), 2, 5)`)
	result, _ = engine.EvalFormula(formula)
	if math.Abs(result.(float64)-2) > EPSILON {
		t.Errorf("Expected: 2\tActual: %v", result)
	}

	engine = NewEngine(xlFile)
	formula = f1Formula.NewFormula(`=IF(FALSE(), 3, 6)`)
	result, _ = engine.EvalFormula(formula)
	if math.Abs(result.(float64)-6) > EPSILON {
		t.Errorf("Expected: 6\tActual: %v", result)
	}
}

func TestArithIf(t *testing.T) {
	var engine *Engine
	var formula *f1Formula.Formula
	var result interface{}

	engine = NewEngine(xlFile)
	formula = f1Formula.NewFormula(`=IF(1, 2)`)
	result, _ = engine.EvalFormula(formula)
	if math.Abs(result.(float64)-2) > EPSILON {
		t.Errorf("Expected: 2\tActual: %v", result)
	}

	engine = NewEngine(xlFile)
	formula = f1Formula.NewFormula(`=IF(0, 2)`)
	result, _ = engine.EvalFormula(formula)
	if result.(bool) != false {
		t.Errorf("Expected: 2\tActual: %v", result)
	}
}

func TestNestedIf(t *testing.T) {
	var engine *Engine
	var formula *f1Formula.Formula
	var result interface{}

	engine = NewEngine(xlFile)
	formula = f1Formula.NewFormula(`=IF(IF(TRUE(), 0, 2), 20, 10)`)
	result, _ = engine.EvalFormula(formula)
	if math.Abs(result.(float64)-10) > EPSILON {
		t.Errorf("Expected: 10\tActual: %v", result)
	}

	engine = NewEngine(xlFile)
	formula = f1Formula.NewFormula(`=IF(1, IF(TRUE(), 3, 20), 30)`)
	result, _ = engine.EvalFormula(formula)
	if math.Abs(result.(float64)-3) > EPSILON {
		t.Errorf("Expected: 3\tActual: %v", result)
	}

	engine = NewEngine(xlFile)
	formula = f1Formula.NewFormula(`=IF(0, 1, IF(TRUE(), 3, 20))`)
	result, _ = engine.EvalFormula(formula)
	if math.Abs(result.(float64)-3) > EPSILON {
		t.Errorf("Expected: 3\tActual: %v", result)
	}
}

func TestBooleanIf(t *testing.T) {
	var engine *Engine
	var formula *f1Formula.Formula
	var result interface{}

	engine = NewEngine(xlFile)
	formula = f1Formula.NewFormula(`=IF(OR(1, 0), TRUE())`)
	result, _ = engine.EvalFormula(formula)
	if result.(bool) != true {
		t.Errorf("Expected: 2\tActual: %v", result)
	}

	engine = NewEngine(xlFile)
	formula = f1Formula.NewFormula(`=IF(AND(1, 3), TRUE())`)
	result, _ = engine.EvalFormula(formula)
	if result.(bool) != true {
		t.Errorf("Expected: 2\tActual: %v", result)
	}
}

func TestAdvancedFunctions(t *testing.T) {
	var engine *Engine
	var formula *f1Formula.Formula
	var result interface{}

	engine = NewEngine(xlFile)
	formula = f1Formula.NewFormula(`=VLOOKUP(3, Discounts!A2:B6, 2, 0)`)
	result, _ = engine.EvalFormula(formula)
	if math.Abs(result.(float64)-2.5) > EPSILON {
		t.Errorf("Expected: 2.5\tActual: %v", result)
	}
}

// func TestActualPricer(t *testing.T) {
// 	localFile, _ := xlsx.OpenFile("../testdocs/dup.xlsx")
// 	var engine *Engine
// 	var formula *f1Formula.Formula

// 	engine = NewEngine(localFile)
// 	engine.activeSheet = localFile.Sheet["Input"]
// 	formula = f1Formula.NewFormula(`=IF(OR(CalculatorNB!$B$12="Decline",CalculatorNB!$B$12="Refer"),CalculatorNB!$B$12,SUM(CalculatorNB!E37:E40))`)
// 	engine.EvalFormula(formula)
// 	if insp := engine.Inspect(); insp["stackHeight"] != "0" {
// 		t.Errorf("Expected: 0\tActual: %s", insp["stackHeight"])
// 	}
// }

func TestExecute(t *testing.T) {
	localFile, _ := xlsx.OpenFile("../testdocs/dup.xlsx")
	var engine *Engine
	var err error
	var inputs map[string]string
	var outputs *map[string]OutParam

	engine = NewEngine(localFile)
	inputs = map[string]string{
		"Input!E18": "1000000.0",
		"Input!E20": "45000.0",
	}

	outputs = &map[string]OutParam{
		"Input!E35": NewOutParam("string"),
	}

	err = engine.Execute(inputs, outputs)
	if err != nil {
		t.Errorf("Unexpected error %v", err)
	}
	if (*outputs)["Input!E35"].Value == "" {
		t.Errorf("Expected: Non-empty\tActual: %s", (*outputs)["Input!E68"])
	}

	engine = NewEngine(localFile)
	inputs = map[string]string{
		"Input!E18": "2000000.0",
		"Input!E20": "0.0",
	}

	outputs = &map[string]OutParam{
		"Input!E35": NewOutParam("string"),
	}

	err = engine.Execute(inputs, outputs)
	if err != nil {
		t.Errorf("Unexpected error %v", err)
	}
	if (*outputs)["Input!E35"].Value == "" {
		t.Errorf("Expected: Non-empty\tActual: %s", (*outputs)["Input!E68"])
	}

	if actual := len(engine.formulaCache); actual == 0 {
		t.Errorf("Expected: len > 0\tActual: %v\n", actual)
	}
}

func TestExecuteRangeRef(t *testing.T) {
	localFile, _ := xlsx.OpenFile("../testdocs/formula1-x1.xlsx")
	var engine *Engine
	var err error
	var inputs map[string]string
	var outputs *map[string]OutParam

	engine = NewEngine(localFile)
	inputs = map[string]string{
		"Input!A1": "1000000.0",
	}

	outputs = &map[string]OutParam{
		"Discounts!A2:B6": NewOutParam("$ref"),
		"Discounts!D2:D4": NewOutParam("$ref"),
	}

	err = engine.Execute(inputs, outputs)
	if err != nil {
		t.Errorf("Unexpected error %v", err)
	}
	if (*outputs)["Discounts!A2:B6"].Value == nil {
		t.Errorf("Expected: Non-nil")
	}
	if (*outputs)["Discounts!D2:D4"].Value == nil {
		t.Errorf("Expected: Non-nil")
	}

	if values, ok := (*outputs)["Discounts!A2:B6"].Value.([][]string); ok {
		if len(values[0][0]) == 0 {
			t.Errorf("Expected: Non-empty\tActual: Empty")
		} else {
			t.Logf("Value: %v\n", values)
		}
	} else {
		t.Errorf("Expected: 2D Array")
	}

	if values, ok := (*outputs)["Discounts!D2:D4"].Value.([]string); ok {
		if len(values[0]) == 0 {
			t.Errorf("Expected: Non-empty\tActual: Empty")
		} else {
			t.Logf("Value: %v\n", values)
		}
	} else {
		t.Errorf("Expected: 1D Array")
	}
}

func TestArithOfLiterals2(t *testing.T) {
	var engine *Engine
	var formula *f1Formula.Formula
	var result interface{}
	engine = NewEngine(xlFile)
	//formula = f1Formula.NewFormula(`=100+300-200+500`)
	//formula = f1Formula.NewFormula(`=100+300-200`)
	formula = f1Formula.NewFormula(`=100+(200/20)`)

	//formula = f1Formula.NewFormula(`=100+300-200+500`)
	//formula = f1Formula.NewFormula(`=100+POWER(2, 3)-200+500`)
	//formula = f1Formula.NewFormula(`=100+300*200+500`)
	//formula = f1Formula.NewFormula(`=100+(300*200-500)`)

	//formula = f1Formula.NewFormula(`=10+3*20*5`)
	//formula = f1Formula.NewFormula(`=10+3*20/10`)
	//formula = f1Formula.NewFormula(`=10*3+60/10`)
	result, _ = engine.EvalFormula(formula)

	//r := float64(0)
	//if r, ok := result.(float64); !ok || math.Abs(r-(-100)) > EPSILON {
	//	t.Errorf("Expected: 0\tActual: %v", result)
	//}

	fmt.Println("result:", result, "r:", 10+3*20/10)

	//engine = NewEngine(xlFile)
	//formula = f1Formula.NewFormula(`=20 - 29 + 10`)
	//result, _ = engine.EvalFormula(formula)
	//if r, ok := result.(float64); !ok || (r-1) > EPSILON {
	//	t.Errorf("Expected: 0\tActual: %v", result)
	//}
	//
	//engine = NewEngine(xlFile)
	//formula = f1Formula.NewFormula(`=10 + 3 * 2`)
	//result, _ = engine.EvalFormula(formula)
	//if r, ok := result.(float64); !ok || (r-16) > EPSILON {
	//	t.Errorf("Expected: 16\tActual: %v", result)
	//}
	//
	//engine = NewEngine(xlFile)
	//formula = f1Formula.NewFormula(`=10 + 3 / 2`)
	//result, _ = engine.EvalFormula(formula)
	//if r, ok := result.(float64); !ok || (r-11.5) > EPSILON {
	//	t.Errorf("Expected: 11.5\tActual: %v", result)
	//}
	//
	//engine = NewEngine(xlFile)
	//formula = f1Formula.NewFormula(`=10 + 20 + 30 - 1`)
	//result, _ = engine.EvalFormula(formula)
	//if r, ok := result.(float64); !ok || (r-59) > EPSILON {
	//	t.Errorf("Expected: 0\tActual: %v", result)
	//}
	//
	//engine = NewEngine(xlFile)
	//formula = f1Formula.NewFormula(`=10 - 1 + 20 + 30`)
	//result, _ = engine.EvalFormula(formula)
	//if r, ok := result.(float64); !ok || (r-59) > EPSILON {
	//	t.Errorf("Expected: 0\tActual: %v", result)
	//}
	//
	//engine = NewEngine(xlFile)
	//formula = f1Formula.NewFormula(`=10 + 20 + 30 - 1 - 2`)
	//result, _ = engine.EvalFormula(formula)
	//if r, ok := result.(float64); !ok || (r-57) > EPSILON {
	//	t.Errorf("Expected: 0\tActual: %v", result)
	//}
	//
	//engine = NewEngine(xlFile)
	//formula = f1Formula.NewFormula(`=10 - 1 - 2 + 20 + 30`)
	//result, _ = engine.EvalFormula(formula)
	//if r, ok := result.(float64); !ok || (r-57) > EPSILON {
	//	t.Errorf("Expected: 0\tActual: %v", result)
	//}
	//
	//engine = NewEngine(xlFile)
	//formula = f1Formula.NewFormula(`=2 * (5 - 1)`)
	//result, _ = engine.EvalFormula(formula)
	//if r, ok := result.(float64); !ok || math.Abs(r-8) > EPSILON {
	//	t.Errorf("Expected: 8\tActual: %v", result)
	//}
	//
	//engine = NewEngine(xlFile)
	//formula = f1Formula.NewFormula(`=(5 - 1) * 2`)
	//result, _ = engine.EvalFormula(formula)
	//if r, ok := result.(float64); !ok || math.Abs(r-8) > EPSILON {
	//	t.Errorf("Expected: 8\tActual: %v", result)
	//}
	//
	//engine = NewEngine(xlFile)
	//formula = f1Formula.NewFormula(`=2 / (5 - 1)`)
	//result, _ = engine.EvalFormula(formula)
	//if r, ok := result.(float64); !ok || math.Abs(r-0.5) > EPSILON {
	//	t.Errorf("Expected: 0.5\tActual: %v", result)
	//}
	//
	//engine = NewEngine(xlFile)
	//formula = f1Formula.NewFormula(`=(5 - 1) / 2`)
	//result, _ = engine.EvalFormula(formula)
	//if r, ok := result.(float64); !ok || math.Abs(r-2) > EPSILON {
	//	t.Errorf("Expected: 2\tActual: %v", result)
	//}
}
