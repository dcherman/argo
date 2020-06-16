package commands

import (
		"fmt"
		wfv1 "github.com/argoproj/argo/pkg/apis/workflow/v1alpha1"
		"github.com/argoproj/pkg/s3"
		"os"

		log "github.com/sirupsen/logrus"
		"github.com/spf13/cobra"

		"github.com/argoproj/argo/workflow/artifacts/gcs"
)

func NewArtifactCommand() *cobra.Command {
		var command = cobra.Command{
				Use:   "resource (get|create|apply|delete) MANIFEST",
				Short: "update a resource and wait for resource conditions",
				Run: func(cmd *cobra.Command, args []string) {
						if len(args) != 1 {
								cmd.HelpFunc()(cmd, args)
								os.Exit(1)
						}
						err := execResource(args[0])
						if err != nil {
								log.Fatalf("%+v", err)
						}
				},
		}
		return &command
}

func deleteArtifact(artifact wfv1.Artifact) error {
		if artifact.ArtifactLocation.S3 != nil {
				s3Client, err := s3.NewS3Client(s3.S3ClientOpts{})

				if err != nil {
						return err
				}

				return s3Client.Delete(artifact.ArtifactLocation.S3.Bucket, artifact.ArtifactLocation.S3.Key)
		}

		if artifact.ArtifactLocation.GCS != nil {
				gcsClient :=
		}
}
