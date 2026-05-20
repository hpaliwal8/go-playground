package exercises

// import "fmt"

// type Transaction struct {
// 	db          map[string]int
// 	localWrites map[string]int
// 	active      bool
// }

// func NewTransaction(db map[string]int) *Transaction {
// 	return &Transaction{
// 		db:          db,
// 		localWrites: make(map[string]int),
// 		active:      true,
// 	}
// }

// func (t *Transaction) Read(key string) (int, error) {
// 	if !t.active {
// 		return 0, fmt.Errorf("transaction is not active")
// 	}

// 	if val, ok := t.localWrites[key]; ok {
// 		return val, nil
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

// 	if _, ok := t.db[key]; !ok {
// 		if _, ok := t.localWrites[key]; !ok {
// 			return fmt.Errorf("value does not already exist in system")
// 		}
// 	}

// 	t.localWrites[key] = value
// 	return nil
// }

// func (t *Transaction) Commit() error {
// 	if !t.active {
// 		return fmt.Errorf("transaction is not active")
// 	}

// 	// TODO:
// 	// 1. apply all localWrites into db
// 	// 2. clear localWrites
// 	// 3. mark transaction inactive

// 	for key, val := range t.localWrites {
// 		t.db[key] = val
// 	}

// 	clear(t.localWrites)

// 	t.active = false

// 	return nil
// }

// func (t *Transaction) Rollback() error {
// 	if !t.active {
// 		return fmt.Errorf("transaction is not active")
// 	}

// 	// TODO:
// 	// 1. discard tentative writes
// 	// 2. mark transaction inactive

// 	clear(t.localWrites)

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

// 	b, err := tx.Read("B")
// 	if err != nil {
// 		panic(err)
// 	}

// 	if err := tx.Write("B", b+100); err != nil {
// 		panic(err)
// 	}

// 	if err := tx.Commit(); err != nil {
// 		panic(err)
// 	}

//		fmt.Println("Final DB:", db)
//	};
