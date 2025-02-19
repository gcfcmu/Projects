package main

// Species: one type of animal in the ecosystem
type Species struct {
  name string
  population float64
  growthRate float64
  naturalDeathRate float64
  carryingCapacity float64
  predator bool
}

// Matrix: the community matrix holding the interaction values between all pairs of species
type Matrix struct {
  interactionValues [][]float64
  speciesKeys map[Species]int
}

// Ecosystem: a set of species (in the ecosystem)
type Ecosystem []*Species
