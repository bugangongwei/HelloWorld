/**
 *题目描述
又到了丰收的季节，恰好小易去牛牛的果园里游玩。 牛牛常说他多整个果园的每个地方都了如指掌，小易不太相信，所以他想考考牛牛。
 在果园里有N堆苹果，每堆苹果的数量为ai,小易希望知道从左往右数第x个苹果是属于哪一堆的。 
 牛牛觉得问题太简单了，所以希望你来替他回答。

输入描述:
第一行一个数n(1<=n<=10^5) 第二行n个数ai(1<=ai<=1000),表示从左往右数第i堆有多少苹果 
第三行一个数m(1<=m<=10^5),表示有m次询问 第四行m个数qi,表示小易希望知道第qi个苹果属于哪一堆。

输出描述:
m行，第i行输出第qi个苹果属于哪一堆。

输入例子1:
5
2 7 3 4 9
3
1 25 11
输出例子1:
1
5
3 
 */
package cn.lsp.wangyi;

import java.util.Scanner;

/**
 * @author LSP
 *
 */
public class Main_10 {

	/**
	 * @param args
	 */
	public static void main(String[] args) {
		// TODO Auto-generated method stub
		Scanner sc = new Scanner(System.in);
		int n = sc.nextInt();
		
		
		int[] a = new int[n];
		for (int i = 0;i < n;i++){
			a[i] = sc.nextInt();
		}
		
		int m = sc.nextInt();
		int[] q = new int[m];
		for (int i = 0;i < m;i++){
			q[i] = sc.nextInt();
		}
		
		
		for(int i = 0;i < m;i++){
			int sum = 0;
			for(int j = 0;j < n;j++){
				sum += a[j];
				if(q[i] <= sum && q[i] >= sum - a[j]){
					System.out.println(j+1);
				}
			}
		}
		
	}

}
