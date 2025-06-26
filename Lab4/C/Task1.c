#include <stdio.h>
#include <stdlib.h>
#include <pthread.h>
#include <unistd.h>

#define MAX_SIZE 10

typedef struct {
    int items[MAX_SIZE];
    int front, rear, size;
    pthread_mutex_t lock;
    pthread_cond_t not_empty;
} ThreadSafeQueue;

void init_queue(ThreadSafeQueue *q) {
    q->front = q->rear = q->size = 0;
    pthread_mutex_init(&q->lock, NULL);
    pthread_cond_init(&q->not_empty, NULL);
}

int is_empty(ThreadSafeQueue *q) {
    return q->size == 0;
}

int is_full(ThreadSafeQueue *q) {
    return q->size == MAX_SIZE;
}

void enqueue(ThreadSafeQueue *q, int item) {
    pthread_mutex_lock(&q->lock);
    if (is_full(q)) {
        printf("Queue is full, dropping item\n");
        pthread_mutex_unlock(&q->lock);
        return;
    }
    q->items[q->rear] = item;
    q->rear = (q->rear + 1) % MAX_SIZE;
    q->size++;
    printf("Enqueued: %d\n", item);
    pthread_cond_signal(&q->not_empty);
    pthread_mutex_unlock(&q->lock);
}

int dequeue(ThreadSafeQueue *q) {
    pthread_mutex_lock(&q->lock);
    while (is_empty(q)) {
        pthread_cond_wait(&q->not_empty, &q->lock);
    }
    int item = q->items[q->front];
    q->front = (q->front + 1) % MAX_SIZE;
    q->size--;
    printf("Dequeued: %d\n", item);
    pthread_mutex_unlock(&q->lock);
    return item;
}

void *producer(void *arg) {
    ThreadSafeQueue *q = (ThreadSafeQueue *)arg;
    for (int i = 0; i < 20; i++) {
        enqueue(q, i);
        usleep(100000); // Simula trabajo
    }
    return NULL;
}

void *consumer(void *arg) {
    ThreadSafeQueue *q = (ThreadSafeQueue *)arg;
    for (int i = 0; i < 10; i++) {
        dequeue(q);
        usleep(150000);
    }
    return NULL;
}

int main() {
    ThreadSafeQueue queue;
    init_queue(&queue);

    pthread_t prod1, cons1, cons2;

    pthread_create(&prod1, NULL, producer, &queue);
    pthread_create(&cons1, NULL, consumer, &queue);
    pthread_create(&cons2, NULL, consumer, &queue);

    pthread_join(prod1, NULL);
    pthread_join(cons1, NULL);
    pthread_join(cons2, NULL);

    return 0;
}
