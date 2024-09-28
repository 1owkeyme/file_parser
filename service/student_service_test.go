package service_test

import (
	"errors"
	"file_parser/service"
	"reflect"
	"testing"
)

type MockReader struct {
	data map[string][]int
	err  error
}

func (mr *MockReader) Read() (map[string][]int, error) {
	return mr.data, mr.err
}

type MockWriter struct {
	writtenData     map[string][]int
	writtenAverages map[string]float64
}

func (mw *MockWriter) Write(data map[string][]int, averages map[string]float64) {
	mw.writtenData = data
	mw.writtenAverages = averages
}

func TestCalculateAverages(t *testing.T) {
	mockData := map[string][]int{
		"Ivan":   {5, 4, 3},
		"Maria":  {4, 4, 5},
		"Alexey": {3, 4},
	}

	expectedAverages := map[string]float64{
		"Ivan":   4.00,
		"Maria":  4.33,
		"Alexey": 3.50,
	}

	studentService := service.NewStudentService(nil, nil)
	averages := studentService.CalculateAverages(mockData)

	if !reflect.DeepEqual(averages, expectedAverages) {
		t.Errorf("Expected %v, got %v", expectedAverages, averages)
	}
}

func TestProcess_Success(t *testing.T) {
	mockReader := &MockReader{
		data: map[string][]int{
			"Ivan":  {5, 4},
			"Maria": {3, 5},
		},
		err: nil,
	}

	mockWriter := &MockWriter{}

	studentService := service.NewStudentService(mockReader, mockWriter)

	studentService.Process()

	expectedData := map[string][]int{
		"Ivan":  {5, 4},
		"Maria": {3, 5},
	}

	expectedAverages := map[string]float64{
		"Ivan":  4.50,
		"Maria": 4.00,
	}

	if !reflect.DeepEqual(mockWriter.writtenData, expectedData) {
		t.Errorf("Expected data %v, got %v", expectedData, mockWriter.writtenData)
	}

	if !reflect.DeepEqual(mockWriter.writtenAverages, expectedAverages) {
		t.Errorf("Expected averages %v, got %v", expectedAverages, mockWriter.writtenAverages)
	}
}

func TestProcess_ReadError(t *testing.T) {
	mockReader := &MockReader{
		err: errors.New("read error"),
	}

	mockWriter := &MockWriter{}

	studentService := service.NewStudentService(mockReader, mockWriter)

	studentService.Process()

	if mockWriter.writtenData != nil || mockWriter.writtenAverages != nil {
		t.Errorf("Expected no data to be written, but got %v and %v", mockWriter.writtenData, mockWriter.writtenAverages)
	}
}
