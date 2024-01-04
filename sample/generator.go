package sample

import (
	"github.com/akagami-harsh/pcbook/pb"
	"github.com/golang/protobuf/ptypes"
)

func NewKeyboard() *pb.Keyboard {
	keyboard := &pb.Keyboard{
		Layout:  randomKeyboardLayout(),
		Backlit: randomBool(),
	}
	return keyboard
}

// NewCPU returns a random CPU.
func NewCPU() *pb.CPU {
	brand := randomCPUBrand()
	name := randomCPUName(brand)

	numberCores := randomInt(2, 8)
	numberThreads := randomInt(numberCores, 12)

	minGhz := randomFloat64(2.0, 4.0)
	maxGhz := randomFloat64(minGhz, 8.0)
	cpu := &pb.CPU{
		Brand:         brand,
		Name:          name,
		NumberCores:   uint32(numberCores),
		NumberThreads: uint32(numberThreads),
		MinGhz:        minGhz,
		MaxGhz:        maxGhz,
	}
	return cpu
}

// NewGPU returns a sample GPU.
func NewGPU() *pb.GPU {
	brand := randomGPUBrand()
	name := randomGPUName(brand)
	minGhz := randomFloat64(1.0, 1.5)
	maxGhz := randomFloat64(minGhz, 4.0)

	memory := &pb.Memory{
		Value: uint64(randomInt(1, 16)),
		Unit:  pb.Memory_GIGABYTE,
	}

	gpu := &pb.GPU{
		Brand:  brand,
		Name:   name,
		MinGhz: minGhz,
		MaxGhz: maxGhz,
		Memory: memory,
	}
	return gpu
}

// NewRam returns a sample RAM.
func NewRam() *pb.Memory {
	ram := &pb.Memory{
		Value: uint64(randomInt(8, 64)),
		Unit:  pb.Memory_GIGABYTE,
	}
	return ram
}

// NewSSD returns a sample SSD.
func NewSSD() *pb.Storage {
	ssd := &pb.Storage{
		Driver: pb.Storage_SSD,
		Memory: &pb.Memory{
			Value: uint64(randomInt(128, 1064)),
			Unit:  pb.Memory_GIGABYTE,
		},
	}
	return ssd

}

// NewSSD returns a sample SSD.
func NewHDD() *pb.Storage {
	hdd := &pb.Storage{
		Driver: pb.Storage_HDD,
		Memory: &pb.Memory{
			Value: uint64(randomInt(1, 6)),
			Unit:  pb.Memory_TERABYTE,
		},
	}
	return hdd

}

// NewScreen returns a sample screen
func NewScreen() *pb.Screen {
	height := randomInt(1080, 4320)
	width := height * 16 / 9
	screen := &pb.Screen{
		SizeInch: randomFloat32(13, 17),
		Resolution: &pb.Screen_Resolution{
			Height: uint32(height),
			Width:  uint32(width),
		},
		Panel:      randomScreenPanel(),
		Multitouch: randomBool(),
	}
	return screen
}

// NewLaptop returns a sample laptop
func NewLaptop() *pb.Laptop {
	brand := randomLaptopBrand()
	name := randomLaptopName(brand)
	laptop := &pb.Laptop{
		Id:       randomID(),
		Brand:    brand,
		Name:     name,
		Cpu:      NewCPU(),
		Ram:      NewRam(),
		Gpu:      []*pb.GPU{NewGPU()},
		Storage:  []*pb.Storage{NewSSD(), NewHDD()},
		Screen:   NewScreen(),
		Keyboard: NewKeyboard(),
		Weight: &pb.Laptop_WeightKg{
			WeightKg: randomFloat64(1.0, 3.0),
		},
		PriceUsd:    randomFloat64(1500, 3000),
		ReleaseYear: uint32(randomInt(2020, 2023)),
		UpdatedAt:   ptypes.TimestampNow(),
	}
	return laptop
}
