/*
 * 题目描述
牛牛以前在老师那里得到了一个正整数数对（x,y），牛牛忘记他们具体是多少了。但是牛牛记得老师告诉他x和y均不大于n,并且x除以y的余数大于等于k。
牛牛希望你能帮他计算一共有多少个可能的数对。

输入描述:
输入包括两个正整数n,k(1<=n<=10^5,0<=k<=n-1)

输出描述：
对于每个测试用例，输出一个正整数表示可能的数对数量。

例子：
输入：
5 2
输出
7
说明
满足条件的数对有：（2,3）（2,4）（2,5）（3,4）（3,5）（4,5）（5,3）

解题思路
余数循环出现
叠加每组循环余数中余数大于k的元素的个数,以及剩下的一次循环余数不足情况下余数大于等于k的元素的个数
举例
n=10 k=2
i= 1  %  3 =  1
i= 2  %  3 =  2
i= 3  %  3 =  0
i= 4  %  3 =  1
i= 5  %  3 =  2
i= 6  %  3 =  0
i= 7  %  3 =  1
i= 8  %  3 =  2
i= 9  %  3 =  0
i=10  %  3 =  1
 * 
 * */


package cn.lsp.wangyi;

import java.util.Scanner;

public class Main_01 {

	/**
	 * @param args
	 */
	public static void main(String[] args) {
		// TODO Auto-generated method stub
		 Scanner in = new Scanner(System.in);
		 int sum = 0;
	        while (in.hasNext()) {
	            int n = in.nextInt();
	            int k = in.nextInt();
	            
	            
	            if(n==1){
	            	if(k==0){
	            		sum = 1; //n==1,k==0的时候，只有(1,1)这种情况
	            	}else{
	            		sum = 0;//
	            	}
	            }
	            
	            if(k==0){
	            	sum = n * n;
	            }
	            	
	            	
	            //在除数y<=k时，余数总小于k，所以不需要考虑y<=k的情况
	            
	            for(int y=k+1;y<=n;y++){  
	            	sum += (n/y) * (y-k);
	            	if(n%y >=k){
	            		sum += (n%y)-(k-1);
	            	}
	            }
	           System.out.println(sum);
	        }
	}

}
