package stripprefix

import (
	"net/http/httptest"
	"testing"

	"github.com/PuerkitoBio/turtles"
	"github.com/stretchr/testify/assert"
)

func TestStripPrefix(t *testing.T) {
	cases := []struct {
		prefix string
		path   string
		want   string
		code   int
	}{
		{"", "/", "/", 200},
		{"/api", "/api/x", "/x", 200},
		{"/api", "/blah/x", "/blah/x", 404},
	}

	for i, c := range cases {
		sp := &StripPrefix{c.prefix}
		h := turtles.Wrap(turtles.StatusHandler(200), sp)
		w := httptest.NewRecorder()
		r := httptest.NewRequest("", c.path, nil)

		h.ServeHTTP(w, r)
		assert.Equal(t, c.code, w.Code, "%d: status", i)
		assert.Equal(t, c.want, r.URL.Path, "%d: path", i)
	}
}
