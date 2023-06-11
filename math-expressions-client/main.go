package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"test-domain.com/math-expressions-client/client"
)

func main() {
	config := LoadConfig()
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Enter 'evaluate', 'validate', or 'errors': ")
		choice, err := reader.ReadString('\n')
		if err != nil {
			return
		}
		choice = strings.TrimSpace(choice)

		switch choice {
		case "evaluate", "validate":
			fmt.Print("Enter an expression: ")
			expression, err := reader.ReadString('\n')
			if err != nil {
				return
			}
			expression = strings.TrimSpace(expression)

			if choice == "evaluate" {
				val := client.SendEvaluateRequest(config.ServiceAddress, expression)
				fmt.Println("Result:", val)
			} else {
				valid, reason := client.SendValidateRequest(config.ServiceAddress, expression)
				if valid {
					fmt.Printf("Valid: %t\n", valid)
				} else {
					fmt.Printf("Valid: %t, Reason: %s\n", valid, reason)
				}
			}
		case "errors":
			errors, err := client.SendErrorsRequest(config.ServiceAddress)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				if len(errors) == 0 {
					fmt.Println("No registred errors yet.")
				} else {
					for _, errJson := range errors {
						fmt.Printf("Expression: %s, Endpoint: %s, Frequency: %d, Type: %s\n",
							errJson.Expression, errJson.Endpoint, errJson.Frequency, errJson.Type)
					}
				}
			}
		default:
			fmt.Println("Invalid choice")
		}

		fmt.Print("Press 'q' to quit or any other key to continue: ")
		quit, err := reader.ReadString('\n')
		if err != nil {
			return
		}
		if strings.TrimSpace(quit) == "q" {
			break
		}
	}
}
