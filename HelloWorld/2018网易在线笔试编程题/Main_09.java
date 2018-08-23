/*
 * 题目描述
小易觉得高数课太无聊了，决定睡觉，不过他对课上的一些内容挺感兴趣的，所以希望你在老师讲到有趣的部分的时候叫醒他一下。 你知道了小易对一堂课每分钟知识点的感兴趣程度，并以分数量化，以及他在这堂课上每分钟是否会睡着，你可以叫醒他一次，这会使他在接下来的k分钟内保持清醒。 你需要选择一种方案最大化小易这堂课听的知识点分值。

输入描述:
第一行n,k(1<=n,k<=10^5)，表示这堂课持续多少分钟，以及叫醒小易一次使他能够保持清醒的时间。
 第二行n个数，a1,a2,...,an(1<=ai<=10^4)表示小易对每分钟知识点的感兴趣评分。
  第三行n个数，t1,t2,...,tn表示每分钟小易是否清醒，1表示清醒。

输出描述:
小易这堂课听到的知识点的最大兴趣值。

输入例子1:
6 3
1 3 5 2 5 4
1 1 0 1 0 0
输出例子1:
16
 * */

package cn.lsp.wangyi;

import java.util.Scanner;

public class Main_09 {

	/**
	 * @param args
	 */
	public static void main(String[] args) {
		// TODO Auto-generated method stub
		Scanner sc = new Scanner(System.in);
		while(sc.hasNext()){
			int n = sc.nextInt();//表示这堂课持续多少分钟
			int k = sc.nextInt();//叫醒小易一次使他能够保持清醒的时间
			
			int[] a = new int[n];
			int[] t = new int[n];
			
			//我要不要用lambda呢？
			//java8：Colletion集合类增加了Lambda Expressions, Streams, and Aggregate Operations
			//数组不支持
			//如果硬将数组转成list，然后用lambda来操作，一方面多线程方式能增加CPU利用率，一方面大量线程上下文切换反而增加消耗
			//所以，在这个规模的应用中，我选择不用lambda
			
			for(int i = 0;i < n;i++){
				a[i] = sc.nextInt();
			}
			
			for(int i = 0;i < n;i++){
				t[i] = sc.nextInt();
			}
			
			int sum_wake = 0;  //清醒时间的得分和
			int max_addition = 0;  //最大增量
			
			for(int i = 0;i < n;i++){
				int addition = 0;//增量
				if(t[i] == 1){
					sum_wake += a[i];
				}else{
					
					for(int j = i;j < Math.min(i + k,n);j++){
						if(t[j] == 0){
							addition += a[j];
						}
					}
				}
				
				max_addition = Math.max(max_addition, addition);
			}
			
			System.out.println(max_addition + sum_wake);
			
		}
		
	}

}
