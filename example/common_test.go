package example_test

import (
	"bytes"
	"github.com/olegshishkin/go-logger"
	glz "github.com/olegshishkin/go-logger-zerolog"
	"github.com/olegshishkin/go-logger-zerolog/example"
	"github.com/rs/zerolog"
	"strings"
	"testing"
)

func TestCommon(t *testing.T) {
	t.Parallel()
	t.Run("Logger", func(t *testing.T) {
		t.Parallel()
		buf := &bytes.Buffer{}
		writer, err := glz.NewLogWriterBuilder().
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
		log := glz.From(example.Common(example.Base(writer, logger.Info)))
		testVar := struct {
			a, b, c string
		}{
			a: "line a",
			b: "line b",
			c: "line c",
		}

		log.Info("test %+v", testVar)

		want := "common_test.go:38 > test {a:line a b:line b c:line c}"
		got := strings.TrimSpace(buf.String())
		if got != want {
			t.Errorf("\ngot:\n%s\nwant:\n%s", got, want)
		}
	})
}
