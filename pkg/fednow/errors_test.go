package fednow_test

import (
	"fmt"
	"testing"

	"github.com/moov-io/fednow20022/pkg/fednow"

	"github.com/stretchr/testify/require"
)

func TestErrorCodes(t *testing.T) {
	ec := fednow.IsError("9910")
	require.NotNil(t, ec)
	require.Equal(t, "9910", ec.Code)
	require.Equal(t, fednow.ErrTemporary, ec.Level)

	ec = fednow.IsError("be07")
	require.NotNil(t, ec)
	require.Equal(t, "BE07", ec.Code)
	require.Equal(t, fednow.ErrFatal, ec.Level)

	ec = fednow.IsError("ac03")
	require.NotNil(t, ec)
	require.Equal(t, "AC03", ec.Code)
	require.Equal(t, fednow.ErrAccountFatal, ec.Level)

	ec = fednow.IsError("AM13")
	require.NotNil(t, ec)
	require.Equal(t, "AM13", ec.Code)
	require.Equal(t, fednow.ErrTemporary, ec.Level)

	ec = fednow.IsError("AM09")
	require.NotNil(t, ec)
	require.Equal(t, "AM09", ec.Code)
	require.Equal(t, fednow.ErrLogic, ec.Level)

	ec = fednow.IsError("")
	require.Nil(t, ec)

	ec = fednow.IsError("9999")
	require.Nil(t, ec)
}

func TestErrorLevel(t *testing.T) {
	// just make sure .Error() doesn't recusively panic
	level := fednow.ErrNetwork
	require.Equal(t, "Network issue", level.Error())
	require.Equal(t, "Network issue", fmt.Sprintf("%v", level))
	require.Equal(t, "Network issue", fmt.Errorf("%w", level).Error())
}
