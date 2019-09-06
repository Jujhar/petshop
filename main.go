package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/manifoldco/promptui"
	"github.com/ziutek/mymysql/mysql"
	_ "github.com/ziutek/mymysql/native" // Native engine
)

var db mysql.Conn
var animalPurchased string // State of animal selected
var month time.Month       // State of current time
var week int               // State of current time
var cats int               // State of cats ready count
var dogs int               // State of dogs ready count
var birds int              // State of birds ready count

func dbConnect() {
	db = mysql.New("tcp", "", "127.0.0.1:3306", user, password, "petshop")
}

func countStock() {
	res, _ := db.Start("SELECT quantity_ready FROM stock")
	v := 0
	for {
		row, _ := res.GetRow()
		if row == nil {
			// No more rows
			break
		}
		for _, col := range row {
			if v == 0 {
				birds, _ = strconv.Atoi(string(col.([]byte)))
			} else if v == 1 {
				cats, _ = strconv.Atoi(string(col.([]byte)))
			} else if v == 2 {
				dogs, _ = strconv.Atoi(string(col.([]byte)))
			}
			v++
		}
	}
}

func displayReport() {
	res, _ := db.Start("select * from reports")

	// print fields names
	fmt.Println()
	i := 0
	for _, field := range res.Fields() {
		if i == 0 {
		} else if i == 1 {
			fmt.Print(" ", field.Name, "     ")
		} else if i == 2 {
			fmt.Print("", field.Name, "  ")
		} else {
			fmt.Print(field.Name, " ")
		}
		i++
	}
	fmt.Println()
	for {
		row, _ := res.GetRow()
		if row == nil {
			// No more rows
			break
		}

		// print all cols
		x := 0
		for _, col := range row {
			if x == 0 {
			} else {
				if col == nil {
					fmt.Print("<NULL>")
				} else {
					os.Stdout.Write(col.([]byte))
				}
				fmt.Print("    ")
			}
			x++
		}
		fmt.Println()
	}
	fmt.Println()
}

func showroomReport() {
	res, _ := db.Start("SELECT `animal`, `quantity_ready` FROM `stock` WHERE 1")
	fmt.Println()
	for {
		row, _ := res.GetRow()
		if row == nil {
			break
		}
		for _, col := range row {
			os.Stdout.Write(col.([]byte))
			fmt.Print(" ")
		}
		fmt.Println()
	}
	fmt.Println()
}

func occupancyReport() {
	res, _ := db.Start("SELECT `animal`, `quantity_available` FROM `stock` WHERE 1")
	fmt.Println()
	for {
		row, _ := res.GetRow()
		if row == nil {
			break
		}
		for _, col := range row {
			os.Stdout.Write(col.([]byte))
			fmt.Print(" ")
		}
		fmt.Println()
	}
	fmt.Println()
}

func deductStock(production bool, qty int) {
	if production == true {
		db.Start(fmt.Sprintf("UPDATE `stock` SET `quantity_ready`= `quantity_ready` - %d WHERE animal = '%vs'", qty, strings.Title(animalPurchased)))
	} else {
		db.Start(fmt.Sprintf("UPDATE `testing_stock` SET `quantity_ready`= `quantity_ready` - %d WHERE animal = '%vs'", qty, strings.Title(animalPurchased)))
	}
}

func logAction(production bool, action string, qty int) {
	if production == true {
		db.Start(fmt.Sprintf("INSERT INTO `transactions_log` (`date`, `log`) VALUES (NOW(), '%v')", fmt.Sprintf("%v X %v", action, qty)))
	} else {
		db.Start(fmt.Sprintf("INSERT INTO `testing_transactions_log` (`date`, `log`) VALUES (NOW(), '%v')", fmt.Sprintf("%v X %v", action, qty)))
	}
}

func addToReport(production bool, qty int) {
	if production == true {
		db.Start(fmt.Sprintf("INSERT INTO `reports` (`id`, `month`, `week`, `units_sold`) VALUES ('%v%d', '%v', '%d', '%d') ON DUPLICATE KEY UPDATE units_sold=units_sold + VALUES(units_sold)", month, week, month, week, qty))
	} else {
		db.Start(fmt.Sprintf("INSERT INTO `testing_reports` (`id`, `month`, `week`, `units_sold`) VALUES ('%v%d', '%v', '%d', '%d') ON DUPLICATE KEY UPDATE units_sold=units_sold + VALUES(units_sold)", month, week, month, week, qty))
	}
}

func purchase(qty int) {
	if qty < 1 {
		return
	}

	deductStock(true, qty)
	logAction(true, fmt.Sprintf("%v bought", animalPurchased), qty)

	if qty == 1 {
		fmt.Println("One pet purchased!")
	} else {
		fmt.Printf("%d %vs purchased!\n", qty, strings.ToLower(animalPurchased))
	}

	addToReport(true, qty)

}

func main() {
	animalsAvailible := []string{
		"Cat",
		"Dog",
		"Bird",
	}

	dbConnect()
	err := db.Connect()
	if err != nil {
		panic(err)
	}

	// identification table
	// chipped animals
	// insurance

	// current time
	dt := time.Now()
	month = dt.Month()
	day := dt.Day()
	week = 1

	switch {
	case day <= 8:
		week = 1
	case day <= 16:
		week = 2
	case day <= 24:
		week = 3
	case day <= 31:
		week = 4
	}

	fmt.Printf("%v %d", month, week)

	if dt.Format("Monday") == "Sunday" {
		fmt.Println("Shop closed!")
	}

	countStock()

	// select action
	prompt := promptui.Select{
		Label: "Select Action",
		Items: []string{"Buy pet",
			"See report (admin)",
			"Showroom animals (admin)",
			"Occupancy report (admin)"},
	}

	_, result, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	fmt.Printf("You chose %q\n", result)
	if result == "Buy pet" {

		// select animal
		prompt2 := promptui.Select{
			Label: "Select pet species",
			Items: animalsAvailible,
		}

		_, animal, err := prompt2.Run()

		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}

		fmt.Printf("You chose %q\n", animal)
		animalPurchased = animal

		// enter qty
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("How many?: ")
		quantity, _ := reader.ReadString('\n')
		quantity = strings.TrimSuffix(quantity, "\n")
		qty, _ := strconv.Atoi(quantity)
		if qty > 5 {
			fmt.Println("Woah! That many huh? No problem!")
		}

		// verify stock available
		if animal == "Cat" {
			if qty <= cats {
				purchase(qty)
			}
		} else if animal == "Dog" {
			if qty <= dogs {
				purchase(qty)
			}
		} else if animal == "Bird" {
			if qty <= birds {
				purchase(qty)
			}
		} else {
			fmt.Println("Sorry out of stock")
		}

		// insurance
		// pay 20 now
	}

	if result == "See report (admin)" {
		displayReport()
	}

	if result == "Showroom animals (admin)" {
		showroomReport()
	}

	if result == "Occupancy report (admin)" {
		occupancyReport()
	}

	// fmt.Println("Alerts: no animals to go to vet\nNo customers need to be contacted.")

}

/* Items: []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday",
"Saturday", "Sunday"}, */
