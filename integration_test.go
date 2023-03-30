package main

import (
	"aprixa/datasource"
	"aprixa/service"
	"context"
	"fmt"
	"os"
	"testing"

 "github.com/testcontainers/testcontainers-go"
_ "github.com/testcontainers/testcontainers-go/modules/compose"

)

var db datasource.Database

func TestMain(m *testing.M) {
	ctx := context.Background()

	// Start the Docker Compose services
	compose := testcontainers.NewLocalDockerCompose([]string{"docker-compose.yml"}, "aprixia-test")
	err := compose.WithContext(ctx).Up()
	if err != nil {
		fmt.Printf("Failed to start Docker Compose: %v\n", err)
		os.Exit(1)
	}
	defer compose.Down()

	// Connect to the PostgreSQL database
	db, err = datasource.NewDB()
	if err != nil {
		fmt.Printf("Failed to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer db.Close()

	// Run the tests
	exitCode := m.Run()

	os.Exit(exitCode)
}

func TestInsertData(t * testing.T){
    exampleUrl := datasource.UrlData{
        RealUrl: "http://www.example.com",
        ShortUrl: "abc123",
        CreateAt: "2023-03-30 17:11:29",
        RedirectCount: 0,
    }
    // Insert test data into the database
    err := db.InsertData(exampleUrl) 
    if err != nil {
        t.Fatalf("Error inserting test data into database: %v", err)
    }
}

func TestGetUrlDataIntegration(t *testing.T) {
        // Call the function being tested
    urlData, err := service.GetUrlData("abc123")

    // Check the results
    if err != nil {
        t.Fatalf("Error getting URL data: %v", err)
    }
    if urlData.RealUrl != "http://www.example.com" {
        t.Errorf("Expected long URL to be %q, but got %q", "http://www.example.com", urlData.RealUrl)
    }
}
