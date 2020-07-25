package week1

import "testing"

type Union [2]S

func TestSuccessor(t *testing.T) {
	cases := []struct {
		n      int
		unions []Union
		findBy S
		want   S
	}{
		{
			5,
			[]Union{
				Union{S(0), S(1)},
				Union{S(1), S(2)},
				Union{S(2), S(3)},
				Union{S(3), S(2)},
				Union{S(0), S(2)},
			},
			S(2),
			S(3),
		},
	}

	for _, c := range cases {
		uf := SuccessorDeleteNew(c.n)

		for _, u := range c.unions {
			uf.Union(u[0], u[1])
		}

		got := uf.Find(c.findBy)

		if got != c.want {
			t.Errorf(
				"Find(%d) == %d, want: %d",
				c.findBy,
				got,
				c.want,
			)
		}
	}
}
