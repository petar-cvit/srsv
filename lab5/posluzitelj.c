#include <stdio.h>
#include <sys/ipc.h>
#include <sys/msg.h>
#include <unistd.h>

#include <stdlib.h>
#include <sys/mman.h>

#include <time.h>

#include <pthread.h>
#include <stdatomic.h>

#include <string.h>

#include <sys/ipc.h>
#include <sys/shm.h>
#include <stdio.h>

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
} id_gen;

typedef enum {false, true} bool;

int taskId = 0;
int sumTime = 0;
int numberOfLiveTasks = 0;
struct mesg_buffer tasksInQueue[10];
int minTimeToFinish = 30; // default value

pthread_mutex_t lock;

pthread_mutex_t lock_process_flag;

bool process = false;

struct mesg_buffer popTask() {
    sumTime -= tasksInQueue[0].duration;

    struct mesg_buffer out = tasksInQueue[0];

    for (int i = 0; i < numberOfLiveTasks; ++i) {
        tasksInQueue[i] = tasksInQueue[i + 1];
    }

    if (numberOfLiveTasks > 0) {
        numberOfLiveTasks--;
    }

    return out;
}

void pushTask(struct mesg_buffer message) {
    printf("P: Zaprimio zadatak %d\n", message.id);
    sumTime += message.duration;
    tasksInQueue[numberOfLiveTasks] = message;
    numberOfLiveTasks++;
}

_Noreturn void* workerThread(void *arg) {
    while(true) {
        pthread_mutex_lock(&lock_process_flag);
        if (!process) {
            pthread_mutex_unlock(&lock_process_flag);
            continue;
        }
        pthread_mutex_unlock(&lock_process_flag);

        sleep(2);
        if (numberOfLiveTasks <= 0) {
            continue;
        }

        pthread_mutex_lock(&lock);
        struct mesg_buffer message = popTask();
        pthread_mutex_unlock(&lock);

        if (message.duration == 0) continue;

        char keyId[80];
        char taskIdStr[50];
        sprintf(taskIdStr, "%d", message.id);

        strcpy(keyId, "dump/lab-5-");
        strcat(keyId, taskIdStr);
        key_t key = ftok(keyId,65);

        int shmid = shmget(key,1024,0666|IPC_CREAT);

        struct task *task = shmat(shmid,(void*)0,0);

        printf("P: task %d\n", task->dataSize);

        printf("P: Doing work\n");
        for (int i = 0; i < message.duration; ++i) {
            printf("P: obrada zadatka %d, obrada podatka %d\n", message.id, task->data[i]);
            sleep(1);
        }

        shmdt(task);
        shmctl(shmid,IPC_RMID,NULL);

        if (message.id == 0) continue;
        printf("P: obradio zadatak %d\n", message.id);
    }
}

_Noreturn void* signalTime(void *arg) {
    (void)arg;

    while(true) {
        sleep(30);

        pthread_mutex_lock(&lock_process_flag);
        if (!process) {
            printf("P: budim radne dretve\n");
        }
        process = true;
        pthread_mutex_unlock(&lock_process_flag);
    }
}

void initTaskIDGen() {
    char keyId[80];
    strcpy(keyId, "dump/lab-5");

    key_t key = ftok(keyId, 65);
    int shmid = shmget(key,1024,0666|IPC_CREAT);

    int *idGen = shmat(shmid,(void*)0,0);
    *idGen = 1;

    printf("P: initialized id generator %d\n", *idGen);

    shmdt(idGen);
}

int main(int _, char *argv[]) {
    int minTimeToFinish = atoi(argv[1]);
    int numberOfTreads = atoi(argv[2]);

    initTaskIDGen();

    pthread_mutex_init(&lock, NULL);
    pthread_mutex_init(&lock_process_flag, NULL);
    pthread_t signalTimeThread;
    pthread_t threads[numberOfTreads];

    for (int i = 0; i < numberOfTreads; ++i) {
        pthread_create(&threads[i], NULL, workerThread, NULL);
        printf("P: started thread\n");
    }
    pthread_create(&signalTimeThread, NULL, signalTime, NULL);

    int msgid;

    key_t key = ftok("progfile", 65);

    while (true) {
        msgid = msgget(key, 0666 | IPC_CREAT);
        msgrcv(msgid, &message, sizeof(message), 1, 0);

        pthread_mutex_lock(&lock);
        pushTask(message);
        pthread_mutex_unlock(&lock);

        if (minTimeToFinish < sumTime) {
            process = true;
        }
    }
}
