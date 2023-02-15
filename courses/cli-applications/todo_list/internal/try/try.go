package try

type Catcher func(err error)

func Try(run func(try Catcher)) error {

	var firstErr error

	catcher := func(err error) {
		if err != nil {
			firstErr = err
			panic(err)
		}
	}

	recoverer := func() {
		recover()
	}

	resolver := func(run func(try Catcher)) {
		defer recoverer()
		run(catcher)
	}

	resolver(run)
	return firstErr
}
