/**
 * 
 */
package cn.lsp.wangyi;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;
import java.util.Scanner;

/**
 * @author LSP
 * 题目描述
小Q正在给一条长度为n的道路设计路灯安置方案。 为了让问题更简单,小Q把道路视为n个方格,需要照亮的地方用'.'表示, 
不需要照亮的障碍物格子用'X'表示。 小Q现在要在道路上设置一些路灯, 对于安置在pos位置的路灯, 
这盏路灯可以照亮pos - 1, pos, pos + 1这三个位置。 小Q希望能安置尽量少的路灯照亮所有'.'区域, 
希望你能帮他计算一下最少需要多少盏路灯。

输入描述:
输入的第一行包含一个正整数t(1 <= t <= 1000), 表示测试用例数 接下来每两行一个测试数据, 
第一行一个正整数n(1 <= n <= 1000),表示道路的长度。 
第二行一个字符串s表示道路的构造,只包含'.'和'X'。

输出描述:
对于每个测试用例, 输出一个正整数表示最少需要多少盏路灯。

输入例子1:
2
3
.X.
11
...XX....XX
输出例子1:

1
3
解题思路
因为路灯只能照亮三个位置，所以将道路方格分成三个一组进行讨论，共有8种情况
如果是上面四种情况，第一个位置是'.'，则路灯数加1，下面四种情况第一个位置是'X'，则将下标向右移一位重新进行分组考虑
...  ..X  .X.  .XX
XXX  XX.  X.X  X..
 *
 */
public class Main_04 {

	/**
	 * @param args
	 */
	public static void main(String[] args) {
		// TODO Auto-generated method stub
		Scanner sc = new Scanner(System.in);
		while(sc.hasNextLine()){
			int t = Integer.parseInt(sc.nextLine()); //用例个数
			
			String[] cases = new String[t];
			for(int i = 0;i < t;i++){
				int n = Integer.parseInt(sc.nextLine());  //道路长度
				String road = sc.nextLine();
				cases[i] = road;
			}

			
			for(int i = 0;i < t;i++){
				int sum = 0;
				int index = 0;
				String s = cases[i];
				while(index < s.length()){
					if(s.charAt(index)=='X'){
						index++;
					}else{
						index += 3;
						sum ++;
					}
				}
				
				System.out.println(sum);
				
			}
			
			
			
		}
	}

}
