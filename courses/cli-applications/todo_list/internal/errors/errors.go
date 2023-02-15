package errors

type ErrorList []error

type ErrorCatcher func(err error)

func Try(try func(err ErrorCatcher), catch func(err error)) {

	var err error

	catcher := func(err error) {
		if err != nil {

		}
	}

	try()
}
