/*
 * 题目描述
小Q得到一个神奇的数列: 1, 12, 123,...12345678910,1234567891011...。 并且小Q对于能否被3整除这个性质很感兴趣。
 小Q现在希望你能帮他计算一下从数列的第l个到第r个(包含端点)有多少个数可以被3整除。

输入描述:
输入包括两个整数l和r(1 <= l <= r <= 1e9), 表示要求解的区间两端。

输出描述:
输出一个整数, 表示区间内能被3整除的数字个数。

输入例子1:
2 5
输出例子1:
3
例子说明1:
1, 12, 123, 1234, 12345, 123456, 1234567, 12345678, 123456789, 12345678910
(0,1,1,0,1,1,0,1,1,0)

解题思路
每个数字加起来能被3整除
周期循环，余数为2或0都能整除
 * */

package cn.lsp.wangyi;

import java.util.Scanner;

public class Main_03 {

	/**
	 * @param args
	 */
	public static void main(String[] args) {
		// TODO Auto-generated method stub
		Scanner sc = new Scanner(System.in);
		int result = 0;
		while(sc.hasNext()){
			int l = sc.nextInt();
			int r = sc.nextInt();
			
			for(int i=l;i<=r;i++){
				if(i%3==0||i%3==2){
					result++;
				}
			}
			System.out.println(result);
		}
		
	}

}
