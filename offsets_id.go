// Code generated by 'gen-objects'; DO NOT EDIT
package klev

type OffsetID string

func ParseOffsetID(id string) (OffsetID, error) {
	if err := validate(id, "off"); err != nil {
		return OffsetID(""), err
	}
	return OffsetID(id), nil
}
