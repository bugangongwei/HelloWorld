/*
 * ��Ŀ����
СQ�õ�һ�����������: 1, 12, 123,...12345678910,1234567891011...�� ����СQ�����ܷ�3����������ʺܸ���Ȥ��
 СQ����ϣ�����ܰ�������һ�´����еĵ�l������r��(�����˵�)�ж��ٸ������Ա�3������

��������:
���������������l��r(1 <= l <= r <= 1e9), ��ʾҪ�����������ˡ�

�������:
���һ������, ��ʾ�������ܱ�3���������ָ�����

��������1:
2 5
�������1:
3
����˵��1:
1, 12, 123, 1234, 12345, 123456, 1234567, 12345678, 123456789, 12345678910
(0,1,1,0,1,1,0,1,1,0)

����˼·
ÿ�����ּ������ܱ�3����
����ѭ��������Ϊ2��0��������
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
