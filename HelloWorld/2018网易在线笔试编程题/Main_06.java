/**
 * ��Ŀ����
ţţ����˯��ͷ�����������˺ܶ����ӣ�ֻ�����������ʱ�����Ż��ѹ������Ҿ������𴲡�
��������������ҪX���ӵ�����ң��Ͽ�ʱ��Ϊ�����AʱB�֣��������������ʲôʱ���𴲡�

��������:
ÿ���������һ������������ ÿ�����������ĵ�һ�а���һ������������ʾ���ӵ�����N(N<=100)��
��������N��ÿ�а���������������ʾ������������ʱ��ΪHi(0<=Hi<24)ʱMi(0<=Mi<60)�֡�
 ��������һ�а���һ����������ʾ������������ҪX(0<=X<=100)���ӵ�����ҡ�
 ��������һ�а���������������ʾ�Ͽ�ʱ��ΪA(0<=A<24)ʱB(0<=B<60)�֡�
  ���ݱ�֤������һ�����ӿ�����ţţ��ʱ������ҡ�

�������:
�������������ʾţţ������ʱ�䡣

��������1:
3 
5 0 
6 0 
7 0 
59 
6 59
�������1:
6 0

����˼·

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
			int N = sc.nextInt(); //��������
			int[][] aclock = new int[N][2];//����
			for(int i = 0;i < N;i++){
				aclock[i][0] = sc.nextInt();
				aclock[i][1] = sc.nextInt();
			}
			
			int X = sc.nextInt(); //��ʰʱ��
			//�Ͽ�ʱ��
			int[] courseT = new int[2];
			courseT[0] = sc.nextInt(); 
			courseT[1] = sc.nextInt();
			
			//����������ʱ��
			int[] last = getUp(X,courseT);
			
			//ѡ����������
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
