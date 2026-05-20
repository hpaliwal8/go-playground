package exercises

// import (
// 	"fmt"
// 	"sync"
// 	"time"
// )

// type Bank struct {
// 	mu sync.Mutex
// 	db map[string]int
// }

// func NewBank() *Bank {
// 	return &Bank{
// 		db: map[string]int{
// 			"A": 100,
// 		},
// 	}
// }

// func (b *Bank) Read(key string) int {
// 	return b.db[key]
// }

// func (b *Bank) Write(key string, value int) {
// 	b.db[key] = value
// }

// func withdrawTen(b *Bank, wg *sync.WaitGroup, name string) {
// 	defer wg.Done()

// 	b.mu.Lock()
// 	defer b.mu.Unlock()

// 	balance := b.Read("A")
// 	fmt.Printf("%s read A = %d\n", name, balance)

// 	time.Sleep(100 * time.Millisecond)

// 	newBalance := balance - 10
// 	b.Write("A", newBalance)
// 	fmt.Printf("%s wrote A = %d\n", name, newBalance)
// }

// func main() {
// 	bank := NewBank()

// 	var wg sync.WaitGroup
// 	wg.Add(2)

// 	go withdrawTen(bank, &wg, "T1")
// 	go withdrawTen(bank, &wg, "T2")

// 	wg.Wait()

// 	fmt.Println("Final A =", bank.Read("A"))
// }
