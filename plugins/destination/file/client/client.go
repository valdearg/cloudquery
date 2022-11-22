package client

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"strings"

	"cloud.google.com/go/storage"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/cloudquery/plugin-sdk/plugins"
	"github.com/cloudquery/plugin-sdk/specs"
	"github.com/rs/zerolog"
)


const batchSize = 1024*1024*4

type Client struct {
	plugins.DefaultReverseTransformer
	logger  zerolog.Logger
	spec    specs.Destination
	csvSpec Spec

	bucket string
	dir string

	awsUploader      *manager.Uploader
	awsDownloader    *manager.Downloader

	gcpStorageClient *storage.Client
	gcpBucket 			 *storage.BucketHandle

	testMode bool

	metrics plugins.DestinationMetrics
}

func New(ctx context.Context, logger zerolog.Logger, spec specs.Destination) (plugins.DestinationClient, error) {
	if spec.WriteMode != specs.WriteModeAppend {
		return nil, fmt.Errorf("file destination only supports append mode")
	}
	c := &Client{
		logger: logger.With().Str("module", "dest-file").Logger(),
		spec:   spec,
	}

	if err := spec.UnmarshalSpec(&c.csvSpec); err != nil {
		return nil, fmt.Errorf("failed to unmarshal postgresql spec: %w", err)
	}
	c.csvSpec.SetDefaults()

	switch c.csvSpec.Backend {
	case BackendTypeLocal:
		if err := os.MkdirAll(c.csvSpec.Directory, 0755); err != nil {
			return nil, fmt.Errorf("failed to create directory: %w", err)
		}
	case BackendTypeS3:
		awsCfg, err := config.LoadDefaultConfig(ctx)
		if err != nil {
			return nil, fmt.Errorf("unable to load AWS SDK config, %w", err)
		}
		awsClient := s3.NewFromConfig(awsCfg)
		c.awsUploader = manager.NewUploader(awsClient)
		c.awsDownloader = manager.NewDownloader(awsClient)

		split := strings.Split(c.csvSpec.Directory, "/")
		if len(split) == 0 {
			return nil, fmt.Errorf("invalid S3 Bucket: %s", c.csvSpec.Directory)
		}
		c.dir = strings.Join(split[1:], "/")
		c.bucket = split[0]
		
		if _, err := c.awsUploader.Upload(ctx, &s3.PutObjectInput{
			Bucket: aws.String(c.csvSpec.Directory),
			Key:    aws.String(c.dir + "/cq-test-file"),
			Body:   bytes.NewReader([]byte("test")),
		}); err != nil {
			return nil, fmt.Errorf("failed to write test file to S3: %w", err)
		}
	case BackendTypeGCS:
		var err error
		c.gcpStorageClient, err = storage.NewClient(ctx)
		if err != nil {
			return nil, fmt.Errorf("failed to create GCP storage client %w", err)
		}
		split := strings.Split(c.csvSpec.Directory, "/")
		if len(split) == 0 {
			return nil, fmt.Errorf("invalid GCS Bucket: %s", c.csvSpec.Directory)
		}
		c.dir = strings.Join(split[1:], "/")
		c.bucket = split[0]
		c.gcpBucket = c.gcpStorageClient.Bucket(c.bucket)
		gcpWriter := c.gcpBucket.Object(c.dir + "/cq-test-file").NewWriter(ctx)
		if _, err := gcpWriter.Write([]byte("test")); err != nil {
			return nil, fmt.Errorf("failed to write test file to GCS: %w", err)
		}
		if err := gcpWriter.Close(); err != nil {
			return nil, fmt.Errorf("failed to close GCS writer: %w", err)
		}
	default:
		return nil, fmt.Errorf("unknown backend: %s", c.csvSpec.Backend)
	}

	return c, nil
}

func (*Client) Close(ctx context.Context) error {
	return nil
}
