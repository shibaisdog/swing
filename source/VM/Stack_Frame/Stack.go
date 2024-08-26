package Stack_Frame

type Stack struct {
	RUN_FILE string
	Stack    map[string]Each_File_Stack
}

func New_Stack() *Stack {
	return &Stack{
		Stack: make(map[string]Each_File_Stack),
	}
}

func (fs Stack) GET_FS(file_name string) Each_File_Stack {
	return fs.Stack[file_name]
}

func (fs *Stack) ADD_Stack(file_name string) {
	fs.RUN_FILE = file_name
	fs.Stack[file_name] = *INIT_Each_File_Stack()
}
