#ifndef MAIN_H_
#define MAIN_H_


#define ERROR_PROB 0.05				
#define THRESHOLD 0.5					
#define MAX_TOKEN_VALUE 10000			
#define TIME_OUT_INTERVAL 3

#define DINHO 0001
#define DINHO_MSG 0002
#define LEADER_ELECTION 0003
#define LEADER_ELECTION_OK 0004
#define I_HAVE_THE_POWER 0005

#define BAR "------------------------------"

#include <bits/stdc++.h>
#include <mpi.h>
#include <sys/time.h>

MPI_Comm comm = MPI_COMM_WORLD;

inline double get_prob() {
    double prob = rand() / (double) RAND_MAX;
    return prob;
}

#endif