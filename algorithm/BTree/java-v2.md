```java


package com.huawei.it.demo.ari.btree.v2;

import java.util.ArrayList;
import java.util.LinkedList;
import java.util.List;

/**
 * <p>这是第一个版本：只实现了插入功能
 * <p>对于一个m阶的b树,假设节点的关键字数量为k:
 * <p>根节点的关键字数量范围: 1 <= k <= m-1
 * <p>非根节点的关键字数量范围: (m-1)/2 <= k <= m-1
 * <p>非叶节点都有k+1个子节点
 * <p>所有的插入操作都涉及对叶节点的插入操作
 * <p>插入时分裂的中间索引值为: (m-1)/2; 如：  3阶中间索引值为1, 4阶中间索引值为1, 5阶中间索引值为2,,,  
 * <p>所有的删除操作都涉及对叶节点的删除操作
 * 
 * @author c84112937
 * @since Apr 24, 2020
 */
public class Btree {

    private int factor;

    private int splitMiddleIndex;

    private Node root;

    public Btree(int factor) {
        super();
        this.factor = factor;
        splitMiddleIndex = (factor - 1) / 2;

        root = new Node(factor);
    }

    // 所有的插入操作都会对叶节点执行插入操作
    // 所有节点执行插入操作后都要判断是否需要分裂
    // 分裂时,判断当前节点是否为根节点,不是根节点时,需要把分裂出来的中间关键字合并到父节点中
    public void insert(int k) {
        Node insertNode = findInsertNode(k);
        System.out.println("inert: " + insertNode);
        doInsert(insertNode, k);
        System.out.println("after inert: " + insertNode);
    }

    private void doInsert(Node insertNode, int newKey) {
        IData newData = Data.newDataObj(newKey);
        while (true) {
            insertNode.addData(newData);

            System.out.println("insertNode.leMaxKeyNum(): " + insertNode.leMaxKeyNum());
            if (insertNode.leMaxKeyNum()) {// 判断是否需要分裂
                // 不需要分裂直接返回
                return;
            }

            List<Data> datas = insertNode.getDatas();
            List<Data> leftDatas = subList(datas, 0, splitMiddleIndex);
            List<Data> rightDatas = subList(datas, splitMiddleIndex + 1, datas.size());

            Node leftNode = new Node(factor, leftDatas);
            Node rightNode = new Node(factor, rightDatas);

            Data leftLink = Data.newLinkObj(leftNode);
            leftLink.rightBrother = rightNode;
            Data rightLink = Data.newLinkObj(rightNode);
            rightLink.leftBrother = leftNode;

            if (insertNode.isRoot()) {
                List<Data> middleDatas = clearExclude(splitMiddleIndex, datas);

                middleDatas.add(0, leftLink);
                middleDatas.add(rightLink);
                insertNode.setDatas(middleDatas);

                leftNode.parent = insertNode;
                rightNode.parent = insertNode;
                return;
            }

            Data data = datas.get(splitMiddleIndex);
            TupleData tupleData = new TupleData();
            tupleData.setHead(leftLink);
            tupleData.setMiddle(data);
            tupleData.setTail(rightLink);

            newData = tupleData; // for next loop
            insertNode = insertNode.parent; // for next loop

            leftNode.parent = insertNode;
            rightNode.parent = insertNode;
        }
    }

    private List<Data> subList(List<Data> datas, int startIndex, int middleIndex) {
        List<Data> res = new ArrayList<Data>();
        for (int i = startIndex; i < middleIndex; i++) {
            res.add(datas.get(i));
        }
        return res;
    }

    private <T> List<T> clearExclude(int middleIndex, List<T> datas) {
        T res = null;
        for (int i = datas.size() - 1; i >= 0; i--) {
            if (i == middleIndex) {
                res = datas.remove(i);
            } else {
                datas.remove(i);
            }
        }
        datas.add(res);
        return datas;
    }

    private Node findInsertNode(int newKey) {
        Node currentNode = root;
        nodeLoop: while (true) {
            List<Data> datas = currentNode.getDatas();
            boolean isLeaf = currentNode.isLeaf();
            int len = datas.size();

            Node nextNode = null;
            Data preData = null;
            dataLoop: for (int i = 0; i < len; i++) {
                Data data = datas.get(i);
                if (data.isLink()) {
                    continue;
                }
                int key = data.key;

                if (newKey == key) {
                    throw new IllegalArgumentException("primary key is duplicate: " + newKey);
                } else if (isLeaf && ((i == 0 && newKey < key) || (len - 1 == i && newKey > key)
                    || (preData != null && preData.key < newKey && newKey < key))) {
                    break nodeLoop;
                } else if (!isLeaf) {
                    if ((i == 1 && newKey < key) || (preData != null && preData.key < newKey && newKey < key)) {
                        nextNode = datas.get(i - 1).child;
                        break dataLoop;
                    } else if (len - 2 == i && newKey > key) {
                        nextNode = datas.get(i + 1).child;
                        break dataLoop;
                    }
                    continue dataLoop;
                } else {
                    preData = data;
                }
            }

            if (nextNode == null) {
                break;
            } else {
                currentNode = nextNode;
            }
        }

        return currentNode;
    }

    // 一. 所有的删除操作都会对叶节点执行删除操作(遵循最小改动原则)
    // 二. 所有非叶节点的删除操作操作都会转化为对叶节点的删除操作
    // 三. 叶节点操作如下:
    // 1. 最右侧叶节点大于最小关键字数时, 直接删除
    // 2. 最右侧叶节点等于最小关键字数时, 转化为对左兄弟叶节点的删除操作
    // 3. 中间叶节点删除操作转化为左兄弟叶节点的删除操作, 当两者都等于最小关键字数时,删除后进行节点合并
    // 4. 叶节点右兄弟为最右侧节点且右兄弟节点大于最小关键字数时, 删除操作后关键字左移即可, 否则按情况3处理
    // 5. 最右侧或最左侧叶节点等于最小关键字数且其兄弟节点也等于最小关键字时, 删除操作后进行节点合并
    // 四. 叶节点发生节点合并后, 可能会导致非叶节点小于最小关键字数, 进行关键字左移, 右移, 或合并即可.(需要注意父子节点的关系更新)
    public void delete(int k) {

    }

    @Override
    public String toString() {
        StringBuilder res = new StringBuilder();
        res.append("Btree [factor=" + factor + "]\n");
        Pchilds pchilds = Pchilds.newObj(root);
        while (pchilds.childcount > 0) {
            Pchilds nextpclids = new Pchilds();
            for (Pchild pc : pchilds.pcs) {
                for (int i = 0; i < pc.childs.size(); i++) {
                    Node currentNode = pc.childs.get(i);
                    Pchild nextpc = new Pchild();
                    nextpc.parent = currentNode;
                    nextpc.childs = new LinkedList<Node>();
                    for (Data d : currentNode.getDatas()) {
                        if (d.isLink()) {
                            nextpc.childs.add(d.child);
                            nextpclids.childcount += d.child.countKey();
                        }
                    }

                    nextpclids.pcs.add(nextpc);
                    res.append(currentNode.toString()).append("<-   ");
                }
                res.append(System.lineSeparator());
            }

            pchilds = nextpclids;
        }

        return res.toString();
    }

}

class Pchild {
    Node parent;

    List<Node> childs;

    public Pchild() {
        this(null);
    }

    public Pchild(Node node) {
        childs = new LinkedList<Node>();
        if (node != null) {
            childs.add(node);
        }
    }
}

class Pchilds {
    List<Pchild> pcs;

    int childcount;

    public Pchilds() {
        this(new LinkedList<Pchild>(), 0);
    }

    public Pchilds(List<Pchild> pcs, int childcount) {
        this.pcs = pcs;
        this.childcount = childcount;
    }

    public static Pchilds newObj(Node node) {
        List<Pchild> pcs = new LinkedList<Pchild>();
        pcs.add(new Pchild(node));
        Pchilds pchilds = new Pchilds(pcs, node.countKey());
        return pchilds;
    }
}
/////////////////////////////////////////////////////////////////////////////////////////////


package com.huawei.it.demo.ari.btree.v2;

public class Data implements IData {

    public Node child;

    public Node leftBrother;

    public Node rightBrother;

    public Integer key;

    public Object value;

    /** 是否为连结对象  */
    private boolean isLink;

    private Data() {
    }

    public static Data newDataObj(Integer key) {
        Data dn = new Data();
        dn.isLink = false;
        dn.key = key;
        return dn;
    }

    public static Data newLinkObj(Node child) {
        Data dn = new Data();
        dn.isLink = true;
        dn.child = child;
        return dn;
    }

    public boolean isLink() {
        return isLink;
    }

    public boolean isData() {
        return !isLink;
    }

    public boolean isEdge() {
        return isLink && (leftBrother == null || rightBrother == null);
    }

    public boolean isSingle() {
        return true;
    }

    public int key() {
        return key;
    }

    @Override
    public String toString() {
        return "" + key;
    }

}

/////////////////////////////////////////////////////////////////////////////////////////////


package com.huawei.it.demo.ari.btree.v2;

import java.util.LinkedList;
import java.util.List;

public class DoubleData {
    public Data left;

    public Data right;

    public List<Data> list = new LinkedList<Data>();

    public void setLeft(Data left) {
        this.left = left;
        if (list.size() > 0) {
            list.add(0, left);
        } else {
            list.add(left);
        }
    }

    public void setRight(Data right) {
        this.right = right;
        list.add(right);
    }
}

/////////////////////////////////////////////////////////////////////////////////////////////


package com.huawei.it.demo.ari.btree.v2;

public interface IData {

    boolean isSingle();

    int key();
}

/////////////////////////////////////////////////////////////////////////////////////////////


package com.huawei.it.demo.ari.btree.v2;

import java.util.LinkedList;
import java.util.List;

public class Node {

    public Node parent;

    public Data parentLeftData;

    public Data parentRightData;

    private int maxKeyNum;

    private int commonMinKeyNum;

    private List<Data> datas;

    public Node(int factor) {
        this(factor, null, new LinkedList<Data>());
    }

    public Node(int factor, List<Data> datas) {
        this(factor, null, datas);
    }

    public Node(int factor, Node parentNode, List<Data> datas) {
        super();
        this.parent = parentNode;
        this.datas = datas;
        maxKeyNum = factor - 1;
        commonMinKeyNum = (factor - 1) / 2;
    }

    public void addData(IData newData) {
        int size = datas.size();

        if (size == 0 && newData.isSingle()) {
            datas.add((Data) newData);
        }

        boolean isLeaf = isLeaf();
        int newKey = newData.key();
        Data preData = null;

        for (int i = 0; i < size; i++) {
            Data data = datas.get(i);
            if (data.isLink()) {
                continue;
            }
            int key = data.key;
            if ((preData == null && newKey < key) || (preData != null && preData.key < newKey && newKey < key)) {
                datasInsert(i, newData);
                return;
            } else if (((isLeaf && size - 1 == i) || (!isLeaf && size - 2 == i)) && newKey > key) {
                datasInsert(size, newData);
                return;
            } else {
                continue;
            }
        }
    }

    private void datasInsert(int index, IData d) {
        if (d.isSingle()) {
            Data data = (Data) d;
            datas.add(index, data);
        } else {
            // eg.
            // link data link data link
            // 0___ 1___ 2___ 3___ 4___
            // index=3
            int i = index - 1;
            Data removeData = datas.remove(i);
            TupleData tupleData = (TupleData) d;

            // eg.
            // link data data link
            // 0___ 1___ 2___ 3___
            // index=3
            tupleData.tail.rightBrother = removeData.rightBrother;
            if (removeData.rightBrother != null) {
                Data data = datas.get(index);
                data.leftBrother = tupleData.tail.child;
            }
            tupleData.head.leftBrother = removeData.leftBrother;
            if (removeData.leftBrother != null) {
                Data data = datas.get(index - 3);
                data.rightBrother = tupleData.head.child;
            }
            datas.addAll(i, tupleData.list);
        }
    }

    /**
     * 关键字数量小于等于最大值，用于判断当前节点是否可以分裂
     * @author c84112937
     * @since Apr 24, 2020
     * @return
     */
    public boolean leMaxKeyNum() {
        return countKey() <= maxKeyNum;
    }

    /**
     * 关键字数量大于最小值
     * <p>用于删除时，判断当前节点是否可以提供关键字给父节点
     * @author c84112937
     * @since Apr 24, 2020
     * @return
     */
    public boolean isRich() {
        return countKey() > commonMinKeyNum;
    }

    /**
     * 关键字数量小于最小值
     * <p>用于当前节点删除节点后，判断关键字移动或节点合并(节点合并也是关键字移动的一种)
     * @author c84112937
     * @since Apr 24, 2020
     * @return
     */
    public boolean isPoor() {
        return countKey() < commonMinKeyNum;
    }

    public Data popupLeafDataTail() {
        return datas.remove(datas.size() - 1);
    }

    public DoubleData popupDoubleDataTail() {
        DoubleData res = new DoubleData();
        res.setLeft(datas.remove(datas.size() - 1));
        res.setRight(datas.remove(datas.size() - 1));
        return res;
    }

    public TupleData popupTupleDataTail() {
        TupleData res = new TupleData();
        res.setHead(datas.remove(datas.size() - 1));
        res.setMiddle(datas.remove(datas.size() - 1));
        res.setTail(datas.remove(datas.size() - 1));
        return res;
    }

    public int countKey() {
        if (datas.size() == 0) {
            return 0;
        }
        if (isLeaf()) {
            return datas.size();
        }
        return (datas.size() - 1) / 2;
    }

    public boolean isRoot() {
        return parent == null;
    }

    public boolean isLeaf() {
        for (Data data : datas) {
            if (data.isLink()) {
                return false;
            }
        }
        return true;
    }

    public List<Data> getDatas() {
        return datas;
    }

    public void setDatas(List<Data> datas) {
        this.datas = datas;
    }

    @Override
    public String toString() {
        StringBuilder res = new StringBuilder();
        res.append("[");

        boolean isInit = false;
        for (int i = 0; i < datas.size(); i++) {
            Data data = datas.get(i);
            if (data.isLink()) {
                continue;
            }
            if (isInit) {
                res.append(",");
            } else {
                isInit = true;
            }
            res.append(datas.get(i).key);
        }
        res.append("]");
        return res.toString();
    }

}

/////////////////////////////////////////////////////////////////////////////////////////////


package com.huawei.it.demo.ari.btree.v2;

public class testBTreeV2 {

    public static void main(String[] args) {
        Btree tree = new Btree(5);
        for (int i = 1; i < 15; i++) {
            tree.insert(i);
            System.out.println(tree);
            System.out.println("===========================");
        }
    }
}

/////////////////////////////////////////////////////////////////////////////////////////////


package com.huawei.it.demo.ari.btree.v2;

import java.util.LinkedList;
import java.util.List;

public class TupleData implements IData {
    public Data head;

    public Data middle;

    public Data tail;

    public List<Data> list = new LinkedList<Data>();

    // public TupleData(Data head, Data middle, Data tail) {
    // super();
    // this.head = head;
    // this.middle = middle;
    // this.tail = tail;
    // list.add(head);
    // list.add(middle);
    // list.add(tail);
    // }

    public void setHead(Data head) {
        this.head = head;
        list.add(head);
    }

    public void setMiddle(Data middle) {
        if (list.size() < 1) {
            throw new RuntimeException("TupleData list.size() < 1");
        }
        this.middle = middle;
        list.add(middle);
    }

    public void setTail(Data tail) {
        if (list.size() < 2) {
            throw new RuntimeException("TupleData list.size() < 2");
        }
        this.tail = tail;
        list.add(tail);
    }

    public boolean isSingle() {
        return false;
    }

    public int key() {
        return middle.key;
    }
}

/////////////////////////////////////////////////////////////////////////////////////////////
```
