/**
 * 
 */
package cn.lsp.wangyi;

/**
 * @author LSP
 *˫������
 *getSize():
 *addFirst(Object):
 *addLast(Object):
 *add(Object):
 *add(index,Object):
 *getValue(index):
 *removeFirst()��
 *removeLast():
 *remove():
 *isEmpty():
 *toString():
 */
public class DoubleLinkedList {  //Test��Ŀ�ڶ����������ͳһ��˫��������
	
	//��Ա�ڲ���
	private class Node{  //�������ֻ࣬��Ҫ������������Ҫ�����ڵ㣬������ý�Node����
		private Node previous = this;
		private Node next = this;
		private Object value;
		
		Node(Object v){
			this.value = v;
		}
		
		public String toString(){ //Node��дObject�е�toString()����������Ȩ�޲���С�ڸ����еķ���Ȩ�ޣ�public��
			return value.toString();
		}
	}
	
	
	//������
	//��Ҫ�õ�Node��Ķ����أ����������޷�����Node�ֻ࣬��ע˫������
	private Node head = new Node(null);
	private int size;
	
	public int getSize(){ //getSize()�Ǹ��������õ�
		return this.size;
	}
	
	/*���صķ���*/
	/*�ڽڵ�ǰ�����һ���ڵ�*/
	private void addBefore(Node newNode, Node node){
		newNode.next = node;
		newNode.previous = node.previous;
		newNode.previous.next = newNode;
		newNode.next.previous = newNode;
		size++;
	}
	
	/*�ڽڵ�������һ���ڵ�*/
	private void addAfter(Node newNode, Node node){
		newNode.previous = node;
		newNode.next = node.next;
		newNode.next.previous = newNode;
		newNode.previous.next = newNode;
		size++;
	}
	
	/*��ȡָ��λ�õĽڵ�*/
	/*��������ֻ�ܴ�ǰ��󰤸��ң�˫������������������򣬸���index��size/2�Ĵ�С�ж��������򣬽�ʡʱ��*/
	private Node getNode(int index) {
		// TODO Auto-generated method stub
		Node node = head;
		
		if(index<0||index>size)
			throw new IndexOutOfBoundsException("IndexOutOfBoundsException");
		
		if(index < size/2){
			//��ǰ�������
			for(int i = 0;i <= index;i++){
				node = node.next;
			}
			return node;
		}else{
			//�Ӻ���ǰ����
			for(int i = size-1;i>=index;i--){
				node = node.previous;
			}
			return node;
		}
		
		//return null;
	}
	
	private void removeNode(Node node) {
		// TODO Auto-generated method stub
		if(size == 0)
			throw new IndexOutOfBoundsException("IndexOutOfBoundsException");
		node.previous.next = node.next;
		node.next.previous = node.previous;
		node.next = null;
		node.previous = null;
		size--;
	}
	
	
	
	
	/*�����ķ���*/
	/*������ͷ���Ԫ��*/
	public boolean addFirst(Object value){
		addAfter(new Node(value),head);
		return false;
	}
	
	/*������β�����Ԫ��*/
	public boolean addLast(Object value){
		addBefore(new Node(value),head);
		return true;
	}

	/*�ں�������Ԫ�أ��൱��addLast*/
	public boolean add(Object value){
		addLast(value);
		return true;
	}
	
	/*��ָ��λ�ò���ڵ�*/
	public boolean add(int index, Object value){
		addAfter(new Node(value), getNode(index));
		return true;
	}

	/*���ָ��λ�ýڵ��ֵ*/
	public Object getValue(int index){
		return getNode(index).value;
	}

	/*ɾ��ͷ�ڵ�*/
	public boolean removeFirst(){
		removeNode(head.next);
		return true;
	}
	
	/*ɾ��ĩβ�ڵ�*/
	public boolean removeLast(){
		removeNode(head.previous);
		return true;
	}
	
	/*ɾ��ָ���ڵ�*/
	public boolean remove(int index){
		removeNode(getNode(index));
		return true;
	}

	/*�ж������Ƿ�Ϊ��*/
	public boolean idEmpty(){
		return size==0;
	}
	
	/*��дtoString()����*/
	public String toString(){
		StringBuilder str = new StringBuilder(">");
		Node node = head;
		for(int i = 0;i < size;i++){
			node = node.next;
			str.append(node.value);
			str.append(";");
		}
		return str.toString();
	}
}
