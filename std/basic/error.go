package basic

import "errors"

func E1(b bool) (string, error) {
	if b {
		return "E1", nil
	} else {
		return "", errors.New("e1")
	}

}
func E2(b bool) (int, error) {
	if b {
		return 100, nil
	} else {
		return 0, errors.New("e2")
	}
}
func E3(b bool) (float32, error) {
	if b {
		return 100.001, nil
	} else {
		return 0, errors.New("e2")
	}
}
