
WIP: `docker` pseudo command to allow bamboo-agents run CI tasks inside Kubernetes PODs
===

The main idea is to emulate real `docker` CLI behaviour of run / exec and cp commands 
by executing them in Kubernetes PODs. 