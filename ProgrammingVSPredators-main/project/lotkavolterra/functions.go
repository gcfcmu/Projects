package main

import (
 //"fmt"
 "math/rand"
)


// StochasticUniform: adds noise to population data (percentage based)
// input: given species population and magnitude of random noise to population value
// output: the population value that has been randomly altered (percentage based)
func StochasticUniform(population float64, magnitude float64) float64{

  percentChange := rand.Float64()
  percentChange = percentChange * magnitude

  sign := rand.Intn(2)
  if sign == 0{
    sign = -1

  } else {

    sign = 1
  }

  percentChange = percentChange * float64(sign)

  return population * (percentChange)/100

}

// StochasticNormal: introduces a small amount of randomization (from normal distribution) to the populaton
// input: given population that will be slightly modified randomly and magnitude of that randomization
// output: the slightly, randomly modified population value
func StochasticNormal(population float64, magnitude float64) float64{

  std := (population / 100) * magnitude

  sample := rand.NormFloat64() * std

  return sample


}

// LotkaVolterra: simulates the population dynamics of all species in the ecosystem over time
// input: timeStep - the time between generations,
//         speciesList - all species in the ecosystem,
//         m - interaction matrix,
//         numGens - number of generations
// output: a map of species' names (string) to a slice of populations from each generation
func StochasticLotkaVolterra(timeStep float64, speciesList Ecosystem, m Matrix, numGens int, randomType string, magnitude float64) map[string][]float64 {

  // create map to keep track of all species' populations over the generations
  speciesPopulations := make(map[string][]float64)

  // populate map with all species in ecosystem, assign their initial populations
  for _, species := range speciesList {
    speciesPopulations[species.name] = make([]float64, numGens)
    speciesPopulations[species.name][0] = species.population
  }

  // for each generation, update each species' population
  for i := 1; i < numGens; i++ {
    for index, species := range speciesList {
      species.population += UpdatePopulation(*species, index, speciesList, timeStep, m)

      if randomType == "normal"{

        species.population += StochasticNormal(species.population, magnitude)


      } else if randomType == "uniform"{

        species.population += StochasticUniform(species.population, magnitude)

      }



      speciesPopulations[species.name][i] = species.population
    }
  }
  return speciesPopulations
}

// LotkaVolterra: simulates the population dynamics of all species in the ecosystem over time
// input: timeStep - the time between generations,
//         speciesList - all species in the ecosystem,
//         m - interaction matrix,
//         numGens - number of generations
// output: a map of species' names (string) to a slice of populations from each generation
func LotkaVolterra(timeStep float64, speciesList Ecosystem, m Matrix, numGens int) map[string][]float64 {

  // create map to keep track of all species' populations over the generations
  speciesPopulations := make(map[string][]float64)

  // populate map with all species in ecosystem, assign their initial populations
  for _, species := range speciesList {
    speciesPopulations[species.name] = make([]float64, numGens)
    speciesPopulations[species.name][0] = species.population
  }

  // for each generation, update each species' population
  for i := 1; i < numGens; i++ {
    for index, species := range speciesList {
      species.population += UpdatePopulation(*species, index, speciesList, timeStep, m)
      if species.population > species.carryingCapacity*10{

        species.population = species.carryingCapacity*10

      }
      speciesPopulations[species.name][i] = species.population
    }
  }

  // return all species and their populations during each generation
  return speciesPopulations
}

// UpdatePopulation: updates the population of target species accounting for natural growth rate and other species in the ecosystem
// input: species - target species whose population will be updated
//        index - index of species in community matrix (interaction matrix)
//        speciesList - list of all species in the ecosystem
//        timeStep - time between each generation
//        interactionMatrix - the community matrix representing each species' effect on one another
// output: how much the current population should increase/decrease (to become the new population)
func UpdatePopulation(species Species, index int, speciesList Ecosystem, timeStep float64, interactionMatrix Matrix) float64 {

  // initialize interaction to 0 (start with "no effect")
  interaction := 0.0

  // Go through the species keys and add each species' effect on target species (including target species' effect on self)
  for i := 0; i < len(speciesList); i++ {
    interaction += interactionMatrix.interactionValues[index][i] * speciesList[i].population
    // interaction += ReturnInteractionValues(interactionMatrix, species, *speciesList[i]) * speciesList[i].population
  }

  // compute the new change in population
  dy := timeStep*species.population*(species.growthRate + interaction)

  // return the new change in population
  return dy
}

// InitalizeInteractionMatrix: sets up the community matrix that holds each species' effect on one another
// input: eco - the ecosystem containing all the species that will affect one another
// output: the community matrix (with length & width = number of species)
func InitializeInteractionMatrix(eco Ecosystem) Matrix {

  // set up matrix
  matrix := make([][]float64, len(eco))

  for i := range(matrix){
    matrix[i] = make([]float64, len(eco))
  }

  // initialize all values in matrix to 0
  for i := range(eco){
    for j := range(eco){
      matrix[i][j] = 0.0
    }
  }

  // set up data structure Matrix
  var m Matrix

  // create map of species to the index they belong to on the community matrix
  mapp := make(map[Species]int)

  // assign species to index on community matrix based on order they appear in the ecosystem's list of species
  for i := range(eco){
    mapp[*eco[i]] = i
  }

  // assign Matrix's field values (the community matrix and species to index map)
  m.interactionValues = matrix
  m.speciesKeys = mapp

  // return the community matrix and species-index map
  return m
}

// SetInteractionValues: places the interaction effect of species 2 on species 1 into the community matrix
// input: value - the interaction value of species 2 on species 1
//         m - the community matrix
//         s1, the species being affected
//         s2, the species causing the interaction effect
// output: none (interaction value will be stored into matrix)
func SetInteractionValues(value float64, m *Matrix, s1 Species, s2 Species) {
  // assigns interaction value to cell in community matrix corrsponding to both species
  // (affected species is row index, affector species is column index)
  m.interactionValues[m.speciesKeys[s1]][m.speciesKeys[s2]] = value
}

// ReturnInteractionValues: acquires the interaction effect that a species has on another species
// input: m - community matrix
//         s1 - the affected species
//         s2 - the affector species
// output: the interaciton value between the two species
func ReturnInteractionValues(m Matrix, s1 Species, s2 Species) float64 {
  return m.interactionValues[m.speciesKeys[s1]][m.speciesKeys[s2]]
}

// InitializeSpecies: sets up a species with the specified fields
// input: name - species' name
//        population - initial population of species
//        growthRate - growth rate of species (percent based)
//        naturalDeathRate - death rate of species due to natural, non-violent causes
//        carryingCapacity - population limit of species in ecosystem
//        predator - whether the species is a predator or not
// output: none (species will be created)
func InitializeSpecies(name string, population, growthRate, naturalDeathRate, carryingCapacity float64, predator bool) Species {
  var animal Species
  animal.name = name
  animal.population = population
  animal.growthRate = growthRate
  animal.naturalDeathRate = naturalDeathRate
  animal.carryingCapacity = carryingCapacity
  animal.predator = predator
  return animal
}
