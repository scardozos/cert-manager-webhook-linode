package main

import (
	"os"
	"testing"

	acmetest "github.com/cert-manager/cert-manager/test/acme"
	"sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"
)

var (
	zone = os.Getenv("TEST_ZONE_NAME")
)

func TestRunsSuite(t *testing.T) {
	/* The manifest path should contain a file named config.json that is a
	   snippet of valid configuration that should be included on the
	   ChallengeRequest passed as part of the test cases.*/
	log.SetLogger(zap.New(zap.UseDevMode(true)))
	fixture := acmetest.NewFixture(&linodeDNSProviderSolver{},
		acmetest.SetManifestPath("testdata/linode"),
		acmetest.SetResolvedZone(zone),
		acmetest.SetAllowAmbientCredentials(false),
		acmetest.SetDNSServer("ns1.linode.com:53"),
	)

	fixture.RunConformance(t)
}
