package exercises

// import "fmt"

// type Transaction struct {
// 	db           map[string]int
// 	local_writes map[string]int
// 	active       bool
// }

// func NewTransaction(db map[string]int) *Transaction {
// 	return &Transaction{
// 		db:           db,
// 		local_writes: make(map[string]int),
// 		active:       true,
// 	}
// }

// func (t *Transaction) Read(key string) (int, error) {
// 	if !t.active {
// 		return 0, fmt.Errorf("transaction is not active")
// 	}

// 	val, ok := t.db[key]
// 	if !ok {
// 		return 0, fmt.Errorf("key %s not found", key)
// 	}

// 	return val, nil
// }

// func (t *Transaction) Write(key string, value int) error {
// 	if !t.active {
// 		return fmt.Errorf("transaction is not active")
// 	}

// 	t.local_writes[key] = value
// 	return nil
// }

// func (t *Transaction) Commit() error {
// 	if !t.active {
// 		return fmt.Errorf("transaction is not active")
// 	}

// 	t.active = false
// 	return nil
// }

// func (t *Transaction) Rollback() error {
// 	if !t.active {
// 		return fmt.Errorf("transaction is not active")
// 	}

// 	// BUG: nothing to undo
// 	t.active = false
// 	return nil
// }

// func main() {
// 	db := map[string]int{
// 		"A": 500,
// 		"B": 300,
// 	}

// 	tx := NewTransaction(db)

// 	a, err := tx.Read("A")
// 	if err != nil {
// 		panic(err)
// 	}

// 	if err := tx.Write("A", a-100); err != nil {
// 		panic(err)
// 	}

// 	// Simulate failure before updating B
// 	if err := tx.Rollback(); err != nil {
// 		panic(err)
// 	}

// 	fmt.Println("Final DB:", db)
// }
