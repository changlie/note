```java


package com.huawei.it.demo.ari.btree.v1;

import java.util.ArrayList;
import java.util.LinkedList;
import java.util.List;

/**
 * <p> 这是第一个版本：只实现了插入功能
 * <p>对于一个m阶的b树,假设节点的关键字数量为k:
 * <p>根节点的关键字数量范围: 1 <= k <= m-1
 * <p>非根节点的关键字数量范围: (m-1)/2 <= k <= m-1
 * <p>非叶节点都有k+1个子节点
 * 
 * @author c84112937
 * @since Apr 24, 2020
 */
public class Btree {
    private int factor; // 阶数

    private int maxKeyAmount;

    private Node root; // 根节点

    private boolean printisWithParentInfo;

    public void setPrintisWithParentInfo(boolean printisWithParentInfo) {
        this.printisWithParentInfo = printisWithParentInfo;
    }

    public Btree(int factor) {
        super();
        this.factor = factor;
        this.maxKeyAmount = factor - 1;
        root = new Node(factor);
    }

    public void delete(int v) {
        Node deleteNode = findDeleteNode(v);

        doDelete(deleteNode, v);
    }

    private void doDelete(Node dNode, int v) {
        if (dNode.isLeaf()) {
            if (dNode.keysLen() > dNode.minKeyAmount()) {
                dNode.delKey(v);
            } else {
                Node parent = dNode.getParent();
                List<Data> keys = parent.getKeys();

                for (int i = 0; i < keys.size(); i++) {
                    Data data = keys.get(i);
                    if (data.getLeftChild() == dNode) {
                        if (data.hasRightChild()) {

                        }
                        break;
                    }
                    if (data.getRightChild() == dNode) {
                        break;
                    }
                }
            }
        }
    }

    private Node findDeleteNode(int v) {
        Node currentNode = root;
        // int count = 0;
        baseLoop: while (true) {
            // System.out.println(v + ":findInsertNode#count: " + (++count));
            List<Data> keys = currentNode.getKeys();

            int len = keys.size();
            // System.out.println(v + ":findInsertNode#keys.size: " + len);
            // if (len < 1) {
            // break baseLoop;
            // }
            Data preData = null;
            Node nextNode = null;
            for (int i = 0; i < len; i++) {
                boolean isEnd = i == len - 1;
                Data data = keys.get(i);
                int key = data.getKey();
                // System.out.println(v + ":findInsertNode#data.getKey: " + key);

                if (i == 0 && v < key && data.hasLeftChild()) {
                    // System.out.println(v + ":findInsertNode#b1");
                    nextNode = data.getLeftChild();
                    break;
                } else if (isEnd && v > key && data.hasRightChild()) {
                    nextNode = data.getRightChild();
                    break;
                } else if (v < key && v > preData.getKey()) {
                    // System.out.println(v + ":findInsertNode#b2");
                    if (data.hasLeftChild()) {
                        // System.out.println(v + ":findInsertNode#b3");
                        nextNode = data.getLeftChild();
                        break;
                    } else if (preData.hasRightChild()) {
                        // System.out.println(v + ":findInsertNode#b4");
                        nextNode = data.getRightChild();
                        break;
                    } else {
                        // System.out.println(v + ":findInsertNode#b5");
                        break baseLoop;
                    }
                } else if (v == key) {
                    // System.out.println(v + ":findInsertNode#b6");
                    // return new DeleteNode(i, currentNode);
                    return currentNode;
                } else {
                    // System.out.println(v + ":findInsertNode#b7");
                    preData = data;
                }
            }

            if (nextNode == null) {
                break;
            } else {
                currentNode = nextNode;
            }

            // System.out.println("findInsertNode#loop#end");
        }

        // System.out.println("findInsertNode#finish!");
        return null;
    }

    static class DeleteNode {
        int keysIndex;

        Node node;

        public DeleteNode(int keysIndex, Node node) {
            super();
            this.keysIndex = keysIndex;
            this.node = node;
        }
    }

    public void insert(int v) {
        Node insertNode = findInsertNode(v);
        // System.out.println("after findInsertNode: " + v);
        addAndSplit(insertNode, v);
        // System.out.println("after " + v + " insert");
        // System.out.println("=========================================");
    }

    private void addAndSplit(Node insertNode, int v) {
        System.out.println("start addAndSplit: " + v);
        Data curData = new Data(v, v + "Data");
        while (insertNode != null) {
            // System.out.println("addAndSplit#enter loop: " + v);
            int keysLen = insertNode.addData(curData);
            // System.out.println("insertNode#indexListString: " + insertNode.getIndexListString());

            if (keysLen <= maxKeyAmount) {
                return;
            }

            List<Data> keys = insertNode.getKeys();
            int middleIndex = keys.size() / 2;

            System.out.println("middleIndex: " + middleIndex);
            List<Data> leftDatas = subList(keys, 0, middleIndex);
            System.out.println("leftDatas: " + leftDatas);
            List<Data> rightDatas = subList(keys, middleIndex + 1, keys.size());
            System.out.println("rightDatas: " + rightDatas);
            Node leftNode = new Node(factor);
            leftNode.setKeys(leftDatas);
            leftNode.setParent(insertNode);
            Node rightNode = new Node(factor);
            rightNode.setKeys(rightDatas);
            rightNode.setParent(insertNode);

            // 对分裂出来的左右节点的子节点重新设置父节点
            resetParentNode(leftNode);
            resetParentNode(rightNode);

            Data data = clearExclude(middleIndex, keys);
            data.setLeftChild(leftNode);
            data.setRightChild(rightNode);

            if (insertNode.isRoot()) {
                return;
            }

            Node parent = insertNode.getParent();
            // 消除前一个父子关联
            clearParentChildRelation(parent, insertNode);

            // 对分裂出来的左右节点重新设置父节点
            leftNode.setParent(parent);
            rightNode.setParent(parent);

            insertNode = parent;
            curData = data;
            System.out.println("addAndSplit#loop#end");
        }
    }

    private void resetParentNode(Node node) {
        List<Data> keys = node.getKeys();
        for (Data data : keys) {
            if (data.hasLeftChild()) {
                data.getLeftChild().setParent(node);
            }
            if (data.hasRightChild()) {
                data.getRightChild().setParent(node);
            }
        }
    }

    private void clearParentChildRelation(Node parent, Node child) {
        List<Data> keys = parent.getKeys();

        for (Data data : keys) {
            if (data.getLeftChild() == child) {
                data.setLeftChild(null);
                return;
            }

            if (data.getRightChild() == child) {
                data.setRightChild(null);
                return;
            }
        }
    }

    private List<Data> subList(List<Data> keys, int startIndex, int middleIndex) {
        List<Data> res = new ArrayList<Data>();
        for (int i = startIndex; i < middleIndex; i++) {
            res.add(keys.get(i));
        }
        return res;
    }

    private Node findInsertNode(int v) {
        Node currentNode = root;
        // int count = 0;
        baseLoop: while (true) {
            // System.out.println(v + ":findInsertNode#count: " + (++count));
            List<Data> keys = currentNode.getKeys();

            int len = keys.size();
            // System.out.println(v + ":findInsertNode#keys.size: " + len);
            // if (len < 1) {
            // break baseLoop;
            // }
            Data preData = null;
            Node nextNode = null;
            for (int i = 0; i < len; i++) {
                boolean isEnd = i == len - 1;
                Data data = keys.get(i);
                int key = data.getKey();
                // System.out.println(v + ":findInsertNode#data.getKey: " + key);

                if (i == 0 && v < key && data.hasLeftChild()) {
                    // System.out.println(v + ":findInsertNode#b1");
                    nextNode = data.getLeftChild();
                    break;
                } else if (isEnd && v > key && data.hasRightChild()) {
                    nextNode = data.getRightChild();
                    break;
                } else if (v < key && v > preData.getKey()) {
                    // System.out.println(v + ":findInsertNode#b2");
                    if (data.hasLeftChild()) {
                        // System.out.println(v + ":findInsertNode#b3");
                        nextNode = data.getLeftChild();
                        break;
                    } else if (preData.hasRightChild()) {
                        // System.out.println(v + ":findInsertNode#b4");
                        nextNode = data.getRightChild();
                        break;
                    } else {
                        // System.out.println(v + ":findInsertNode#b5");
                        break baseLoop;
                    }
                } else if (v == key) {
                    // System.out.println(v + ":findInsertNode#b6");
                    throw new RuntimeException("primary key is duplicate: " + v);
                } else {
                    // System.out.println(v + ":findInsertNode#b7");
                    preData = data;
                }
            }

            if (nextNode == null) {
                break;
            } else {
                currentNode = nextNode;
            }

            // System.out.println("findInsertNode#loop#end");
        }

        // System.out.println("findInsertNode#finish!");
        return currentNode;
    }

    private <T> T clearExclude(int middleIndex, List<T> keys) {
        T res = null;
        for (int i = keys.size() - 1; i >= 0; i--) {
            if (i == middleIndex) {
                res = keys.remove(i);
            } else {
                keys.remove(i);
            }
        }
        keys.add(res);
        return res;
    }

    @Override
    public String toString() {
        StringBuilder res = new StringBuilder();
        res.append("Btree [factor=" + factor + ", maxKeyAmount=" + maxKeyAmount + "]").append(System.lineSeparator());

        // List<Node> currentNodes = new LinkedList<Node>();
        // currentNodes.add(root);
        List<Pchild> pList = new LinkedList<Btree.Pchild>();
        Pchild pchild = new Pchild();
        pchild.childs.add(root);
        pList.add(pchild);
        Pchilds pchilds = new Pchilds();
        pchilds.pcs = pList;
        pchilds.childcount = 1;

        while (pchilds != null) {
            // System.out.println("---------------enter loop: " + currentNodes.size());
            Pchilds nextPchilds = new Pchilds();
            for (int x = 0; x < pchilds.pcs.size(); x++) {
                Pchild pc = pchilds.pcs.get(x);
                List<Node> currentNodes = pc.childs;

                for (int i = 0; i < currentNodes.size(); i++) {
                    Node node = currentNodes.get(i);
                    List<Data> keys = node.getKeys();
                    Pchild nextPc = new Pchild();
                    nextPc.parent = node;

                    // res.append("|");
                    for (int j = 0; j < keys.size(); j++) {
                        Data data = keys.get(j);
                        Node leftChild = data.getLeftChild();
                        Node rightChild = data.getRightChild();
                        // Integer key = data.getKey();

                        if (leftChild != null) {
                            nextPc.childs.add(leftChild);
                            nextPchilds.childcount += 1;
                            // System.out.println("toString#leftChild: " + leftChild.getKeys());
                        }
                        if (rightChild != null) {
                            nextPc.childs.add(rightChild);
                            nextPchilds.childcount += 1;
                            // System.out.println("toString#rightChild: " + rightChild.getKeys());
                        }
                        // res.append(key).append((j == keys.size() - 1) ? "|" : ",");
                    }
                    res.append(i + ":" + keys.toString());

                    String nodeSeparator = null;
                    if (printisWithParentInfo) {
                        nodeSeparator = (pc.parent == null ? "{root" : pc.parent) + "  ";
                    } else {
                        nodeSeparator = "<-   ";
                    }

                    res.append(nodeSeparator);

                    nextPchilds.pcs.add(nextPc);
                }

                if (x == pchilds.pcs.size() - 1) {
                    res.append(System.lineSeparator());
                }
            }

            pchilds = nextPchilds.childcount > 0 ? nextPchilds : null;
        }

        return res.toString();
    }

    static class Pchild {
        Node parent;

        List<Node> childs = new LinkedList<Node>();
    }

    static class Pchilds {
        List<Pchild> pcs = new LinkedList<Pchild>();

        int childcount = 0;
    }

}

/////////////////////////////////////////////////////////////////////////////////////////////

package com.huawei.it.demo.ari.btree.v1;

public class BtreeTest {

}

/////////////////////////////////////////////////////////////////////////////////////////////


package com.huawei.it.demo.ari.btree.v1;

public class ChildIndex {

    private int index; // 关键字数组索引

    private boolean left; // 是否为左字节树

    public int getIndex() {
        return index;
    }

    public void setIndex(int index) {
        this.index = index;
    }

    public boolean isLeft() {
        return left;
    }

    public void setLeft(boolean left) {
        this.left = left;
    }

}

/////////////////////////////////////////////////////////////////////////////////////////////


package com.huawei.it.demo.ari.btree.v1;

public class Data {
    private Integer key;

    private Object value;

    private Node leftChild;

    private Node rightChild;

    public Data(Integer key, Object value) {
        super();
        this.key = key;
        this.value = value;
    }

    public Integer getKey() {
        return key;
    }

    public void setKey(Integer key) {
        this.key = key;
    }

    public Object getValue() {
        return value;
    }

    public void setValue(Object value) {
        this.value = value;
    }

    public Node getLeftChild() {
        return leftChild;
    }

    public void setLeftChild(Node leftChild) {
        this.leftChild = leftChild;
    }

    public Node getRightChild() {
        return rightChild;
    }

    public void setRightChild(Node rightChild) {
        this.rightChild = rightChild;
    }

    public boolean hasLeftChild() {
        return this.leftChild != null;
    }

    public boolean hasRightChild() {
        return this.rightChild != null;
    }

    @Override
    public String toString() {
        return "" + key;
    }

}

/////////////////////////////////////////////////////////////////////////////////////////////


package com.huawei.it.demo.ari.btree.v1;

import java.util.ArrayList;
import java.util.List;

public class Demo {

    public static void main(String[] args) {
        // testBtree();

        // Btree btree = new Btree(5);
        // btree.insert(99);
        // btree.insert(99);

        // System.out.println("before end");

        // System.out.println("ending!!!");

        List<Integer> list = new ArrayList<Integer>();
        list.add(-1);
        list.add(0, 998);
        list.add(1, 9527);
        list.add(list.size(), 99999);
        list.add(23);
        list.add(47);
        System.out.println(list);
        // System.out.println(list.subList(0, 2));
        // System.out.println(list.subList(0, 3));
        // System.out.println(list.remove(4));
        // System.out.println(list.remove(3));
        // System.out.println(list.remove(1));
        // System.out.println(list.remove(0));
        // System.out.println(list);
        // System.out.println(5 / 2);
        // System.out.println(7 / 2);
    }

    static void testBtree() {
        Btree btree = new Btree(5);
        // btree.setPrintisWithParentInfo(true);

        for (int i = 0; i < 100; i++) {
            btree.insert(i);
            System.out.println(btree.toString());
            System.out.println("=========================");
        }
    }

}

/////////////////////////////////////////////////////////////////////////////////////////////


package com.huawei.it.demo.ari.btree.v1;

import java.util.LinkedList;
import java.util.List;

// m 必须为单数
public class Node {
    private Node parent;

    private int factor;

    private int maxKeyAmount;

    private List<ChildIndex> childIndexs;

    private List<Data> keys;

    public Node(int factor) {
        super();
        this.factor = factor;
        this.maxKeyAmount = factor - 1;
        childIndexs = new LinkedList<ChildIndex>();
        keys = new LinkedList<Data>();
    }

    public int childCount() {
        return childIndexs.size();
    }

    public boolean hasChild() {
        return childIndexs.size() > 0;
    }

    public int keysLen() {
        return keys.size();
    }

    public int minKeyAmount() {
        if (isRoot()) {
            return 1;
        } else {
            return factor / 2;
        }
    }

    public int maxKeyAmount() {
        return maxKeyAmount;
    }

    public Node getParent() {
        return parent;
    }

    public void setParent(Node parent) {
        this.parent = parent;
    }

    public List<ChildIndex> getChildIndexs() {
        return childIndexs;
    }

    public void setChildIndexs(List<ChildIndex> childIndexs) {
        this.childIndexs = childIndexs;
    }

    public List<Data> getKeys() {
        return keys;
    }

    public void setKeys(List<Data> keys) {
        this.keys = keys;
    }

    public boolean isRoot() {
        return parent == null;
    }

    public boolean isLeaf() {
        for (Data data : keys) {
            if (data.hasLeftChild() || data.hasRightChild()) {
                return false;
            }
        }
        return true;
    }

    // public int addKey(int key) {
    //
    //
    // return addData(d);
    // }

    int addData(Data d) {
        if (keys == null) {
            keys = new LinkedList<Data>();
        }

        int key = d.getKey();
        if (keys.size() == 0) {
            keys.add(d);
            return keys.size();
        }
        // System.out.println("addData:step1:" + key);

        if (keys.size() == 1) {
            Data data = keys.get(0);

            if (key < data.getKey()) {
                keys.add(0, d);
            } else {
                keys.add(d);
            }
            return keys.size();
        }
        // System.out.println("addData:step2:" + key);

        for (int i = 0; (i + 1) < keys.size(); i++) {
            // System.out.println(i + ":addData:step3:" + key);
            int nextIndex = i + 1;
            Data data = keys.get(i);
            Integer tmp = data.getKey();
            Data data1 = keys.get(nextIndex);
            Integer tmp1 = data1.getKey();
            boolean isEnd = nextIndex == keys.size() - 1;
            if ((!isEnd && tmp < key && key < tmp1)) {
                keys.add(nextIndex, d);
                break;
            } else if (tmp == key) {
                keys.remove(i);
                keys.add(i, d);
                break;
            } else if (tmp1 == key) {
                keys.remove(nextIndex);
                keys.add(nextIndex, d);
                break;
            } else if (isEnd && tmp1 < key) {
                keys.add(keys.size(), d);
                break;
            }
        }

        System.out.println("addData finish![" + key);
        return keys.size();
    }

    public int leftKey() {
        Data start = keys.get(0);
        Integer startIndex = start.getKey();
        return startIndex;
    }

    public int rightKey() {
        Data end = keys.get(keysLen() - 1);
        Integer endIndex = end.getKey();
        return endIndex;
    }

    public String getIndexListString() {
        StringBuilder res = new StringBuilder();
        for (Data data : keys) {
            res.append(data.getKey()).append(", ");
        }
        return res.toString();
    }

    @Override
    public String toString() {
        return "N" + keys + "";
    }

    public void delKey(int v) {
        for (int i = 0; i < keysLen(); i++) {
            Data data = keys.get(i);
            int key = data.getKey();
            if (key == v) {
                keys.remove(i);
                return;
            }
        }
    }

}

/////////////////////////////////////////////////////////////////////////////////////////////
```
