package db

import (
	"math/rand"
	"time"
)

type SkipNode struct {
	member  string
	score   string
	forward []*SkipNode
}

type SkipList struct {
	head   *SkipNode
	levels int
}

const MaxLevel = 16

func NewSkipList() *SkipList {
	head := &SkipNode{
		forward: make([]*SkipNode, MaxLevel),
	}
	return &SkipList{head: head, levels: 1}
}

func (s *SkipList) ZAdd(score string, member string) {
	update := make([]*SkipNode, MaxLevel)
	x := s.head

	// Find the position to insert the new node
	for i := s.levels - 1; i >= 0; i-- {
		for x.forward[i] != nil && compareScores(x.forward[i].score, score) < 0 {
			x = x.forward[i]
		}
		update[i] = x
	}

	level := randomLevel()
	if level > s.levels {
		for i := s.levels; i < level; i++ {
			update[i] = s.head
		}
		s.levels = level
	}

	newNode := &SkipNode{
		member:  member,
		score:   score,
		forward: make([]*SkipNode, level),
	}

	for i := 0; i < level; i++ {
		newNode.forward[i] = update[i].forward[i]
		update[i].forward[i] = newNode
	}
}

func (s *SkipList) ZRange() []string {
	result := []string{}
	x := s.head.forward[0]
	for x != nil {
		result = append(result, x.member)
		x = x.forward[0]
	}
	return result
}

func (s *SkipList) ZScore(member string) (string, bool) {
	x := s.head.forward[0]
	for x != nil {
		if x.member == member {
			return x.score, true
		}
		x = x.forward[0]
	}
	return "", false
}

func (s *SkipList) ZRank(member string) (int, bool) {
	rank := 0
	x := s.head.forward[0]
	for x != nil {
		if x.member == member {
			return rank, true
		}
		x = x.forward[0]
		rank++
	}
	return 0, false
}

func compareScores(score1, score2 string) int {
	// Compare scores as strings
	if score1 < score2 {
		return -1
	}
	if score1 > score2 {
		return 1
	}
	return 0
}

func randomLevel() int {
	rand.NewSource(time.Now().UnixNano())
	level := 1
	for rand.Float64() < 0.5 && level < MaxLevel {
		level++
	}
	return level
}
