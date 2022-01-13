#include <stdio.h>
#include <sys/ipc.h>
#include <sys/msg.h>
#include <fcntl.h>

#include <stdio.h>
#include <stdlib.h>
#include <sys/mman.h>
#include <sys/stat.h>
#include <unistd.h>
#include <string.h>

#include <time.h>

#include <sys/shm.h>


// structure for message queue
struct mesg_buffer {
    long mesg_type;
    int  duration;
    int id;
} message;

struct task {
    int data[20];
    int dataSize;
} task_struct;

struct idGen {
    int current;
};

void send(struct mesg_buffer message) {
    key_t key;
    int msgid;

    key = ftok("progfile", 65);
    msgid = msgget(key, 0666 | IPC_CREAT);
    message.mesg_type = 1;

    printf("G: saljem %d %d\n", message.id, message.duration);

    msgsnd(msgid, &message, sizeof(message), 0);
}

int getLastId() {
    key_t key = ftok("./dump/lab-5",65);
    int shmid = shmget(key,1024,0666|IPC_CREAT);

    struct idGen *idGen = shmat(shmid,(void*)0,0);

    int out = idGen->current;

    shmdt(idGen);
    shmctl(shmid,IPC_RMID,NULL);

    return out;
}

int storeId(int id) {
    key_t key = ftok("./dump/lab-5", 65);
    int shmid = shmget(key,1024,0666|IPC_CREAT);

    struct idGen *idGen = (struct idGen*) shmat(shmid,(void*)0,0);
    idGen->current = id;

    shmdt(idGen);
}

int main(int argc, char *argv[])
{
    int numberOfTasks = atoi(argv[1]);
    int maxDuration = atoi(argv[2]);
    srand(time(0));

    int last = getLastId();


    for (int i = last; i < numberOfTasks + last; ++i) {
        int duration = rand() % (maxDuration + 1) + maxDuration / 2;
        message.duration = duration;
        message.id = i;

        char keyId[80];
        char taskIdStr[50];
        sprintf(taskIdStr, "%d", i);
        strcpy(keyId, "./dump/lab-5-");
        strcat(keyId, taskIdStr);

        __attribute__((unused)) FILE *fp = fopen (keyId, "w");

        key_t key = ftok(keyId,65);
        int shmid = shmget(key,1024,0666|IPC_CREAT);

        struct task *task = (struct task*) shmat(shmid,(void*)0,0);
        task->dataSize = message.duration;
        for (int j = 0; j < duration; ++j) {
            task->data[j] = (rand() % 400) + 100;
        }

        printf("G: data to do: ");
        for (int j = 0; j < message.duration; ++j) {
            printf("%d ", task->data[j]);
        }
        printf("\n");

        shmdt(task);

        send(message);
    }

    storeId(last + numberOfTasks);

    return 0;
}
