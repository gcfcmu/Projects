#!/bin/bash
PORT=$1
PORT="${PORT:-8888}"
conda activate base
jupyter notebook --no-browser --port=$PORT --ip=0.0.0.0
