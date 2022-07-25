#include <stdio.h>
#include "main.h"

int MAX_ROUNDS = 6;
double TX_PROB = 1.0 - ERROR_PROB;

unsigned long int get_PRNG_seed() {
	struct timeval tv;
	gettimeofday(&tv,NULL);
	unsigned long time_in_micros = 1000000 * tv.tv_sec + tv.tv_usec + getpid();

	return time_in_micros;
}

bool try_leader_elect() {

	double prob = rand() / (double) RAND_MAX;
	bool leader_elect = (prob > THRESHOLD);

	return leader_elect;
}


int main(int argc, char *argv[]) {

	int myrank, np;
	int leader = 0;
	int *recv_buf = (int *) malloc(sizeof(int));
	int dinho = DINHO;

	// user input argv[1]: designated initial leader
	leader = atoi(argv[1]);
	// user input argv[2]: how many rounds to run the algorithm
	MAX_ROUNDS = atoi(argv[2]);
	// user input argv[3]: packet trasnmission success/failure probability
	TX_PROB = atoi(argv[3]);


    printf("\n%s%s", BAR, BAR);
    printf("\n%s%s", BAR, BAR);
	printf("\n Initialize dinhoverse:: \n\tMAX_ROUNDS = %d \n\tinitial leader = %d \n\tTX_PROB = %f\n", MAX_ROUNDS, leader, TX_PROB);
	printf("\n%s%s", BAR, BAR);
	printf("\n%s%s\n\n", BAR, BAR);

  // MPI Initiliazation
	MPI_Init(&argc, &argv);
	MPI_Comm_rank(comm, &myrank);
	MPI_Comm_size(comm, &np);

    MPI_Status status;

	srand(get_PRNG_seed());

    /*arranging nodes in ring since bully is hard as hell!*/

	int succ, pred;

	if (myrank == 0) {
		pred = np - 1;
	}
	else {
		pred = myrank - 1;
	}

	if (myrank == (np - 1)) {
		succ = 0;
	}
	else {
			succ = myrank + 1;
	}

	// Main loop, each iteration is a round
	for (int round = 0; round < MAX_ROUNDS; round++) {
		printf("\n%s ROUND %d %s\n", BAR, round, BAR);

        if (leader == myrank) {
            // if random function activates, start election, else just ping
            if (try_leader_elect()) {
                recv_buf[0] = myrank;

                MPI_Send( recv_buf , 1 , MPI_INT , succ , LEADER_ELECTION , comm);
                printf("\ndinho[ rank %d][ round %d]: SENT LEADER ELECTION MESSAGE TO NODE %d, TAG %d\n", myrank, round, succ, LEADER_ELECTION);
                fflush(stdout);
            }
            else {
                MPI_Send( &dinho , 1 , MPI_INT , succ , DINHO_MSG , comm);
                printf("\ndinho[ rank %d][ round %d]: SENT LEADER NORMAL MESSAGE TO NODE %d, TAG %d\n", myrank, round, succ, DINHO_MSG);
                fflush(stdout);
            }
            MPI_Request request;
            int recv_buf[1];
            int flag = 0;
            MPI_Irecv( recv_buf , 1 , MPI_INT , pred , MPI_ANY_TAG , comm , &request);
            while (!flag) {
                MPI_Test( &request , &flag , &status);
            }
            if (flag) {
                switch (status.MPI_TAG)
                {
                    case DINHO_MSG:     
                        printf("\ndinho[ rank %d][ round %d]: RECEIVED LAST DINHO MSG, RING COMPLETE!\n", myrank, round);
                        fflush(stdout);
                        break;
                
                    case LEADER_ELECTION:
                        leader = recv_buf[0];
                        MPI_Send( recv_buf , 1 , MPI_INT , succ , LEADER_ELECTION_OK , comm);
                        MPI_Recv( recv_buf , 1 , MPI_INT , pred , LEADER_ELECTION_OK , comm , MPI_STATUS_IGNORE);
                        printf("\ndinho[ rank %d][ round %d]: new dinhoverse leader: dinho[ rank %d ]\n", myrank, round, leader);
                        fflush(stdout);
                        break;
                
                    default:
                        break;
                }
            }
        }
        else {
            int flag = 0;
            while (!flag) {
                MPI_Iprobe( MPI_ANY_SOURCE , MPI_ANY_TAG , comm , &flag , &status);
            }
            if (flag) {
                MPI_Recv( recv_buf , 1 , MPI_INT , pred , MPI_ANY_TAG , comm , &status);

                switch (status.MPI_TAG)
                {
                    case DINHO_MSG:
                        MPI_Send( &dinho , 1 , MPI_INT , succ , DINHO_MSG , comm);
                        printf("\ndinho[ rank %d][ round %d]: REPASSING DINHO MSG! ENJOY!\n", myrank, round);
                        fflush(stdout);
                        break;

                    case LEADER_ELECTION:
                        // check if it will function for this election
                        if (try_leader_elect()) {
                            if(myrank > recv_buf[0]) {
                                recv_buf[0] = myrank;
                            }
                            printf("\ndinho[ rank %d][ round %d]: I WANT TO BE LEADER\n", myrank, round);
                        }
                        else {
                            printf("\ndinho[ rank %d][ round %d]: DINHO IS SLEEPING! GO MAKE THIS ELECTION YOURSELVES\n", myrank, round);
                        }
                        // send answer to next dinho
                        MPI_Send( recv_buf , 1 , MPI_INT , succ , LEADER_ELECTION , comm);

                        // waiting for dinhoverse`s election results
                        MPI_Recv( recv_buf , 1 , MPI_INT , MPI_ANY_SOURCE , LEADER_ELECTION_OK , comm , &status);

                        leader = recv_buf[0];
                        printf("\ndinho[ rank %d][ round %d]: NEW DINHO LEADER: %d\n", myrank, round, leader);
                        fflush(stdout);

                        MPI_Send( recv_buf , 1 , MPI_INT , succ , LEADER_ELECTION_OK , comm);
                        break;
                
                    default:
                        break;
                }
            }
        }
        MPI_Barrier( comm );
	}
	printf("\n DINHO Leader for NODE %d = %d\n", myrank, leader);

	MPI_Finalize();
	return 0;
}