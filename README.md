# undo-redo with command pattern

undo redo implementation in go

## Interface
~~~Go
type State interface{}

type Command interface {
	Commit() State
	Rollback() State
}

type History interface {
	State() State
	Save(command Command)
	Undo()
	Redo()
	Clear()
}
~~~
