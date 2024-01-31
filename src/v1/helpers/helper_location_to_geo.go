package helpers

import (
	"encoding/json"
	"fmt"
	"os/exec"
)

func GetGeolocation(city string) (map[string]float64, error) {
	cmd := exec.Command("python3", "src/v1/helpers/geolocation_conversor/conversor.py", city)
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	var geo map[string]float64
	err = json.Unmarshal(output, &geo)
	if err != nil {
		return nil, err
	}

	if _, err := geo["error"]; err {
		return nil, fmt.Errorf("City not found\n")
	}

	return geo, nil
}
