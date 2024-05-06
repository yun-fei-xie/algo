package _7_design

import (
	"fmt"
	"sort"
	"testing"
)

// https://leetcode.cn/problems/design-twitter/

var startSeq = 0

type tweet struct {
	id       int
	sequence int
}

type Twitter struct {
	newsMap    map[int][]*tweet
	followsMap map[int]map[int]struct{}
}

func Constructor() Twitter {
	return Twitter{
		newsMap:    make(map[int][]*tweet),
		followsMap: make(map[int]map[int]struct{}),
	}
}

func nextSequence() int {
	startSeq++
	return startSeq
}

func (this *Twitter) PostTweet(userId int, tweetId int) {
	if tweets, exists := this.newsMap[userId]; exists {
		tweets = append(tweets, &tweet{
			id:       tweetId,
			sequence: nextSequence(),
		})
		this.newsMap[userId] = tweets
	} else {
		tweets = make([]*tweet, 0)
		tweets = append(tweets, &tweet{
			id:       tweetId,
			sequence: nextSequence(),
		})
		this.newsMap[userId] = tweets
	}
}

// 这一步是瓶颈 ,应该使用合并多个有序链表的思路
func (this *Twitter) GetNewsFeed(userId int) []int {
	users := make([]int, 0)
	users = append(users, userId)
	followUsers := this.followsMap[userId]
	for id, _ := range followUsers {
		users = append(users, id)
	}
	alltweets := make([]*tweet, 0)

	for _, u := range users {
		tweets := this.newsMap[u]
		alltweets = append(alltweets, tweets...)
	}

	sort.Slice(alltweets, func(i, j int) bool {
		if alltweets[i].sequence > alltweets[j].sequence {
			return true
		}
		return false
	})

	res := make([]int, 0)

	for i := 0; i < len(alltweets) && i < 10; i++ {
		res = append(res, alltweets[i].id)
	}

	return res
}

func (this *Twitter) Follow(followerId int, followeeId int) {
	if follows, exists := this.followsMap[followerId]; exists {
		follows[followeeId] = struct{}{}
		this.followsMap[followerId] = follows
	} else {
		follows = make(map[int]struct{})
		follows[followeeId] = struct{}{}
		this.followsMap[followerId] = follows
	}
}

func (this *Twitter) Unfollow(followerId int, followeeId int) {

	if follows, exists := this.followsMap[followerId]; exists {
		for id, _ := range follows {
			if id == followeeId {
				delete(follows, id)
				this.followsMap[followerId] = follows
			}
		}
	}
}

/**
 * Your Twitter object will be instantiated and called as such:
 * obj := Constructor();
 * obj.PostTweet(userId,tweetId);
 * param_2 := obj.GetNewsFeed(userId);
 * obj.Follow(followerId,followeeId);
 * obj.Unfollow(followerId,followeeId);
 */

func TestTwitter(t *testing.T) {
	obj := Constructor()
	obj.PostTweet(1, 5)
	getNewFeeds := obj.GetNewsFeed(1)
	fmt.Println(getNewFeeds)

	obj.Follow(1, 2)
	obj.PostTweet(2, 6)
	getNewFeeds = obj.GetNewsFeed(1)
	fmt.Println(getNewFeeds)

	obj.Unfollow(1, 2)
	getNewFeeds = obj.GetNewsFeed(1)
	fmt.Println(getNewFeeds)
}

func TestTwitter2(t *testing.T) {

	obj := Constructor()
	obj.PostTweet(1, 5)
	obj.PostTweet(1, 3)
	feeds := obj.GetNewsFeed(1)
	fmt.Println(feeds)

}
