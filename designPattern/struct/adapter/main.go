package main

func main() {
	hdmi := NewHDMI()
	usb := NewUSB()
	lighting := NewLighting()
	mac := NewMac()

	LightingAdapter, _ := NewLightingAdapter(hdmi)
	mac.ReadPorts(LightingAdapter)

	LightingAdapter, _ = NewLightingAdapter(usb)
	mac.ReadPorts(LightingAdapter)

	mac.ReadPorts(lighting)

}
