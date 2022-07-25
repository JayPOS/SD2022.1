all: compile 
	mpirun -np 4 bin 0 10 0.5

compile:
	mpic++ -o bin main.cpp