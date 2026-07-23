package persistence

func (a *AOF) FileClose() error {
	err := a.File.Close()
	if err != nil {
		return err
	}
	return nil
}
