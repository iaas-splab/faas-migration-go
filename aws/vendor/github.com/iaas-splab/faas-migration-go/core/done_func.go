package core

import "time"

func Done(req IDRequest, repo Repository) (*ToDoItem, error) {
	if err := repo.Init(); err != nil {
		return nil, err
	}

	item, err := repo.Get(req.ID)
	if err != nil {
		return nil, err
	}

	item.Done = true
	item.DoneTimestamp = time.Now().Unix()

	return repo.Put(item)
}
