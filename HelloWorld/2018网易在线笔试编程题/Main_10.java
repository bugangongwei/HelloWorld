/**
 *��Ŀ����
�ֵ��˷��յļ��ڣ�ǡ��С��ȥţţ�Ĺ�԰�����档 ţţ��˵����������԰��ÿ���ط�������ָ�ƣ�С�ײ�̫���ţ��������뿼��ţţ��
 �ڹ�԰����N��ƻ����ÿ��ƻ��������Ϊai,С��ϣ��֪��������������x��ƻ����������һ�ѵġ� 
 ţţ��������̫���ˣ�����ϣ�����������ش�

��������:
��һ��һ����n(1<=n<=10^5) �ڶ���n����ai(1<=ai<=1000),��ʾ������������i���ж���ƻ�� 
������һ����m(1<=m<=10^5),��ʾ��m��ѯ�� ������m����qi,��ʾС��ϣ��֪����qi��ƻ��������һ�ѡ�

�������:
m�У���i�������qi��ƻ��������һ�ѡ�

��������1:
5
2 7 3 4 9
3
1 25 11
�������1:
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
