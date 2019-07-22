package core

type IDRequest struct {
	ID string `json:"id"`
}

func Get(req IDRequest, repo Repository) (*ToDoItem, error) {
	if err := repo.Init(); err != nil {
		return nil, err
	}

	return repo.Get(req.ID)
}
