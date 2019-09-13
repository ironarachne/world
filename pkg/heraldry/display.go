package heraldry

// SimplifiedDevice is a simplified version of an heraldic device meant for JSON display
type SimplifiedDevice struct {
	Blazon string
	ImageURL string
}

// Simplify returns a simplified device
func (device Device) Simplify() SimplifiedDevice {
	sd := SimplifiedDevice{
		Blazon: device.Blazon,
		ImageURL: device.ImageURL,
	}

	return sd
}
