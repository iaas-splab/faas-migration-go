package core

func Delete(req IDRequest, repo Repository) error {
	if err := repo.Init(); err != nil {
		return err
	}

	_, err := repo.Get(req.ID)
	if err != nil {
		return err
	}

	return repo.Delete(req.ID)
}
