## Goroutines:

`Goroutines` are lightweight threads managed by the Go runtime. They allow concurrent execution of functions. You can start a new goroutine using the go keyword.

Example:

```go
package main

import (
	"fmt"
	"time"
)

func myFunc() {
	for i := 0; i < 5; i++ {
		fmt.Println("Hello from Goroutine")
		time.Sleep(time.Millisecond * 500)
	}
}

func main() {
	go myFunc()

	for i := 0; i < 3; i++ {
		fmt.Println("Hello from Main")
		time.Sleep(time.Millisecond * 1000)
	}
}

```

## Channels:

`Channels` are a way for goroutines to communicate with each other and synchronize their execution. They can be used to send and receive data between goroutines.

Example:

```go
package main

import (
	"fmt"
	"time"
)

func sendData(ch chan string) {
	ch <- "Hello, Channel!"
}

func main() {
	myChannel := make(chan string)

	go sendData(myChannel)

	// Receiving data from channel
	msg := <-myChannel
	fmt.Println(msg)
}


```

## Select:

The select statement in Go allows you to wait on multiple channel operations. It blocks until one of its cases can proceed, then it executes that case.

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		ch1 <- "one"
	}()

	go func() {
		time.Sleep(1 * time.Second)
		ch2 <- "two"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Println("received", msg1)
		case msg2 := <-ch2:
			fmt.Println("received", msg2)
		}
	}
}

```

## Atomic Operations with sync/atomic:

The sync/atomic package provides low-level atomic memory operations that are atomic with respect to concurrent executions. It's used for operations such as atomic counters, flags, and more.

```go
package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

func main() {
	var counter int64 = 0

	for i := 0; i < 1000; i++ {
		go func() {
			atomic.AddInt64(&counter, 1)
		}()
	}

	time.Sleep(time.Second) // Wait for goroutines to finish

	fmt.Println("Counter:", atomic.LoadInt64(&counter))
}
```

### Studi Kasus 1: Concurrent Data Processing

Deskripsi:

Anda memiliki daftar angka yang ingin dihitung rata-ratanya secara konkuren.
Gunakan goroutines untuk membagi pekerjaan, hitung total, dan rata-ratanya.
Gunakan channel untuk mengumpulkan hasil dari goroutines.
Tampilkan hasil total dan rata-rata di akhir.

```go
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func calculate(data []int, resultCh chan float64) {
	total := 0
	for _, num := range data {
		total += num
	}
	resultCh <- float64(total) / float64(len(data)) // Kirim rata-rata ke channel
}

func main() {
	rand.Seed(time.Now().UnixNano())

	dataSize := 1000
	data := make([]int, dataSize)
	for i := 0; i < dataSize; i++ {
		data[i] = rand.Intn(100) // Data angka acak
	}

	numGoroutines := 5
	resultCh := make(chan float64)

	for i := 0; i < numGoroutines; i++ {
		go calculate(data[i*(dataSize/numGoroutines):(i+1)*(dataSize/numGoroutines)], resultCh)
	}

	var total float64
	for i := 0; i < numGoroutines; i++ {
		total += <-resultCh
	}

	average := total / float64(numGoroutines)
	fmt.Printf("Total: %.2f\nAverage: %.2f\n", total, average)
}


```

### Studi Kasus 2: Concurrent Data Processing (Fan-In & Fan-Out)

Deskripsi:

Anda memiliki banyak pekerjaan yang ingin diproses secara konkuren (misalnya, pemrosesan file teks, komputasi intensif, dll.).
Gunakan fan-out untuk membagi pekerjaan ke beberapa goroutines.
Gunakan fan-in untuk mengumpulkan hasil dari goroutines yang berbeda.
Implementasikan select untuk menunggu dan mengumpulkan hasil dari goroutines secara non-blokir.

```go
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Printf("Worker %d processing job %d\n", id, job)
		time.Sleep(time.Millisecond * time.Duration(rand.Intn(1000))) // Simulasi pekerjaan
		results <- job * 2 // Contoh: hasil operasi pekerjaan
	}
}

func fanOut(numWorkers int, jobs <-chan int, results chan<- int) {
	var wg sync.WaitGroup
	wg.Add(numWorkers)

	for i := 0; i < numWorkers; i++ {
		go func(workerID int) {
			defer wg.Done()
			worker(workerID, jobs, results)
		}(i)
	}

	go func() {
		wg.Wait()
		close(results)
	}()
}

func fanIn(results ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)

	output := func(c <-chan int) {
		defer wg.Done()
		for n := range c {
			out <- n
		}
	}

	wg.Add(len(results))
	for _, r := range results {
		go output(r)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func main() {
	rand.Seed(time.Now().UnixNano())
	numJobs := 10
	numWorkers := 5

	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	for i := 0; i < numJobs; i++ {
		jobs <- i // Generate jobs
	}
	close(jobs)

	fanOut(numWorkers, jobs, results)

	// Fan-in: Mengumpulkan hasil dari goroutines
	finalResults := fanIn(results)

	for res := range finalResults {
		fmt.Println("Result:", res)
	}
}

```

### Studi Kasus 4: Concurreny Patterns & Advanced Tooling

Deskripsi:

Implementasikan pola fan-in dan fan-out untuk memproses data secara paralel.
Gunakan context untuk mengelola pembatalan dan timeout pada sekelompok tugas.
Gunakan sync.Pool untuk mengelola dan memanfaatkan objek yang dibutuhkan secara berulang, seperti objek yang besar atau kompleks secara komputasi.
Manfaatkan sync/atomic untuk operasi penghitungan yang aman secara konkuren.

```go
package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Printf("Worker %d processing job %d\n", id, job)
		time.Sleep(time.Millisecond * time.Duration(rand.Intn(1000))) // Simulasi pekerjaan
		results <- job * 2 // Contoh: hasil operasi pekerjaan
	}
}

func fanOut(ctx context.Context, numWorkers int, jobs <-chan int, results chan<- int) {
	var wg sync.WaitGroup
	wg.Add(numWorkers)

	for i := 0; i < numWorkers; i++ {
		go func(workerID int) {
			defer wg.Done()
			worker(workerID, jobs, results)
		}(i)
	}

	go func() {
		wg.Wait()
		close(results)
	}()
}

func fanIn(results ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)

	output := func(c <-chan int) {
		defer wg.Done()
		for n := range c {
			out <- n
		}
	}

	wg.Add(len(results))
	for _, r := range results {
		go output(r)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func main() {
	rand.Seed(time.Now().UnixNano())
	numJobs := 10
	numWorkers := 5

	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*500)
	defer cancel()

	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	for i := 0; i < numJobs; i++ {
		jobs <- i // Generate jobs
	}
	close(jobs)

	fanOut(ctx, numWorkers, jobs, results)

	finalResults := fanIn(results)

	var totalCount int64 // Penggunaan atomic untuk operasi aman secara konkuren
	for res := range finalResults {
		atomic.AddInt64(&totalCount, int64(res))
	}

	fmt.Println("Total Result:", totalCount)
}

```

Studi Kasus 5: Konkurensi dengan Sinkronisasi Data dan Keadaan
Deskripsi:

Buat aplikasi sederhana yang mengelola akun pengguna secara konkuren (menambah saldo, menarik saldo, dan melihat saldo).
Gunakan goroutines untuk melakukan operasi pada akun pengguna dan sinkronisasi akses ke data akun menggunakan sync.Mutex.
Implementasikan manajemen keadaan akun untuk memastikan operasi penambahan dan penarikan saldo berjalan dengan benar.

```go
package main

import (
	"fmt"
	"sync"
)

type Account struct {
	mutex sync.Mutex
	balance int
}

func (acc *Account) Deposit(amount int) {
	acc.mutex.Lock()
	defer acc.mutex.Unlock()
	acc.balance += amount
}

func (acc *Account) Withdraw(amount int) bool {
	acc.mutex.Lock()
	defer acc.mutex.Unlock()
	if acc.balance >= amount {
		acc.balance -= amount
		return true
	}
	return false
}

func (acc *Account) Balance() int {
	acc.mutex.Lock()
	defer acc.mutex.Unlock()
	return acc.balance
}

func main() {
	userAccount := &Account{balance: 100}

	var wg sync.WaitGroup

	// Goroutines to simulate multiple transactions
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			userAccount.Deposit(10) // Menambah saldo
			fmt.Printf("Deposit. New Balance: %d\n", userAccount.Balance())

			if success := userAccount.Withdraw(20); success { // Menarik saldo
				fmt.Printf("Withdrawal Successful. New Balance: %d\n", userAccount.Balance())
			} else {
				fmt.Println("Withdrawal Failed. Insufficient Balance.")
			}
		}()
	}

	wg.Wait()
	fmt.Println("Final Balance:", userAccount.Balance())
}


```

### Studi Kasus 6: Sinkronisasi Data dalam Struktur yang Kompleks

Deskripsi:

Buat aplikasi untuk mengelola buku pesanan pada toko online secara konkuren.
Gunakan struktur data kompleks yang melibatkan buku pesanan dengan informasi barang, jumlah, dan status pengiriman.
Manfaatkan sync.Mutex untuk mengamankan akses konkuren terhadap buku pesanan.
Implementasikan manajemen status pesanan untuk melacak dan memperbarui status pengiriman.

```go
package main

import (
	"fmt"
	"sync"
)

type Order struct {
	ID       int
	Item     string
	Quantity int
	Status   string
}

type OrderManager struct {
	orders []*Order
	mutex  sync.Mutex
}

func (om *OrderManager) AddOrder(order *Order) {
	om.mutex.Lock()
	defer om.mutex.Unlock()
	om.orders = append(om.orders, order)
}

func (om *OrderManager) UpdateStatus(orderID int, status string) {
	om.mutex.Lock()
	defer om.mutex.Unlock()
	for _, order := range om.orders {
		if order.ID == orderID {
			order.Status = status
			return
		}
	}
}

func (om *OrderManager) ListOrders() {
	om.mutex.Lock()
	defer om.mutex.Unlock()
	for _, order := range om.orders {
		fmt.Printf("OrderID: %d, Item: %s, Quantity: %d, Status: %s\n", order.ID, order.Item, order.Quantity, order.Status)
	}
}

func main() {
	orderManager := &OrderManager{}

	for i := 1; i <= 5; i++ {
		order := &Order{
			ID:       i,
			Item:     fmt.Sprintf("Item %d", i),
			Quantity: i * 2,
			Status:   "Pending",
		}
		orderManager.AddOrder(order)
	}

	go func() {
		orderManager.UpdateStatus(3, "Shipped") // Menandai pengiriman order dengan ID 3
	}()

	orderManager.ListOrders()
}


```

Studi Kasus 7: Konkurensi dengan Pipelines dan Kontrol Melalui Context
Deskripsi:

Buat sistem pipeline konkurensi sederhana yang mengolah data sensor secara bersamaan.
Terapkan konkurensi menggunakan goroutines untuk membaca, memproses, dan menyimpan data sensor.
Gunakan context untuk mengontrol pembatalan atau timeout pada operasi konkuren.

```go
package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

type SensorData struct {
	ID        int
	Timestamp time.Time
	Value     float64
}

func readSensor(ctx context.Context, sensorID int, dataCh chan<- SensorData) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Sensor %d: Shutting down\n", sensorID)
			return
		default:
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(500))) // Simulasi pembacaan sensor
			dataCh <- SensorData{
				ID:        sensorID,
				Timestamp: time.Now(),
				Value:     rand.Float64() * 100, // Data sensor acak
			}
		}
	}
}

func processSensorData(ctx context.Context, dataCh <-chan SensorData) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Processing: Shutting down")
			return
		case data := <-dataCh:
			// Proses data sensor (contoh: cetak data)
			fmt.Printf("Received Data: SensorID %d, Value %.2f, Timestamp %s\n", data.ID, data.Value, data.Timestamp.Format("15:04:05"))
		}
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	numSensors := 3
	dataCh := make(chan SensorData)

	// Goroutines untuk membaca data sensor
	for i := 0; i < numSensors; i++ {
		go readSensor(ctx, i+1, dataCh)
	}

	// Goroutine untuk memproses data sensor
	go processSensorData(ctx, dataCh)

	// Tunggu hingga waktu timeout atau pembatalan konteks
	<-ctx.Done()
}

```

### Studi Kasus 8: Konkurensi dengan Buffered Channel dan Goroutine Pool

Deskripsi:

Buat sistem pengelolaan pesan konkuren dengan pool goroutines.
Gunakan buffered channel untuk mengelola antrian pesan.
Implementasikan goroutine pool untuk pengelolaan pengiriman pesan secara konkuren.

```go
package main

import (
	"fmt"
	"sync"
)

type Message struct {
	ID   int
	Body string
}

func messageWorker(id int, messageCh <-chan Message, wg *sync.WaitGroup) {
	defer wg.Done()
	for msg := range messageCh {
		fmt.Printf("Worker %d: Processing message %d - %s\n", id, msg.ID, msg.Body)
	}
}

func main() {
	numWorkers := 3
	numMessages := 10

	messageCh := make(chan Message, numMessages)

	var wg sync.WaitGroup
	wg.Add(numWorkers)

	// Goroutine Pool
	for i := 0; i < numWorkers; i++ {
		go messageWorker(i+1, messageCh, &wg)
	}

	// Mengirim pesan ke buffered channel
	for i := 1; i <= numMessages; i++ {
		messageCh <- Message{ID: i, Body: fmt.Sprintf("Message %d", i)}
	}
	close(messageCh)

	wg.Wait()
	fmt.Println("All messages processed.")
}

```

Studi Kasus 9: Konkurensi dengan sync.Cond untuk Sinyal dan Sinkronisasi
Deskripsi:

Implementasikan struktur data dengan sinkronisasi penggunaan sync.Cond.
Gunakan sync.Cond untuk sinkronisasi dan sinyal antar goroutines untuk mengatur proses pembacaan dan penulisan data.

```go
package main

import (
	"fmt"
	"sync"
	"time"
)

type DataStore struct {
	sync.Mutex
	cond  *sync.Cond
	valid bool
	data  int
}

func (ds *DataStore) writeData() {
	ds.Lock()
	defer ds.Unlock()

	// Menunggu data menjadi valid
	for !ds.valid {
		ds.cond.Wait()
	}

	// Menulis data
	ds.data = 10
	fmt.Println("Data ditulis.")
}

func (ds *DataStore) readData() {
	ds.Lock()
	defer ds.Unlock()

	// Mengatur data menjadi valid
	ds.valid = true
	fmt.Println("Data menjadi valid.")
	// Memberitahu goroutine lain bahwa data sudah valid
	ds.cond.Broadcast()
}

func main() {
	dataStore := &DataStore{
		cond: sync.NewCond(&sync.Mutex{}),
	}

	// Goroutine untuk menulis data
	go dataStore.writeData()

	// Goroutine untuk membaca data
	go func() {
		time.Sleep(time.Second)
		dataStore.readData()
	}()

	time.Sleep(3 * time.Second)
}

```

### Studi Kasus 10: Konkurensi dengan sync.WaitGroup dan Komunikasi Channel

Deskripsi:

Buat aplikasi konkuren yang menghitung jumlah kemunculan kata dalam beberapa teks.
Gunakan sync.WaitGroup untuk menunggu selesainya goroutines sebelum menampilkan hasil.
Gunakan channel untuk berbagi dan mengumpulkan hasil dari goroutines yang berjalan secara konkuren.

```go
package main

import (
	"fmt"
	"strings"
	"sync"
)

func wordCounter(text string, word string, wg *sync.WaitGroup, resultCh chan<- int) {
	defer wg.Done()
	words := strings.Fields(text)
	count := 0

	for _, w := range words {
		if strings.ToLower(w) == strings.ToLower(word) {
			count++
		}
	}

	resultCh <- count
}

func main() {
	texts := []string{
		"Lorem ipsum dolor sit amet, consectetur adipiscing elit.",
		"Sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.",
		"Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.",
	}

	wordToFind := "ipsum"
	var wg sync.WaitGroup
	resultCh := make(chan int, len(texts))

	for _, text := range texts {
		wg.Add(1)
		go wordCounter(text, wordToFind, &wg, resultCh)
	}

	go func() {
		wg.Wait()
		close(resultCh)
	}()

	totalCount := 0
	for count := range resultCh {
		totalCount += count
	}

	fmt.Printf("Total occurrences of '%s' across texts: %d\n", wordToFind, totalCount)
}

```
