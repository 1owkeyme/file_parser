package reader_test

import (
	"file_parser/reader"
	"os"
	"reflect"
	"testing"
)

func TestFileReader_Read_ValidData(t *testing.T) {
	file, err := os.CreateTemp("", "testdata.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(file.Name())

	testData := "Ivan 5\nMaria 4\nIvan 3\nAlexey 5\nMaria 5\nAlexey 4\n"
	file.WriteString(testData)
	file.Close()

	fr := reader.NewFileReader(file.Name())
	data, err := fr.Read()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expectedData := map[string][]int{
		"Ivan":   {5, 3},
		"Maria":  {4, 5},
		"Alexey": {5, 4},
	}

	if !reflect.DeepEqual(data, expectedData) {
		t.Errorf("Expected %v, got %v", expectedData, data)
	}
}

func TestFileReader_Read_InvalidData(t *testing.T) {
	file, err := os.CreateTemp("", "invalid_testdata.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(file.Name())

	invalidData := "Ivan 5\nInvalidLine\nMaria\nAlexey 4\n"
	file.WriteString(invalidData)
	file.Close()

	fr := reader.NewFileReader(file.Name())
	data, err := fr.Read()
	if err == nil {
		t.Fatal("Expected an error, got nil")
	}

	if len(data) != 0 {
		t.Errorf("Expected empty data, got %v", data)
	}
}
