package enumeration

type void interface{}

type constant struct {
	Name        string //Name of the field
	Value       void   //Holds the value of the field
	Description string //Provides the description of the enum
}

type Enum struct {
	Name        string     //Name of the enum
	Description string     //Provides the description of the enum
	Fields      []constant //Holds the fields of the enum
	Source      string     //Holds the source of enum
}
