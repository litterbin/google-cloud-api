package main

import (
	"golang.org/x/net/context"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/compute/v1"
)

import "testing"

// export GOOGLE_APPLICATION_CREDENTIALS=./a.json

func Test_Auth(t *testing.T) {
	// Use oauth2.NoContext if there isn't a good context to pass in.
	ctx := context.TODO()

	client, err := google.DefaultClient(ctx, compute.ComputeScope)
	if err != nil {
		t.Error(err)
	}
	computeService, err := compute.New(client)
	if err != nil {
		t.Error(err)
	}

	t.Logf("computeService: %v", computeService)
}
