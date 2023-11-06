package meteorology

import "fmt"

type TemperatureUnit int

func (t *TemperatureUnit) String() string {
	if *t == Celsius {
		return "°C"
	}

	if *t == Fahrenheit {
		return "°F"
	}

	return ""
}

const (
	Celsius    TemperatureUnit = 0
	Fahrenheit TemperatureUnit = 1
)

type Stringer interface {
	String() string
}

type Temperature struct {
	degree int
	unit   TemperatureUnit
}

func (t *Temperature) String() string {
	return fmt.Sprintf("%v %v", t.degree, t.unit.String())
}

type SpeedUnit int

func (t *SpeedUnit) String() string {
	unit := ""
	if *t == KmPerHour {
		unit = "km/h"
	}
	if *t == MilesPerHour {
		unit = "mph"
	}
	return unit
}

const (
	KmPerHour    SpeedUnit = 0
	MilesPerHour SpeedUnit = 1
)

type Speed struct {
	magnitude int
	unit      SpeedUnit
}

func (t *Speed) String() string {
	return fmt.Sprintf("%v %v", t.magnitude, t.unit.String())
}

type MeteorologyData struct {
	location      string
	temperature   Temperature
	windDirection string
	windSpeed     Speed
	humidity      int
}

func (m *MeteorologyData) String() string {
	return fmt.Sprintf("%v: %v, Wind %v at %v, %v%% Humidity", m.location, m.temperature.String(), m.windDirection, m.windSpeed.String(), m.humidity)
}
