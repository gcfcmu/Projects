package main

import (
    //"bytes"
	"os"
	// "math"
	// "math/rand"

	//chart "github.com/wcharczuk/go-chart/v2"
    chart "github.com/wcharczuk/go-chart" //exposes "chart"
	"fmt"
)

func CreateGraph(generations, predPops1, predPops2, preyPops1, midPops1 []float64) {
	fmt.Println("start")
	xValues, yValues := generations, predPops1
	xValues2, yValues2 := generations, predPops2
	xValues3, yValues3 := generations, preyPops1
	xValues4, yValues4 := generations, midPops1
	// graph := chart.Chart{
	// 	XAxis: chart.XAxis{
	// 		Name: "Little Thomas",
	// 	},
	// 	YAxis: chart.YAxis {
	// 		Name: "Big Thomas",
	// 	},
	// 	Title: "Thomas Chart",
	// 	TitleStyle: getTitleStyle(),
	// 	Series: []chart.Series{
	// 		chart.ContinuousSeries{
	// 			// Style: chart.Style{
	// 			// 	StrokeColor: chart.GetDefaultColor(0).WithAlpha(64),
	// 			// 	FillColor:   chart.GetDefaultColor(0).WithAlpha(64),
	// 			// },
	// 			XValues: xValues,
	// 			YValues: yValues,
	// 		},
	// 	},
	// }
	graph := chart.Chart{
		XAxis: chart.XAxis{
			Name: "The XAxis",
			Style: GetAxisStyle(),
		},
		YAxis: chart.YAxis{
			Name: "The YAxis",
			Style: GetAxisStyle(),
		},
		Title: "Pop Dynamics Chart",
		TitleStyle: GetTitleStyle(),
		Series: []chart.Series{
			chart.ContinuousSeries{
				// Style: chart.Style{
				// 	StrokeColor: chart.GetDefaultColor(0).WithAlpha(64),
				// 	FillColor:   chart.GetDefaultColor(0).WithAlpha(64),
				// },
				XValues: xValues,
				YValues: yValues,
			},
			chart.ContinuousSeries{
				XValues: xValues2,
				YValues: yValues2,
			},
			chart.ContinuousSeries{
				XValues: xValues3,
				YValues: yValues3,
			},
			chart.ContinuousSeries{
				XValues: xValues4,
				YValues: yValues4,
			},
		},
	}

	filename := "popdynamics.png"
	f, err := os.Create(filename)
	if err != nil {
		panic("unable to create file")
	}

	defer f.Close()

	err = graph.Render(chart.PNG, f)
	if err != nil {
		panic("unable to render graph")
	}
}

func CreatePreyGraph(generations, preyPops []float64) {
	fmt.Println("start")
	// xValues, yValues := generations, pred1Pops
	// xValues2, yValues2 := generations, pred2Pops
	xValues3, yValues3 := generations, preyPops
	// graph := chart.Chart{
	// 	XAxis: chart.XAxis{
	// 		Name: "Little Thomas",
	// 	},
	// 	YAxis: chart.YAxis {
	// 		Name: "Big Thomas",
	// 	},
	// 	Title: "Thomas Chart",
	// 	TitleStyle: getTitleStyle(),
	// 	Series: []chart.Series{
	// 		chart.ContinuousSeries{
	// 			// Style: chart.Style{
	// 			// 	StrokeColor: chart.GetDefaultColor(0).WithAlpha(64),
	// 			// 	FillColor:   chart.GetDefaultColor(0).WithAlpha(64),
	// 			// },
	// 			XValues: xValues,
	// 			YValues: yValues,
	// 		},
	// 	},
	// }
	graph := chart.Chart{
		XAxis: chart.XAxis{
			Name: "The XAxis",
			Style: GetAxisStyle(),
		},
		YAxis: chart.YAxis{
			Name: "The YAxis",
			Style: GetAxisStyle(),
		},
		Title: "Pop Dynamics Chart",
		TitleStyle: GetTitleStyle(),
		Series: []chart.Series{
			// chart.ContinuousSeries{
			// 	// Style: chart.Style{
			// 	// 	StrokeColor: chart.GetDefaultColor(0).WithAlpha(64),
			// 	// 	FillColor:   chart.GetDefaultColor(0).WithAlpha(64),
			// 	// },
			// 	XValues: xValues,
			// 	YValues: yValues,
			// },
			// chart.ContinuousSeries{
			// 	XValues: xValues2,
			// 	YValues: yValues2,
			// },
			chart.ContinuousSeries{
				XValues: xValues3,
				YValues: yValues3,
			},
		},
	}

	filename := "popdynamics.png"
	f, err := os.Create(filename)
	if err != nil {
		panic("unable to create file")
	}

	defer f.Close()

	err = graph.Render(chart.PNG, f)
	if err != nil {
		panic("unable to render graph")
	}
}

func GetTitleStyle()chart.Style{
	return chart.Style{
		Show: true,
		FontSize: 20,
		//FontColor: drawing.ColorBlue
		//FontColor: drawing.ColorBlue,
	}
}

func GetAxisStyle()chart.Style{
	return chart.Style{
		Show: true,
		FontSize: 14,
		//FontColor: drawing.ColorBlue,
	}
}
