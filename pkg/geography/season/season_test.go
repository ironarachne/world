package season

import "testing"

func TestGetDaylightHours(t *testing.T) {
	hours := getDaylightHours(99, 1)
	if hours != 21 {
		t.Error("failed to get 21 hours of daylight at north pole in January: ", hours)
	}

	hours = getDaylightHours(3, 1)
	if hours != 12 {
		t.Error("failed to get 12 hours of daylight at equator in January: ", hours)
	}

	hours = getDaylightHours(-99, 1)
	if hours != 3 {
		t.Error("failed to get 3 hours of daylight at south pole in January: ", hours)
	}

	hours = getDaylightHours(99, 7)
	if hours != 3 {
		t.Error("failed to get 3 hours of daylight at north pole in July: ", hours)
	}

	hours = getDaylightHours(0, 7)
	if hours != 12 {
		t.Error("failed to get 12 hours of daylight at equator in July: ", hours)
	}

	hours = getDaylightHours(-99, 7)
	if hours != 21 {
		t.Error("failed to get 21 hours of daylight at south pole in July: ", hours)
	}
}
