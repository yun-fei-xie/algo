package _88_double_test

/*
2424. 最长上传前缀
https://leetcode.cn/problems/longest-uploaded-prefix/description/

给你一个 n 个视频的上传序列，每个视频编号为 1 到 n 之间的 不同 数字，你需要依次将这些视频上传到服务器。请你实现一个数据结构，在上传的过程中计算 最长上传前缀 。
如果 闭区间 1 到 i 之间的视频全部都已经被上传到服务器，那么我们称 i 是上传前缀。最长上传前缀指的是符合定义的 i 中的 最大值 。
请你实现 LUPrefix 类：
LUPrefix(int n) 初始化一个 n 个视频的流对象。
void upload(int video) 上传 video 到服务器。
int longest() 返回上述定义的 最长上传前缀 的长度。

方法：用一个变量x追踪最长前缀
这个变量只有在查询的时候会进行更新
*/
type LUPrefix struct {
	video map[int]bool
	x     int
}

func Constructor(n int) LUPrefix {
	return LUPrefix{
		video: map[int]bool{},
		x:     1,
	}
}

func (this *LUPrefix) Upload(video int) {
	this.video[video] = true
}

func (this *LUPrefix) Longest() int {
	for this.video[this.x] == true {
		this.x++
	}
	return this.x - 1
}

/**
 * Your LUPrefix object will be instantiated and called as such:
 * obj := Constructor(n);
 * obj.Upload(video);
 * param_2 := obj.Longest();
 */
