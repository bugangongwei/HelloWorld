/**
 * 题目描述
牛牛去犇犇老师家补课，出门的时候面向北方，但是现在他迷路了。虽然他手里有一张地图，但是他需要知道自己面向哪个方向，
请你帮帮他。

输入描述:
每个输入包含一个测试用例。 每个测试用例的第一行包含一个正整数，表示转方向的次数N(N<=1000)。
 接下来的一行包含一个长度为N的字符串，由L和R组成，L表示向左转，R表示向右转。

输出描述:
输出牛牛最后面向的方向，N表示北，S表示南，E表示东，W表示西。

输入例子1:
3
LRR
输出例子1:
E

解题思路
‘东西南北’看出一个环状的存储形式，L-->逆时针，R-->顺时针
环状存储定位元素用余数去找就可以；
 */
package cn.lsp.wangyi;

import java.util.LinkedList;
import java.util.Scanner;

/**
 * @author LSP
 *
 */
public class Main_05 {

	/**
	 * @param args
	 */
	public static void main(String[] args) {
		// TODO Auto-generated method stub
//		DoubleLinkedList dlist = new DoubleLinkedList();
//		dlist.addFirst('N');
//		dlist.add('E');
//		dlist.add('S');
//		dlist.addLast('W');
		
		LinkedList<Character> list = new LinkedList<Character>();
		list.addFirst('N');
		list.add('E');
		list.add('S');
		list.addLast('W');
		
		Scanner sc = new Scanner(System.in);
		while(sc.hasNextLine()){
			int N = Integer.parseInt(sc.nextLine());
			String str = sc.nextLine();
			
			int index = 0;
			int i = 0;
			while(i<str.length()){
				if(str.charAt(i)=='L'){
					index = (index - 1 + 4)%4;
				}else{
					index = (index + 1 + 4)%4;
				}
				i++;
			}
			System.out.println(list.get(index));
		}
	}

}
