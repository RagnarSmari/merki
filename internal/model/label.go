package model 

type Label struct {
	Version string `json:"version"`
	Size Size `json:"size"`
	Dpi int `json:"dpi"`
	Elements []Element `json:"elements"`
}

type Size struct {
	Width  int  `json:"width"`
	Height int  `json:"height"`
	Unit   Unit `json:"unit"`
}

type Unit string

const (
	UnitMM   Unit = "mm"
	UnitDots Unit = "dots"
)

type Element struct {
	Type      ElementType `json:"type"`
	X         int         `json:"x"`
	Y         int         `json:"y"`
	Font      string      `json:"font"`
	Size      int         `json:"size"`
	Width     int         `json:"width"`
	Height    int         `json:"height"`
	Value     string      `json:"value"`
	Thickness float32     `json:"thickness"`
	Symbology string      `json:"symbology"`
}

type ElementType string

const (
	ElementText    ElementType = "text"
	ElementBarcode ElementType = "barcode"
	ElementQR      ElementType = "qr"
	ElementBox     ElementType = "box"
)
