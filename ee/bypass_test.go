package ee

import (
	"database/sql"
	"fmt"
	"io"
	"testing"
)

type ComplexErrStuff struct {
	byPass *ByPass
}

func (s *ComplexErrStuff) Next() ComplexErrStuff2 {
	if s.byPass.Try(fmt.Errorf("a weirdo error right at the beginning: %w", io.EOF)) {
		panic("unreachable")
	}
	return ComplexErrStuff2{s.byPass}
}

type ComplexErrStuff2 struct {
	byPass *ByPass
}

func (s ComplexErrStuff2) Next2() error {
	if s.byPass.Try(fmt.Errorf("another follow-up error: %w", sql.ErrNoRows)) {
		panic("unreachable")
	}

	if s.byPass.Try(sql.ErrTxDone) {
		panic("unreachable")
	}
	return s.byPass.Error()
}

func TestByPass(t *testing.T) {
	stuff := &ComplexErrStuff{byPass: &ByPass{}}
	err := stuff.Next().Next2()
	if err == nil {
		t.Fatal("must not be nil")
	}
	t.Log(err)
}
