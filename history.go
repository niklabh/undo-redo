package history

// State can be anything
type State interface{}

// Command interface with Commit and Rollback method
type Command interface {
	Commit() State
	Rollback() State
}

// History interface with functionality of the undo.
type History interface {
	State() State
	Save(command Command)
	Undo()
	Redo()
	Clear()
}

// HistImpl type contains the history of the changing states.
type HistImpl struct {
	undos   *Stack
	current State
	redos   *Stack
	limit   int
}

// NewHistory creates new history with the undo/redo capacity of `limit`
func NewHistory(limit int) History {
	undos := NewStack(limit)
	redos := NewStack(limit)
	h := &HistImpl{
		undos:   undos,
		current: nil,
		redos:   redos,
		limit:   limit,
	}
	return h
}

// State returns current state of the history.
// Notice that initial state is nil
func (h *HistImpl) State() State {
	return h.current
}

// Save method changes the current state and saves previous state for undo purposes.
func (h *HistImpl) Save(command Command) {
	h.undos.Push(command)
	h.current = command.Commit()
	h.redos.Clear()
}

// Undo method undos the state.
func (h *HistImpl) Undo() {
	command := h.undos.Pop()
	if command == nil {
		// no undo avaliable
		return
	}
	h.redos.Push(command)
	h.current = command.Rollback()
}

// Redo method redos the state.
func (h *HistImpl) Redo() {
	command := h.redos.Pop()
	if command == nil {
		return
	}
	h.undos.Push(command)
	h.current = command.Commit()
}

// Clear method clears history
func (h *HistImpl) Clear() {
	h.undos.Clear()
	h.redos.Clear()
}
