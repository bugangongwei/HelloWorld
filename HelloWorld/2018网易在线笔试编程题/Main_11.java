/**
 * 题目描述
给你n个a,m个z组成所有可能的字符串，并将字符串按字典序从小到大排列，输出第k个字符串。

输入描述:
第一行为三个数，分别为a的个数n,z的个数m,第k个字符串。

输出描述:
第k个字符串

输入例子1:
2 2 6
输出例子1:
zzaa
解题思路
刚开始深度优先存储所有字符串，然后取出第k个字符串，通过30%；
通过计算不同等级的组合数，看k在哪一种情况下，比较考验数学分析。
 */
package cn.lsp.wangyi;

import java.util.ArrayList;
import java.util.Scanner;

/**
 * @author LSP
 *
 */
public class Main_11 {

	/**
	 * @param args
	 */
	public static void main(String[] args) {
		// TODO Auto-generated method stub
		
		Scanner in = new Scanner(System.in);
        int n = in.nextInt();
        int m = in.nextInt();
        int k = in.nextInt();
        ArrayList<String> map = new ArrayList<String>();
        DFS(n,m,map,"");
        if(k>map.size()) System.out.println(-1);
        else System.out.println(map.get(k-1));
		
		
	}

	 private static void DFS(int n,int m,ArrayList<String> map,String s){
	        if(n==0&&m==0){
	            map.add(s);
	            return;
	        }
	        if(n>0){
	            DFS(n-1,m,map,s+"a");
	        }
	        if(m>0){
	            DFS(n,m-1,map,s+"z");
	        }
	    }
	
}
