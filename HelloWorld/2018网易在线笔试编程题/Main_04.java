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
 * ��Ŀ����
СQ���ڸ�һ������Ϊn�ĵ�·���·�ư��÷����� Ϊ�����������,СQ�ѵ�·��Ϊn������,��Ҫ�����ĵط���'.'��ʾ, 
����Ҫ�������ϰ��������'X'��ʾ�� СQ����Ҫ�ڵ�·������һЩ·��, ���ڰ�����posλ�õ�·��, 
��յ·�ƿ�������pos - 1, pos, pos + 1������λ�á� СQϣ���ܰ��þ����ٵ�·����������'.'����, 
ϣ�����ܰ�������һ��������Ҫ����յ·�ơ�

��������:
����ĵ�һ�а���һ��������t(1 <= t <= 1000), ��ʾ���������� ������ÿ����һ����������, 
��һ��һ��������n(1 <= n <= 1000),��ʾ��·�ĳ��ȡ� 
�ڶ���һ���ַ���s��ʾ��·�Ĺ���,ֻ����'.'��'X'��

�������:
����ÿ����������, ���һ����������ʾ������Ҫ����յ·�ơ�

��������1:
2
3
.X.
11
...XX....XX
�������1:

1
3
����˼·
��Ϊ·��ֻ����������λ�ã����Խ���·����ֳ�����һ��������ۣ�����8�����
��������������������һ��λ����'.'����·������1���������������һ��λ����'X'�����±�������һλ���½��з��鿼��
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
			int t = Integer.parseInt(sc.nextLine()); //��������
			
			String[] cases = new String[t];
			for(int i = 0;i < t;i++){
				int n = Integer.parseInt(sc.nextLine());  //��·����
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
