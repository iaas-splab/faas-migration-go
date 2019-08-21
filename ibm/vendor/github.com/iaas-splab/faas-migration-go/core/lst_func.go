package core

func List(repo Repository) ([]ToDoItem, error) {
	if err := repo.Init(); err != nil {
		return nil, err
	}

	return repo.List()
}
