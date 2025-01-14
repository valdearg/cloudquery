package client

import (
	"bytes"
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"

	"github.com/cloudquery/filetypes/csv"
	"github.com/cloudquery/filetypes/json"
	"github.com/cloudquery/plugin-sdk/plugins/destination"
	"github.com/cloudquery/plugin-sdk/specs"
	"github.com/rs/zerolog"
)

type Client struct {
	destination.UnimplementedUnmanagedWriter
	logger     zerolog.Logger
	spec       specs.Destination
	pluginSpec Spec

	s3Client   *s3.Client
	uploader   *manager.Uploader
	downloader *manager.Downloader

	csvTransformer         *csv.Transformer
	csvReverseTransformer  *csv.ReverseTransformer
	jsonTransformer        *json.Transformer
	jsonReverseTransformer *json.ReverseTransformer
}

func New(ctx context.Context, logger zerolog.Logger, spec specs.Destination) (destination.Client, error) {
	if spec.WriteMode != specs.WriteModeAppend {
		return nil, fmt.Errorf("destination only supports append mode")
	}
	c := &Client{
		logger:                 logger.With().Str("module", "s3").Logger(),
		spec:                   spec,
		csvTransformer:         &csv.Transformer{},
		jsonTransformer:        &json.Transformer{},
		csvReverseTransformer:  &csv.ReverseTransformer{},
		jsonReverseTransformer: &json.ReverseTransformer{},
	}

	if err := spec.UnmarshalSpec(&c.pluginSpec); err != nil {
		return nil, fmt.Errorf("failed to unmarshal postgresql spec: %w", err)
	}
	if err := c.pluginSpec.Validate(); err != nil {
		return nil, err
	}
	c.pluginSpec.SetDefaults()

	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return nil, fmt.Errorf("unable to load AWS SDK config: %w", err)
	}

	c.s3Client = s3.NewFromConfig(cfg)
	c.uploader = manager.NewUploader(c.s3Client)
	c.downloader = manager.NewDownloader(c.s3Client)

	// we want to run this test because we want it to fail early if the bucket is not accessible
	if _, err := c.uploader.Upload(ctx, &s3.PutObjectInput{
		Bucket: aws.String(c.pluginSpec.Bucket),
		Key:    aws.String("/tmp/.cq-test-file"),
		Body:   bytes.NewReader([]byte("")),
	}); err != nil {
		return nil, fmt.Errorf("failed to write test file to S3: %w", err)
	}
	return c, nil
}

func (*Client) Close(ctx context.Context) error {
	return nil
}
