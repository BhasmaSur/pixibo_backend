package brand

import "fmt"

type BrandsStruct struct {
	Brands []Brand `json:"brands"`
}

type Brand struct {
	Name            string                `json:"name"`
	Gender          string                `json:"gender"`
	BrandSizeCharts BrandSizeChartsStruct `json:"brandSizeCharts"`
	Id              string                `json:"id"`
}

type BrandSizeChartsStruct struct {
	Region          string                `json:"region"`
	SubType         string                `json:"subType"`
	SizeCharts      []SizeChartsStruct    `json:"sizeCharts"`
	AdjustmentTable AdjustmentTableStruct `json:"adjustmentTable"`
}

type SizeChartsStruct struct {
	Size       string `json:"size"`
	Bust       int    `json:"bust"`
	BustUpper  int    `json:"bustUpper"`
	Waist      int    `json:"waist"`
	WaistUpper int    `json:"waistUpper"`
	Hip        int    `json:"hip"`
	HipUpper   int    `json:"hipUpper"`
}

type AdjustmentTableStruct struct {
	Bust  []BustStruct  `json:"bust"`
	Waist []WaistStruct `json:"waist"`
	Hip   []HipStruct   `json:"hip"`
}

type BustStruct struct {
	TooTight int     `json:"tooTight"`
	SnugFit  float64 `json:"snugFit"`
	LooseFit int     `json:"looseFit"`
	TooLoose int     `json:"tooLoose"`
	Median   int     `json:"median"`
	Size     string  `json:"size"`
}

type WaistStruct struct {
	TooTight float64 `json:"tooTight"`
	SnugFit  int     `json:"snugFit"`
	LooseFit float64 `json:"looseFit"`
	TooLoose float64 `json:"tooLoose"`
	Median   int     `json:"median"`
	Size     string  `json:"size"`
}

type HipStruct struct {
	TooTight float64 `json:"tooTight"`
	SnugFit  int     `json:"snugFit"`
	LooseFit float64 `json:"looseFit"`
	TooLoose int     `json:"tooLoose"`
	Median   int     `json:"median"`
	Size     string  `json:"size"`
}

func (b *Brand) setInitialValue() {
	b.Name = "mradul"
	b.Gender = "23"
}

func PrintInitialValue(b Brand) {
	fmt.Println(b)
}
