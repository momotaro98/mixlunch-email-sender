package main

import (
	"fmt"
	"testing"
)

func TestParseDatetimeString_date_Success(t *testing.T) {
	// Arrange
	s := "2019-05-01"
	// Act
	act, err := ParseDatetimeString(s)
	// Assert
	if err != nil {
		t.Errorf("fail. err: %+v", err)
	}
	if d := fmt.Sprintf(act.Format("2006-01-02")); d != s {
		t.Errorf("fail. got=%s, expected=%s", d, s)
	}
}

func TestParseDatetimeString_datetime_Success(t *testing.T) {
	// Arrange
	expected := "2019-05-01"
	s := "2019-05-01T07:03:04Z"
	// Act
	act, err := ParseDatetimeString(s)
	// Assert
	if err != nil {
		t.Errorf("fail. err: %+v", err)
	}
	if d := fmt.Sprintf(act.Format("2006-01-02")); d != expected {
		t.Errorf("fail. got=%s, expected=%s", d, expected)
	}
}

func TestParseDatetimeString_InvalidString_Fail(t *testing.T) {
	// Arrange
	s := "2019/05/01T07:03:04Z"
	// Act
	_, err := ParseDatetimeString(s)
	// Assert
	if err == nil {
		t.Errorf("fail. got=%+v, expected=nil", err)
	}
}

func TestSquashDateString_ValidInput_Success(t *testing.T) {
	// Arrange
	s := "2019-05-01"
	// Act
	act, err := SquashDateString(s)
	// Assert
	if err != nil {
		t.Errorf("fail. There should not be error: %v", err)
	}
	if act != "20190501" {
		t.Errorf("fail. Actual result: %s", act)
	}
}

func TestSquashDateString_InvalidInput_Fail(t *testing.T) {
	// Arrange
	s := "I'm-your-father"
	// Act
	_, err := SquashDateString(s)
	// Assert
	if err == nil {
		t.Errorf("fail. There should be error.")
	}
}

func TestShiftTimeStringUtfToJst_ValidInput_Success(t *testing.T) {
	// Arrange
	s := "03:30"
	// Act
	act, err := ShiftTimeStringUtfToJst(s)
	// Assert
	if err != nil {
		t.Errorf("fail. There should not be error: %v", err)
	}
	if act != "12:30" {
		t.Errorf("fail. Actual result: %s", act)
	}
}

func TestShiftTimeStringUtfToJst_ValidNoneZeroInput_Success(t *testing.T) {
	// Arrange
	s := "2:00"
	// Act
	act, err := ShiftTimeStringUtfToJst(s)
	// Assert
	if err != nil {
		t.Errorf("fail. There should not be error: %v", err)
	}
	if act != "11:00" {
		t.Errorf("fail. Actual result: %s", act)
	}
}

func TestShiftTimeStringUtfToJst_InvalidInput_Success(t *testing.T) {
	// Arrange
	s := "0400"
	// Act
	_, err := ShiftTimeStringUtfToJst(s)
	// Assert
	if err == nil {
		t.Errorf("fail. There should be error.")
	}
}

func TestShiftTimeStringUtfToJst_InvalidInput2_Success(t *testing.T) {
	// Arrange
	s := "three:30"
	// Act
	_, err := ShiftTimeStringUtfToJst(s)
	// Assert
	if err == nil {
		t.Errorf("fail. There should be error.")
	}
}
