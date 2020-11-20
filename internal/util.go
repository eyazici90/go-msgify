package internal

func ReturnOnErr(actions ...func() error) error {
	for _, act := range actions {
		if e := act(); e != nil {
			return e
		}
	}
	return nil
}

func Must(action func() error) {
	err := action()
	if err != nil {
		panic(err)
	}
}
