package main

import (
  "os"
  "log"
  "bufio"
  "strconv"
  "fmt"
  "strings"
  "io/ioutil"
)


// ReadGensFromFile: reads the number of generations from the given file name
// input: directory and filename are both strings that provide the path to the file and the file itself
// output: obtain the number of generations from the file
func ReadGensFromFile(directory string, filename string) int {

  // track number of generations
  var gens int

  // try to access the file
  file, err := os.Open(filename)
  if err != nil {
    log.Fatal(err)
  }
  defer file.Close()

  // read the file contents
  fileContents, err := ioutil.ReadFile(directory + filename)
	if err != nil {
		panic(err)
	}

  // read in the contents line by line and 
  inputLines := strings.Split(strings.TrimSpace(strings.Replace(string(fileContents), "\r\n", "\n", -1)), "\n")

  fmt.Println(inputLines)

  // return the number of generation
  return gens
}

// WriteGensToFile: write the number of generations (user input) to a text file
// input: generations as a slice of strings to write to file, filename of file we are creating
func WriteGensToFile(gens []string, filename string) {
  // create file
	file, err := os.Create(filename)
	if err != nil {
		log.Fatal("Error creating numGens.txt", err)
	}
	defer file.Close()

  // dump generations values
	fmt.Fprintln(file, gens[0])
}

// WriteTimeToFile: write the time frame values to text file
// input: time steps as slice of strings to write to file, filename of file we are creating
func WriteTimeToFile(time []string, filename string) {
  // create file
  file, err := os.Create(filename)
  if err != nil {

    log.Fatal("Error creating timeStep.txt", err)

  }
  defer file.Close()

  // dump time step values
  fmt.Fprintln(file, time[0])
}

// ReadNumGens: reads the number of generations from the given file
// input: the file to read from
func ReadNumGens(filename string) int {

  var numGens int

  // access file
  file, err := os.Open(filename)
  if err != nil{

    log.Fatal("Error reading gens", err)

  }

  defer file.Close()

  // go through line by line and read numGens value
  scanner := bufio.NewScanner(file)
  scanner.Scan()
  currentLine := scanner.Text()

  numGens, err2 := strconv.Atoi(currentLine)

  if err2 != nil{
    panic("Could not read numGens")
  }

  if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

  return numGens
}

// ReadTime: reads the time value from given file
// input: given filename for file
func ReadTime(fileName string) float64 {


	//now, read out the file
	fileContents, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	//trim out extra space and store as a slice of strings, each containing one line.
	outputLines := strings.Split(strings.TrimSpace(strings.Replace(string(fileContents), "\r\n", "\n", -1)), "\n")

	//parse the float
	answer, err := strconv.ParseFloat(outputLines[0], 64)

	if err != nil {
		panic(err)
	}

	return answer
}


// WriteParametersToFile: writing the species and their parameters to text file with given name
// input: species' name, carrying capacity, growthRate and initial population (all characteristics of species population), given filename to write values to
func WriteParametersToFile(species string, carryingCapacity string, growthRate string, initialPopulation string, filename string){
 
  // create file to write values to
  file, err := os.Create(filename)
  if err != nil {

    log.Fatal("Error creating name.txt", err)

  }
  defer file.Close()

  // write all the species population characteristics to the new file
  fmt.Fprintln(file, species)
  fmt.Fprintln(file, initialPopulation)
  fmt.Fprintln(file, growthRate)
  fmt.Fprintln(file, carryingCapacity)
}

// ReadParameters: reads the given filename and construct a map of species names' to their initial population, growth rate and carrying capacity values
// input: the given name of the file to read values from
func ReadParameters(fileName string) map[string][]float64 {

  // set up map of species names' to their population characteristics
  output := make(map[string][]float64)
  values := make([]float64, 3)

  //now, read out the file
  fileContents, err := ioutil.ReadFile(fileName)
  if err != nil {
    panic(err)
  }

  //trim out extra space and store as a slice of strings, each containing one line.
  outputLines := strings.Split(strings.TrimSpace(strings.Replace(string(fileContents), "\r\n", "\n", -1)), "\n")
  fmt.Println("outputlines", outputLines)

  // populate each species key in map with that species' population characteristics
  for i := 1; i < 4; i++{

    // parse in the population characteristics float64 values
    curVal, err := strconv.ParseFloat(outputLines[i], 64)
    fmt.Println("curval", curVal)

    if err != nil{
      panic(err)
    } else {
      values[i-1] = curVal
    }
  }

  // put the population characteristics values into the output map
  output[outputLines[0]] = values

  return output
}

// ReadMatrixFromWeb: converts a map of all species in ecosystem and their characteristics
//                    into a super long slice of just their characterstics (numerical values)
func ReadMatrixFromWeb(m map[string][]string) []float64 {


  firstCol := m["interact"]

  length := len(firstCol)

  // set up super long slice that will contain matrix values (values from the map of species to their characteristics)
  var matrixValuesSlice []float64
  matrixValuesSlice = make([]float64, length*length)

  // populate the super long slice with all numerical values from the matrix/map
  for k := range(firstCol) {

    parsedFloat, err := strconv.ParseFloat(firstCol[k], 64)

    matrixValuesSlice[k*length] = parsedFloat

    if err != nil{

      panic(err)

    }


  }

  i := 0
  curCol := 1

  for i < length {

      key := ("col" + strconv.Itoa(i))

      if m[key] != nil {

          for j := range(m[key]){


              parsedFloat, err := strconv.ParseFloat(m[key][j], 64)

              if err != nil {

                panic(err)


              }

              matrixValuesSlice[j*length + curCol] = parsedFloat


          }


      }

      curCol = curCol + 1

      i = i + 1

  }

  // return the final super long slice with every species' population characteristics
  return matrixValuesSlice

}
