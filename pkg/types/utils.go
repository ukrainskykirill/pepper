package types

import "encoding/json"

type UnsetString struct {
	Value string
	Valid bool
	Set   bool
}

func (i *UnsetString) UnmarshalJSON(data []byte) error {
	i.Set = true
  
	if string(data) == "null" {
	  i.Valid = false
	  return nil
	}
	var temp string
	if err := json.Unmarshal(data, &temp); err != nil {
	  return err
	}
	i.Value = temp
	i.Valid = true
	return nil
  }
