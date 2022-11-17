package example

import (
	"bytes"
	"github.com/olegshishkin/go-logger"
	glz "github.com/olegshishkin/go-logger-zerolog"
	"github.com/rs/zerolog"
	"strings"
	"testing"
)

func TestCommon(t *testing.T) {
	t.Run("Logger", func(t *testing.T) {
		buf := &bytes.Buffer{}
		w, err := glz.NewLogWriterBuilder().
			WithConsole(zerolog.ConsoleWriter{
				Out:           buf,
				PartsExclude:  []string{zerolog.TimestampFieldName, zerolog.LevelFieldName},
				FieldsExclude: []string{glz.DefaultLogSourceKey},
				NoColor:       true,
			}).
			Build()
		if err != nil {
			panic(err)
		}
		log := glz.From(Common(Base(w, logger.Info)))
		testVar := struct {
			a, b, c string
		}{
			a: "line a",
			b: "line b",
			c: "line c",
		}

		log.Info("test %+v", testVar)

		if got, want := strings.TrimSpace(buf.String()), "common_test.go:35 > test {a:line a b:line b c:line c}"; got != want {
			t.Errorf("\ngot:\n%s\nwant:\n%s", got, want)
		}
	})
}
