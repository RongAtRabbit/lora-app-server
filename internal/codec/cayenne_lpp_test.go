package codec

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestCayenneLPPDecode(t *testing.T) {
	Convey("Given a set of tests", t, func() {
		tests := []struct {
			Name   string
			Bytes  []byte
			Struct CayenneLPP
		}{
			{
				Name:  "2 digital input",
				Bytes: []byte{3, 0, 100, 5, 0, 210},
				Struct: CayenneLPP{
					DigitalInput: map[byte]uint8{
						3: 100,
						5: 210,
					},
				},
			},
			{
				Name:  "2 digital output",
				Bytes: []byte{3, 1, 100, 5, 1, 210},
				Struct: CayenneLPP{
					DigitalOutput: map[byte]uint8{
						3: 100,
						5: 210,
					},
				},
			},
			{
				Name:  "2 analog input",
				Bytes: []byte{3, 2, 0, 10, 5, 2, 3, 232},
				Struct: CayenneLPP{
					AnalogInput: map[byte]float64{
						3: 0.1,
						5: 10,
					},
				},
			},
			{
				Name:  "2 analog output",
				Bytes: []byte{3, 3, 0, 10, 5, 3, 3, 232},
				Struct: CayenneLPP{
					AnalogOutput: map[byte]float64{
						3: 0.1,
						5: 10,
					},
				},
			},
			{
				Name:  "2 illuminance sensors",
				Bytes: []byte{3, 101, 0, 10, 5, 101, 3, 232},
				Struct: CayenneLPP{
					IlluminanceSensor: map[byte]uint16{
						3: 10,
						5: 1000,
					},
				},
			},
			{
				Name:  "2 presence sensors",
				Bytes: []byte{3, 102, 5, 5, 102, 3},
				Struct: CayenneLPP{
					PresenceSensor: map[byte]uint8{
						3: 5,
						5: 3,
					},
				},
			},
			{
				Name:  "2 temperature sensors",
				Bytes: []byte{3, 103, 1, 16, 5, 103, 0, 255},
				Struct: CayenneLPP{
					TemperatureSensor: map[byte]float64{
						3: 27.2,
						5: 25.5,
					},
				},
			},
			{
				Name:  "2 humidity sensors",
				Bytes: []byte{3, 104, 41, 5, 104, 150},
				Struct: CayenneLPP{
					HumiditySensor: map[byte]float64{
						3: 20.5,
						5: 75,
					},
				},
			},
			{
				Name:  "2 accelerometers",
				Bytes: []byte{3, 113, 0, 1, 0, 2, 0, 3, 5, 113, 3, 234, 7, 211, 11, 187},
				Struct: CayenneLPP{
					Accelerometer: map[byte]Accelerometer{
						3: {X: 0.001, Y: 0.002, Z: 0.003},
						5: {X: 1.002, Y: 2.003, Z: 3.003},
					},
				},
			},
			{
				Name:  "2 barometers",
				Bytes: []byte{3, 115, 4, 31, 5, 115, 9, 196},
				Struct: CayenneLPP{
					Barometer: map[byte]float64{
						3: 105.5,
						5: 250,
					},
				},
			},
			{
				Name:  "2 gyrometer sensors",
				Bytes: []byte{3, 134, 0, 1, 0, 2, 0, 3, 5, 134, 3, 233, 7, 210, 11, 187},
				Struct: CayenneLPP{
					Gyrometer: map[byte]Gyrometer{
						3: {X: 0.01, Y: 0.02, Z: 0.03},
						5: {X: 10.01, Y: 20.02, Z: 30.03},
					},
				},
			},
			{
				Name:  "gps sensor",
				Bytes: []byte{1, 136, 6, 118, 95, 242, 150, 10, 0, 3, 232},
				Struct: CayenneLPP{
					GPSLocation: map[byte]GPSLocation{
						1: {Latitude: 42.3519, Longitude: -87.9094, Altitude: 10},
					},
				},
			},
		}

		for i, test := range tests {
			Convey(fmt.Sprintf("Testing: %s [%d]", test.Name, i), func() {
				Convey("Decoding", func() {
					lpp := CayenneLPP{}
					So(lpp.UnmarshalBinary(test.Bytes), ShouldBeNil)
					So(lpp, ShouldResemble, test.Struct)
				})
				Convey("Encoding", func() {
					b, err := test.Struct.MarshalBinary()
					So(err, ShouldBeNil)

					lpp := CayenneLPP{}
					So(lpp.UnmarshalBinary(b), ShouldBeNil)
					So(lpp, ShouldResemble, test.Struct)
				})
			})
		}
	})
}
