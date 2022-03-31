package goclient

import (
	"testing"

	"github.com/Jeffail/gabs/v2"
	"github.com/stretchr/testify/require"
)

const (
	session1 = "session1"
	session2 = "session2"
)

/**
* This is a the pre-interview test to make sure everything will run on the interview day
 */

// TestEverything does an end to end test of the server
// Using gabs to pluck values out of json to save us from having to create a bunch of structs
// you can read more here https://github.com/Jeffail/gabs
// In general the client will return a *gabs.Container as an api response. Fields within the response
// can be accessed via a jsonp or jq syntax. e.g. if response is `{ "foo": "bar", "fooz": { "id": 1 } }` you can access the "id"
// with an expression like `response.Path("fooz.id").Data().(string)`
func TestEverything(t *testing.T) {

	var client = New()
	var err error

	var require = require.New(t)

	// establish 3 socket connections
	err = client.Connect("session1")
	require.NoError(err)
	err = client.Connect("session2")
	require.NoError(err)

	var (
		session1ID, session2ID float64
		response               *gabs.Container
		method                 string
	)

	t.Run("create", func(t *testing.T) {
		// Create Session 1
		method = "test.create"
		response, err = client.SendRequest(session1, method, map[string]interface{}{"name": "session1", "random_string": "ses11"})
		require.NoError(err, "%s failed as %s", method, session1)

		require.True(response.Path("ok").Data().(bool), "%s failed as %s", method, session1)
		session1ID = response.Path("test.id").Data().(float64)
		require.Greater(session1ID, float64(0))

		// Create Session 2
		method = "test.create"
		response, err = client.SendRequest(session2, method, map[string]interface{}{"name": "session2", "random_string": "ses22"})
		require.NoError(err, "%s failed as %s", method, session2)

		require.True(response.Path("ok").Data().(bool), "%s failed as %s", method, session2)
		session2ID = response.Path("test.id").Data().(float64)
		require.Greater(session2ID, float64(0))

		// Get by session 1
		method = "test.info"
		response, err = client.SendRequest(session1, method, map[string]interface{}{"test_id": session1ID})
		require.NoError(err, "%s failed as %s", method, session1)
		id := response.Path("test.id").Data().(float64)
		require.Equal(id, session1ID)
		name := response.Path("test.name").Data().(string)
		require.Equal(name, "session1")
		randomString := response.Path("test.random_string").Data().(string)
		require.Equal(randomString, "ses11")

		// Get by session 2
		method = "test.info"
		response, err = client.SendRequest(session2, method, map[string]interface{}{"test_id": session2ID})
		require.NoError(err, "%s failed as %s", method, session2)
		id = response.Path("test.id").Data().(float64)
		require.Equal(id, session2ID)
		name = response.Path("test.name").Data().(string)
		require.Equal(name, "session2")
		randomString = response.Path("test.random_string").Data().(string)
		require.Equal(randomString, "ses22")

		// Broadcast by session 1
		method = "test.broadcast"
		response, err = client.SendRequest(session1, method, map[string]interface{}{"test_id": session1ID})
		require.NoError(err, "%s failed as %s", method, session1)
		id = response.Path("test.id").Data().(float64)
		require.Equal(id, session1ID)

		// Broadcast by session 2
		method = "test.broadcast"
		response, err = client.SendRequest(session2, method, map[string]interface{}{"test_id": session2ID})
		require.NoError(err, "%s failed as %s", method, session2)
		id = response.Path("test.id").Data().(float64)
		require.Equal(id, session2ID)

	})

	t.Cleanup(func() {
		client.Disconnect("session1")
		// require.NoError(err)
		client.Disconnect("session2")
		// require.NoError(err)
	})
}
