/**
 * ��Ŀ����
Ϊ���ҵ��Լ�����Ĺ�����ţţ�ռ���ÿ�ֹ������ѶȺͱ��ꡣţţѡ�����ı�׼�����ѶȲ�������������ֵ������£�ţţѡ�񱨳���ߵĹ�����
��ţţѡ�����Լ��Ĺ�����ţţ��С���������ţţ��æѡ������ţţ��Ȼʹ���Լ��ı�׼������С����ǡ�
ţţ��С���̫���ˣ�������ֻ�ð�������񽻸����㡣

��������:
ÿ���������һ������������ ÿ�����������ĵ�һ�а����������������ֱ��ʾ����������N(N<=100000)��С��������M(M<=100000)��
 ��������N��ÿ�а����������������ֱ��ʾ��������Ѷ�Di(Di<=1000000000)�ͱ���Pi(Pi<=1000000000)��
  ��������һ�а���M�����������ֱ��ʾM��С��������ֵAi(Ai<=1000000000)�� ��֤������������ı�����ͬ��

�������:
����ÿ��С��飬�ڵ�����һ�����һ����������ʾ���ܵõ�����߱��ꡣһ���������Ա������ѡ��

��������1:
3 3
1 100
10 1000
1000000000 1001
9 10 1000000000

�������1:
100
1000
1001

����˼·
ά��һ����N+M����dp[N+M]�����飬��¼��ͬ�����Ͳ�ͬ�Ѷ��µ����н��
���Ӷ� MAX��O(NlogN),O(MlogM),O(N+M))
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
			int N = sc.nextInt();//��������
			int M = sc.nextInt();//С�������
			
			int[][] mission = new int[N][2];//
			for(int i = 0;i < N;i++){
				for(int j = 0;j < 2;j++){
					mission[i][j] = sc.nextInt();
				}
			}
			
			int[] worker = new int[M]; //worker�Ĺ�������
			for(int i = 0;i < M;i++){
				worker[i] = sc.nextInt();
			}
			
			/*��worker��������ȥƥ����ʤ�εĹ����е�н����ߵĹ���*/
			for(int i = 0;i < M;i++){
				int max = 0; //��ǰworker����ʤ�εĹ��������н��
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
