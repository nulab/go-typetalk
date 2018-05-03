package v1

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"testing"

	. "github.com/nulab/go-typetalk/typetalk/internal"
)

func Test_OrganizationsService_GetMyOrganizations_should_get_my_organizations(t *testing.T) {
	setup()
	defer teardown()
	b, _ := ioutil.ReadFile(fixturesPath + "get-my-organizations.json")
	mux.HandleFunc("/spaces",
		func(w http.ResponseWriter, r *http.Request) {
			TestMethod(t, r, "GET")
			TestQueryValues(t, r, Values{"excludesGuest": true})
			fmt.Fprint(w, string(b))
		})

	result, _, err := client.Organizations.GetMyOrganizations(context.Background(), true)
	if err != nil {
		t.Errorf("Returned error: %v", err)
	}
	var want *struct {
		MySpaces []*Organization `json:"mySpaces"`
	}
	json.Unmarshal(b, &want)
	if !reflect.DeepEqual(result, want.MySpaces) {
		t.Errorf("Returned result:\n result  %v,\n want %v", result, want)
	}
}

func Test_OrganizationsService_GetOrganizationMembers_should_get_some_organization_members(t *testing.T) {
	setup()
	defer teardown()
	spaceKey := "test"
	b, _ := ioutil.ReadFile(fixturesPath + "get-organization-members.json")
	mux.HandleFunc(fmt.Sprintf("/spaces/%s/members", spaceKey),
		func(w http.ResponseWriter, r *http.Request) {
			TestMethod(t, r, "GET")
			fmt.Fprint(w, string(b))
		})

	result, _, err := client.Organizations.GetOrganizationMembers(context.Background(), spaceKey)
	if err != nil {
		t.Errorf("Returned error: %v", err)
	}
	want := &OrganizationMembers{}
	json.Unmarshal(b, want)
	if !reflect.DeepEqual(result, want) {
		t.Errorf("Returned result:\n result  %v,\n want %v", result, want)
	}
}
