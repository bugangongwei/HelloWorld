/**
 * ��Ŀ����
ţţȥ�Ġ���ʦ�Ҳ��Σ����ŵ�ʱ�����򱱷���������������·�ˡ���Ȼ��������һ�ŵ�ͼ����������Ҫ֪���Լ������ĸ�����
����������

��������:
ÿ���������һ������������ ÿ�����������ĵ�һ�а���һ������������ʾת����Ĵ���N(N<=1000)��
 ��������һ�а���һ������ΪN���ַ�������L��R��ɣ�L��ʾ����ת��R��ʾ����ת��

�������:
���ţţ�������ķ���N��ʾ����S��ʾ�ϣ�E��ʾ����W��ʾ����

��������1:
3
LRR
�������1:
E

����˼·
�������ϱ�������һ����״�Ĵ洢��ʽ��L-->��ʱ�룬R-->˳ʱ��
��״�洢��λԪ��������ȥ�ҾͿ��ԣ�
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
