package game

import "testing"

func TestNewGrid(t *testing.T) {
	expect := newGrid()

	if expect != nil {
		t.Logf("Created grid successfully.")
	} else {
		t.Errorf("Failed to create grid.")
	}
}

func TestIsCellBlank(t *testing.T) {
	c := Coordinate{row: 0, column: 0}
	testGrid := newGrid()

	// test isCellBlank() expect return true
	expect := testGrid.isCellBlank(c)
	if expect != true {
		t.Errorf("Expected true - actual: %v", expect)
	}

	RenderGrid(testGrid)
	// test isCellBlank() expect return false
	c.row = 1
	c.column = 1

	t.Logf("change row: %d and column: %d", c.row, c.column)

	_, err := testGrid.insertAtCoordinate(c, x)
	if err != nil {
		t.Errorf("could not insert valuet in test grid")
	}

	RenderGrid(testGrid)
	expect = testGrid.isCellBlank(c)
	if expect != false {
		t.Errorf("isCellBlank() expected false - actual: %v", expect)
	}

}

func TestInsertAtCoordinate(t *testing.T) {
	c := Coordinate{row: 0, column: 0}
	testGrid := newGrid()

	// test (0,0) valid insert
	ok, err := testGrid.insertAtCoordinate(c, x)
	if err != nil {
		t.Errorf("could not insert valuet in test grid")
	}
	if ok {
		t.Logf("Entry made")
	}

	// test isCellBlank() expect return false
	c.row = 1
	c.column = 1

	ok, err = testGrid.insertAtCoordinate(c, x)
	if err != nil {
		t.Errorf("could not insert valuet in test grid")
	}
	if !ok {
		t.Logf("Did not make entry")
	}
}

func TestGetCellValue(t *testing.T) {
	testGrid := newGrid()
	c := Coordinate{row: 0, column: 0}

	// test when value is present
	testGrid.insertAtCoordinate(c, x)
	expect, err := testGrid.getCellValue(c)
	if err != nil {
		t.Errorf("Failed to get cell value")
	}
	if expect != x {
		t.Errorf("Value obtained does not match expected")
	} else {
		t.Logf("successfully got value from coordinate")
	}

	// test when value is blank
	c.row = 1
	c.column = 1
	expect, err = testGrid.getCellValue(c)
	if err != nil {
		t.Errorf("Failed to get cell value")
	}
	if expect != 0 {
		t.Errorf("Expected blank value at coordinate - actual received non zero value")
	} else {
		t.Logf("Cell is blank as expected")
	}

	// test when coordinate is out of bounds
	c.row = 4
	c.column = -5
	expect, err = testGrid.getCellValue(c)
	if err == nil {
		t.Logf(err.Error())
		if expect == 0 {
			t.Logf("Expected blank out of bound error")
		}
	}
}
