package grab

import (
	"fmt"
	"testing"
)

type Instagram struct {
	followers map[int]map[int]struct{}
	followees map[int]map[int]struct{}

	feeds      map[int][]int // userId -> []PhotoId
	feedsCache map[int][]int
}

func NewInstagram() *Instagram {
	ins := &Instagram{
		followers:  make(map[int]map[int]struct{}),
		followees:  make(map[int]map[int]struct{}),
		feeds:      make(map[int][]int),
		feedsCache: make(map[int][]int),
	}
	return ins
}

func (ins *Instagram) PostPhoto(userId int, photoId int) {
	ins.feeds[userId] = append(ins.feeds[userId], photoId)
	ins.feedsCache[userId] = append(ins.feedsCache[userId], photoId)
	followees := ins.followees[userId]
	for followee := range followees {
		ins.feedsCache[followee] = append(ins.feedsCache[followee], photoId)
		size := len(ins.feedsCache[followee])
		if size > 10 {
			ins.feedsCache[followee] = ins.feedsCache[followee][size-10:]
		}
	}
}

// List<Integer> getFeed(int userId)
func (ins *Instagram) GetFeed(userId int) []int {
	return ins.feedsCache[userId]
}

func (ins *Instagram) Follow(followerId, followeeId int) {
	if followeeId == followerId {
		return
	}
	_, ok := ins.followers[followerId]
	if !ok {
		ins.followers[followerId] = map[int]struct{}{}
	}
	ins.followers[followerId][followeeId] = struct{}{}

	_, ok = ins.followees[followeeId]
	if !ok {
		ins.followees[followeeId] = map[int]struct{}{}
	}
	ins.followees[followeeId][followerId] = struct{}{}
}

func (ins *Instagram) Unfollow(followerId, followeeId int) {
	delete(ins.followers[followerId], followeeId)
	delete(ins.followees[followeeId], followerId)

	followeeFeedIds := ins.feeds[followeeId]
	set := make(map[int]struct{})
	for _, id := range followeeFeedIds {
		set[id] = struct{}{}
	}

	// remove feedCache
	newFeedCache := make([]int, 0)
	for _, feedId := range ins.feedsCache[followerId] {
		_, ok := set[feedId]
		if !ok {
			newFeedCache = append(newFeedCache, feedId)
		}
	}
	ins.feedsCache[followerId] = newFeedCache
}

func Test_Instagram(t *testing.T) {
	insta := NewInstagram()

	insta.PostPhoto(1, 101)
	fmt.Println(insta.GetFeed(1))

	// Instagram insta = new Instagram();

	// insta.postPhoto(1, 101); // User 1 posts photo 101
	// insta.getFeed(1);         // Returns [101]
	// 1 -> 2
	insta.Follow(1, 2)

	fmt.Println("----", insta.followers, insta.followees)

	insta.PostPhoto(2, 102)

	fmt.Println("----", insta.feeds, insta.feedsCache)

	fmt.Println(insta.GetFeed(1))

	// insta.follow(1, 2);       // User 1 follows User 2
	// insta.postPhoto(2, 102);  // User 2 posts photo 102
	// insta.getFeed(1);         // Returns [102, 101]

	insta.Unfollow(1, 2)
	fmt.Println(insta.GetFeed(1))

	// insta.unfollow(1, 2);     // User 1 unfollows User 2
	// insta.getFeed(1);         // Returns [101]

}
