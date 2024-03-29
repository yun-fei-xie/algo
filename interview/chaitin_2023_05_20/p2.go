package main

/*
Description

C 公司安全团队为了提升企业网络安全防护能力，会经常组织网络安全攻防演练，在一次演练中，Alice 作为攻方选手，成功获取了某台内网主机的权限。不过防守方也提前做了很多准备，为网络中的各台主机配置了复杂的拓扑关系，但是这些信息已经被机智的 Alice 通过物理渗透获取到了。Alice 发现，她攻击的内网中有 N 台主机，其中只有部分机器能够互通，并且攻击方需要窃取的目标信息被分别放在了其中 K 台主机上。为了尽可能避免被防守方发现，Alice 需要让自己攻击过程中在主机之间跳转的次数尽可能少，请你帮她分析最优的攻击路径至少需要多少次主机之间的跳转。

PS: Alice 作为一名经验丰富的白帽子，只要两台主机互相连通，她便可以从其中一台跳转到另一台。

PPS:需要注意的是，由于特殊的网络配置，从主机 a 跳转到主机 b，再从 b 跳转到 a，这是两次跳转行为。不要理解为从 a 机器 ssh 到 b 机器，然后从 b 机器上Ctrl ^D回来到 a 只算一次。


Input
第 1 行输入 1 个整数
�
T (
1
≤
�
≤
10
1≤T≤10)，表示数据组数。

对于每组输入:

第 1 行输入 2 个整数
�
N
�
K，表示有
�
N 台主机，
�
K 台上有目标信息 (
1
≤
�
≤
500
，
1
≤
�
≤
10
，
1
≤
�
≤
�
1≤N≤500，1≤K≤10，1≤K≤N，主机编号从 1 到 N)，Alice 当前在 1 号主机上；
第 2 行输入空格分隔的
�
K 个整数，表示有目标信息的主机编号；
接下来一个整数
�
M，表示有演练环境拓扑中有
�
M 条边 (
0
≤
�
≤
3000
0≤M≤3000)；
接下里输入
�
M 行，每行两个整数
a
b，表示
a 和b 之间有一条边 (双向边)。

Output
对于每组输入，输出一个整数，表示要获取全部目标信息，最少需要多少次主机之间的跳转，如果无法获取全部目标信息，则输出
−
1
−

*/
/*
回退也算跳转次数
*/
func main2() {

}
