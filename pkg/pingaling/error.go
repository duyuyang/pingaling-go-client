package pingaling

import "fmt"

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

// ErrNotExpectedJSON is returned when the API response isn't expected JSON
type ErrNotExpectedJSON struct {
	OriginalBody string
	Err          error
}

// ErrBadStatusCode is returned when the API returns a non 200 error code
type ErrBadStatusCode struct {
	OriginalBody string
	Code         int
}

func (e *ErrBadStatusCode) Error() string {
	return fmt.Sprintf("Invalid status code: %d", e.Code)
}

func (e *ErrNotExpectedJSON) Error() string {
	return fmt.Sprintf("Unexpected JSON: %s from %s", e.Err.Error(), e.OriginalBody)
}
