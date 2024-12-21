package serror

import (
	"errors"
	"fmt"
	"log/slog"
	"path/filepath"

	"runtime"
	"strings"
)

type SError struct {
	err error
	at  string
}

const (
	prefix          = "(("
	sufix           = "))"
	separator       = "+"
	sourceSeparator = ":"
	prefixSize      = len(prefix)
)

func (err *SError) Error() string {
	return fmt.Sprintf("((%s+%s))", err.err, err.at)
}

func New(s string) *SError {
	return &SError{
		err: errors.New(s),
		at:  caller(2),
	}
}

func WrapSkip(err error, skip int) *SError {
	skip += 2
	if skip < 2 {
		skip = 2
	}
	return &SError{
		err: err,
		at:  caller(skip),
	}
}

func Wrap(err error) *SError {
	return &SError{
		err: err,
		at:  caller(2),
	}
}

func caller(skip int) string {
	pc, file, no, ok := runtime.Caller(skip)
	if ok {
		b := filepath.Base(file)
		f := filepath.Base(runtime.FuncForPC(pc).Name())
		return fmt.Sprintf("%s:%d:%s", b, no, f)
	}
	return ""
}

func DecodeMessage(s string) (msg string, attrs []slog.Attr) {
	if s == "" {
		return "", []slog.Attr{}
	}

	serrorFrom := strings.Index(s, prefix)
	serrorTo := strings.Index(s, sufix)

	if serrorFrom == serrorTo {
		return s, []slog.Attr{}
	}

	serrorMsg := s[serrorFrom+2 : serrorTo]

	elem := strings.Split(serrorMsg, separator)
	if 2 != len(elem) {
		return s, []slog.Attr{}
	}

	sources := strings.Split(elem[1], sourceSeparator)
	if 3 != len(sources) {
		return s, []slog.Attr{}
	}

	s = strings.Replace(s, s[serrorFrom:serrorTo+2], elem[0], 1)

	return s, []slog.Attr{
		slog.String("file", sources[0]),
		slog.String("line", sources[1]),
		slog.String("func", sources[2]),
	}
}
