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

func (rdca *RefDataCorporateActions) UnmarshalJSON(b []byte) error {
	var corporateActions []CorporateActionDTO

	if err := json.Unmarshal(b, &corporateActions); err != nil {
		return err
	}

	rdca.CorporateActions = corporateActions
	return nil
}

func (rdd *RefDataDividends) UnmarshalJSON(b []byte) error {
	var dividends []DividendDTO

	if err := json.Unmarshal(b, &dividends); err != nil {
		return err
	}

	rdd.Dividends = dividends
	return nil
}

func (rdd *RefDataNextDayExDates) UnmarshalJSON(b []byte) error {
	var nextDayExDates []NextDayExDateDTO

	if err := json.Unmarshal(b, &nextDayExDates); err != nil {
		return err
	}

	rdd.NextDayExDates = nextDayExDates
	return nil
}

func (rdsd *RefDataSymbolDirectories) UnmarshalJSON(b []byte) error {
	var symbolDirectories []SymbolDirectoryDTO

	if err := json.Unmarshal(b, &symbolDirectories); err != nil {
		return err
	}

	rdsd.SymbolDirectories = symbolDirectories
	return nil
}

func (s *SymbolDTO) UnmarshalJSON(b []byte) error {
	var symbolMap map[string]interface{}

	if err := json.Unmarshal(b, &symbolMap); err != nil {
		return err
	}

	switch id := symbolMap["iexId"].(type) {
	case string:
		if intID, err := strconv.Atoi(id); err != nil {
			return err
		} else {
			s.IexId = intID
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
