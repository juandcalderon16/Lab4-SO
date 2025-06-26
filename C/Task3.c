#include <stdio.h>
#include <stdlib.h>
#include <pthread.h>
#include <semaphore.h>
#include <unistd.h>

#define N 5
pthread_mutex_t forks[N];
sem_t room;

void think(int id) {
    printf("Philosopher %d is thinking\n", id);
    usleep(rand() % 500000);
}

void eat(int id) {
    printf("Philosopher %d is eating\n", id);
    usleep(rand() % 500000);
}

void *philosopher(void *arg) {
    int id = *(int *)arg;
    while (1) {
        think(id);
        sem_wait(&room); // Máximo 4 filósofos al mismo tiempo
        pthread_mutex_lock(&forks[id]);
        pthread_mutex_lock(&forks[(id + 1) % N]);
        eat(id);
        pthread_mutex_unlock(&forks[id]);
        pthread_mutex_unlock(&forks[(id + 1) % N]);
        sem_post(&room);
    }
    return NULL;
}

int main() {
    pthread_t threads[N];
    int ids[N];

    sem_init(&room, 0, N - 1);
    for (int i = 0; i < N; i++) {
        pthread_mutex_init(&forks[i], NULL);
        ids[i] = i;
        pthread_create(&threads[i], NULL, philosopher, &ids[i]);
    }

    for (int i = 0; i < N; i++) {
        pthread_join(threads[i], NULL); // Nunca termina
    }

    return 0;
}
