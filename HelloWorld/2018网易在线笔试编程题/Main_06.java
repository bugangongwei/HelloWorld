/**
 * 题目描述
牛牛总是睡过头，所以他定了很多闹钟，只有在闹钟响的时候他才会醒过来并且决定起不起床。
从他起床算起他需要X分钟到达教室，上课时间为当天的A时B分，请问他最晚可以什么时间起床。

输入描述:
每个输入包含一个测试用例。 每个测试用例的第一行包含一个正整数，表示闹钟的数量N(N<=100)。
接下来的N行每行包含两个整数，表示这个闹钟响起的时间为Hi(0<=Hi<24)时Mi(0<=Mi<60)分。
 接下来的一行包含一个整数，表示从起床算起他需要X(0<=X<=100)分钟到达教室。
 接下来的一行包含两个整数，表示上课时间为A(0<=A<24)时B(0<=B<60)分。
  数据保证至少有一个闹钟可以让牛牛及时到达教室。

输出描述:
输出两个整数表示牛牛最晚起床时间。

输入例子1:
3 
5 0 
6 0 
7 0 
59 
6 59
输出例子1:
6 0

解题思路

 */
package cn.lsp.wangyi;

import java.util.Scanner;

/**
 * @author LSP
 *
 */
public class Main_06 {

	/**
	 * @param args
	 */
	public static void main(String[] args) {
		// TODO Auto-generated method stub
		Scanner sc = new Scanner(System.in);
		while(sc.hasNext()){
			int N = sc.nextInt(); //闹钟数量
			int[][] aclock = new int[N][2];//闹钟
			for(int i = 0;i < N;i++){
				aclock[i][0] = sc.nextInt();
				aclock[i][1] = sc.nextInt();
			}
			
			int X = sc.nextInt(); //收拾时间
			//上课时间
			int[] courseT = new int[2];
			courseT[0] = sc.nextInt(); 
			courseT[1] = sc.nextInt();
			
			//计算最晚起床时间
			int[] last = getUp(X,courseT);
			
			//选择最优闹钟
			for(int i = N-1;i >= 0;i--){
				if(last[0] == aclock[i][0]){
					if(last[1] >= aclock[i][1]){
						System.out.print(aclock[i][0]);
						System.out.print(" ");
						System.out.println(aclock[i][1]);
						break;
					}
				}
				
				if(last[0] > aclock[i][0]){
					System.out.print(aclock[i][0]);
					System.out.print(" ");
					System.out.println(aclock[i][1]);
					break;
				}
				
			}
		}
		
		
		
	}
	
	private static int[] getUp(int X,int[] arr){
		int[] r = new int[2];
		if(X<=arr[1]){
			r[0] = arr[0];
			r[1] = arr[1] - X;
		}else{
			r[0] = arr[0] - 1;
			r[1] = arr[1] + 60 - X;
		}
		return r;
	}

}
