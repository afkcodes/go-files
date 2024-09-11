package sayhello

import "testing"

func TestSayHello(t *testing.T) {
	subtests := []struct {
		items  []string
		result string
	}{
		{
			result: "Hello, world!",
		},
		{
			items:  []string{"afkcodes"},
			result: "Hello, afkcodes!",
		},
	}

	for _, st := range subtests {
		if s := Say(st.items); s != st.result {
			t.Errorf("Wanted %s (%v), got %s", st.result, st.items, s)
		}
	}

	// want := "Hello, test!"
	// got := Say([]string{"test"})

	// if want != got {
	// 	t.Errorf("Wanted %s got %s", want, got)
	// }
}
