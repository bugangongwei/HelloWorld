/**
 * ��Ŀ����
����n��a,m��z������п��ܵ��ַ����������ַ������ֵ����С�������У������k���ַ�����

��������:
��һ��Ϊ���������ֱ�Ϊa�ĸ���n,z�ĸ���m,��k���ַ�����

�������:
��k���ַ���

��������1:
2 2 6
�������1:
zzaa
����˼·
�տ�ʼ������ȴ洢�����ַ�����Ȼ��ȡ����k���ַ�����ͨ��30%��
ͨ�����㲻ͬ�ȼ������������k����һ������£��ȽϿ�����ѧ������
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
