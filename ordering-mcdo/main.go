package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
)

type MenuItem struct {
	ID			int
	Name		string
	Price		float64
	Category	string
}

type OrderItem struct {
	Item		MenuItem
	Quantity	int
}

var menu = []MenuItem{
	// Burgers
	{1, "Burger McDo", 55.00, "Burgers"},
	{2, "Cheesy Burger McDo", 65.00, "Burgers"},
	{3, "Big Mac", 195.00, "Burgers"},
	{4, "Quarter Pounder with Cheese", 185.00, "Burgers"},
	{5, "McCrispy Chicken Fillet Sandwich", 105.00, "Burgers"},
	{6, "Filet-O-Fish", 120.00, "Burgers"},
	{7, "McChicken", 110.00, "Burgers"},
	
	// Chicken Meals
	{8, "1-pc Chicken McDo with Rice", 95.00, "Chicken"},
	{9, "2-pc Chicken McDo with Rice", 175.00, "Chicken"},
	{10, "6-pc Chicken McNuggets", 135.00, "Chicken"},
	{11, "McWings (2pcs)", 85.00, "Chicken"},
	{12, "McCrispy Chicken Fillet Ala King with Rice", 105.00, "Chicken"},
	
	// Rice Bowls
	{13, "McSpaghetti", 55.00, "Rice Bowls"},
	{14, "Chicken McDo with McSpaghetti", 130.00, "Rice Bowls"},
	{15, "Burger McDo with Fries", 85.00, "Rice Bowls"},
	
	// Drinks
	{16, "Coke (Regular)", 45.00, "Drinks"},
	{17, "Sprite (Regular)", 45.00, "Drinks"},
	{18, "Iced Tea (Regular)", 40.00, "Drinks"},
	{19, "Orange Juice", 50.00, "Drinks"},
	{20, "Pineapple Juice", 50.00, "Drinks"},
	
	// Sides
	{21, "French Fries (Medium)", 55.00, "Sides"},
	{22, "French Fries (Large)", 75.00, "Sides"},
	{23, "Apple Pie", 35.00, "Sides"},
	
	// Desserts
	{24, "McFlurry Oreo", 65.00, "Desserts"},
	{25, "Hot Fudge Sundae", 35.00, "Desserts"},
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var cart []OrderItem

	fmt.Println("\n\n\n")

	for {
		clearScreen()
		fmt.Println("╔════════════════════════════════════════════════╗")
		fmt.Println("║     McDONALD'S PHILIPPINES ORDERING SYSTEM     ║")
		fmt.Println("╚════════════════════════════════════════════════╝")
		fmt.Println("\n[1] View Menu")
		fmt.Println("[2] Add Item to Cart")
		fmt.Println("[3] View Cart")
		fmt.Println("[4] Remove Item from Cart")
		fmt.Println("[5] Checkout")
		fmt.Println("[6] Exit")
		fmt.Print("\nSelect an option: ")

		scanner.Scan()
		choice := scanner.Text()

		switch choice {
		case "1":
			clearScreen()
			displayMenu()
			waitForEnter(scanner)
		case "2":
			clearScreen()
			cart = addItemToCart(scanner, cart)
		case "3":
			viewCart(cart)
			waitForEnter(scanner)
		case "4":
			clearScreen()
			cart = removeFromCart(scanner, cart)
		case "5":
			checkout(cart)
			cart = []OrderItem{}
			waitForEnter(scanner)
		case "6":
			fmt.Println("\nThank you for visiting McDonald's! Goodbye! 👋")
			return
		default:
			fmt.Println("\n❌Invalid input. Please try again")
			waitForEnter(scanner)
		}
	}
}

func displayMenu() {
	categories := []string{"Burgers", "Chicken", "Rice Bowls", "Drinks", "Sides", "Desserts"}

	fmt.Println("\n╔════════════════════════════════════════════════════════════╗")
	fmt.Println("║                    McDONALD'S MENU                         ║")
	fmt.Println("╚════════════════════════════════════════════════════════════╝")

	for _, cat := range categories {
		fmt.Printf("\n\n %s\n", cat)
		fmt.Println(strings.Repeat("-", 60))

		for _, item := range menu {
			if cat == item.Category {
				fmt.Printf("\n[%d] %-40s ₱%.2f\n", item.ID, item.Name, item.Price)
			}
		}
	}
}

func addItemToCart(scanner *bufio.Scanner, cart []OrderItem) []OrderItem {
	displayMenu()

	fmt.Print("\nEnter item number to order (or 0 to cancel): ")
	scanner.Scan()

	idStr := scanner.Text()
	no, err := strconv.Atoi(idStr)
	
	if err != nil || no <= 0 {
		fmt.Println("\nCancelled!")
		waitForEnter(scanner)
		return cart
	}

	var orderItem *MenuItem
	for _, item := range menu {
		if item.ID == no {
			orderItem = &item
			break
		}
	}

	if orderItem == nil {
		fmt.Println("❌ Unavailable item.")
		waitForEnter(scanner)
		return cart
	}

	// --- quantity

	fmt.Print("Enter the quantity: ")
	scanner.Scan()

	qtyStr := scanner.Text()
	qty, err := strconv.Atoi(qtyStr)

	if err != nil || qty <= 0 {
		fmt.Println("Invalid quantity")
		waitForEnter(scanner)
		return cart
	}

	// --- increase quantity if item is already exists
	// I'm using indexing here to reflect the new value to existing object(cart)
	found := false
	for i := range cart {
		if orderItem.ID == cart[i].Item.ID {
			cart[i].Quantity += qty; // <--
			found = true
			break
		}
	}

	// --- new item
	if !found {
		cart = append(cart, OrderItem{
			Item: *orderItem,
			Quantity: qty,
		})
	}

	fmt.Printf("\nAdded %dx %s to cart\n", qty, orderItem.Name)
	waitForEnter(scanner)
	return cart
}

func viewCart(cart []OrderItem) {
	clearScreen()

	fmt.Println("\n╔════════════════════════════════════════════════════════════╗")
	fmt.Println("║                         YOUR CART                          ║")
	fmt.Println("╚════════════════════════════════════════════════════════════╝")

	if len(cart) == 0 {
		fmt.Println("🛒 Your cart is empty!")
		return
	}

	total := 0.0
	for i, item := range cart {
		subtotal := item.Item.Price * float64(item.Quantity)
		fmt.Printf("\n[%d] %dx %40s ₱%.2f\n",
			i+1, item.Quantity, item.Item.Name, subtotal)
		total += subtotal
	}

	fmt.Println(strings.Repeat("-", 60))
	fmt.Printf("Total: ₱%.2f\n", total)
}

func removeFromCart(scanner *bufio.Scanner, cart []OrderItem) []OrderItem {
	viewCart(cart)

	if len(cart) == 0 {
		waitForEnter(scanner)
		return cart
	}

	fmt.Print("Enter item number to remove (or 0 to cancel): ")
	scanner.Scan()

	noStr := scanner.Text()
	no, err := strconv.Atoi(noStr)

	if err != nil || no == 0 {
		fmt.Println("❌ Cancelled!")
		waitForEnter(scanner)
		return cart
	}

	if no < 1 || no > len(cart) {
		fmt.Println("Invalid item number")
		return cart
	}

	removedItem := cart[no-1].Item.Name
	cart = append(cart[:no-1], cart[no-1+1:]...)

	fmt.Printf("\n✅ Removed %s from cart!\n", removedItem)
	waitForEnter(scanner)
	return cart
}

func checkout(cart []OrderItem) {
	clearScreen()

	fmt.Println("╔════════════════════════════════════════════════════════════╗")
	fmt.Println("║                          CHECKOUT                          ║")
	fmt.Println("╚════════════════════════════════════════════════════════════╝\n")

	if len(cart) == 0 {
		fmt.Println("🛒 Your cart is empty!")
		return
	}

	fmt.Println("Order Summary")
	fmt.Println(strings.Repeat("-", 60))

	total := 0.0
	for _, item := range cart {
		subtotal := item.Item.Price * float64(item.Quantity)
		fmt.Printf("\n%dx %-40s ₱%.2f\n", item.Quantity, item.Item.Name, subtotal)
		total += subtotal
	}

	fmt.Println(strings.Repeat("=", 60))
	fmt.Printf("Total Amount: ₱%.2f\n", total)
	fmt.Println(strings.Repeat("=", 60))

	fmt.Println("\n✅ Order placed successfully!")
	fmt.Println("🍔 Your order is being prepared. Thank you!")
}

func waitForEnter(scanner *bufio.Scanner) {
	fmt.Print("\nPress Enter to continue...")
	scanner.Scan()
}

func clearScreen() {
	fmt.Print("\033[H\033[2J")
}