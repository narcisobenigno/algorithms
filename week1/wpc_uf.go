package week1

type S int

type SuccessorDelete struct {
	sites []S
	sizes []int
}

func (uf *SuccessorDelete) rootOf(site S) S {
	for site != uf.sites[site] {
		uf.sites[site] = uf.sites[uf.sites[site]]
		site = uf.sites[site]
	}
	return site
}

func (uf *SuccessorDelete) Union(p S, q S) {
	rp := uf.rootOf(p)
	rq := uf.rootOf(q)

	if rp == rq {
		return
	}

	if uf.sizes[rp] > uf.sizes[rq] {
		uf.sites[rq] = rp
		uf.sizes[rp] = uf.sizes[rq]
	} else {
		uf.sites[rp] = rq
		uf.sizes[rq] += uf.sizes[rp]
	}
}

func (uf *SuccessorDelete) Find(site S) S {
	return uf.largest[uf.rootOf(site)]
}

func SuccessorDeleteNew(m int) *SuccessorDelete {
	items := make([]S, m+1)
	sizes := make([]int, m+1)
	for i := range items {
		site := S(i)
		items[i] = site
	}
	for i := range sizes {
		sizes[i] = 1
	}
	uf := &SuccessorDelete{items, sizes}
	for _, i := range items {
		uf.Union(i, S(m))
	}
	return uf
}
