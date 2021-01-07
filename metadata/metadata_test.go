package metadata

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"

	"testing"

	"github.com/stretchr/testify/assert"
)

func ServerMock() (baseURL string, mux *http.ServeMux, teardownFn func()) {
	mux = http.NewServeMux()
	srv := httptest.NewServer(mux)
	return srv.URL, mux, srv.Close
}

func Test_Deserialization(t *testing.T) {
	baseURL, mux, teardown := ServerMock()
	defer teardown()

	mux.HandleFunc("/metadata", func(w http.ResponseWriter, req *http.Request) {
		b, err := ioutil.ReadFile("testdata/" + "device.json")
		if err != nil {
			t.Fatal(err)
		}
		if _, err := w.Write(b); err != nil {
			t.Fatal(err)
		}
	})

	device, err := GetMetadataFromURL(baseURL)
	assert.Nil(t, err)
	assert.NotNil(t, device)

	assert.Equal(t, "a583c581-73a2-4101-b9f7-55c02ecf99bf", device.ID)
	assert.Equal(t, "ewr1-t1.small.x86-01", device.Hostname)

	volumes := device.Volumes
	assert.Equal(t, 2, len(volumes))
	assert.Equal(t, "volume-df0579d9", volumes[0].Name)
	assert.Equal(t, "iqn.2013-05.com.daterainc:tc:01:sn:e0651c43a0f7dd90", volumes[0].IQN)
	assert.Equal(t, 2, len(volumes[0].IPs))
	assert.Equal(t, "10.144.32.122", volumes[0].IPs[0].String())
	assert.Equal(t, "10.144.48.122", volumes[0].IPs[1].String())
	assert.Equal(t, 100, volumes[0].Capacity.Size)
	assert.Equal(t, "gb", volumes[0].Capacity.Unit)

}
