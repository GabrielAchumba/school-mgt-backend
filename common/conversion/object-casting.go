package conversion

import (
	"encoding/json"
	networkingerrors "networkingAPIS/common/networkingerrors"
)

func Convert(source interface{}, destination interface{}) error {
	result, err := json.Marshal(source)
	if err != nil {

		networkingerrors.Error("message 1: " + err.Error())
		return networkingerrors.Error("Error in matching model.")
	}
	//////////er := json.NewDecoder(source).Decode(destination)

	err = json.Unmarshal([]byte(result), destination)
	if err != nil {
		networkingerrors.Error("message 2: " + err.Error())
		return networkingerrors.Error("Error in matching model.")
	}
	return nil
}
