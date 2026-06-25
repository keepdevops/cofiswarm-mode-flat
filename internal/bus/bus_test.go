package bus

import (
	"encoding/json"
	"testing"

	"github.com/keepdevops/cofiswarm-observer-sdk/pkg/servicecomponent"
)

func TestInfoRouteReturnsMode(t *testing.T) {
	subj := servicecomponent.Prefix + ".mode.flat.info"
	out, err := Routes("flat")[subj](nil)
	if err != nil {
		t.Fatal(err)
	}
	if r := out.(infoReply); !r.OK || r.Mode != "flat" {
		t.Fatalf("got %+v", r)
	}
}

func TestHealthRouteOK(t *testing.T) {
	subj := servicecomponent.Prefix + ".mode.flat.health"
	out, _ := Routes("flat")[subj](nil)
	if r := out.(healthReply); !r.OK || r.Status != "ok" {
		t.Fatalf("got %+v", r)
	}
}

func TestReplyCarriesSchemaVersion(t *testing.T) {
	subj := servicecomponent.Prefix + ".mode.flat.info"
	out, _ := Routes("flat")[subj](nil)
	b, _ := json.Marshal(out)
	var m map[string]any
	_ = json.Unmarshal(b, &m)
	if m["schema_version"] != servicecomponent.SchemaVersion {
		t.Fatalf("schema_version = %v, want %s", m["schema_version"], servicecomponent.SchemaVersion)
	}
}
