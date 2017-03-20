package ctrl

import (
	"reflect"
	"testing"

	"github.com/goadesign/goa"
	core "github.com/ihcsim/bluelens"
	"github.com/ihcsim/bluelens/server/app/test"
	"github.com/ihcsim/bluelens/server/store"
)

const maxCount = 20

func TestRecommendationsController(t *testing.T) {
	// mock the store() function
	storeFunc := store.Instance
	storeFuncMock := func() core.Store {
		s, err := core.NewFixtureStore()
		if err != nil {
			t.Fatal("Unexpected error: ", err)
		}
		return s
	}
	store.Instance = storeFuncMock
	defer func() {
		store.Instance = storeFunc
	}()
	fixtureStore := store.Instance().(*core.FixtureStore)

	// retrieve the expected recommendations from the fixture store
	recommendations, err := fixtureStore.Recommendations(maxCount)
	if err != nil {
		t.Fatal("Unexpected error: ", err)
	}

	service := goa.New("goatest")
	ctrl := NewRecommendationsController(service)

	// test the recommendations of all users in the store
	users, err := store.Instance().ListUsers(maxCount)
	if err != nil {
		t.Fatal("Unexpected error: ", err)
	}
	for _, user := range users {
		_, actual := test.RecommendRecommendationsOK(t, nil, nil, ctrl, user.ID, maxCount)
		expected := mediaTypeRecommendations(recommendations[user.ID])
		if !reflect.DeepEqual(expected, actual) {
			t.Errorf("Recommendations response mismatch. Expected:\n%+v\nBut got:\n%+v", expected, actual)
		}
	}
}