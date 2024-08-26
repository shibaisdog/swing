package Heap

type Type struct {
	INT_16   bool
	INT_32   bool
	INT_64   bool
	FLOAT_32 bool
	FLOAT_64 bool
	STRING   bool
	MAP      bool
	LIST     bool
	BOOL     bool
	FUNC     bool
	NULL     bool
}

type Variable struct {
	Name  string
	Type  Type
	Value interface{}
	Const bool
}

type ML_Variable struct {
	Type  Type
	Value interface{}
}
