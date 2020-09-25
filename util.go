package msgify

func returnOnErr(actions ...func() error) error {
	for _, act := range actions {
		if e := act(); e != nil {
			return e
		}
	}
	return nil
}

func must(action func() error) {
	err := action()
	if err != nil {
		panic(err)
	}
}
