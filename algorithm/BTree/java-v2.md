```java


package com.huawei.it.demo.ari.btree.v2;

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

    public Btree(int factor) {
        super();
        this.factor = factor;
        splitMiddleIndex = (factor - 1) / 2;
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
        return "Btree [factor=" + factor + "]";
    }
}
/////////////////////////////////////////////////////////////////////////////////////////////


package com.huawei.it.demo.ari.btree.v2;

public class Data {

    public Node child;

    public Node leftBrother;

    public Node rightBrother;

    public Integer key;

    public Object value;

    /** 是否为连结对象  */
    private boolean isLink;

    private Data() {
    }

    public Data newDataObj() {
        Data dn = new Data();
        dn.isLink = false;
        return dn;
    }

    public Data newLinkObj() {
        Data dn = new Data();
        dn.isLink = true;
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

import java.util.LinkedList;
import java.util.List;

public class Node {

    private Node parentNode;

    public Data parentLeftData;

    public Data parentRightData;

    private int maxKeyNum;

    private int commonMinKeyNum;

    private List<Data> datas;

    public Node(int factor) {
        this(factor, null, new LinkedList<Data>());
    }

    public Node(int factor, Node parentNode, List<Data> datas) {
        super();
        this.parentNode = parentNode;
        this.datas = datas;
        maxKeyNum = factor - 1;
        commonMinKeyNum = (factor - 1) / 2;
    }

    /**
     * 关键字数量大于最大值，用于判断当前节点是否可以分裂
     * @author c84112937
     * @since Apr 24, 2020
     * @return
     */
    public boolean isMaxKeyNum() {
        return countKey() > maxKeyNum;
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
        return (datas.size() - 1) / 2;
    }

    public boolean isRoot() {
        return parentNode == null;
    }

    public boolean isLeaf() {
        for (Data data : datas) {
            if (data.isLink()) {
                return false;
            }
        }
        return true;
    }
}

/////////////////////////////////////////////////////////////////////////////////////////////


package com.huawei.it.demo.ari.btree.v2;

import java.util.LinkedList;
import java.util.List;

public class TupleData {
    public Data head;

    public Data middle;

    public Data tail;

    public List<Data> list = new LinkedList<Data>();

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
}

/////////////////////////////////////////////////////////////////////////////////////////////
```
