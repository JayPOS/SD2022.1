#include "main.h"

#define MAX_ROUNDS 10

using namespace std;

int randMessageTag() {
    return rand() % 10;
}

int timeToElection() {
    if(randMessageTag() == 1) {
        return LEADER_ELECTION;
    }
    return NORMAL_MESSAGE;
}

int main(int argc, char *argv[]) {

    int rank;
    int leader = 0;
    int rounds = atoi(argv[1]);
    int np;

    int *buf = (int *)malloc(sizeof(int));
    int *recv_buf;
    MPI_Status status;
    MPI_Request request;
    MPI_Request sent;

    int client = 0;
    int maior = -1;


    MPI_Init(&argc, &argv);
    MPI_Comm_rank(comm, &rank);
    MPI_Comm_size(comm , &np);

    int ans = -1;

    int i;
    while (rounds--) {
        // cout << "not done\n";
        if (leader == np-1)
            break;
        if (rank == leader) {
            cout << "where is dinho3\n";
            buf[0] = rank;
            if (client != leader) {
                MPI_Isend( buf , 1 , MPI_INT , client , LEADER_ELECTION , comm , &request);
            }
            if (client < np-1 ) client++;
            int flag = 0;
            MPI_Irecv( recv_buf , 1 , MPI_INT , MPI_ANY_SOURCE , MPI_ANY_TAG , comm , &request);
            while (!flag) {
                MPI_Test( &request , &flag , &status);
            }
            if (flag) {
                if (status.MPI_SOURCE > rank && LEADER_ELECTION_OK) {
                    maior = max(maior, status.MPI_SOURCE);
                    cout << "where is dinho2\n";
                }
            }
            else {
                leader = maior;
                MPI_Isend( buf , 1 , MPI_INT , leader , LEADER_GRANT , comm , &request);
                cout << "where is dinho\n";
            }

        }
        else {
            int flag = 0;
            cout << "where is dinho5\n";
            while (!flag) {
                cout << "where is dinho6\n";
                MPI_Iprobe( MPI_ANY_SOURCE , MPI_ANY_TAG , comm , &flag , &status);
            }
            if (flag) {
                cout << "where is dinho4\n";
                MPI_Recv( recv_buf , 1 , MPI_INT , MPI_ANY_SOURCE , MPI_ANY_TAG , comm , &status);
                if (status.MPI_SOURCE < rank && status.MPI_TAG == LEADER_ELECTION) {
                    buf[0] = rank;
                    MPI_Isend( buf , 1 , MPI_INT , recv_buf[0] , LEADER_ELECTION_OK , comm , &request);
                }
            }
        }
    }

    cout << leader << "\n";

    MPI_Finalize();
    return 0;
}

