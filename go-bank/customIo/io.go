package customIo

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func WriteContent(value float64, filename string) error {
	valueString := fmt.Sprint(value)
	err := os.WriteFile(filename, []byte(valueString), 0644)
	if err != nil {
		return err
	}
	return nil

}
func ReadContent(filename string) (value float64, err error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Print("ERROR: ")
		err = errors.New("failed to locate file")
		return 1000, err
	}
	dataString := string(data)
	value, err = strconv.ParseFloat(dataString, 64)
	if err != nil {
		fmt.Print("ERROR: ")
		err = errors.New("error parsing value")
		return 1000, err
	}
	return value, nil

}
