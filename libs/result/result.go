package result

type Result[T any] struct {
	value T
	error error
}

func Ok[T any](value T) Result[T] {
	return Result[T]{
		value: value,
	}
}

func Err[T any](err error) Result[T] {
	return Result[T]{
		error: err,
	}
}

func Try[T any](f func() (T, error)) Result[T] {
	value, err := f()

	if err != nil {
		return Err[T](err)
	}

	return Ok(value)
}

func Fold[T any, R any](result Result[T], onSuccess func(T) R, onError func(err error) R) R {
	if result.error != nil {
		return onError(result.error)
	}
	return onSuccess(result.value)
}

func Map[T any, R any](result Result[T], mapper func(t T) R) Result[R] {
	return FlatMap(result, func(t T) Result[R] {
		return Result[R]{
			value: mapper(t),
		}
	})
}

func FlatMap[T any, R any](result Result[T], mapper func(t T) Result[R]) Result[R] {
	return Fold(result, func(t T) Result[R] {
		return mapper(t)
	}, func(err error) Result[R] {
		return Result[R]{
			error: err,
		}
	})
}
