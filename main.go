package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/cli/go-gh/v2/pkg/api"
)

func updateStatus(client *api.GraphQLClient, msg, emoji, busy string) {
	query := `
        mutation($message: String!, $emoji: String!, $limitedAvailability: Boolean!) {
            changeUserStatus(input: {
                message: $message
                emoji: $emoji
                limitedAvailability: $limitedAvailability
            }) {
                clientMutationId
            }
        }
    `

	variables := map[string]interface{}{
		"message":             msg,
		"emoji":               emoji,
		"limitedAvailability": busy == "true",
	}
	err := client.Do(query, variables, nil)
	if err != nil {
		fmt.Printf("❌ GitHub status update failed: %v\n", err)
	} else if busy == "true" {
		fmt.Printf("✅ GitHub status updated successfully: (%s)\n", msg)
	} else {
		fmt.Printf("✅ GitHub status cleared successfully\n")
	}
}

func viewStatus(client *api.GraphQLClient) {
	var res struct {
		Viewer struct {
			Status struct {
				Message                      string
				IndicatesLimitedAvailability bool
			}
		}
	}

	query := `{ viewer { status { message indicatesLimitedAvailability } } }`
	_ = client.Do(query, nil, &res)
	fmt.Printf("Message: %s\nBusy: %v\n", res.Viewer.Status.Message, res.Viewer.Status.IndicatesLimitedAvailability)
}

func main() {
	// Defaults
	msgString := flag.String("msg", "Busy", "Message will appear in your GitHub status")
	emojiString := flag.String("emoji", ":no_entry:", "Emoji will appear in your GitHub status")

	// Specific events
	// vacationBool := flag.Bool("vacation", false, "Set status to vacation")
	// sickBool := flag.Bool("sick", false, "Set status to sick")
	// busyBool := flag.Bool("busy", false, "Set status to busy")

	// Duration
	// For now, we won't implement duration parsing and we will set expiresAt to Never for simplicity, but this is where it would go
	//forString := flag.String("for", "", "Duration for status (e.g., '1h', '30m')")
	flag.Parse()

	client, err := api.DefaultGraphQLClient()
	if err != nil {
		log.Fatal(err)
	}

	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go [on|off|view|version]")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "on":
		updateStatus(client, *msgString, *emojiString, "true")
	case "off":
		updateStatus(client, "", "", "false")
	case "view":
		viewStatus(client)
	case "version":
		fmt.Println("gh-afk extension v0.1.0")
	default:
		fmt.Println("Usage: go run main.go [on|off|view|version]")
	}
}
