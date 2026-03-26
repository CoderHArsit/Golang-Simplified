package main

import "fmt"

// ============================================================
// METHOD RECEIVERS: The Complete Guide
// ============================================================
// A "method" is just a function with a RECEIVER.
// The receiver tells Go: "this function belongs to this type."
//
// Syntax:
//   func (receiverVar ReceiverType) MethodName() ReturnType { ... }
//
// Two kinds:
//   (u User)   → Value Receiver  (gets a copy)
//   (u *User)  → Pointer Receiver (gets the address)
// ============================================================

type BankAccount struct {
	Owner   string
	Balance float64
}

// ----- VALUE RECEIVER -----
// Gets a COPY of BankAccount. Like taking a photocopy of a document.
// Any changes inside this method are made to the COPY, not the original.
func (b BankAccount) GetBalance() float64 {
	return b.Balance
}

// This method TRIES to modify, but it modifies the COPY. The original is SAFE.
func (b BankAccount) FakeDeposit(amount float64) {
	b.Balance += amount
	fmt.Printf("  [Inside FakeDeposit] Balance is now: %.2f (but this is the COPY!)\n", b.Balance)
}

// ----- POINTER RECEIVER -----
// Gets the ADDRESS of BankAccount. Like getting the key to a locker.
// Any changes inside this method affect the ORIGINAL.
func (b *BankAccount) Deposit(amount float64) {
	b.Balance += amount
	fmt.Printf("  [Inside Deposit] Balance is now: %.2f (this is the REAL one!)\n", b.Balance)
}

func (b *BankAccount) Withdraw(amount float64) error {
	if amount > b.Balance {
		return fmt.Errorf("insufficient funds: have %.2f, want %.2f", b.Balance, amount)
	}
	b.Balance -= amount
	return nil
}

// ============================================================
// INTERFACE BEHAVIOR WITH RECEIVERS
// ============================================================
// If an interface requires a method with a POINTER receiver,
// only a pointer can satisfy the interface.
// If it requires a VALUE receiver, both value and pointer work.

type Stringer interface {
	String() string
}

type Person struct {
	Name string
	Age  int
}

// Value receiver → Both Person and *Person satisfy Stringer
func (p Person) String() string {
	return fmt.Sprintf("%s (age %d)", p.Name, p.Age)
}

// ============================================================
// METHOD SETS: The Rule
// ============================================================
// Type T's method set    → only value receivers
// Type *T's method set   → both value AND pointer receivers
//
// This means:
//   var p Person = ...   → can only call value receiver methods
//   var p *Person = ...  → can call BOTH value and pointer methods
//
// But don't worry! Go auto-converts for you in most cases.
// ============================================================

func main() {
	fmt.Println("===== METHOD RECEIVERS DEEP DIVE =====")

	// 1. Value Receiver: Changes are LOST
	fmt.Println("\n--- 1. Value Receiver (FakeDeposit) ---")
	acc := BankAccount{Owner: "Harshit", Balance: 1000}
	fmt.Printf("Before FakeDeposit: %.2f\n", acc.Balance)
	acc.FakeDeposit(500)
	fmt.Printf("After FakeDeposit:  %.2f ← Unchanged! The 500 was lost.\n", acc.Balance)

	// 2. Pointer Receiver: Changes STICK
	fmt.Println("\n--- 2. Pointer Receiver (Deposit) ---")
	fmt.Printf("Before Deposit: %.2f\n", acc.Balance)
	acc.Deposit(500)
	fmt.Printf("After Deposit:  %.2f ← Changed! The 500 was added.\n", acc.Balance)

	// 3. Go's Auto-Conversion
	// Even though 'acc' is NOT a pointer, Go auto-converts acc.Deposit()
	// to (&acc).Deposit() for you. This is syntactic sugar!
	fmt.Println("\n--- 3. Auto-Conversion ---")
	accPtr := &acc
	accPtr.Deposit(200) // Pointer calling pointer method → normal
	acc.Deposit(100)    // Value calling pointer method → Go does (&acc).Deposit(100)
	fmt.Printf("Final Balance: %.2f\n", acc.Balance)

	// 4. Withdraw with Error Handling
	fmt.Println("\n--- 4. Pointer Receiver with Errors ---")
	err := acc.Withdraw(5000)
	if err != nil {
		fmt.Printf("Withdraw failed: %v\n", err)
	}
	err = acc.Withdraw(300)
	if err == nil {
		fmt.Printf("Withdrew 300. New Balance: %.2f\n", acc.Balance)
	}

	// 5. Interfaces with Value Receivers
	fmt.Println("\n--- 5. Interfaces & Method Sets ---")
	p := Person{Name: "Harshit", Age: 25}
	var s Stringer = p   // Works! Value receiver → both value and pointer satisfy
	var s2 Stringer = &p // Also works!
	fmt.Println("Value:", s)
	fmt.Println("Pointer:", s2)
}
