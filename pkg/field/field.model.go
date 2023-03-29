package field

type void interface{}

type Field struct {
	Name        string //Field name.
	Value       void   //Holds the value of field of any type. enum type only has values.
	CType       string //Data typ of field represented in c language Ex: int -> __s32, unsigned int -> __u32
	GoType      string //Data type of field represented in go language Ex: int -> int32 or int64
	Source      string //Holds the source line from where other field information was extracted.
	Description string //provides infomation about field like what it is and what it does.
	IsBitField  bool   //True -> if field is a bit field, False -> if not a bit field.
	BitLength   int    //This field is set if the IsBitField <- True. It holds number of bits the field might hold.
	IsPointer   bool   //True -> if field is a pointer, False -> if not a pointer
	IsArray     bool   //True -> if field is a array, False -> if not an array
	Length      int    //This field is set if the IsArray <- True. It hold size or length of the varibale.
}
