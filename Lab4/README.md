# Práctica 4 de Sistemas Operativos – Mecanismos de Sincronización

**Universidad:** Universidad de Antioquia  
**Curso:** Sistemas Operativos  
**Estudiantes:** Juan Diego Calderón Bermeo y Ana Maria Vega Angarita  
**Lenguajes:** C++ y Go

## 🧠 Objetivos

Implementar y comparar mecanismos de sincronización:
- Mutexes (exclusión mutua)
- Variables de condición
- Semáforos

Cada tarea fue implementada en C++ y también en Go.

---

## ✅ Tareas Implementadas

### 1. Cola Segura para Hilos

- Una cola protegida por exclusión mutua y variables de condición.
- Los consumidores esperan cuando la cola está vacía.

**C++:** `queue.cpp` usando `std::mutex` y `std::condition_variable`  
**Go:** `queue.go` usando `sync.Mutex` y `sync.Cond`

**Probado con:** 1 productor y 2 consumidores.

---

### 2. Productor–Consumidor (Búfer Acotado)

- Problema clásico con un búfer de tamaño fijo.
- Semáforos controlan los espacios disponibles y ocupados.
- Un mutex asegura el acceso seguro al búfer.

**C++:** `producer_consumer.cpp` usando `std::counting_semaphore` y `std::mutex`  
**Go:** `producer_consumer.go` usando canales con búfer y `sync.Mutex`

**Probado con:** 1 productor y 2 consumidores.

---

### 3. Filósofos Comensales

- Cinco filósofos alternan entre pensar y comer.
- Los tenedores están protegidos por mutexes.
- Un semáforo (o canal con búfer) limita los comensales simultáneos para evitar interbloqueos.

**C++:** `dining_philosophers.cpp` usando `std::mutex` y `std::counting_semaphore`  
**Go:** `philosophers.go` usando `sync.Mutex` y canal con búfer

**Probado con:** 5 hilos de filósofos; 3 ciclos en Go, bucle infinito en C++.

---

## 🖥️ Cómo Compilar y Ejecutar

### C++
```bash
g++ -std=c++20 -pthread queue.cpp -o queue.x
g++ -std=c++20 -pthread producer_consumer.cpp -o producer_consumer.x
g++ -std=c++20 -pthread dining_philosophers.cpp -o dining_philosophers.x

./queue.x
./producer_consumer.x
./dining_philosophers.x
