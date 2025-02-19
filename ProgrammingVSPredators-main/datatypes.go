package main

type Ecosystem struct {

  time float64
  population []*Organism
  grazers []*Organism
  producers []*Organism
  predators []*Organism
  decomposers []*Organism
  sun bool
  temperature float64

}

type Organism struct {

  name string
  predators []string
  prey []string
  energy float64
  producer bool
  reproduction float64
  circadian string
  habitat string
  diet []*Organism
  velocity OrderedPair
  mass float64
  vision float64


}

type OrderedPair struct {
	x float64
	y float64
}
