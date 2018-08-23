/*
 * 题目描述
牛牛准备参加学校组织的春游, 出发前牛牛准备往背包里装入一些零食, 牛牛的背包容量为w。
 牛牛家里一共有n袋零食, 第i袋零食体积为v[i]。 牛牛想知道在总体积不超过背包容量的情况下,
 他一共有多少种零食放法(总体积为0也算一种放法)。

输入描述:
输入包括两行 第一行为两个正整数n和w(1 <= n <= 30, 1 <= w <= 2 * 10^9),表示零食的数量和背包的容量。 第二行n个正整数v[i](0 <= v[i] <= 10^9),表示每袋零食的体积。

输出描述:
输出一个正整数, 表示牛牛一共有多少种零食放法。

输入例子1:
3 10
1 2 4
输出例子1:
8
例子说明1:
三种零食总体积小于10,于是每种零食有放入和不放入两种情况，一共有222 = 8种情况。
 * */
package cn.lsp.wangyi;

import java.util.Scanner;

public class Main_07 {

	/**
	 * @param args
	 */
	public static void main(String[] args) {
		// TODO Auto-generated method stub
		Scanner sc= new Scanner(System.in);
		while(sc.hasNext()){
			int n = sc.nextInt();//物品数量
			int w = sc.nextInt();//背包容量
			
			int[] v = new int[n];
			for(int i=0;i<n;i++){
				v[i] = sc.nextInt();
			}//物品体积
			
			int sum = 0;
			int num = 0;
			for(int i = 0;i < n;i++){
				sum += v[i];
				if(sum >= w){
					break;
				}
				num = i + 1;
			}
			
			System.out.println(Math.pow(2, num));
			
			
		}
	}

}
