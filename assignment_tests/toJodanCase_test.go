package assignmenttests_test

import (
	"testing"

	"github.com/mintcolorfuls/go-codewar/assignments"
)

func TestToJodanCase(t *testing.T) {
	t.Run("most trees are blue", func(t *testing.T) {
		want := "Most Trees Are Blue"

		got := assignments.ToJodanCase(want)

		if got != want {
			t.Errorf("want: %s, but got: %s", want, got)
		}
	})

	t.Run("All the rules in this world were made by someone no smarter than you. So make your own.", func(t *testing.T) {
		want := "All The Rules In This World Were Made By Someone No Smarter Than You. So Make Your Own."

		got := assignments.ToJodanCase(want)

		if want != got {
			t.Errorf("want: %s, but got: %s", want, got)
		}
	})
}
