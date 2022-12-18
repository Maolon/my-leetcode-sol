/*
 * @lc app=leetcode id=743 lang=typescript
 *
 * [743] Network Delay Time
 */

// @lc code=start

type node = {
  to: number;
  weight: number;
};

function networkDelayTime(times: number[][], n: number, k: number): number {
  const graph = new Map<number, node[]>();
  //构造图
  for (let edge of times) {
    // from
    const from = edge[0];
    const to = edge[1];
    const weight = edge[2];
    if (graph.has(from)) {
      graph.set(edge[0], [
        {
          to,
          weight,
        },
      ]);
    } else {
      const curNode = graph.get(from);
      if (curNode) {
        curNode.push({
          to,
          weight,
        });
        graph.set(from, curNode);
      }
    }
  }
}

const dijkstra = (start: number, graph: Map<number, node[]>) => {};
// @lc code=end
