package goiex

import (
	"encoding/json"
)

func (c *Chart) UnmarshalJSON(b []byte) error {
	var charts []chart

	if err := json.Unmarshal(b, &charts); err != nil {
		return err
	}

	c.Charts = charts
	return nil
}
