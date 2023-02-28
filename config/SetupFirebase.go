package config

import (
	"context"
	"log"
	"path/filepath"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"google.golang.org/api/option"
)

func SetupFirebase() (*auth.Client, *firestore.Client, error) {
	serviceAccountKeyFilePath, err := filepath.Abs("./credentials.json")
	if err != nil {
		log.Fatalf("Unable to load service account credentials: %v", err)
	}
	opt := option.WithCredentialsFile(serviceAccountKeyFilePath)

	// Firebase app initialization
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf("Error initializing Firebase app: %v", err)
	}

	// Firebase Auth initialization
	auth, err := app.Auth(context.Background())
	if err != nil {
		log.Fatalf("Error initializing Firebase Auth client: %v", err)
	}

	// Firestore initialization
	client, err := app.Firestore(context.Background())
	if err != nil {
		log.Fatalf("Error initializing Firestore client: %v", err)
	}

	return auth, client, nil
}
