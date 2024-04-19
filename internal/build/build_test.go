package build

import (
	"testing"
)

func Test_getDependencyInfoFromData(t *testing.T) {
	content := `
{
        "Path": "github.com/gregfurman/genver",
        "Main": true,
        "Dir": "/genver",
        "GoMod": "/genver/go.mod",
        "GoVersion": "1.21.5"
}
{
        "Path": "golang.org/x/crypto",
        "Version": "v0.21.0",
        "Time": "2024-03-04T18:29:30Z",
        "Indirect": true
}
{
        "Path": "golang.org/x/mod",
        "Version": "v0.8.0",
        "Time": "2023-02-02T20:50:06Z",
        "Indirect": true
}
{
        "Path": "golang.org/x/net",
        "Version": "v0.22.0",
        "Time": "2024-03-04T19:59:26Z",
        "Dir": "/go/pkg/mod/golang.org/x/net@v0.22.0",
        "GoMod": "/go/pkg/mod/cache/download/golang.org/x/net/@v/v0.22.0.mod",
        "GoVersion": "1.18"
}
{
        "Path": "golang.org/x/sys",
        "Version": "v0.18.0",
        "Time": "2024-02-16T13:31:09Z",
        "Indirect": true
}
{
        "Path": "golang.org/x/term",
        "Version": "v0.18.0",
        "Time": "2024-03-04T17:33:41Z",
        "Indirect": true
}
{
        "Path": "golang.org/x/text",
        "Version": "v0.14.0",
        "Time": "2023-10-11T21:58:48Z",
        "Dir": "/go/pkg/mod/golang.org/x/text@v0.14.0",
        "GoMod": "/go/pkg/mod/cache/download/golang.org/x/text/@v/v0.14.0.mod",
        "GoVersion": "1.18"
}
{
        "Path": "golang.org/x/tools",
        "Version": "v0.6.0",
        "Time": "2023-02-08T22:43:44Z",
        "Indirect": true
}`

	expMap := map[string]string{
		nameFromPath("golang.org/x/crypto"): "v0.21.0",
		nameFromPath("golang.org/x/mod"):    "v0.8.0",
		nameFromPath("golang.org/x/net"):    "v0.22.0",
		nameFromPath("golang.org/x/sys"):    "v0.18.0",
		nameFromPath("golang.org/x/term"):   "v0.18.0",
		nameFromPath("golang.org/x/text"):   "v0.14.0",
		nameFromPath("golang.org/x/tools"):  "v0.6.0",
	}

	deps := getDependencyInfoFromData(content)
	if len(deps) != len(expMap) {
		t.Errorf("incorrect number of dependencies found: want %d, found %d", len(expMap), len(deps))
		t.FailNow()
	}

	for _, dep := range deps {
		expVersion, ok := expMap[dep.Name]
		if !ok {
			t.Errorf("Could not find dependency [%s] found in expected dependency list", dep.Name)
			continue
		}

		if expVersion != dep.Version {
			t.Errorf("Incorrect version found for dependency [%s]: want %s, found %s", dep.Name, expVersion, dep.Version)
			continue
		}

		delete(expMap, dep.Name)
	}

	if len(expMap) > 0 {
		t.Errorf("did not find all expected dependencies: remaining [%v]", expMap)
	}

}

func Test_nameFromPath(t *testing.T) {
	pathNames := []string{
		"golang.org/x/crypto",
		"golang.org/x/mod",
		"golang.org/x/net",
		"golang.org/x/sys",
		"golang.org/x/term",
		"golang.org/x/text",
		"golang.org/x/tools",
		"github.com/aws/aws-sdk-go-v2",
		"github.com/aws/aws-sdk-go",
		"github.com/aws/smithy-go",
	}

	expRenderedName := []string{
		"XCrypto_Golang",
		"XMod_Golang",
		"XNet_Golang",
		"XSys_Golang",
		"XTerm_Golang",
		"XText_Golang",
		"XTools_Golang",
		"AwsAwsSdkGoV2_Github",
		"AwsAwsSdkGo_Github",
		"AwsSmithyGo_Github",
	}

	for i, v := range pathNames {
		if want, got := nameFromPath(v), expRenderedName[i]; want != got {
			t.Errorf("Rendered name from path is incorrect: got %s, want %s", got, want)
		}
	}

}
