/*
 * ��Ŀ����
С�׾��ø�����̫�����ˣ�����˯�����������Կ��ϵ�һЩ����ͦ����Ȥ�ģ�����ϣ��������ʦ������Ȥ�Ĳ��ֵ�ʱ�������һ�¡� ��֪����С�׶�һ�ÿ�ÿ����֪ʶ��ĸ���Ȥ�̶ȣ����Է����������Լ��������ÿ���ÿ�����Ƿ��˯�ţ�����Խ�����һ�Σ����ʹ���ڽ�������k�����ڱ������ѡ� ����Ҫѡ��һ�ַ������С�����ÿ�����֪ʶ���ֵ��

��������:
��һ��n,k(1<=n,k<=10^5)����ʾ���ÿγ������ٷ��ӣ��Լ�����С��һ��ʹ���ܹ��������ѵ�ʱ�䡣
 �ڶ���n������a1,a2,...,an(1<=ai<=10^4)��ʾС�׶�ÿ����֪ʶ��ĸ���Ȥ���֡�
  ������n������t1,t2,...,tn��ʾÿ����С���Ƿ����ѣ�1��ʾ���ѡ�

�������:
С�����ÿ�������֪ʶ��������Ȥֵ��

��������1:
6 3
1 3 5 2 5 4
1 1 0 1 0 0
�������1:
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
			int n = sc.nextInt();//��ʾ���ÿγ������ٷ���
			int k = sc.nextInt();//����С��һ��ʹ���ܹ��������ѵ�ʱ��
			
			int[] a = new int[n];
			int[] t = new int[n];
			
			//��Ҫ��Ҫ��lambda�أ�
			//java8��Colletion������������Lambda Expressions, Streams, and Aggregate Operations
			//���鲻֧��
			//���Ӳ������ת��list��Ȼ����lambda��������һ������̷߳�ʽ������CPU�����ʣ�һ��������߳��������л�������������
			//���ԣ��������ģ��Ӧ���У���ѡ����lambda
			
			for(int i = 0;i < n;i++){
				a[i] = sc.nextInt();
			}
			
			for(int i = 0;i < n;i++){
				t[i] = sc.nextInt();
			}
			
			int sum_wake = 0;  //����ʱ��ĵ÷ֺ�
			int max_addition = 0;  //�������
			
			for(int i = 0;i < n;i++){
				int addition = 0;//����
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
