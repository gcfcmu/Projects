package main

import (
    //"bytes"
	"os"

	//chart "github.com/wcharczuk/go-chart/v2"
    chart "github.com/wcharczuk/go-chart" //exposes "chart"
	"fmt"
)

// CreateGraphMultipleSpecies: creates a population dynamics graph of all species in the ecosystem over time (generations)
// input: generations - time ecosystem runs for, allPopulations - map of species names to their populations throughout the generation,
//        and given file name to create file that will hold the graph
// output: (no direct output of this function, but the population dynamics graph is created here)
func CreateGraphMultipleSpecies(generations []float64, allPopulations map[string][]float64, filename string) {
	fmt.Println("start")

	//  sets up the slice of trend lines that will represent each species populations over time in the graph
	var allCurves []chart.ContinuousSeries
	allCurves = make([]chart.ContinuousSeries, len(allPopulations))

	var allCurvesPhase []chart.ContinuousSeries
	allCurvesPhase = make([]chart.ContinuousSeries, len(allPopulations))

	curveIndex := 0

	firstIteration := true
	var firstPopulation []float64
	var firstSpeciesName string

	for species, i := range(allPopulations){

		if firstIteration == true {
			firstPopulation = i
			firstSpeciesName = species
			firstIteration = false
		}

		// creates curves for each species populations through time
		var currentCurve chart.ContinuousSeries
		currentCurve.XValues = generations
		currentCurve.YValues = i

		var currentCurvePhase chart.ContinuousSeries
		currentCurvePhase.XValues = firstPopulation
		currentCurvePhase.YValues = i

		allCurves[curveIndex] = currentCurve
		allCurvesPhase[curveIndex] = currentCurvePhase
		curveIndex += 1

	}

	// sets up axes and title of the population dynamics graph
	var graph chart.Chart

	graph.XAxis = chart.XAxis{

		Name: "Generations",
		Style: GetAxisStyle(),

	}

	graph.YAxis = chart.YAxis{

		Name: "Species Populations",
		Style: GetAxisStyle(),

	}

	graph.Title = "Lotka Volterra Population Dynamics"
	graph.TitleStyle = GetTitleStyle()

	var graphPhase chart.Chart
	
	graphPhase.XAxis = chart.XAxis {
		Name: firstSpeciesName,
		Style: GetAxisStyle(),
	}

	graphPhase.YAxis = chart.YAxis {
		Name: "Species Populations",
		Style: GetAxisStyle(),
	}

	graphPhase.Title = "Phase Portraits of Ecosystem Species"
	graphPhase.TitleStyle = GetTitleStyle()
	

	for i := range(allCurves){
		graph.Series = append(graph.Series, allCurves[i])
		graphPhase.Series = append(graphPhase.Series, allCurvesPhase[i])
	}

	OutputGraph(graph, filename)
	OutputGraph(graphPhase, "phase.png")
}

// CreateGraph: creates graph of ecosystem with hard-coded initialized species and their corresponding values (initial population, interaction values... etc)
// input: generations - time throughout ecosystem, pred1Pops & pred2Pops & preyPops - populations of the 3 species over time in the ecosystem,
//        given filename of file to write graph to (string)
func CreateGraph(generations, pred1Pops, pred2Pops, preyPops []float64, filename string) {
	fmt.Println("start")
	xValues, yValues := generations, pred1Pops
	xValues2, yValues2 := generations, pred2Pops
	xValues3, yValues3 := generations, preyPops

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
		},
	}

	OutputGraph(graph, filename)
}

// OutputGraph: creates a file, renders the chart and puts that in the file
// input: graph - the given chart to be placed into the file, given filename
func OutputGraph(graph chart.Chart, filename string) {
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

// getTitleStyle: customizing style for chart title
func GetTitleStyle()chart.Style{
	return chart.Style{
		//Show: true,
		FontSize: 20,
		//FontColor: drawing.ColorBlue
		//FontColor: drawing.ColorBlue,
	}
}

// getAxisStyle: customizing style for chart axis
func GetAxisStyle()chart.Style{
	return chart.Style{
		//Show: true,
		FontSize: 14,
		//FontColor: drawing.ColorBlue,
	}
}
