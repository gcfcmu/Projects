package main
import "testing"
import "fmt"

func TestInitializeInteractionMatrix(t *testing.T){

  type test struct{

    numRows int
    ecosystem Ecosystem
    answer int
    vals []float64


  }


  var Test test

  Test.answer = 3

  Test.numRows = 3

  Test.vals = make([]float64, 9)

  for i := range(Test.vals){

    Test.vals[i] = 0.0 + float64(i)

  }

  fmt.Println(Test.vals)

  var s1 Species
  s1.population = 10.0
  var s2 Species
  s2.population = 20.0
  var s3 Species
  s3.population = 30.0

  var eco Ecosystem
  eco = make([]*Species, Test.numRows)

  eco[0] = &s1
  eco[1] = &s2
  eco[2] = &s3


  outcome := InitializeInteractionMatrix(eco)
  x := 0.0




  for i := 0; i < len(outcome.interactionValues); i++{
    for j := 0; j < len(outcome.interactionValues[i]); j++{
      fmt.Println("x=", x)
      fmt.Println("outcome", outcome.interactionValues[i][j])
      outcome.interactionValues[i][j] += x
      x += 1.0
    }
  }

  outcome1 := len(outcome.interactionValues)

  outcome2 := len(outcome.interactionValues[1])

  pass := true

  if outcome1 != Test.answer || outcome2 != Test.answer {
    t.Error("Error! the correct number of rows is", Test.answer, "and your function gives", outcome1, "Rows and", outcome2, "cols")
    pass = false
  }

  k := 0

  for i := 0; i < Test.answer; i++{
    for j := 0; j < Test.answer; j++{

      if outcome.interactionValues[i][j] != Test.vals[k]{
        t.Error("Error! incorrect value given in matrix")
        pass = false
      }

      k += 1

    }
  }


  if pass{
    fmt.Println(outcome)
    fmt.Println("Correct! The matrix has lenght", len(outcome.interactionValues), "and the values are equal to", Test.vals)
}
}


func TestSetInteractionValues(t *testing.T){

  type test struct {

    matrix Matrix
    answer float64

  }

  var Test test

  var dog Species
  dog.name = "dog"
  dog.population = 10.0
  dog.growthRate = 1.0
  dog.carryingCapacity = 50.0

  var cat Species
  cat.name = "cat"
  cat.population = 15.0
  cat.growthRate = 1.0
  cat.carryingCapacity = 50.0

  eco := Ecosystem{&dog, &cat}

  Test.matrix = InitializeInteractionMatrix(eco)

  Test.answer = 1.0

  SetInteractionValues(1.0, &Test.matrix, dog, cat)

  if Test.answer != ReturnInteractionValues(Test.matrix, dog, cat){
    t.Error("Error! Set interaction and return interaction don't match up")
  } else {

    fmt.Println("Correct! Here is the interaction matrix:", Test.matrix)
  }





}


func TestUpdatePopulation(t *testing.T){



  type test struct {

    species Species
    time float64
    interactions Matrix
    answer float64

  }

  var Test test

  var dog Species
  dog.name = "dog"
  dog.population = 10.0
  dog.growthRate = 1.0
  dog.carryingCapacity = 50.0

  var cat Species
  cat.name = "cat"
  cat.population = 15.0
  cat.growthRate = 1.0
  cat.carryingCapacity = 50.0

  var rat Species
  rat.name = "rat"
  rat.population = 5.0
  rat.growthRate = 1.0
  rat.carryingCapacity = 50.0

  eco := Ecosystem{&dog, &cat, &rat}

  Test.species = dog

  Test.time = 1.0

  interact := InitializeInteractionMatrix(eco)


  SetInteractionValues(1.0, &interact, dog, cat)
  SetInteractionValues(1.1, &interact, dog, rat)
  SetInteractionValues(1.2, &interact, cat, dog)
  SetInteractionValues(1.3, &interact, cat, rat)
  SetInteractionValues(1.4, &interact, rat, dog)
  SetInteractionValues(1.5, &interact, rat, cat)

  fmt.Println(interact)




  Test.interactions = interact


  dog.population = UpdatePopulation(Test.species, 1, eco, Test.time, Test.interactions)

  fmt.Println(dog.population)


  //Next, calculate by hand correct values and verify























}

func TestReadNumGens(t *testing.T){

  type test struct {

    answer int


  }

  var Test test

  Test.answer = 600


  outcome := ReadNumGens("tests/numGensTestStatic.txt")


  if Test.answer != outcome {
    t.Error("Error! the number in numGens.txt is", outcome, "and it should be", Test.answer)
  }




}


func TestReadWriteGens(t *testing.T){

  type test struct {

    answer int
    answer2 int
    filename string


  }

  var Test test

  Test.answer = 7000
  Test.answer2 = 5000

  Test.filename = "tests/numGensTest.txt"


  WriteGensToFile([]string{"7000"}, Test.filename)
  outcome1 := ReadNumGens(Test.filename)
  WriteGensToFile([]string{"5000"}, Test.filename)
  outcome2 := ReadNumGens(Test.filename)


  if outcome1 != Test.answer || outcome2 != Test.answer2{

    t.Error("Inproper reading and writing functions. Outcomes are", outcome1, "and", outcome2)


  } else if outcome1 == Test.answer && outcome2 == Test.answer2 {

    fmt.Println("correctly read in and wrote files for numgens")

  }




}


func TestReadWriteTime(t *testing.T){

  type test struct {

    answer float64
    answer2 float64
    filename string


  }

  var Test test

  Test.answer = 0.5
  Test.answer2 = 0.24

  Test.filename = "tests/timeTest.txt"


  WriteTimeToFile([]string{"0.5"}, Test.filename)
  outcome1 := ReadTime(Test.filename)
  WriteTimeToFile([]string{"0.24"}, Test.filename)
  outcome2 := ReadTime(Test.filename)


  if outcome1 != Test.answer || outcome2 != Test.answer2{

    t.Error("Inproper reading and writing functions. Outcomes are", outcome1, "and", outcome2)


  } else if outcome1 == Test.answer && outcome2 == Test.answer2 {

    fmt.Println("correctly read in and wrote files for timestep")

  }

}



func TestNormal(t *testing.T){

  type test struct {

    population float64
    draws []float64

  }


  var Test test
  Test.population = 100.0
  Test.draws = make([]float64, 5)

  for i := range(Test.draws){

    Test.draws[i] = StochasticNormal(Test.population)

  }

  fmt.Println("Now printing testing draws for normal distribution")

  for i := range(Test.draws){

    fmt.Println(Test.draws[i])

  }

}

func TestUniform(t *testing.T){

  type test struct {

    population float64
    draws []float64
    pass bool


  }

  var Test test

  Test.population = 100.0
  Test.draws = make([]float64, 5)
  Test.pass = true


  for i := range(Test.draws){

    Test.draws[i] = StochasticUniform(Test.population)

    if Test.draws[i] < 0 {

      if Test.draws[i] < -1 {
        t.Error("draw is outside the range -1%")
        fmt.Println(Test.draws[i])
        Test.pass = false
      }


    } else {

      if Test.draws[i] > 1 {
        t.Error("draw is outside range +1%")
        fmt.Println(Test.draws[i])
        Test.pass = false

      }


    }






  }


  if Test.pass{
    fmt.Println("successfully passed Uniform test")

    fmt.Println("The values for uniform draw with population =", Test.population, "are:")
    for i := range(Test.draws){
      fmt.Println(Test.draws[i])
    }

  }



}



func TestWriteParams(t *testing.T){


  type test struct{

    names []string
    growths []string
    init []string
    carrying []string
    answer map[string][]float64
    answerVals []float64

  }



  var Test test

  Test.names = []string{"cow", "crow", "crane"}
  Test.growths = []string{"100.0", "200.0", "300.0"}
  Test.init = []string{"1000.0", "2000.0", "3000.0"}
  Test.carrying = []string{"1500.0", "4000.0", "6000.0"}
  Test.answer = make(map[string][]float64, 0)

  Test.answerVals = []float64{1500.0, 100.0, 1000.0}

  Test.answer["cow"] = Test.answerVals





  WriteParametersToFile(Test.names[0], Test.carrying[0], Test.growths[0], Test.init[0], Test.names[0] + ".txt")


  output := ReadParameters(Test.names[0] + ".txt")

  fmt.Println(output)





}


func TestReadMatrixFromWeb(t *testing.T){


  type test struct{

    answer []float64
    form map[string][]string

  }



  var Test test

  Test.answer = []float64{0.0, 1.1, 2.2, 3.3, 4.4, 5.5, 6.6,
    7.7, 8.8, 9.9, 10.0, 11.1, 12.2, 13.3, 14.4, 15.5}

  Test.form = make(map[string][]string)

  Test.form["interact"] = []string{"0.0", "4.4", "8.8", "12.2"}

  Test.form["col0"] = []string{"1.1", "5.5", "9.9", "13.3"}

  Test.form["col1"] = []string{"2.2", "6.6", "10.0", "14.4"}

  Test.form["col2"] = []string{"3.3", "7.7", "11.1", "15.5"}


  outcome := ReadMatrixFromWeb(Test.form)


  pass := true

  for i := range(outcome) {

    if outcome[i] != Test.answer[i]{

      t.Error("Error reading in matrix. The test dataset is", Test.answer[i], "but the outcome is", outcome[i])
      pass = false
    }

  }


  if pass == true {

    fmt.Println("Correctly read in the matrix and stored it as a slice")

  }







}


func TestMultipleSpeciesGraph(t *testing.T){


  type test struct {

    allSpecies map[string][]float64
    Generations []float64


  }

  var Test test

  Test.Generations = make([]float64, 0)


  for i := 0; i < 10; i++{

    Test.Generations = append(Test.Generations, float64(i))


  }

  var species1 []float64
  var species2 []float64
  var species3 []float64

  species1 = make([]float64, 10)
  species2 = make([]float64, 10)
  species3 = make([]float64, 10)

  population := 0.0

  for i := range(species1){

    species1[i] = population
    species2[i] = population + 1.0
    species3[i] = population + 2.0


  }

  species3[8] = 1000.0
  species2[3] = 1000.0

  Test.allSpecies = make(map[string][]float64, 0)

  Test.allSpecies["bear"] = species1
  Test.allSpecies["wolf"] = species2
  Test.allSpecies["rabbit"] = species3

  CreateGraphMultipleSpecies(Test.Generations, Test.allSpecies, "graph_test.png")


  //Check the png for correct output





}
