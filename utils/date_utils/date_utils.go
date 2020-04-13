package date_utils

import "time"

const (
	apiDateLayout = "2006-01-02T15:04:05Z"
	apiDbLayout   = "2006-01-02 15:04:05"
)

// GetNow - Return current now in UTC format
func GetNow() time.Time {
	return time.Now().UTC()
}

// GetNowString - Returns the current time in apiDateLayout format.
func GetNowString() string {
	return GetNow().Format(apiDateLayout)

}

// GetNowDBFormat - Returns the current time in apiDbLayout format.
func GetNowDBFormat() string {
	return GetNow().Format(apiDbLayout)

}
