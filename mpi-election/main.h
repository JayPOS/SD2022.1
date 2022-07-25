#ifndef MAIN_H_
#define MAIN_H_

#define NORMAL_MESSAGE 0001
#define LEADER_ELECTION 0002
#define LEADER_ELECTION_OK 0003
#define LEADER_GRANT 0004
#define I_HAVE_THE_POWER 0005

#include <bits/stdc++.h>
#include <mpi.h>

MPI_Comm comm = MPI_COMM_WORLD;

int randMessageTag();
int timeToElection();

#endif