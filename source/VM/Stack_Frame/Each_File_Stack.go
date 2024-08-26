package Stack_Frame

type Each_File_Stack struct {
	Doing_Func bool
	Memory     *Var_Stack
	Func       *Var_Stack
}

func INIT_Each_File_Stack() *Each_File_Stack {
	return &Each_File_Stack{
		Memory: INIT_Var_Stack(),
		Func:   INIT_Var_Stack(),
	}
}

func (fs Each_File_Stack) GET_STACK() *Var_Stack {
	if fs.Doing_Func {
		return fs.Func
	} else {
		return fs.Memory
	}
}
