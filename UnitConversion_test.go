package dyn4go

import (
	"math"
	"testing"
)

const CLOSE_ENOUGH float64 = 1.0e-9

func TestFootMeter(t *testing.T) {
	if math.Abs(FOOT_TO_METER*METER_TO_FOOT-1) > CLOSE_ENOUGH {
		t.Error("FOOT_TO_METER * METER_TO_FOOT is not close enough to 1")
	}
	if math.Abs(FeetToMeters(MetersToFeet(2.5))-2.5) > CLOSE_ENOUGH {
		t.Error("FeetToMeters(MetersToFeet(2.5)) is not close enough to 2.5")
	}
}

func TestSlugKilogram(t *testing.T) {
	if math.Abs(SLUG_TO_KILOGRAM*KILOGRAM_TO_SLUG-1) > CLOSE_ENOUGH {
		t.Error("SLUG_TO_KILOGRAM * KILOGRAM_TO_SLUG is not close enough to 1")
	}
	if math.Abs(SlugsToKilograms(KilogramsToSlugs(2.5))-2.5) > CLOSE_ENOUGH {
		t.Error("SlugsToKilograms(KilogramsToSlugs(2.5)) is not close enough to 2.5")
	}
}

func TestPoundKilogram(t *testing.T) {
	if math.Abs(POUND_TO_KILOGRAM*KILOGRAM_TO_POUND-1) > CLOSE_ENOUGH {
		t.Error("POUND_TO_KILOGRAM * KILOGRAM_TO_POUND is not close enough to 1")
	}
	if math.Abs(PoundsToKilograms(KilogramsToPounds(2.5))-2.5) > CLOSE_ENOUGH {
		t.Error("PoundsToKilograms(KilogramsToPounds(2.5)) is not close enough to 2.5")
	}
}

func TestMPSToFPS(t *testing.T) {
	if math.Abs(MetersPerSecondToFeetPerSecond(FeetPerSecondToMetersPerSecond(2.5))-2.5) > CLOSE_ENOUGH {
		t.Error("MetersPerSecondToFeetPerSecond(FeetPerSecondToMetersPerSecond(2.5)) is not close enough to 2.5")
	}
}

func TestPoundNewton(t *testing.T) {
	if math.Abs(POUND_TO_NEWTON*NEWTON_TO_POUND-1) > CLOSE_ENOUGH {
		t.Error("POUND_TO_NEWTON * NEWTON_TO_POUND is not close enough to 1")
	}
	if math.Abs(PoundsToNewtons(NewtonsToPounds(2.5))-2.5) > CLOSE_ENOUGH {
		t.Error("PoundsToNewtons(NewtonsToPounds(2.5)) is not close enough to 2.5")
	}
}

func TestFootPoundNewtonMeter(t *testing.T) {
	if math.Abs(FOOT_POUND_TO_NEWTON_METER*NEWTON_METER_TO_FOOT_POUND-1) > CLOSE_ENOUGH {
		t.Error("FOOT_POUND_TO_NEWTON_METER * NEWTON_METER_TO_FOOT_POUND is not close enough to 1")
	}
	if math.Abs(NewtonMetersToFootPounds(FootPoundsToNewtonMeters(2.5))-2.5) > CLOSE_ENOUGH {
		t.Error("NewtonMetersToFootPounds(FootPoundsToNewtonMeters(2.5)) is not close enough to 2.5")
	}
}
