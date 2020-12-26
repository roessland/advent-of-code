package main

type Shortest struct {
	dist map[string]map[string]int
	absent int
}

func NewShortest(absent int) Shortest {
	as := Shortest{
		dist: make(map[string]map[string]int),
		absent: absent,
	}
	return as
}

func (as Shortest) GetDist(uid, vid string) int {
	if as.dist[uid] == nil {
		as.dist[uid] = make(map[string]int)
	}
	dist, ok := as.dist[uid][vid]
	if !ok {
		return as.absent
	}
	return dist
}

func (as Shortest) SetDist(uid, vid string, dist int) {
	if as.dist[uid] == nil {
		as.dist[uid] = make(map[string]int)
	}
	as.dist[uid][vid] = dist
}