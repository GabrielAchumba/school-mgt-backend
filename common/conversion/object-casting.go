package conversion

import (
	"encoding/json"

	errors "github.com/GabrielAchumba/school-mgt-backend/common/errors"
)

func Convert(source interface{}, destination interface{}) error {
	result, err := json.Marshal(source)
	if err != nil {

		errors.Error("message 1: " + err.Error())
		return errors.Error("Error in matching model.")
	}
	//////////er := json.NewDecoder(source).Decode(destination)

	err = json.Unmarshal([]byte(result), destination)
	if err != nil {
		errors.Error("message 2: " + err.Error())
		return errors.Error("Error in matching model.")
	}
	return nil
}
