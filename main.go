package main

import (
	"fmt"
  "time"

	"github.com/manifoldco/promptui"
)

func main() {

  // Data
  AnimalsAvailible := []string{
		"Cat",
		"Dog",
		"Bird",
	}

  // Current time
  dt := time.Now()
  fmt.Println(dt.Format("Monday"))
  if (dt.Format("Sunday") == "Wednesday") {
    fmt.Println("Shop closed!\n")
  }

  // Select action
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

	fmt.Printf("You choose %q\n", result)

  // Select animal

  prompt2 := promptui.Select{
		Label: "Select pet species",
		Items: AnimalsAvailible,
	}

	_, animal, err := prompt2.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	fmt.Printf("You choose %q\n", animal)

  // Select animal

  // On Buy
  // Add to report

  fmt.Println("Alerts: no animals to go to vet\nNo customers need to be contacted.")

}



/* Items: []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday",
  "Saturday", "Sunday"}, */
