//- pkgs/date/date.go

package date

import (
	"database/sql/driver"
	"fmt"
	"strings"
	"time"
)

// 1. Define the custom type
type Date time.Time

const DateFormat = "2006-01-02"

// 2. Format JSON output to "YYYY-MM-DD"
func (d Date) MarshalJSON() ([]byte, error) {
	return fmt.Appendf(nil, `"%s"`, time.Time(d).Format(DateFormat)), nil
}

// 3. Parse incoming JSON "YYYY-MM-DD"
func (d *Date) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	if s == "null" || s == "" {
		return nil
	}
	t, err := time.Parse(DateFormat, s)
	if err != nil {
		return err
	}
	*d = Date(t)
	return nil
}

// 4. Scan database value into this type
func (d *Date) Scan(value interface{}) error {
	if value == nil {
		return nil
	}
	t, ok := value.(time.Time)
	if !ok {
		return fmt.Errorf("failed to scan date: %v", value)
	}
	*d = Date(t)
	return nil
}

// 5. Convert this type for database queries
func (d Date) Value() (driver.Value, error) {
	return time.Time(d).Format(DateFormat), nil
}

// DaysBefore returns the date N days before t in YYYY-MM-DD format.
func DateBefore(t time.Time, n int) time.Time {
	tdate := t.AddDate(0, 0, -n).Format(DateFormat)
	result, _ := time.Parse(DateFormat, tdate)
	return result
}

// MonthRange returns the first and last day of the month for a given date.
func MonthRange(t time.Time) (firstDay, lastDay time.Time) {
	firstDay = time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, t.Location())
	lastDay = firstDay.AddDate(0, 1, -1)
	return
}

func FormatDateTypesense(d *Date) string {
	if d == nil {
		return ""
	}
	// Karena date.Date = time.Time (atau embed time.Time), akses langsung:
	t := time.Time(*d) // jika type alias: type Date = time.Time
	if t.IsZero() {
		return ""
	}
	return t.UTC().Format(time.RFC3339)
}
