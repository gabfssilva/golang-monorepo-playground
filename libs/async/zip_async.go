package async

func ZipEffectAsync[F comparable, S comparable](f1 func() (F, error), f2 func() (S, error)) (F, S, error) {
	ch1 := make(chan F, 1)
	ch2 := make(chan S, 1)
	chErr := make(chan error, 1)

	go func() {
		result, err := f1()
		if err != nil {
			chErr <- err
			return
		}
		ch1 <- result
	}()

	go func() {
		result, err := f2()
		if err != nil {
			chErr <- err
			return
		}
		ch2 <- result
	}()

	var resF F
	var resS S
	var count int

	for count < 2 {
		select {
		case err := <-chErr:
			return *new(F), *new(S), err
		case res := <-ch1:
			resF = res
			count++
		case res := <-ch2:
			resS = res
			count++
		}
	}

	return resF, resS, nil
}

func ZipEffectAsync3[F comparable, S comparable, T comparable](
	f1 func() (F, error),
	f2 func() (S, error),
	f3 func() (T, error)) (F, S, T, error) {

	ch1 := make(chan F, 1)
	ch2 := make(chan S, 1)
	ch3 := make(chan T, 1)
	chErr := make(chan error, 1)

	go func() {
		result, err := f1()
		if err != nil {
			chErr <- err
			return
		}
		ch1 <- result
	}()

	go func() {
		result, err := f2()
		if err != nil {
			chErr <- err
			return
		}
		ch2 <- result
	}()

	go func() {
		result, err := f3()
		if err != nil {
			chErr <- err
			return
		}
		ch3 <- result
	}()

	var resF F
	var resS S
	var resT T
	var count int

	for count < 3 {
		select {
		case err := <-chErr:
			return *new(F), *new(S), *new(T), err
		case res := <-ch1:
			resF = res
			count++
		case res := <-ch2:
			resS = res
			count++
		case res := <-ch3:
			resT = res
			count++
		}
	}

	return resF, resS, resT, nil
}

func ZipEffectAsync4[F comparable, S comparable, T comparable, U comparable](
	f1 func() (F, error),
	f2 func() (S, error),
	f3 func() (T, error),
	f4 func() (U, error)) (F, S, T, U, error) {

	ch1 := make(chan F, 1)
	ch2 := make(chan S, 1)
	ch3 := make(chan T, 1)
	ch4 := make(chan U, 1)
	chErr := make(chan error, 1)

	go func() {
		result, err := f1()
		if err != nil {
			chErr <- err
			return
		}
		ch1 <- result
	}()

	go func() {
		result, err := f2()
		if err != nil {
			chErr <- err
			return
		}
		ch2 <- result
	}()

	go func() {
		result, err := f3()
		if err != nil {
			chErr <- err
			return
		}
		ch3 <- result
	}()

	go func() {
		result, err := f4()
		if err != nil {
			chErr <- err
			return
		}
		ch4 <- result
	}()

	var resF F
	var resS S
	var resT T
	var resU U
	var count int

	for count < 4 {
		select {
		case err := <-chErr:
			return *new(F), *new(S), *new(T), *new(U), err
		case res := <-ch1:
			resF = res
			count++
		case res := <-ch2:
			resS = res
			count++
		case res := <-ch3:
			resT = res
			count++
		case res := <-ch4:
			resU = res
			count++
		}
	}

	return resF, resS, resT, resU, nil
}
