package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"pixiboParser/brand"
	"strconv"
	"strings"
)

func main() {
	//fmt.Println("its working")
	file := os.Args[2]
	//fmt.Println(file)
	//file := "brands_sample.csv"
	f, err := os.Open(file)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(len(lines[0]))
	parseCsvToJson(lines)

}

func parseCsvToJson(lines [][]string) {
	var b brand.Brand
	var bS brand.BrandsStruct
	allSizes := map[string]bool{
		"XS":  true,
		"S":   true,
		"M":   true,
		"L":   true,
		"XL":  true,
		"XXL": true,
	}
	// b.Name = "mradul"
	// b.Gender = "23"
	// //obj.setInitialValue()
	// brand.PrintInitialValue(b)

	firstRecord := true
	numRows := len(lines)
	numColumns := len(lines[0])
	var i, j, k int
	for i = 0; i < numRows; i++ {
		switch strings.Trim(lines[i][0], " ") {
		case "*":
			//fmt.Println(i)
			if !firstRecord {
				//fmt.Println(b)
				val := append(bS.Brands, b)
				bS.Brands = val
			}
			if i == numRows-1 {
				break
			}
			firstRecord = false
			var newB brand.Brand
			b = newB
			b.Name = strings.Trim(lines[i+1][0], " ")
			i++
			break
		case "Gender":
			b.Gender = strings.Trim(lines[i][1], " ")
			break
		case "Region":
			b.BrandSizeCharts.Region = strings.Trim(lines[i][1], " ")
			break
		case "SubType":
			b.BrandSizeCharts.SubType = strings.Trim(lines[i][1], " ")
			break
		case "Id":
			b.Id = strings.Trim(lines[i][1], " ")
			break
		case "Size":
			sizeCount := 0
			for k = i + 1; k < numRows; k++ {
				var emptySizeCharts brand.SizeChartsStruct
				if allSizes[strings.Trim(lines[k][0], " ")] {
					emptySizeCharts.Size = strings.Trim(lines[k][0], " ")
					for j = 0; j < numColumns; j++ {
						//fmt.Println("here 1")
						//fmt.Println(strings.Trim(lines[i][j], " "))
						if strings.Trim(lines[i][j], " ") != "" {
							//fmt.Println("here 2")
							//fmt.Println(strings.Trim(lines[i][j], " "))
							switch strings.Trim(lines[i][j], " ") {
							case "Size":
								break
							case "Chest":
								emptySizeCharts.Bust, _ = strconv.Atoi(strings.Trim(lines[k][j], " "))
								break
							case "Chest (U)":
								emptySizeCharts.BustUpper, _ = strconv.Atoi(strings.Trim(lines[k][j], " "))
								break
							case "Waist":
								emptySizeCharts.Waist, _ = strconv.Atoi(strings.Trim(lines[k][j], " "))
								break
							case "Waist (U)":
								emptySizeCharts.WaistUpper, _ = strconv.Atoi(strings.Trim(lines[k][j], " "))
								break
							case "Hip":
								emptySizeCharts.Hip, _ = strconv.Atoi(strings.Trim(lines[k][j], " "))
								break
							case "Hip (U)":
								emptySizeCharts.HipUpper, _ = strconv.Atoi(strings.Trim(lines[k][j], " "))
								break
							}
						}
					}
					val := append(b.BrandSizeCharts.SizeCharts, emptySizeCharts)
					b.BrandSizeCharts.SizeCharts = val
					sizeCount = sizeCount + 1
				} else {
					break
				}
			}
			i = i + sizeCount
		case "Chest":
			bustCount := 0
			for k = i + 1; k < numRows; k++ {
				var emptyBust brand.BustStruct
				if strings.Trim(lines[k][0], " ") == "" {
					for j = 1; j < numColumns; j++ {
						if strings.Trim(lines[k][j], " ") != "" {
							switch strings.Trim(lines[i][j], " ") {
							case "Size":
								emptyBust.Size = strings.Trim(lines[k][j], " ")
								break
							case "Too Tight":
								emptyBust.TooTight, _ = strconv.Atoi(strings.Trim(lines[k][j], " "))
								break
							case "Snug Fit":
								emptyBust.SnugFit, _ = strconv.ParseFloat(strings.Trim(lines[k][j], " "), 64)
								break
							case "Loose Fit":
								emptyBust.LooseFit, _ = strconv.Atoi(strings.Trim(lines[k][j], " "))
								break
							case "Too Loose":
								emptyBust.TooLoose, _ = strconv.Atoi(strings.Trim(lines[k][j], " "))
								break
							case "Median":
								emptyBust.Median, _ = strconv.Atoi(strings.Trim(lines[k][j], " "))
								break
							}
						} else {
							break
						}
					}
					val := append(b.BrandSizeCharts.AdjustmentTable.Bust, emptyBust)
					b.BrandSizeCharts.AdjustmentTable.Bust = val
					bustCount = bustCount + 1
				} else {
					break
				}
			}
			i = i + bustCount
		case "Waist":
			waistCount := 0
			for k = i + 1; k < numRows; k++ {
				var emptyWaist brand.WaistStruct
				if strings.Trim(lines[k][0], " ") == "" {
					for j = 1; j < numColumns; j++ {
						if strings.Trim(lines[k][j], " ") != "" {
							switch strings.Trim(lines[i][j], " ") {
							case "Size":
								emptyWaist.Size = strings.Trim(lines[k][j], " ")
								break
							case "Too Tight":
								emptyWaist.TooTight, _ = strconv.ParseFloat(strings.Trim(lines[k][j], " "), 64)
								break
							case "Snug Fit":
								emptyWaist.SnugFit, _ = strconv.Atoi(strings.Trim(lines[k][j], " "))
								break
							case "Loose Fit":
								emptyWaist.LooseFit, _ = strconv.ParseFloat(strings.Trim(lines[k][j], " "), 64)
								break
							case "Too Loose":
								emptyWaist.TooLoose, _ = strconv.ParseFloat(strings.Trim(lines[k][j], " "), 64)
								break
							case "Median":
								emptyWaist.Median, _ = strconv.Atoi(strings.Trim(lines[k][j], " "))
								break
							}
						} else {
							break
						}
					}
					val := append(b.BrandSizeCharts.AdjustmentTable.Waist, emptyWaist)
					b.BrandSizeCharts.AdjustmentTable.Waist = val
					waistCount = waistCount + 1
				} else {
					break
				}
			}
			i = i + waistCount
		case "Hip":
			hipCount := 0
			for k = i + 1; k < numRows; k++ {
				var emptyHip brand.HipStruct
				if strings.Trim(lines[k][0], " ") == "" {
					for j = 1; j < numColumns; j++ {
						if strings.Trim(lines[k][j], " ") != "" {
							switch strings.Trim(lines[i][j], " ") {
							case "Size":
								emptyHip.Size = strings.Trim(lines[k][j], " ")
								break
							case "Too Tight":
								emptyHip.TooTight, _ = strconv.ParseFloat(strings.Trim(lines[k][j], " "), 64)
								break
							case "Snug Fit":
								emptyHip.SnugFit, _ = strconv.Atoi(strings.Trim(lines[k][j], " "))
								break
							case "Loose Fit":
								emptyHip.LooseFit, _ = strconv.ParseFloat(strings.Trim(lines[k][j], " "), 64)
								break
							case "Too Loose":
								emptyHip.TooLoose, _ = strconv.Atoi(strings.Trim(lines[k][j], " "))
								break
							case "Median":
								emptyHip.Median, _ = strconv.Atoi(strings.Trim(lines[k][j], " "))
								break
							}
						} else {
							break
						}
					}
					val := append(b.BrandSizeCharts.AdjustmentTable.Hip, emptyHip)
					b.BrandSizeCharts.AdjustmentTable.Hip = val
					hipCount = hipCount + 1
				} else {
					break
				}
			}
			i = i + hipCount
		}

	}
	jsonB, err := json.Marshal(bS)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(jsonB))
	fileErr := os.WriteFile("output.file", jsonB, 0644)
	if fileErr != nil {
		fmt.Println("Not able to save data into the file")
	}
}
