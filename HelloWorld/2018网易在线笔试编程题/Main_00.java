/**
 ��Ŀ������
ƽ������n�����Σ����½�����(x1[i],y1[i])�����Ͻ�����(x2[i],y2[i])���ж��ص����������Ŀ�Ƕ��١�
������������������й�����������Ϊ�������ص��ģ������Ǳ߽�ͽ��䡣������ƽ�����ص��������ĵط��ľ�����Ŀ��
 
 ��������:
����������У���һ�б�ʾ������Ŀn���ڶ���x1-->���½Ǻ����꣬������y1-->���½������꣬������x2-->���ϽǺ����꣬������y2-->���Ͻ������ꡣ

�������:
�������ص��ľ�����Ŀ����������ص����1.

����˼·��
�������һ�����������ķ�ʽ�������Ե�һ��Ϊ��׼���Ժ���ľ��ν����������жϣ�������룬���֦����������
����ص�������Ҫ�������Ƿ�������ص������ڵģ������ߵݹ��������ֵ���ɡ�
ͬʱ����Ҫ��ÿ����Ϊ��׼����������һ�Σ�����̰��˼�뱣�����ֵ��

for i 0�� n:
	��׼���� x11,y11,x22,y22=x1[i],y1[i],x2[i],y2[i]
		for j 0�� n:  //�������׼�ص�����������
			if x1[j],y1[j],x2[j],y2[j]���׼������������������
				continue��
			else:  //�ص�
				��1�����ص�������Ϊ�»�׼��ȥʣ�µľ�����Ѱ����󷽰���
				��2���ھɻ�׼�������£���ʣ�µľ���������Ѱ���ص��򣬷�����һ����󷽰���
				���أ�1����2�����ֵ��


���룺
7
0 2 2 3 3 6 6
0 0 0 0 0 0 0
1 3 4 5 5 7 8
1 1 2 1 3 1 2
ÿһ�������⣺
1,2,3,3,3,2,2
���⣺
3
 *
 */
package cn.lsp.wangyi;

import java.util.Scanner;

/**
 * @author LSP
 *
 */
public class Main_00 {

	/**
	 * @param args
	 */
	
	public static int solve(int[] x1, int[] x2, int[] y1, int[] y2, int k,
            int xa, int ya, int xb, int yb) {
        if (k == x1.length)
            return 0;
        else {
            if (x1[k] >= xb || y1[k] >= yb || xa >= x2[k] || ya >= y2[k])// �ų�����������
                return solve(x1, x2, y1, y2, k + 1, xa, ya, xb, yb);
            else {
                int xa1 = Math.max(xa, x1[k]);
                int ya1 = Math.max(ya, y1[k]);
                int xb1 = Math.min(xb, x2[k]);
                int yb1 = Math.min(yb, y2[k]);// �����ǰ���μ������ڣ�������������
                return Math.max(solve(x1, x2, y1, y2, k + 1, xa, ya, xb, yb),
                        solve(x1, x2, y1, y2, k + 1, xa1, ya1, xb1, yb1) + 1);// ��ǰ�����㲻��
            }
        }
    }
	
	
	public static void main(String[] args) {
		// TODO Auto-generated method stub
		Scanner sc = new Scanner(System.in);
        while (sc.hasNext()) {
            int n = sc.nextInt();
            int[] x1 = new int[n];
            int[] y1 = new int[n];
            int[] x2 = new int[n];
            int[] y2 = new int[n];
            for (int i = 0; i < n; i++) {
                x1[i] = sc.nextInt();
            }
            for (int i = 0; i < n; i++) {
                y1[i] = sc.nextInt();
            }
            for (int i = 0; i < n; i++) {
                x2[i] = sc.nextInt();
            }
            for (int i = 0; i < n; i++) {
                y2[i] = sc.nextInt();
            }
            int ans = 1;
            for (int i = 0; i < x1.length; i++)// �����Բ�ͬ����Ϊ�������м��㣬ͬʱ����Ӧ�������Ϊ��׼ֵ�������ֵ��
            {
                int num = solve(x1, x2, y1, y2, 0, x1[i], y1[i], x2[i], y2[i]);
                ans = Math.max(num, ans);
                //System.out.println(num);
            }
           // System.out.println("====");
           // System.out.println(ans);
        }
	}

}
