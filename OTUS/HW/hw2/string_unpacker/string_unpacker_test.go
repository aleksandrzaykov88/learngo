package stringunpacker

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUnpackString(t *testing.T) {
	require.Equal(t, "aaaabccddddde", UnpackString("a4bc2d5e"))
	require.Equal(t, "abcd", UnpackString("abcd"))
	require.Equal(t, "", UnpackString("45"))
	require.Equal(t, `qwe45`, UnpackString(`qwe\4\5`))
	require.Equal(t, `qwe44444`, UnpackString(`qwe\45`))
	require.Equal(t, `qwe\\\\\`, UnpackString(`qwe\\5`))
	require.Equal(t, "", UnpackString(""))
}
