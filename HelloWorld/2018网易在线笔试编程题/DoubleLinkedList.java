/**
 * 
 */
package cn.lsp.wangyi;

/**
 * @author LSP
 *双向链表
 *getSize():
 *addFirst(Object):
 *addLast(Object):
 *add(Object):
 *add(index,Object):
 *getValue(index):
 *removeFirst()：
 *removeLast():
 *remove():
 *isEmpty():
 *toString():
 */
public class DoubleLinkedList {  //Test项目内都可以用这个统一的双向链表类
	
	//成员内部类
	private class Node{  //其他的类，只需要操作链表，不需要操作节点，所以最好将Node隐藏
		private Node previous = this;
		private Node next = this;
		private Object value;
		
		Node(Object v){
			this.value = v;
		}
		
		public String toString(){ //Node重写Object中的toString()方法，访问权限不能小于父类中的访问权限（public）
			return value.toString();
		}
	}
	
	
	//链表类
	//需要用到Node类的都隐藏，让其他类无法操作Node类，只关注双链表本身
	private Node head = new Node(null);
	private int size;
	
	public int getSize(){ //getSize()是给其他类用的
		return this.size;
	}
	
	/*隐藏的方法*/
	/*在节点前面插入一个节点*/
	private void addBefore(Node newNode, Node node){
		newNode.next = node;
		newNode.previous = node.previous;
		newNode.previous.next = newNode;
		newNode.next.previous = newNode;
		size++;
	}
	
	/*在节点后面插入一个节点*/
	private void addAfter(Node newNode, Node node){
		newNode.previous = node;
		newNode.next = node.next;
		newNode.next.previous = newNode;
		newNode.previous.next = newNode;
		size++;
	}
	
	/*获取指定位置的节点*/
	/*单向链表只能从前向后挨个找，双向链表可以有两个方向，根据index和size/2的大小判断搜索方向，节省时间*/
	private Node getNode(int index) {
		// TODO Auto-generated method stub
		Node node = head;
		
		if(index<0||index>size)
			throw new IndexOutOfBoundsException("IndexOutOfBoundsException");
		
		if(index < size/2){
			//从前向后搜索
			for(int i = 0;i <= index;i++){
				node = node.next;
			}
			return node;
		}else{
			//从后向前搜索
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
	
	
	
	
	/*公开的方法*/
	/*在链表头添加元素*/
	public boolean addFirst(Object value){
		addAfter(new Node(value),head);
		return false;
	}
	
	/*在链表尾部添加元素*/
	public boolean addLast(Object value){
		addBefore(new Node(value),head);
		return true;
	}

	/*在后面增加元素，相当于addLast*/
	public boolean add(Object value){
		addLast(value);
		return true;
	}
	
	/*在指定位置插入节点*/
	public boolean add(int index, Object value){
		addAfter(new Node(value), getNode(index));
		return true;
	}

	/*获得指定位置节点的值*/
	public Object getValue(int index){
		return getNode(index).value;
	}

	/*删除头节点*/
	public boolean removeFirst(){
		removeNode(head.next);
		return true;
	}
	
	/*删除末尾节点*/
	public boolean removeLast(){
		removeNode(head.previous);
		return true;
	}
	
	/*删除指定节点*/
	public boolean remove(int index){
		removeNode(getNode(index));
		return true;
	}

	/*判断链表是否为空*/
	public boolean idEmpty(){
		return size==0;
	}
	
	/*重写toString()函数*/
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
