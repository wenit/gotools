package date

import "testing"

func TestCurrentDate(t *testing.T) {
	t.Log(CurrentDate())
}

func TestCurrentDateTime(t *testing.T) {
	t.Log(CurrentDateTime())
}

func TestCurrentTime(t *testing.T) {
	t.Log(CurrentTime())
}

func TestFormatDate(t *testing.T) {
	t.Log(FormatDate("20060102"))
	t.Log(FormatDate("150405"))
	t.Log(FormatDate("2006-01-02 15:04:05.000"))
}

func TestConvert(t *testing.T) {
	t.Log(Convert("2019-12-30 17:29:42.123", "2006-01-02 15:04:05.000", "20060102 150405.000"))

	t.Log(Convert("2019-1230 17:29:42.123", "2006-01-02 15:04:05.000", "20060102 150405.000"))
}
