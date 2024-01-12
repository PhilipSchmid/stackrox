package ecr

import (
	"encoding/base64"
	"net/http"
	"strings"
	"time"

	awsECR "github.com/aws/aws-sdk-go/service/ecr"
	"github.com/heroku/docker-registry-client/registry"
	"github.com/pkg/errors"
	"github.com/stackrox/rox/pkg/registries/docker"
	"github.com/stackrox/rox/pkg/sync"
)

type awsTransport struct {
	registry.Transport
	config    *docker.Config
	client    *awsECR.ECR
	expiresAt *time.Time
	mutex     sync.Mutex
}

func newAWSTransport(config *docker.Config, client *awsECR.ECR) *awsTransport {
	transport := &awsTransport{config: config, client: client}
	if err := transport.refreshNoLock(); err != nil {
		log.Error("Failed to refresh ECR token: ", err)
	}
	return transport
}

func (t *awsTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	t.mutex.Lock()
	defer t.mutex.Unlock()
	if t.isExpiredNoLock() {
		if err := t.refreshNoLock(); err != nil {
			return nil, err
		}
	}
	return t.Transport.RoundTrip(req)
}

func (t *awsTransport) isExpiredNoLock() bool {
	return t.expiresAt == nil || t.expiresAt.After(time.Now())
}

func (t *awsTransport) refreshNoLock() error {
	authToken, err := t.client.GetAuthorizationToken(&awsECR.GetAuthorizationTokenInput{})
	if err != nil {
		return errors.Wrap(err, "failed to get authorization token")
	}
	if len(authToken.AuthorizationData) == 0 {
		return errors.New("received empty authorization data in token")
	}
	authData := authToken.AuthorizationData[0]
	decoded, err := base64.StdEncoding.DecodeString(*authData.AuthorizationToken)
	if err != nil {
		return errors.Wrap(err, "failed to decode authorization token")
	}
	username, password, ok := strings.Cut(string(decoded), ":")
	if !ok {
		return errors.New("malformed basic auth response from AWS")
	}
	t.expiresAt = authData.ExpiresAt
	t.config.Username = username
	t.config.Password = password
	t.Transport = docker.DefaultTransport(t.config)
	return nil
}