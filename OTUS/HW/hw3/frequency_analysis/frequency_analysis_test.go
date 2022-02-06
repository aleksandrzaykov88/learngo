package frequencyanalysis

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var text1 = "я тебя люблю люблю"
var res1 = []string{"люблю", "тебя", "я"}

func TestTopTenWords(t *testing.T) {
	require.Equal(t, res1, TopTenWords(text1))
}
