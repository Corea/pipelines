package integration

import (
	"flag"
	"time"
)

var (
	namespace           = flag.String("namespace", "kubeflow", "The namespace ml pipeline deployed to")
	initializeTimeout   = flag.Duration("initializeTimeout", 2*time.Minute, "Duration to wait for test initialization")
	runIntegrationTests = flag.Bool("runIntegrationTests", false, "Whether to also run integration tests that call the service")
	runUpgradeTests     = flag.Bool("runUpgradeTests", false, "Whether to run upgrade tests")
)

/**
 * Differences in dev mode:
 * 1. Resources are not cleaned up when a test finishes, so that developer can debug manually.
 * 2. One step that doesn't work locally is skipped.
 */
var isDevMode = flag.Bool("isDevMode", false, "Dev mode helps local development of integration tests")

var isDebugMode = flag.Bool("isDebugMode", false, "Whether to enable debug mode. Debug mode will log more diagnostics messages.")

var (
	isKubeflowMode    = flag.Bool("isKubeflowMode", false, "Runs tests in full Kubeflow mode")
	resourceNamespace = flag.String("resourceNamespace", "", "The namespace that will store the test resources in Kubeflow mode")
)
