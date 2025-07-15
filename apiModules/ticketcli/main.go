package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

type Ticket struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

var apiURL string

func main() {
	var rootCmd = &cobra.Command{Use: "ticketcli"}
	rootCmd.PersistentFlags().StringVar(&apiURL, "api", "http://localhost:8080", "Ticket API base URL")

	rootCmd.AddCommand(createCmd, getCmd, updateCmd, deleteCmd, listCmd)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new ticket",
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		typeVal, _ := cmd.Flags().GetString("type")
		ticket := Ticket{Name: name, Type: typeVal}
		data, _ := json.Marshal(ticket)
		resp, err := http.Post(apiURL+"/tickets", "application/json", bytes.NewBuffer(data))
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)
		fmt.Println(string(body))
	},
}

func init() {
	createCmd.Flags().String("name", "", "Ticket name")
	createCmd.Flags().String("type", "", "Ticket type (kindA, kindB, kindC)")
	createCmd.MarkFlagRequired("name")
	createCmd.MarkFlagRequired("type")
}

var getCmd = &cobra.Command{
	Use:   "get [id]",
	Short: "Get a ticket by ID",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		resp, err := http.Get(fmt.Sprintf("%s/tickets/%s", apiURL, args[0]))
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)
		fmt.Println(string(body))
	},
}

var updateCmd = &cobra.Command{
	Use:   "update [id]",
	Short: "Update a ticket by ID",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		typeVal, _ := cmd.Flags().GetString("type")
		ticket := Ticket{Name: name, Type: typeVal}
		data, _ := json.Marshal(ticket)
		req, _ := http.NewRequest("PUT", fmt.Sprintf("%s/tickets/%s", apiURL, args[0]), bytes.NewBuffer(data))
		req.Header.Set("Content-Type", "application/json")
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)
		fmt.Println(string(body))
	},
}

func init() {
	updateCmd.Flags().String("name", "", "Ticket name")
	updateCmd.Flags().String("type", "", "Ticket type (kindA, kindB, kindC)")
	updateCmd.MarkFlagRequired("name")
	updateCmd.MarkFlagRequired("type")
}

var deleteCmd = &cobra.Command{
	Use:   "delete [id]",
	Short: "Delete a ticket by ID",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		req, _ := http.NewRequest("DELETE", fmt.Sprintf("%s/tickets/%s", apiURL, args[0]), nil)
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		defer resp.Body.Close()
		fmt.Println("Deleted ticket", args[0])
	},
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all tickets",
	Run: func(cmd *cobra.Command, args []string) {
		resp, err := http.Get(apiURL + "/tickets")
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)
		fmt.Println(string(body))
	},
}
