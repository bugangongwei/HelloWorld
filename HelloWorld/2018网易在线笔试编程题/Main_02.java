/**
 * 题目描述
为了找到自己满意的工作，牛牛收集了每种工作的难度和报酬。牛牛选工作的标准是在难度不超过自身能力值的情况下，牛牛选择报酬最高的工作。
在牛牛选定了自己的工作后，牛牛的小伙伴们来找牛牛帮忙选工作，牛牛依然使用自己的标准来帮助小伙伴们。
牛牛的小伙伴太多了，于是他只好把这个任务交给了你。

输入描述:
每个输入包含一个测试用例。 每个测试用例的第一行包含两个正整数，分别表示工作的数量N(N<=100000)和小伙伴的数量M(M<=100000)。
 接下来的N行每行包含两个正整数，分别表示该项工作的难度Di(Di<=1000000000)和报酬Pi(Pi<=1000000000)。
  接下来的一行包含M个正整数，分别表示M个小伙伴的能力值Ai(Ai<=1000000000)。 保证不存在两项工作的报酬相同。

输出描述:
对于每个小伙伴，在单独的一行输出一个正整数表示他能得到的最高报酬。一个工作可以被多个人选择。

输入例子1:
3 3
1 100
10 1000
1000000000 1001
9 10 1000000000

输出例子1:
100
1000
1001

解题思路
维护一个（N+M）的dp[N+M]的数组，记录不同能力和不同难度下的最大薪酬
复杂度 MAX（O(NlogN),O(MlogM),O(N+M))
 */
package cn.lsp.wangyi;

import java.util.Scanner;

/**
 * @author LSP
 *
 */
public class Main_02 {

	/**
	 * @param args
	 */
	public static void main(String[] args) {
		// TODO Auto-generated method stub
		Scanner sc = new Scanner(System.in);
		
		while(sc.hasNext()){
			int N = sc.nextInt();//工作数量
			int M = sc.nextInt();//小伙伴数量
			
			int[][] mission = new int[N][2];//
			for(int i = 0;i < N;i++){
				for(int j = 0;j < 2;j++){
					mission[i][j] = sc.nextInt();
				}
			}
			
			int[] worker = new int[M]; //worker的工作能力
			for(int i = 0;i < M;i++){
				worker[i] = sc.nextInt();
			}
			
			/*用worker的能力，去匹配能胜任的工作中的薪酬最高的工作*/
			for(int i = 0;i < M;i++){
				int max = 0; //当前worker可以胜任的工作的最大薪酬
				//int capacity = worker[i];
				for(int j = 0;j < N;j++){
					//int cap = mission[0][j];
					if(mission[j][0] <= worker[i]){
						if(mission[j][1] > max){
							max = mission[j][1];
						}
					}
				}
				
				System.out.println(max);
				
			}
			
		}
		
		
	}

}
