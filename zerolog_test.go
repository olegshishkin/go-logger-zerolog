package zerolog

import (
	"bytes"
	"github.com/olegshishkin/go-logger"
	"github.com/rs/zerolog"
	"strings"
	"testing"
)

func TestLogger(t *testing.T) {
	t.Run("Logger", func(t *testing.T) {
		buf := &bytes.Buffer{}
		w, err := NewLogWriterBuilder().
			WithConsoleWriter(zerolog.ConsoleWriter{
				Out:           buf,
				PartsExclude:  []string{"time", "level"},
				FieldsExclude: []string{"logger"},
				NoColor:       true,
			}).
			Build()
		if err != nil {
			panic(err)
		}
		log, err := Logger(w, logger.Info)
		if err != nil {
			panic(err)
		}
		testVar := struct {
			a, b, c string
		}{
			a: "line a",
			b: "line b",
			c: "line c",
		}

		log.Info("test %+v", testVar)

		if got, want := strings.TrimSpace(buf.String()), "zerolog_test.go:37 > test {a:line a b:line b c:line c}"; got != want {
			t.Errorf("\ngot:\n%s\nwant:\n%s", got, want)
		}
	})
}
