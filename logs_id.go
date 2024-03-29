// Code generated by 'gen-objects'; DO NOT EDIT
package klev

type LogID struct{ string }

var nilLogID = LogID{}

func ParseLogID(id string) (LogID, error) {
	if err := validate(id, "log"); err != nil {
		return nilLogID, err
	}
	return LogID{id}, nil
}

func (id LogID) String() string {
	return id.string
}

func (id LogID) IDValue() string {
	return id.string
}

func (id LogID) MarshalText() ([]byte, error) {
	return []byte(id.string), nil
}

func (id *LogID) UnmarshalText(text []byte) error {
	str := string(text)
	if len(str) == 0 {
		return nil
	}

	if err := validate(str, "log"); err != nil {
		return err
	}

	id.string = str
	return nil
}
