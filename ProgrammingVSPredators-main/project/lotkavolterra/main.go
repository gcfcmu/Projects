package main

import (
  "fmt"
  "net/http"
  "html/template"
  "os"
  "path"
  "encoding/json"
  "hash/fnv"
  "strconv"
)

func handlerFunc(w http.ResponseWriter, r *http.Request){

  fmt.Fprint(w, "<h1>Lotka Volterra Model for Population Dynamics<h1>")



  if r.Method == "GET" {
		t, err := template.ParseFiles("stuff/webpage.html")
		if err != nil {
			panic(err)
		}
		t.Execute(w, nil) // we could pass a struct in to apply formatting if we wanted
	}
}


func ImageHandler(w http.ResponseWriter, r *http.Request) {

    http.ServeFile(w, r, "./thomas.png")
}

type PostRequest struct {
	Key1 string
	Key2 string
}

type PostResponse struct {
	Msg  string
	Keys string
}

type Page struct {
	Title    string
	Contents template.HTML
}

// servePostRequest a helper method to encapsulate JSON shenanigans
func servePostRequest(w http.ResponseWriter, r *http.Request) {

  if r.Method == "POST"{// parse the post request
	var req PostRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// create response
	response := PostResponse{
		Msg:  "Hello! Thanks for the POST request.",
		Keys: fmt.Sprintf("%+v", req),
	}

	// issue response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}


}

func ParameterHandler(w http.ResponseWriter, r *http.Request){

  fmt.Println("parameter", r.Method)

  if r.Method == "GET" {

    r.ParseForm()

    t, err := template.ParseFiles("stuff/params.html")
    if err != nil {
      panic(err)
    }

    numGens, numGensCheck := r.Form["numGens"]
    timeStep, timeStepCheck := r.Form["timeStep"]

    speciesNames, speciesNamesCheck := r.Form["name"]
    carryingCapacity, carryingCapacityCheck := r.Form["capacity"]
    growthRate, growthRateCheck := r.Form["growth"]
    initialPopulation, initialCheck := r.Form["initial"]

    stochastic, stochasticCheck := r.Form["stoch"]
    magnitude, magnitudeCheck := r.Form["var"]

    fmt.Println(magnitudeCheck)
    fmt.Println(magnitude[0])

    if magnitude[0] == "" {

      magnitude = []string{"0.0"}

    }

    fmt.Println("MAGNITUDE", magnitude)

    matrixValuesSlice := ReadMatrixFromWeb(r.Form)

    //fmt.Println(matrixValuesSlice)



    for i := 0; i < len(r.Form["name"]); i++{

      if numGensCheck && timeStepCheck && speciesNamesCheck && growthRateCheck && initialCheck && carryingCapacityCheck{

      WriteParametersToFile(speciesNames[i], carryingCapacity[i], growthRate[i], initialPopulation[i], speciesNames[i] + ".txt")

    }

    }

    fmt.Println(r.Form)
    fmt.Println(numGens)
    fmt.Println(timeStep)




  if stochasticCheck && numGensCheck && timeStepCheck && numGens != nil && timeStep != nil{

    WriteGensToFile(numGens, "numGens.txt")
    WriteTimeToFile(timeStep, "timeStep.txt")

    unique := hash(numGens[0] + timeStep[0])

    plotFile := fmt.Sprintf("%d.png", unique)

    plotFile = "thomas.png"

    parsedMagnitude, err := strconv.ParseFloat(magnitude[0], 64)


    if err != nil{
      panic(err)
    }


    UpdateGraphParameters(plotFile, speciesNames, matrixValuesSlice, stochastic[0], parsedMagnitude)


    //fmt.Println("successfully created thomas")
    fmt.Println("things:", r.Form)



    //page := Page{
    //  Title: "Lotka Volterra Model",
    //  Contents: template.HTML(fmt.Sprintf(
    //    "<img src='/stuff' alt='graph.png' width='1024' height='400'>", plotFile,
    //  )),
    //}

    http.ServeFile(w, r, plotFile)


    t.Execute(w, nil)
  }



	}




}

func hash(s string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32()
}

func UpdateGraphParameters(filename string, speciesInputs []string, interactionVals []float64, stochastic string, magnitude float64){

  numGens := ReadNumGens("numGens.txt")
  timeStep := ReadTime("timeStep.txt")
  parameterMaps := make([]map[string][]float64, 0)

  var speciesList Ecosystem
  speciesList = make([]*Species, 0)

  for i := range(speciesInputs){

    parameterMaps = append(parameterMaps, ReadParameters(speciesInputs[i] + ".txt"))



  }

  for i := range(parameterMaps){

    curMap := parameterMaps[i]

    for key, _ := range(curMap){

      curName := key

      curSpecies := InitializeSpecies(curName, curMap[key][0], curMap[key][1], 0, curMap[key][2], true)

      speciesList = append(speciesList, &curSpecies)

    }


  }

  fmt.Println("parameter maps:", parameterMaps)
  fmt.Println("species List:", speciesList)

  //bears := InitializeSpecies("bear", 70.0, -40.0, 0, 500, true)
  //wolves := InitializeSpecies("wolf", 50.0, -40.0, 0, 500, true)
  //rabbits := InitializeSpecies("rabbit", 600.0, 80.0, 0, 1000, true)

  //var speciesList Ecosystem

  //speciesList = append(speciesList, &bears)
  //speciesList = append(speciesList, &wolves)
  //speciesList = append(speciesList, &rabbits)

  //TODO read in from webpage COMPLETE, TESTING
  //interactionVals := []float64{-0.030, -0.001, 0.095, -0.003, -0.03, 0.095, -0.82,-0.82,-0.0172}

  fmt.Println(interactionVals)

  // set up the interaction matrix
  interactionMatrix := InitializeInteractionMatrix(speciesList)

  fmt.Println("matrix before", interactionMatrix)

  for i := range(speciesList){
    for j := range(speciesList){

      SetInteractionValues(interactionVals[i*len(speciesList)  +j], &interactionMatrix, *(speciesList)[i], *(speciesList)[j])
      fmt.Println((speciesList)[i].name, (speciesList)[j].name, interactionVals[i*len(speciesList)  +j])

    }
  }

  //fmt.Println("matrix:", interactionMatrix)
  /*
  // set interaction values between bears and wolves (competition)
  SetInteractionValues(-0.001, &interactionMatrix, bears, wolves)
  SetInteractionValues(-0.003, &interactionMatrix, wolves, bears)

  // set interaction values between bears and rabbits (predation)
  SetInteractionValues(0.095, &interactionMatrix, bears, rabbits)
  SetInteractionValues(-0.82, &interactionMatrix, rabbits, bears)

  // set interaction values between wolves and rabbits (predation)
  SetInteractionValues(0.095, &interactionMatrix, wolves, rabbits)
  SetInteractionValues(-0.82, &interactionMatrix, rabbits, wolves)

  // set negative self-interaction values (serves as "carrying capacity")
  SetInteractionValues(-0.030, &interactionMatrix, bears, bears)
  SetInteractionValues(-0.030, &interactionMatrix, wolves, wolves)
  SetInteractionValues(-0.0172, &interactionMatrix, rabbits, rabbits)
  */

  var updatedPopulations map[string][]float64

  if stochastic == "none"{

    updatedPopulations = LotkaVolterra(timeStep, speciesList, interactionMatrix, numGens)

  } else {

    updatedPopulations = StochasticLotkaVolterra(timeStep, speciesList, interactionMatrix, numGens, stochastic, magnitude)

  }



  fmt.Println("successfully created simulation")

  // acquire the slice of generations
  generations := make([]float64, numGens)

  for i := 0; i < numGens; i++ {
    generations[i] = float64(i)
  }

  //fmt.Println(updatedPopulations[speciesInputs[0]], updatedPopulations[speciesInputs[1]], updatedPopulations[speciesInputs[2]], filename)
  // create the population dynamics graph
  //CreateGraph(generations, updatedPopulations[speciesInputs[0]], updatedPopulations[speciesInputs[1]], updatedPopulations[speciesInputs[2]], filename)

  CreateGraphMultipleSpecies(generations, updatedPopulations, filename)

}

func getUrlHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method, "on", "get_url") //get request method
	if r.Method == "GET" {
		t, err := template.ParseFiles(path.Join("stuff", "dowload.jpg"))
		if err != nil {
			panic(err)
		}
		t.Execute(w, nil) // we could pass a struct in to apply formatting if we wanted
	}
}

func infoHandler(w http.ResponseWriter, r *http.Request){

  if r.Method == "GET" {
		t, err := template.ParseFiles("stuff/info.html")
		if err != nil {
			panic(err)
		}

    http.FileServer(http.Dir("/stuff"))



		t.Execute(w, nil) // we could pass a struct in to apply formatting if we wanted
	}

}


func movieHandler(w http.ResponseWriter, r *http.Request){

      http.ServeFile(w, r, "/stuff/tutorialnumgens.png")
}


func main(){


  // setting up species - case 1: 60, 60, 600 rabbit should go down, preds go up...
  bears := InitializeSpecies("bear", 70.0, -40.0, 0, 500, true)
  wolves := InitializeSpecies("wolf", 50.0, -40.0, 0, 500, true)
  rabbits := InitializeSpecies("rabbit", 600.0, 80.0, 0, 1000, true)

  // case 2: 100, 100, 200 need rabbits to go up, preds go down...
  // bears := InitializeSpecies("bear", 100.0, -40.0, 0, 500, true)
  // wolves := InitializeSpecies("wolf", 100.0, -40.0, 0, 500, true)
  // rabbits := InitializeSpecies("rabbit", 200.0, 200.0, 0, 1000, true)

  // case 3: prey and preds equal pop...
  // bears := InitializeSpecies("bear", 200.0, -40.0, 0, 500, true)
  // wolves := InitializeSpecies("wolf", 200.0, -40.0, 0, 500, true)
  // rabbits := InitializeSpecies("rabbit", 200.0, 200.0, 0, 1000, true)

  // create an ecosystem, which contains a slice of species
  var speciesList Ecosystem

  // add existing species to our ecosystem
  speciesList = append(speciesList, &bears)
  speciesList = append(speciesList, &wolves)
  speciesList = append(speciesList, &rabbits)

  // set up the interaction matrix
  interactionMatrix := InitializeInteractionMatrix(speciesList)

  // set interaction values between bears and wolves (competition)
  SetInteractionValues(-0.001, &interactionMatrix, bears, wolves)
  SetInteractionValues(-0.003, &interactionMatrix, wolves, bears)

  // set interaction values between bears and rabbits (predation)
  SetInteractionValues(0.095, &interactionMatrix, bears, rabbits)
  SetInteractionValues(-0.82, &interactionMatrix, rabbits, bears)

  // set interaction values between wolves and rabbits (predation)
  SetInteractionValues(0.095, &interactionMatrix, wolves, rabbits)
  SetInteractionValues(-0.82, &interactionMatrix, rabbits, wolves)

  // set negative self-interaction values (serves as "carrying capacity")
  SetInteractionValues(-0.030, &interactionMatrix, bears, bears)
  SetInteractionValues(-0.030, &interactionMatrix, wolves, wolves)
  SetInteractionValues(-0.0172, &interactionMatrix, rabbits, rabbits)



  // set up generations and time parameters
  numGens := 8000 // 8000
  timeStep := 0.0001 //.0001
  // run LotkaVolterra model
  updatedPopulations := LotkaVolterra(timeStep, speciesList, interactionMatrix, numGens)

  // acquire the slice of generations
  generations := make([]float64, numGens)
  for i := 0; i < numGens; i++ {
    generations[i] = float64(i)
      if i % 1 == 0 {
        fmt.Println("Generation: ", i)
        fmt.Println("Total Population (original 400): ", updatedPopulations["bear"][i] + updatedPopulations["wolf"][i] + updatedPopulations["rabbit"][i])
        fmt.Println(updatedPopulations["bear"][i])
        fmt.Println(updatedPopulations["wolf"][i])
        fmt.Println(updatedPopulations["rabbit"][i])
        fmt.Println()
      }
  }

  // create the population dynamics graph
  CreateGraph(generations, updatedPopulations["bear"], updatedPopulations["wolf"], updatedPopulations["rabbit"], "thomas.png")

  fmt.Println(speciesList[0].name, speciesList[1].name, speciesList[2].name)

  http.HandleFunc("/tutorialnumgens.png", movieHandler)

  http.HandleFunc("/", handlerFunc)
  http.HandleFunc("/get_url/", getUrlHandler)
  http.HandleFunc("/parameters/", ParameterHandler)
  http.HandleFunc("/info/", infoHandler)




  fs := http.FileServer(http.Dir("morestuff"))
  http.Handle("/morestuff/", http.StripPrefix("/morestuff/", fs))

  http.HandleFunc("/thomas.png", ImageHandler)





  _ = os.MkdirAll("morestuff", 0o755)
	//http.Handle("/morestuff/", http.StripPrefix("/morestuff/", http.FileServer(http.Dir("./morestuff"))))
	http.Handle("/stuff/", http.StripPrefix("/stuff/", http.FileServer(http.Dir("./stuff"))))



  fmt.Println("Starting server on :3000...")
  err := http.ListenAndServe(":3000", nil)

  if err != nil{
    panic(err)
  }






}
