package week1

import "testing"

type Pair [2]Site

func TestTotalConnected(t *testing.T) {
	cases := []struct {
		n          int
		unions     []Pair
		wantSizeOf []int
	}{
		{
			5,
			[]Pair{
				Pair{Site(0), Site(1)},
				Pair{Site(1), Site(2)},
				Pair{Site(2), Site(3)},
				Pair{Site(3), Site(2)},
				Pair{Site(0), Site(2)},
			},
			[]int{4, 4, 4, 4, 1},
		},
	}

	for _, c := range cases {
		uf := NewWpcUf(c.n)

		for _, u := range c.unions {
			uf.Union(u[0], u[1])
		}

		for i, _ := range make([]int, c.n) {
			want := c.wantSizeOf[i]
			got := uf.TotalConnected(Site(i))

			if got != want {
				t.Errorf(
					"TotalConnected(%d) == %d, want: %d",
					i,
					got,
					want,
				)
			}
		}
	}
}

func TestFind(t *testing.T) {
	cases := []struct {
		n      int
		unions []Pair
		findBy Site
		want   Site
	}{
		{
			5,
			[]Pair{
				Pair{Site(0), Site(1)},
				Pair{Site(1), Site(2)},
				Pair{Site(2), Site(3)},
				Pair{Site(3), Site(2)},
				Pair{Site(0), Site(2)},
			},
			Site(2),
			Site(3),
		},
	}

	for _, c := range cases {
		uf := NewWpcUf(c.n)

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
