package async

import (
	"errors"
	"testing"
	"time"
)

func TestZipEffectAsync(t *testing.T) {
	start := time.Now()
	first, second, err := ZipEffectAsync(testFunc1, testFunc2)
	end := time.Now()

	if end.Sub(start) >= 210*time.Millisecond {
		t.Fatalf("Expected 200ms, actual %v", end.Sub(start))
	}

	if first != 100 {
		t.Fail()
	}
	if second != "200" {
		t.Fail()
	}
	if err != nil {
		t.Fatalf(err.Error())
	}
}

func TestZipEffectAsyncWithError(t *testing.T) {
	start := time.Now()
	_, _, err := ZipEffectAsync(testFuncError1, testFunc2)
	end := time.Now()

	if end.Sub(start) >= 110*time.Millisecond {
		t.Fatalf("Expected 100ms, actual %v", end.Sub(start))
	}

	if err == nil {
		t.Fatalf("This should have raised an error")
	}
}

func TestZipEffectAsync3(t *testing.T) {
	start := time.Now()
	first, second, third, err := ZipEffectAsync3(testFunc1, testFunc2, testFunc3)
	end := time.Now()

	if end.Sub(start) >= 210*time.Millisecond {
		t.Fatalf("Expected 200ms, actual %v", end.Sub(start))
	}

	if first != 100 {
		t.Fail()
	}
	if second != "200" {
		t.Fail()
	}
	if third != 300 {
		t.Fail()
	}
	if err != nil {
		t.Fatalf(err.Error())
	}
}

func TestZipEffectAsync3WithError(t *testing.T) {
	start := time.Now()
	_, _, _, err := ZipEffectAsync3(testFunc1, testFuncError1, testFunc2)
	end := time.Now()

	if end.Sub(start) >= 110*time.Millisecond {
		t.Fatalf("Expected 100ms, actual %v", end.Sub(start))
	}

	if err == nil {
		t.Fatalf("This should have raised an error")
	}
}

func TestZipEffectAsync4(t *testing.T) {
	start := time.Now()
	first, second, third, fourth, err := ZipEffectAsync4(testFunc1, testFunc2, testFunc3, testFunc4)
	end := time.Now()

	if end.Sub(start) >= 210*time.Millisecond {
		t.Fatalf("Expected 200ms, actual %v", end.Sub(start))
	}

	if first != 100 {
		t.Fail()
	}
	if second != "200" {
		t.Fail()
	}
	if third != 300 {
		t.Fail()
	}
	if fourth != 400 {
		t.Fail()
	}
	if err != nil {
		t.Fatalf(err.Error())
	}
}

func TestZipEffectAsync4WithError(t *testing.T) {
	start := time.Now()
	_, _, _, _, err := ZipEffectAsync4(testFunc1, testFunc2, testFuncError1, testFunc3)
	end := time.Now()

	if end.Sub(start) >= 110*time.Millisecond {
		t.Fatalf("Expected 100ms, actual %v", end.Sub(start))
	}

	if err == nil {
		t.Fatalf("This should have raised an error")
	}
}

func testFunc1() (int, error) {
	time.Sleep(100 * time.Millisecond)
	return 100, nil
}

func testFunc2() (string, error) {
	time.Sleep(200 * time.Millisecond)
	return "200", nil
}

func testFunc3() (int, error) {
	time.Sleep(50 * time.Millisecond)
	return 300, nil
}

func testFunc4() (int, error) {
	time.Sleep(150 * time.Millisecond)
	return 400, nil
}

func testFuncError1() (int, error) {
	time.Sleep(100 * time.Millisecond)
	return 0, errors.New("ups, something went wrong")
}
