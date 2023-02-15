package main

func Try(run func(try func(err error))) error {

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

	resolver := func(run func(try func(err error))) {
		defer recoverer()
		run(catcher)
	}

	resolver(run)
	return firstErr
}
