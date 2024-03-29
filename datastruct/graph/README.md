# 图相关数据结构和算法

## 拓扑排序

1. 建图的时候，构建各个顶点的入度集合（用一个数组记录每个顶点的入度）。
2. 将入度为0的顶点都加入队列中。
3. 当队列不为空时，从队列中取出一个顶点，访问该顶点所有的出边。将出边的to顶点的入度--。 如果to顶点的入度为0，那么将该顶点加入到队列中。
4. 不断重复上述过程，直到队列为空。
5. 如果最终还有顶点入度不为0，则图中包含环。

## 二分图问题

二分图判定算法：染色法、

二分图相关算法：匈牙利算法、最大匹配、完美匹配
https://www.renfei.org/blog/bipartite-matching.html

## 最短路径问题

### 迪杰斯特拉算法

Dijkstra（/ˈdikstrɑ/或/ˈdɛikstrɑ/）算法由荷兰计算机科学家 E. W. Dijkstra 于 1956 年发现，1959 年公开发表。
是一种求解非负权图上单源最短路径的算法。

单源：起点是固定的

局限性：权不能为负数

### Bellman-Ford算法

## 桥

对于无向图来说，如果删除了某条边，图的连通分量的个数发生了变化。
那么这条边就是图中的一个桥。（从字面意思上理解比较形象）
桥是图中最脆弱的关系。一个图中可以有多个桥。

极端情况下，如果桥是一颗树，那么这棵树的所有的边都是桥。

应用：交通系统、社交网络

### 桥的识别算法

使用深度优先遍历可以解决。
桥和环有着紧密的联系。
假设无向边A->B，判断A->B这条边是不是桥，那么看，从B出发，不走回头路(A->B这条边)，能不能再回到A。

1. 如果能回到，那么A->B不是桥。这是因为，如果能够回到。那么顶点A和顶点B之间形成了一个环。
   删掉A->B这条边，图的连通性和没有删掉之前相比，并没有发生改变。
2. 如果不能回到，那么A->B是桥。这是因为，删掉A->B这条边，B无法再访问到A。也就是之前还能访问，现在无法访问。
   图的连通性发生了改变。

如何定义回到和回不到。对于回到来说，如果B能够回到A之前的点，那么也相当于间接回到A。
例如，在深度优先遍历中，如果遍历的顺序是D->E->C->A->B
如果知道了B可以回到D，那么B肯定可以顺着D回到A。更加通俗一点，如果B可以回到A的任意一个祖先节点（D、E、C）
那么，B就可以回到A。

所以，在深度优先遍历中，需要使用两个数组。
数组ord[i]表示顶点i被遍历到的顺序，也叫遍历序号。
数组low[i]表示顶点i不通过父亲节点（不走回头路）可以访问到的遍历序号最小的顶点。

ord数组和low数组的初始化和更新过程是怎样的？
用一个全局计数器id记录每一个顶点被访问的顺序。
ord数组只有初始没有更新。当dfs(i)第一次访问一个没有被访问的顶点时，顶点i被访问。
ord[i]= id++
同时，visited[i] = true

当一个顶点第一次被访问时，low[i]= ord[i]。因为low[i]表示顶点i不同过父亲节点，可以访问到的遍历序号最小的顶点。
初始时，这个顶点还没有向外访问，所以它能够访问到的遍历序号最小的顶点就是它自己。
当顶点i访问到了一个已经被访问过的相邻的顶点时（visited[j] == true），可以更新一下low[i]。
low[i] = min(low[i] , low[j])

更新low[i]这一步更加细节一些，
如果顶点i的邻接点j没有被访问过，就递归进去访问。 递归回来更新一下low[i]。
如果顶点i的邻接点j已经被访问过，就直接更新一下low[i]。

如果发现low[j]是大于ord[i]的，那么边i->j就是一座桥。
这是因为，顶点j不走回头路（j->i），怎么访问，都无法回到顶点i。

## 割点

## 哈密尔顿回路

## 哈密尔顿路径


