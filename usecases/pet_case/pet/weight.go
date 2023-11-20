package pet

type Weight uint8

var (
	VerySmall Weight = 0b00001
	Small     Weight = 0b00010
	Medium    Weight = 0b00100
	Large     Weight = 0b01000
	VeryLarge Weight = 0b10000
)
