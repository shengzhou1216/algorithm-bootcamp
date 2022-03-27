//911. 在线选举
package week5

type TopVotedCandidate struct {
    tops []int // 与times对应的票数最高的人
    times []int
}


func Constructor(persons []int, times []int) TopVotedCandidate {
    h := map[int]int{}
    top_vote := 0 // 最高票数
    top_p := 0 // 最高票数的人
    tops := []int{}
    for _,p := range persons {
        h[p]++
        if h[p] >= top_vote {
            top_vote = h[p]
            top_p = p
        }
        tops = append(tops,top_p)
    }
    return TopVotedCandidate{
        tops: tops,
        times: times,
    }
}


func (this *TopVotedCandidate) Q(t int) int {
    // 1. 求t对应的times中的索引
    left := 0
    right := len(this.times) - 1
    for left < right {
        mid := (left + right + 1) / 2
        if this.times[mid] <= t {
            left = mid
        }else {
            right = mid - 1
        }
    }
    // 2. 求票数最高的人
    return this.tops[right]
}


/**
 * Your TopVotedCandidate object will be instantiated and called as such:
 * obj := Constructor(persons, times);
 * param_1 := obj.Q(t);
 */