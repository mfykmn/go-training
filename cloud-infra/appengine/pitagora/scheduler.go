package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"golang.org/x/oauth2/google"
	"google.golang.org/api/cloudscheduler/v1"
)

func timeToUnixCron(t time.Time) string {
	_, m, d := t.Date()
	return fmt.Sprintf("%d %d %d %d %d", t.Minute(), t.Hour(), d, m, t.Weekday())
}

type Scheduler struct {
	projectID string
	locationID string

	*cloudscheduler.Service
}

func newScheduler(ctx context.Context, projectID, locationID string) (*Scheduler, error){
	c, err := google.DefaultClient(ctx, cloudscheduler.CloudPlatformScope)
	if err != nil {
		return nil, err
	}
	service, err := cloudscheduler.New(c)
	if err != nil {
		return nil, err
	}
	return &Scheduler{
		projectID: projectID,
		locationID: locationID,
		Service: service}, nil
}

func (s *Scheduler) Reserve(ctx context.Context, t time.Time) error {
	parent := "projects/"+s.projectID+"/locations/"+s.locationID
	rb := &cloudscheduler.Job{
		Description: "Created by GAE",
		AppEngineHttpTarget: &cloudscheduler.AppEngineHttpTarget{
			HttpMethod: http.MethodGet,
			RelativeUri: "/db",
		},
		TimeZone: "Asia/Tokyo",
		Schedule: timeToUnixCron(t),
	}

	_, err := s.Projects.Locations.Jobs.Create(parent, rb).Context(ctx).Do()
	return err
}