// Copyright © 2022-2023 Obol Labs Inc. Licensed under the terms of a Business Source License 1.1

package errors_test

import (
	"io"
	"reflect"
	"testing"

	"github.com/omni-network/omni/lib/errors"

	"github.com/stretchr/testify/require"
)

var errSentinel = errors.NewSentinel("test sentinel")

func TestComparable(t *testing.T) {
	t.Parallel()
	require.False(t, reflect.TypeOf(errors.New("x")).Comparable())
}

func TestIs(t *testing.T) {
	t.Parallel()
	errX := errors.New("x")

	err1 := errors.New("1", "1", "1")
	err11 := errors.Wrap(err1, "w1")
	err111 := errors.Wrap(err11, "w2")
	errWrapSent := errors.Wrap(errSentinel, "w1")

	require.Equal(t, "x", errX.Error())
	require.Equal(t, "1", err1.Error())
	require.Equal(t, "w1: 1", err11.Error())
	require.Equal(t, "w2: w1: 1", err111.Error())
	require.Equal(t, "w1: test sentinel", errWrapSent.Error())

	require.True(t, errors.Is(err1, err1))
	require.True(t, errors.Is(err11, err1))
	require.True(t, errors.Is(err111, err1))
	require.False(t, errors.Is(err1, err11))
	require.True(t, errors.Is(err11, err11))
	require.True(t, errors.Is(err111, err11))
	require.False(t, errors.Is(err1, err111))
	require.False(t, errors.Is(err11, err111))
	require.True(t, errors.Is(err111, err11))
	require.True(t, errors.Is(errWrapSent, errSentinel))

	errIO1 := errors.Wrap(io.EOF, "w1")
	errIO11 := errors.Wrap(errIO1, "w2")

	require.Equal(t, "w1: EOF", errIO1.Error())
	require.Equal(t, "w2: w1: EOF", errIO11.Error())

	require.True(t, errors.Is(io.EOF, io.EOF))
	require.True(t, errors.Is(errIO1, io.EOF))
	require.True(t, errors.Is(errIO11, io.EOF))
	require.False(t, errors.Is(io.EOF, errIO1))
	require.True(t, errors.Is(errIO1, errIO1))
	require.True(t, errors.Is(errIO11, errIO1))
	require.False(t, errors.Is(io.EOF, errIO11))
	require.False(t, errors.Is(errIO1, errIO11))
	require.True(t, errors.Is(errIO11, errIO11))
	require.False(t, errors.Is(err111, errX))
}

func TestFormat(t *testing.T) {
	t.Parallel()

	err1 := errors.New("1", "1", "1")
	err11 := errors.Wrap(err1, "w1", "11", "11")
	err111 := errors.Wrap(err11, "w2", "111", "this is a long wrapped string")
	errWrapSent := errors.Wrap(errSentinel, "w1")

	require.Equal(t, "<nil>", errors.Format(nil))
	require.Equal(t, "1 [1=1]", errors.Format(err1))
	require.Equal(t, "w1: 1 [11=11, 1=1]", errors.Format(err11))
	require.Equal(t, "w2: w1: 1 [111=this is a long wrapped string, 11=11, 1=1]", errors.Format(err111))
	require.Equal(t, "w1: test sentinel", errors.Format(errWrapSent))
}
