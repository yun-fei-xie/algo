package main

/*



节点权值
时间限制： 3000MS
内存限制： 589824KB
题目描述：
现在有若干节点。每个节点上有能量塔。所有节点构成一棵树。

某个节点u可以为和u距离不超过给定值的节点各提供一点能量。

此处距离的定义为两个节点之间经过的边的数量。特别的，节点u到本身的距离为零。



现在给出每个节点上的能量塔可以为多远的距离内的点提供能量。

小美想要探究每个节点上的能量值具体是多少。你的任务是帮助小美计算得到，并依次输出。



输入描述
第一行一个整数N，表示节点的数量。

接下来一行N个以空格分开的整数，依次表示节点1，节点2，…，节点N的能量塔所能提供能量的最远距离。

接下来N-1行，每行两个整数，表示两个点之间有一条边。


1≤N≤500，节点上能量塔所能到达的最远距离距离不会大于 500.

输出描述
一行N个整数，依次表示节点1，节点2，…，节点N上的能量值。


样例输入
3
1 1 1
1 2
2 3
样例输出
2 3 2

提示
样例解释1

节点1上的能量塔可以为节点1和2提供能量。

节点2上的能量塔可以为节点1、2和3提供能量。

节点3上的能量塔可以为节点2和3提供能量。



输入样例2

10

1 2 1 1 1 2 3 1 1 1

1 2

2 3

1 4

2 5

4 6

3 7

5 8

1 9

2 10

输出样例2

6 6 3 4 4 2 3 3 3 3
*/

func main() {

}
