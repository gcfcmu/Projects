# gcf_dsp_sop
SOP validation and possible improvements for DSP

Setting up conda environments on the PSC

for the QC notebook:
requires seaborn
scipy
sklearn

conda install -c conda-forge 


# Step 1 - Log into the PSC

ssh -X username@bridges2.psc.edu

and input your password 

interact -p RM 

# Step 2 - create the environment 

conda create -n name -c conda-forge python3.9

# Step 3 - install

conda install -c conda-forge ipykernel
conda install -c conda-forge seaborn

Specifically for the PSC, the packages should be installed onto the $PROJECT directory 

After installing conda, one can move the conda folder to $PROJECT with the following commands:

mv ~/.conda $PROJECT/ 
$PROJECT/.conda ~/.conda

This sets up a pointer to the conda folder, and installation can be done under $PROJECT (which has more space) 

module load anaconda3

conda activate 

conda create -n envname seaborn numpy pandas scikit-learn scipy matplotlib

Replace "envname" with the desired name of the environment. 

# Step 4 - Activate 

run

conda activate envname

interact -p RM

conda deactivate

to start an interactive session


# Step 5 - Running the notebook 

on the PSC server, run 

srun miniconda3/envs/envname/bin/jupyter-lab --no-browser --ip=0.0.0.0

or any other binary file for jupyter lab 

If token authentication is necessary, copy and paste from the terminal window.

# Step 6 - Connect through local SSH port

On a separate terminal window, connect using ssh -L 

for example, ssh -L 8888:r001.ib.bridges2.psc.edu:8888 bridges2.psc.edu -l username

replacing 8888 with the local host port and r001 with the node given in the interactive job
(it should be displayed in the terminal window after running the interact command)
