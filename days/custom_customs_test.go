package days

import "testing"

func TestCustomsAnswers(t *testing.T) {
	groups := ParseCustomsGroups([]byte("abc\n\na\nb\nc\n\nab\nac\n\na\na\na\na\n\nb\n"))
	t.Run("CombinedAnswers", func(t *testing.T) {
		expects := []int{3, 3, 3, 1, 1}

		for i, expect := range expects {
			got := groups[i].CombinedAnswers()
			if got != expect {
				t.Errorf("{%v}.CombinedAnswers() = %v; want %v", groups[i], got, expect)
			}
		}
	})

	t.Run("CommonAnswers", func(t *testing.T) {
		expects := []int{3, 0, 1, 1, 1}

		for i, expect := range expects {
			got := groups[i].CommonAnswers()
			if got != expect {
				t.Errorf("{%v}.CommonAnswers() = %v; want %v", groups[i], got, expect)
			}
		}
	})
}
