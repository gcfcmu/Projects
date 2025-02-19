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

  dog.population = UpdatePopulation(Test.species, Test.time, Test.interactions)

  fmt.Println(dog.population)


  //Next, calculate by hand correct values and verify























}
