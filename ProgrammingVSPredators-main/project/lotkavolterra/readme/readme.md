## ReadMe

Hello! This is our population dynamics simulator that is within a web application. This application will allow users to set up their own ecosystem with their own specified species, parameters and interaction values between the species.

#### How to start the web application

Unzip all the files in the project folder directly into /Users/user/go/src. The github folder should be placed in /Users/user/go/src. The rest of the code can be left as is, and the working directory is /Users/user/go/src/project/lotka.

After bringing the files of this project (including this file) into your go/src directory, you are now ready to start our population dynamics web application.

* First, use a terminal to navigate to this project directory (lotkavolterra)

* Enter `go build` into the terminal to compile all the code

* Enter `go run .` to run all the files in the lotkavolterra folder

* Then, open up a web browser and type in "localhost:3000" to run the web application!

### Using the Web Application

Here is a link to coding demonstration. [demo link](https://youtu.be/aM8zO8EbFk0)

#### Step 1: Choose simulation type
You can either chose to simulate a system in which each population will follow the model directly, or introduce a randomization parameter (either normally or uniformly distributed) that can model systems such as those with immigration and emigration.

If you choose to select either stochastic method, an input box will appear asking for the magnitude of that change. We recommend values around 1%. If nothing is input, the regular model will be used instead of the stochastic one.

#### Step 2: Input the number of generations and timestep parameters
For the number of generations, make sure to enter at least 2 generations so that the model runs. A good number of generations will depend on your initial conditions as well as the time step.

For time step, you should enter a value less than 1 (but more than 0) for best results (e.g. around 0.0001). A smaller number here will mean a more precise simulation. However, keep in mind that each generation will update populations one time step, so if you lower timestep your model will appear to run for less time.

Some good parameters for number of generations and time step are 8000 and 0.0001, for example.

#### Step 3: Input species information
Next, we want to enter the information for each species. Two important things to note:

Each species should have their own name. If two species are given the same name, the model will still run, but will overwrite previously named species.

Predators should have a negative growth value. Prey should have positive growth. This will ensure that the predator-prey system functions given the version of the model implemented here.


#### Step 4: Insert values for the interaction matrix
It can be difficult to know what values will work for the matrix to produce a stable model. However, a good rule of thumb is to put negative values along the diagonals (self-interacting values), negative values for predators actin on prey, and positive values for prey acting on predators.

The interaction values are given by the order of input from the species parameters. That is, if the first row is "Bear" and the second is "Rabbit", then the matrix will look as follows:

Clicking on the "Add more species" button will create a new row in the species parameters and will change the size of the matrix accordingly. It is best to press this button before inputting parameters into the matrix.

#### Step 5: Submitting
Once all the parameters are set, the Submit button will create a PNG file and display the graph. Click on the browser's back arrows to input new parameters or change existing ones.
