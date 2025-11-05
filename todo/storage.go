package todo

import "sync"

type Storage struct {
	mu     sync.Mutex
	tasks  []Task
	nextID int
}

func NewStorage() *Storage {
	return &Storage{
		tasks:  []Task{},
		nextID: 1,
	}
}

func (s *Storage) Add(title string) Task {
	s.mu.Lock()
	defer s.mu.Unlock()

	task := Task{
		ID:    s.nextID,
		Title: title,
	}
	s.nextID++
	s.tasks = append(s.tasks, task)
	return task
}

func (s *Storage) List() []Task {
	s.mu.Lock()
	defer s.mu.Unlock()
	return append([]Task{}, s.tasks...) // копия, чтобы не менять снаружи
}

func (s *Storage) ToggleDone(id int) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	for i, t := range s.tasks {
		if t.ID == id {
			s.tasks[i].Completed = !t.Completed
			return true
		}
	}
	return false
}
