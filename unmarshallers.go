package goiex

import (
	"encoding/json"
	"strconv"
)

func (c *Chart) UnmarshalJSON(b []byte) error {
	var charts []ChartDTO

	if err := json.Unmarshal(b, &charts); err != nil {
		return err
	}

	c.Charts = charts
	return nil
}

func (rds *RefDataSymbols) UnmarshalJSON(b []byte) error {
	var symbols []SymbolDTO

	if err := json.Unmarshal(b, &symbols); err != nil {
		return err
	}

	rds.Symbols = symbols
	return nil
}

func (s *SymbolDTO) UnmarshalJSON(b []byte) error {
	var symbolMap map[string]interface{}

	if err := json.Unmarshal(b, &symbolMap); err != nil {
		return err
	}

	switch id := symbolMap["iexId"].(type) {
	case string:
		if intId, err := strconv.Atoi(id); err != nil {
			return err
		} else {
			s.IexId = intId
		}
	case float32:
		s.IexId = int(id)
	case float64:
		s.IexId = int(id)
	default:
	}

	s.Symbol = symbolMap["symbol"].(string)
	s.Name = symbolMap["name"].(string)
	s.Date = symbolMap["date"].(string)
	s.IsEnabled = symbolMap["isEnabled"].(bool)
	s.Type = symbolMap["type"].(string)

	return nil
}
