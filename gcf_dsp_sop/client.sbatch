#!/bin/bash
SSH_PORT=$1
HTC_NODE=$2

#Kill any existing connections on this port:
kill $(lsof -t -i:$SSH_PORT)

#Attach this port to CRC using same port (note: could also be different ports)
ssh -L 8888:r001.ib.bridges2.psc.edu:8888 bridges2.psc.edu -l gfonseca
