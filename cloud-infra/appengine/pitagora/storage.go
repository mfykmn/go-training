package main

import (
	"bytes"
	"context"
	"fmt"
	"io"

	"cloud.google.com/go/storage"
)

type Storage struct {
	projectID string

	*storage.Client
}

type Object struct {
	name string
	data string
}

func newStorage(ctx context.Context, projectID string) (*Storage, error) {
	c, err := storage.NewClient(ctx)
	if err != nil {
		return nil, err
	}
	return &Storage{projectID: projectID, Client: c}, nil
}

func (s *Storage) Download(ctx context.Context, bucket, objectName string) (*Object, error) {
	o := s.Bucket(bucket).Object(objectName)
	r, err := o.NewReader(ctx)
	if err != nil {
		return nil, err
	}
	defer r.Close()

	b := &bytes.Buffer{}
	if _, err := io.Copy(b, r); err != nil {
		return nil, err
	}
	return &Object{
		name: objectName,
		data: b.String(),
	}, nil
}

func (s *Storage) Upload(ctx context.Context, bucket string, obj *Object) error {
	o := s.Bucket(bucket).Object(obj.name)
	w := o.NewWriter(ctx)

	if _, err := fmt.Fprintf(w, obj.data); err != nil {
		return err
	}
	return w.Close()
}