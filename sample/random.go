package sample

import (
	"math/rand"
	"time"

	"github.com/akagami-harsh/pcbook/pb"
	"github.com/google/uuid"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func randomScreenPanel() pb.Screen_Panel {
	if rand.Intn(2) == 1 {
		return pb.Screen_IPS
	}
	return pb.Screen_LCD
}

func randomKeyboardLayout() pb.Keyboard_Layout {
	switch rand.Intn(3) {
	case 1:
		return pb.Keyboard_QWERTY
	case 2:
		return pb.Keyboard_QWERTZ
	default:
		return pb.Keyboard_AZERTY
	}
}

func randomCPUBrand() string {
	return randomStringFromSet("Intel", "AMD", "ARM")
}

func randomGPUBrand() string {
	return randomStringFromSet("NVIDIA", "AMD")
}

func randomGPUName(brand string) string {
	switch brand {
	case "NVIDIA":
		return randomStringFromSet("GTX 980", "GTX 1080", "GTX 1070")
	case "AMD":
		return randomStringFromSet("RX 480", "RX 580", "RX 570")
	default:
		return ""
	}
}

func randomLaptopBrand() string {
	return randomStringFromSet("Apple", "Dell", "Lenovo", "Asus")
}

func randomLaptopName(brand string) string {
	switch brand {
	case "Apple":
		return randomStringFromSet("MacBook", "MacBook Pro", "MacBook Air")
	case "Dell":
		return randomStringFromSet("Inspiron", "XPS", "Latitude")
	case "Lenovo":
		return randomStringFromSet("ThinkPad", "ThinkStation", "ThinkCentre")
	case "Asus":
		return randomStringFromSet("ZenBook", "X540U", "X441U")
	default:
		return ""
	}
}

func randomStringFromSet(a ...string) string {
	n := len(a)
	if n == 0 {
		return ""
	}
	return a[rand.Intn(n)]
}

func randomCPUName(brand string) string {
	switch brand {
	case "Intel":
		return randomStringFromSet("i3", "i5", "i7")
	case "AMD":
		return randomStringFromSet("A6", "A8", "A10")
	case "ARM":
		return randomStringFromSet("A53", "A72", "A73")
	default:
		return ""
	}
}

func randomBool() bool {
	return rand.Intn(2) == 1
}

func randomInt(min int, max int) int {
	return min + rand.Intn(max-min+1)
}

func randomFloat64(min float64, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

func randomFloat32(min float32, max float32) float32 {
	return min + rand.Float32()*(max-min)
}

func randomID() string {
	return uuid.New().String()
}
