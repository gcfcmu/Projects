package main

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"text/tabwriter"
	"time"
)
package main


//Ecosystem is a structure with the following fields
type EcoSystem struct {

  width float64
  time float64 // Current year of ecosystem, set to 0 initially and update by 1 each timestep
  preys []*Prey
  plants []*Plant
  predators []*Predator

}

//Organism is a structure with the following fields
//It has all fields common to both animal and plants
type Organism struct {

  name string
  reproduction float64 ///
  age float64
  position OrderedPair
  mass float64
  energy float64

}

//Animal is a structure with the following fields
//It has all the organism fields (OrganismProperties) as well as
//fields unique to animals
type Animal struct {

   OrganismProperties Organism
   vision float64
   velocity OrderedPair
   accel OrderedPair
   maxSpeed float64
   speed  float64
}

//Plant is a structure with all the organism fields (OrganismProperties)
//It has no unique plant specific fields
type Plant struct {
OrganismProperties  Organism
}

//Predator is a structure with all animal fields,
//as well as an array of prey which forms its diet.
type Predator struct {
       anAnimal  Animal
       diet []*Prey
       //preys []string
}
//Predator is a structure with all animal fields,
//as well as a plant field which forms its diet.
type Prey struct {
      anAnimal Animal
      aPlant Plant
}

//OrderedPair for acceleration, velocity and positions.
type OrderedPair struct {
	x float64
	y float64
}


// Sample initialization values
var energyConstant float64     = 0.8    // Predefined energy constant
var canvasWidth int            = 1000   // Width of the ecosystem
var numPlants int              = 500    // Initial Number of plants
var numPreys int               = 100    // Initial Number of prey
var numPredators int           = 50     // Initial Number of predators
var preyDecayConstant int      = 1      // Constant decay rate of prey energy  (Figure out what to set)
var predatorDecayConstant int   = 0.5   // Constant decay rate of predator energy (Figure out what to set)
var initialEcosystem EcoSystem


// CalcDistance: calculates the distance between two living species
func CalcDistance(org1, org2 Organism) float64 {

      var SqrOfdeltaX =  (org1.position.x - org2.position.x) * (org1.position.x - org2.position.x)
      var SqrOfdeltaY  = (org1.position.y - org2.position.y) * (org1.position.y - org2.position.y)

      var distance = math.Sqrt(SqrOfdeltaX +SqrOfdeltaY)

      return distance
}

// CheckMaxSpeed: checks speed of a sinlge animal and recalibrates it to max speed when crossed
func CheckMaxSpeed(anAnimal Animal) float64 {

  var currentSpeed = math.Sqrt((anAnimal.velocity.x *anAnimal.velocity.x) + (anAnimal.velocity.y*anAnimal.velocity.y))

	if currentSpeed > anAnimal.maxSpeed {
		anAnimal.speed = anAnimal.maxSpeed
	}
	return anAnimal.speed

}

//Function for when the prey eats up the plant
func HandleConsumptionPreyPlant(aPrey Prey, aPlant Plant) {

	//Prey gains energy proportional to plants mass
  aPrey.OrganismProperties.energy += energyConstant * aPlant.OrganismProperties.mass

	//Plant dies
  aPlant.OrganismProperties.mass = 0

}

//Function for when the predator eats up the prey
func HandleConsumptionPredatorPrey(aPredator Predator, aPrey Prey) {

	// Predator gains energy proportional to preys mass
  aPredator.OrganismProperties.energy += energyConstant * aPrey.OrganismProperties.mass
  // Prey dies
  aPrey.OrganismProperties.mass = 0 // Marked as 0 for deletion

}

// Function for constant decay of energy of a predator over time
func UpdateEnergyPredator(allPredators []*Predator){

			for i:=0;i<len(allPredator);i++ {
			  aPredator := allPredator[i]
			  aPredator.OrganismProperties.energy =  aPredator.OrganismProperties.energy - predatorDecayConstant
			}
}

// Function for constant decay of energy of a prey over time
func UpdateEnergyPrey(allPreys []*Prey){

			for i:=0;i<len(allPreys);i++ {
			  aPrey := allPreys[i]
			  aPrey.OrganismProperties.energy =  aPrey.OrganismProperties.energy - preyDecayConstant
			}
}

// Function to set up the initial ecosystem
func InitializeEcosystem(anEcoSystem EcoSystem) {
	anEcoSystem.width         = 500
	anEcoSystem.time 			    = 0
	anEcoSystem.plants		    = InitializePlants(numPlants, canvasWidth)
	anEcoSystem.predators 		= InitializePredators(numPredators, canvasWidth)
	anEcoSystem.preys   	    = InitializePrey(numPreys, canvasWidth)
}

// Function to place plant objects on the initial ecosystem
func InitializePlants(numPlants int, canvasWidth int ) []*Plant{

    rand.Seed(time.Now().UnixNano())
	  minPos := 0

    var aPlant *Plant
    var allPlants []*Plant = make([]*Plant,numPlants)
    for i:=0; i<numPlants; i++ {
	  aPlant = new(Plant)
      aPlant.OrganismProperties .age = 1
      aPlant.OrganismProperties .mass = 1
      aPlant.OrganismProperties .position.x = (float64) (minPos + rand.Intn((canvasWidth - minPos)))
      aPlant.OrganismProperties .position.y = (float64) (minPos + rand.Intn((canvasWidth - minPos)))

      allPlants[i] = aPlant
    }
  return allPlants
}

// Function to place prey objects on the initial ecosystem
func InitializePrey(numPreys int, canvasWidth int) []*Prey{

  rand.Seed(time.Now().UnixNano())
  minPos := 0

  var aPrey *Prey
  var allPreys []*Prey = make([]*Prey,numPreys)
  for i:=0; i<numPreys; i++ {
	aPrey = new(Prey)

    aPrey.OrganismProperties.energy = 100   // Can randomize too
    aPrey.OrganismProperties.position.x = (float64) (minPos + rand.Intn((canvasWidth - minPos)))
    aPrey.OrganismProperties.position.y = (float64) (minPos + rand.Intn((canvasWidth - minPos)))

  	allPreys[i] = aPrey
  }
  return allPreys
}

// Function to place predator objects on the initial ecosystem
func InitializePredators(numPredators int, canvasWidth int) []*Predator{

  rand.Seed(time.Now().UnixNano())

  minPos := 0

  var aPredator *Predator
  var allPredators []*Predator
  for i:=0; i<numPredators; i++ {

	aPredator = new(Predator)
    aPredator.OrganismProperties.energy = 100   // Can randomize too
    aPredator.OrganismProperties.position.x = (float64) (minPos + rand.Intn((canvasWidth - minPos)))
    aPredator.OrganismProperties.position.y = (float64) (minPos + rand.Intn((canvasWidth - minPos)))

    allPredators[i] = aPredator
  }
  return allPredators
}

// CopyEcosystem: copies the ecosystem so that a series of ecosystems can be generated for simulating population dynamics
func CopyEcosystem(currentEcosystem Ecosystem) EcoSystem {

  var newEcoSystem EcoSystem

	newEcoSystem.width = currentEcoSystem.width
  newEcoSystem.time = currentEcoSystem.time

	//let's make the new ecosystems slice of plant objects
	numPlants := len(currentEcoSystem.plants)
	newEcoSystem.plants = make([]*Plant, numPlants)

	//now, copy all of the ecosystem fields into our new ecosystem
  for i := range currentEcoSystem.plants {

    newEcoSystem.plants[i] = new(Plant)
  	newEcoSystem.plants[i].OrganismProperties .age = currentEcoSystem.plants[i].OrganismProperties .age
    newEcoSystem.plants[i].OrganismProperties .mass = currentEcoSystem.plants[i].OrganismProperties .mass
    newEcoSystem.plants[i].OrganismProperties .position = currentEcoSystem.plants[i].OrganismProperties .position

	}

  numPrey := len(currentEcoSystem.prey)
	newEcoSystem.prey = make([]*Prey, numPrey)

	//now, copy all of the ecosystem fields into our new ecosystem
  for i := range currentEcoSystem.prey {

    newEcoSystem.prey[i] = new(Prey)
  	newEcoSystem.prey[i].OrganismProperties .age = currentEcoSystem.prey[i].OrganismProperties .age
    newEcoSystem.prey[i].OrganismProperties .mass = currentEcoSystem.prey[i].OrganismProperties .mass
    newEcoSystem.prey[i].OrganismProperties .position = currentEcoSystem.prey[i].OrganismProperties .position
    newEcoSystem.prey[i].OrganismProperties .energy = currentEcoSystem.prey[i].OrganismProperties .energy

    newEcoSystem.prey[i].accel = currentEcoSystem.prey[i].accel
    newEcoSystem.prey[i].vision = currentEcoSystem.prey[i].vision
    newEcoSystem.prey[i].velocity = currentEcoSystem.prey[i].velocity
    newEcoSystem.prey[i].maxSpeed = currentEcoSystem.prey[i].maxSpeed
    newEcoSystem.prey[i].speed = currentEcoSystem.prey[i].speed
    newEcoSystem.prey[i].diet = currentEcoSystem.prey[i].diet
	}

  numPredator := len(currentEcoSystem.predator)
  newEcoSystem.predator = make([]*Predator, numPredator)

  //now, copy all of the ecosystem fields into our new ecosystem
  for i := range currentEcoSystem.predator {

    newEcoSystem.predator[i] = new(Predator)
    newEcoSystem.predator[i].OrganismProperties.age = currentEcoSystem.predator[i].OrganismProperties.age
    newEcoSystem.predator[i].OrganismProperties.mass = currentEcoSystem.predator[i].OrganismProperties.mass
    newEcoSystem.predator[i].OrganismProperties.position = currentEcoSystem.predator[i].OrganismProperties.position
    newEcoSystem.predator[i].OrganismProperties.energy = currentEcoSystem.predator[i].OrganismProperties.energy

    newEcoSystem.predator[i].accel = currentEcoSystem.predator[i].accel
    newEcoSystem.predator[i].vision = currentEcoSystem.predator[i].vision
    newEcoSystem.predator[i].velocity = currentEcoSystem.predator[i].velocity
    newEcoSystem.predator[i].maxSpeed = currentEcoSystem.predator[i].maxSpeed
    newEcoSystem.predator[i].speed = currentEcoSystem.predator[i].speed
    newEcoSystem.predator[i].diet = currentEcoSystem.predator[i].diet
  }

	return newEcosystem
}


//Incomplete Functions
/*+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++


// UpdateEcosystem: updates kinematics of all living organisms within the ecosystem
func UpdateEcosystem(anEcosystem Ecosystem)   {
/*
	range over all plants and animals (predators and preys separately)

	define consumptionRange to be some small value
	ranging through predators
	  lower its energy by some pre-defined constant (small value like 1)
	  if energy <= 0, predator dies
	  ranging through preys (for each predator)
	if predator can consume prey (calcDistance(pred, prey)) <= consumption range && predator.energy < 80 {
		// this stuff is handled in HandleConsumption()
		predator eats prey
		predator gains like 40 energy or so
		prey disappears from the slice of preys
	}
	if CanChasePrey(specific predator) == true {
		update predator's accel, vel and pos accordingly (connect this to ChasePrey())
		// may want to make loss of energy rate even higher when chasing.... (optional)

		// these are updated in ChasePrey(), not in UpdateEcosystem()
		accel := a relatively high value
		vel: updated based on new accel
		pos: updated based on new accel
	} else {
		Update acceleration, velocity and position in a manner that represents "random walking"
		accel := random low value
		vel: updated based on new accel
		pos: updated based on new accel
	}
	// ranging through preys
	// 	lower its energy by some pre-defined constant (small value like 1)
	//   	if energy <= 0, prey dies

	if prey can consume plant and energy < 80, it does so... and gains like 40 energy or something

	    ranging through predators
	if NeedsToRunFromPredator(specific prey) == true { //this can be based of closest predator
		update prey's accel, vel and pos accordingly (connect this to ChasePrey())
		// two methods for updating kinematics
		1) prey keeps track of all repulsion forces from predators and runs according to that
		2) prey just runs directly opposite from closest predator ()
		--prey.accel.x = scalar(prey.pos.x - pred.pos.x) --> make this slightly randomized
		--prey.accel.y = scalar(prey.pos.x - pred.pos.x) --> make this slightly randomized
		// these are updated in RunAwayFromPredator(), not this function
		accel := a relatively high value
		vel: updated based on new accel
		pos: updated based on new vel
	} else if (canConsumePlant == true) { // if prey is not in danger from predator, and near plant, prey goes to eat plant
		accel := medium value?
		vel: updated based on new accel
		pos: updated based on new vel
	} else {
		Update acceleration, velocity and position in a manner that represents "random walking"
		accel := random low value
		vel: updated based on new accel
		pos: updated based on new accel
	}


	// might not want to call updateAcceleration, velocity and position here
	// we can just handle updating these via the chase, consumption and run away functions

		i.e chase() will speed up predator's kinematics
		runaway() will speed up prey's kinematics
		consumption() removes prey from the board and returns predator's kinematics to "normal" (random walking)

	//Update their kinematics (accel, vel, pos) accordingly
*/
}

// UpdateAcceleration: updates the acceleration of the animal
func UpdateAcceleration() {
	// incor
}
// UpdateVelocity: updates the velocity of the animal
func UpdateVelocity() {
}
// UpdatePosition: updates the position of the animal
func UpdatePosition() {
}
// CanChasePrey: determines if a predator is able to chase prey
func CanChasePrey() {
	// uses sight radius to determine if predator will chase the closest prey to it
/*
	Range through all prey
		Keep track of all prey that are within sight radius of selected predator
			For all prey within sight radius, make predator chase closest prey
				call ChasePrey() on predator and set it to move towards the prey, accelerating till max speed
		If there's no prey within sight radius, predator continues moving around randomly and slowly
	*/

}


// CalcDistance: calculates the distance between two living species
func CalcDistance(org1, org2 Organism) float64 {

      var SqrOfdeltaX = (org1.position.x - org2.position.x) * (org1.position.x - org2.position.x)
      var SqrOfdeltaY  = (org1.position.y - org2.position.y) * (org1.position.y - org2.position.y)

      var distance = math.Sqrt(SqrOfdeltaX +SqrOfdeltaY)

      return distance
}

// ChasePrey: predator accelerates till max speed, moving in direction of closest prey
func ChasePrey(prey Prey, pred Predator)  {


	// prey.accel.x = scalar(prey.OrgProps.position.x - pred.OrgProps.position.x) //--> make this slightly randomized, to allow for mix of successful and failed chases
	// prey.accel.y = scalar(prey.OrgProps.position.y - pred.OrgProps.position.y)// --> make this slightly randomized

// make this slightly randomized, to allow for mix of successful and failed chases

   var animalDistance = CalcDistance(prey.OrgProps,pred.OrgProps)

   if animalDistance > pred.vision {
        return
   }

    // Update predators position and acceleration
    // Position to follow prey position
    // Acceleration to increase

}

// RunFromPredator: prey will run from closest predator within proximal distance
func RunFromPredator(prey, pred Animal) {


  var animalDistance = CalcDistance(prey.OrganismProperties ,pred.OrganismProperties )

  if animalDistance > prey.vision{
       return
  }

  // Update prey position and acceleration
  // Position to move in direction opposite to predator position
  // Acceleration to increase

	// prey.accel.x = scalar(prey.OrgProps.position.x - pred.OrgProps.position.x) //--> make this slightly randomized
	// prey.accel.y = scalar(prey.OrgProps.position.y - pred.OrgProps.position.y) //--> make this slightly randomized

}



// this function might not be needed...
func UpdateEnergy() {

}


// SimulateEcosytem: simulates the ecosystem and its inhabitants over a series of generations (time)
func SimulateEcosystem() {
	InitializeEcosystem(initialEcosystem);

}
