package commons

/*
Organization bla bla bla
*/
type Organization struct {
	Name string
	ID   string
}

// IsValid bla bla bla
func (o Organization) IsValid() bool {
	return len(o.Name) > 4
}

const (
	typeText  = iota
	typeInt   = iota
	typeImage = iota
)

type typeField struct {
	Name      string
	FieldType int
	IsID      bool

	// typeText modifiers
	minLength int
	maxLength int
	format    string

	// typeInt modifiers
	minIntValue int
	maxIntValue int
}

func (f typeField) isValid() bool {
	ft := f.FieldType
	return (ft == typeImage || ft == typeText || ft == typeInt) &&
		len(f.Name) > 0
}

/*
Project bla bla bla
*/
type Project struct {
	Name   string
	ID     string
	Fields []typeField
}

func (p Project) isValid() bool {
	return len(p.Name) > 4 && p.fieldsAreValid()
}

func (p Project) fieldsAreValid() bool {
	hasID := false
	for _, f := range p.Fields {
		if f.IsID {
			if hasID {
				return false
			}
			hasID = true
		}

		if !f.isValid() {
			return false
		}

	}

	return hasID
}

/*
Credential bla bla bla
*/
type Credential struct {
	ExternalID string
	ID         string
	Fields     map[string]interface{}
}

// User bla bla bla
type User struct {
	Email string
	ID    string
}
