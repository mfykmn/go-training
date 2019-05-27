package main

import (
	"context"
	"net/http"
	"os"

	"golang.org/x/oauth2/google"
	"google.golang.org/api/cloudscheduler/v1"
)

var (
	projectID = os.Getenv("PROJECT_ID")
	locationID = os.Getenv("LOCATION_ID")
)

type Scheduler struct {
	*cloudscheduler.Service
}

func newScheduler(ctx context.Context) (*Scheduler, error){
	c, err := google.DefaultClient(ctx, cloudscheduler.CloudPlatformScope)
	if err != nil {
		return nil, err
	}
	service, err := cloudscheduler.New(c)
	if err != nil {
		return nil, err
	}
	return &Scheduler{service}, nil
}

func (s *Scheduler) Reserve(ctx context.Context) error {
	parent := "projects/"+projectID+"/locations/"+locationID
	rb := &cloudscheduler.Job{
		Description: "Created by GAE",
		AppEngineHttpTarget: &cloudscheduler.AppEngineHttpTarget{
			HttpMethod: http.MethodGet,
			RelativeUri: "/db",
		},
		TimeZone: "Asia/Tokyo",
		Schedule: "* * * * *",
	}

	_, err := s.Projects.Locations.Jobs.Create(parent, rb).Context(ctx).Do()
	return err
}