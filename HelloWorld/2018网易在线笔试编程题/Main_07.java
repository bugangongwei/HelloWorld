/*
 * ��Ŀ����
ţţ׼���μ�ѧУ��֯�Ĵ���, ����ǰţţ׼����������װ��һЩ��ʳ, ţţ�ı�������Ϊw��
 ţţ����һ����n����ʳ, ��i����ʳ���Ϊv[i]�� ţţ��֪������������������������������,
 ��һ���ж�������ʳ�ŷ�(�����Ϊ0Ҳ��һ�ַŷ�)��

��������:
����������� ��һ��Ϊ����������n��w(1 <= n <= 30, 1 <= w <= 2 * 10^9),��ʾ��ʳ�������ͱ����������� �ڶ���n��������v[i](0 <= v[i] <= 10^9),��ʾÿ����ʳ�������

�������:
���һ��������, ��ʾţţһ���ж�������ʳ�ŷ���

��������1:
3 10
1 2 4
�������1:
8
����˵��1:
������ʳ�����С��10,����ÿ����ʳ�з���Ͳ��������������һ����222 = 8�������
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
			int n = sc.nextInt();//��Ʒ����
			int w = sc.nextInt();//��������
			
			int[] v = new int[n];
			for(int i=0;i<n;i++){
				v[i] = sc.nextInt();
			}//��Ʒ���
			
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
