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
public class Btree implements Config {

    private int factor;

    private boolean printWithParent;

    private boolean isDebug;

    private int splitMiddleIndex;

    // eg.
    // 0 is link, 1 is data
    // datas ==> m->splitMiddleIndex: internalSplitMiddleIndex
    // 0101010 ==> 3->[1]{(3-1)/2}: 3=1*2+1
    // 010101010 ==> 4->[1]{(4-1)/2}: 3=1*2+1
    // 01010101010 ==> 5->[2]{(5-1)/2}: 5=2*2+1
    // 0101010101010 ==> 6->[2]{(6-1)/2}: 5=2*2+1
    // 010101010101010 ==> 7->[3]{(7-1)/2}: 7=3*2+1
    // 01010101010101010 ==> 8->[3]{(8-1)/2}: 7=3*2+1
    private int internalSplitMiddleIndex;

    private Node root;

    public Btree(int factor) {
        LogUtil.config = this;
        this.factor = factor;
        splitMiddleIndex = (factor - 1) / 2;
        internalSplitMiddleIndex = splitMiddleIndex * 2 + 1;

        root = new Node(factor);
    }

    public void setDebug(boolean isDebug) {
        this.isDebug = isDebug;
    }

    public void setPrintWithParent(boolean printWithParent) {
        this.printWithParent = printWithParent;
    }

    // 所有的插入操作都会对叶节点执行插入操作
    // 所有节点执行插入操作后都要判断是否需要分裂
    // 分裂时,判断当前节点是否为根节点,不是根节点时,需要把分裂出来的中间关键字合并到父节点中
    public void insert(int k) {
        Node insertNode = findInsertNode(k);
        LogUtil.a("inert " + k + " to " + insertNode + " isLeaf: " + insertNode.isLeaf());
        doInsert(insertNode, k);
    }

    private void doInsert(Node insertNode, int newKey) {
        IData newData = Data.newDataObj(newKey);
        while (true) {
            insertNode.addData(newData);

            LogUtil.a("doInsert# %s leMaxKeyNum: %s", insertNode, insertNode.leMaxKeyNum());
            LogUtil.a("doInsert# keyNum: %s maxKeyNum: %s cminKeyNum: %s",
                insertNode.countKey(),
                insertNode.maxKeyNum,
                insertNode.commonMinKeyNum);
            LogUtil.a("doInsert# " + "datas: " + insertNode.getDatas().size() + ", " + insertNode.getDatas());
            if (insertNode.leMaxKeyNum()) {// 判断是否需要分裂
                // 不需要分裂直接返回
                return;
            }

            int middleIndex = insertNode.isLeaf() ? splitMiddleIndex : internalSplitMiddleIndex;
            List<Data> datas = insertNode.getDatas();
            List<Data> leftDatas = subList(datas, 0, middleIndex);
            List<Data> rightDatas = subList(datas, middleIndex + 1, datas.size());

            Node leftNode = new Node(factor, leftDatas);
            Node rightNode = new Node(factor, rightDatas);

            // 为新节点的子节点重新指定父节点
            resetParent(leftDatas, leftNode);
            resetParent(rightDatas, rightNode);

            Data leftLink = Data.newLinkObj(leftNode);
            leftLink.rightBrother = rightNode;
            Data rightLink = Data.newLinkObj(rightNode);
            rightLink.leftBrother = leftNode;

            if (insertNode.isRoot()) {
                List<Data> middleDatas = clearExclude(middleIndex, datas);

                middleDatas.add(0, leftLink);
                middleDatas.add(rightLink);
                insertNode.setDatas(middleDatas);

                leftNode.parent = insertNode;
                rightNode.parent = insertNode;
                return;
            }

            Data data = datas.get(middleIndex);
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

    private void resetParent(List<Data> datas, Node newParent) {
        for (Data data : datas) {
            if (data.isLink()) {
                data.child.parent = newParent;
            }
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
                    LogUtil.a("findInsertNode#link: " + i);
                    continue;
                }
                int key = data.key;

                if (newKey == key) {
                    throw new IllegalArgumentException("primary key is duplicate: " + newKey);
                } else if (isLeaf && ((i == 0 && newKey < key) || (len - 1 == i && newKey > key)
                    || (preData != null && preData.key < newKey && newKey < key))) {
                    LogUtil.a("findInsertNode#find it: %s", currentNode);
                    break nodeLoop;
                } else if (!isLeaf) {
                    LogUtil.a("findInsertNode#newKey: %s, preData: %s, key: %s <-- checkNotLeaf", newKey, preData, key);
                    if ((i == 1 && newKey < key) || (preData != null && preData.key < newKey && newKey < key)) {
                        nextNode = datas.get(i - 1).child;
                        break dataLoop;
                    } else if (len - 2 == i && newKey > key) {
                        nextNode = datas.get(i + 1).child;
                        break dataLoop;
                    }
                }
                LogUtil.a("findInsertNode#step: else");
                preData = data;
            }

            if (nextNode == null) {
                break;
            } else {
                currentNode = nextNode;
            }
        }
        if (!currentNode.isLeaf()) {
            throw new RuntimeException("findInsertNode fail!!! " + newKey);
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
        DelInfo info = findDeleteNode(k);
        if (info == null) {
            System.out.println("nothing for delete");
            return;
        }

        doDelete(info, k);
    }

    private void doDelete(DelInfo info, int k) {
        if (info.delNode.isLeaf()) {
            leafNodeDataDel(info.delNode, info.dataIndex);
        } else {

        }
    }

    // 三. 叶节点删除操作如下:
    // 1. 最右侧叶节点大于最小关键字数时, 直接删除
    // 2. 最右侧或最左侧叶节点等于最小关键字数且兄弟节点大于最小关键字数时, 删除操作后进行关键字右移或左移
    // 3. 中间叶节点删除操作转化为左兄弟叶节点的删除操作, 当两者都等于最小关键字数时,删除后进行节点合并
    // 4. 叶节点右兄弟为最右侧节点且右兄弟节点大于最小关键字数时, 删除操作后关键字左移即可, 否则按情况3处理
    // 5. 最右侧或最左侧叶节点等于最小关键字数且其兄弟节点也等于最小关键字时, 删除操作后进行节点合并
    private DelInfo leafNodeDataDel(Node delNode, int dataIndex) {

        // 删除节点为根节点时，直接删除后返回。
        // 1. 最右侧叶节点大于最小关键字数时, 直接删除
        if (delNode.isRoot() || delNode.isRich()) {
            delNode.delete(dataIndex);
            return null;
        }

        LinkInfo pl = getParentLink(delNode);

        if (pl.index == pl.datas.size() - 1 && delNode.isPoor()) {
            Node leftBrother = pl.datas.get(pl.index - 2).child;
            if (leftBrother.isRich()) {
                keyDeleteAndRightMove(delNode, dataIndex, pl, leftBrother);
                return null;
            } else if (leftBrother.isPoor()) {
                // 5. 最右侧或最左侧叶节点等于最小关键字数且其兄弟节点也等于最小关键字时, 删除操作后进行节点合并
                delNode.delete(dataIndex);

                Data midData = pl.datas.get(pl.index - 1);
                leftBrother.tailInsert(midData);
                leftBrother.tailInsert(delNode.getDatas());

                delNode.parent.popupDoubleDataTail();

                if (delNode.parent.isRoot()) {
                    leftBrother.parent = null;
                }
                return null;
            }
        }

        if (pl.index == 0 && delNode.isPoor()) {
            Node rightBrother = pl.datas.get(pl.index + 2).child;
            if (rightBrother.isRich()) {
                keyDeleteAndLeftMove(delNode, dataIndex, pl, rightBrother);
                return null;
            } else if (rightBrother.isPoor()) {
                // 5. 最右侧或最左侧叶节点等于最小关键字数且其兄弟节点也等于最小关键字时, 删除操作后进行节点合并
                delNode.delete(dataIndex);

                Data midData = pl.datas.get(pl.index + 1);
                rightBrother.headInsert(midData);
                rightBrother.headInsert(delNode.getDatas());

                delNode.parent.popupDoubleDataHead();

                if (delNode.parent.isRoot()) {
                    rightBrother.parent = null;
                }
                return null;
            }
        }

        if (0 < pl.index && pl.index < pl.datas.size() - 1) {
            int rightLastTwo = pl.datas.size() - 3;
            Node leftBrother = pl.datas.get(pl.index - 2).child;
            if (pl.index < rightLastTwo) {
                // 3. 中间叶节点删除操作时，左兄弟叶节点大于最小关键字数，删除操作后关键字右移; 当两者都等于最小关键字数时,删除后进行节点合并
                if (leftBrother.isRich()) {
                    keyDeleteAndRightMove(delNode, dataIndex, pl, leftBrother);
                    return null;
                } else if (leftBrother.isPoor()) {
                    keyDeleteAndMergeNodeInternal(delNode, dataIndex, pl, leftBrother);
                    return null;
                }
            }
            Node rightBrother = pl.datas.get(pl.index + 2).child;
            if (pl.index == rightLastTwo) {
                // 4. 叶节点右兄弟为最右侧节点且右兄弟节点大于最小关键字数时, 删除操作后关键字左移即可, 否则按情况3处理
                if (leftBrother.isRich()) {
                    keyDeleteAndRightMove(delNode, dataIndex, pl, leftBrother);
                    return null;
                } else if (leftBrother.isPoor() && rightBrother.isRich()) {
                    keyDeleteAndLeftMove(delNode, dataIndex, pl, rightBrother);
                    return null;
                } else if (leftBrother.isPoor() && rightBrother.isPoor()) {
                    keyDeleteAndMergeNodeInternal(delNode, dataIndex, pl, leftBrother);
                    return null;
                }
            }
        }

        throw new RuntimeException("leafNodeDataDel: unknow delete " + delNode);
    }

    private void keyDeleteAndLeftMove(Node delNode, int dataIndex, LinkInfo pl, Node rightBrother) {
        delNode.delete(dataIndex);
        Data rightData = rightBrother.popupLeafDataHead();
        Data midData = pl.datas.get(pl.index + 1);

        delNode.parent.replace(pl.index + 1, rightData);
        delNode.tailInsert(midData);
    }

    private void keyDeleteAndMergeNodeInternal(Node delNode, int dataIndex, LinkInfo pl, Node leftBrother) {
        delNode.delete(dataIndex);

        Data midData = pl.datas.get(pl.index - 1);
        leftBrother.tailInsert(midData);
        leftBrother.tailInsert(delNode.getDatas());

        delNode.parent.popupDoubleData(pl.index - 1);
    }

    private void keyDeleteAndRightMove(Node delNode, int dataIndex, LinkInfo pl, Node leftBrother) {
        delNode.delete(dataIndex);

        Data leftData = leftBrother.popupLeafDataTail();
        Data midData = pl.datas.get(pl.index - 1);

        delNode.parent.replace(pl.index - 1, leftData);
        delNode.headInsert(midData);
    }

    private LinkInfo getParentLink(Node delNode) {
        List<Data> datas = delNode.parent.getDatas();
        int len = datas.size();
        for (int i = 0; i < len; i++) {
            Data data = datas.get(i);
            if (data.isLink() && data.child == delNode) {
                return new LinkInfo(i, data, datas);
            }
        }
        throw new RuntimeException("getParentLink Exception: " + delNode);
    }

    private DelInfo findDeleteNode(int delKey) {
        Node currentNode = root;
        int dataIndex = -1;
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

                if (delKey == key) {
                    dataIndex = i;
                    break nodeLoop;
                } else if (isLeaf && ((i == 0 && delKey < key) || (len - 1 == i && delKey > key)
                    || (preData != null && preData.key < delKey && delKey < key))) {
                    currentNode = null;
                    break nodeLoop;
                } else if (!isLeaf) {
                    if ((i == 1 && delKey < key) || (preData != null && preData.key < delKey && delKey < key)) {
                        nextNode = datas.get(i - 1).child;
                        break dataLoop;
                    } else if (len - 2 == i && delKey > key) {
                        nextNode = datas.get(i + 1).child;
                        break dataLoop;
                    }
                }
                preData = data;
            }

            if (nextNode == null) {
                currentNode = null;
                break;
            } else {
                currentNode = nextNode;
            }
        }

        if (currentNode == null) {
            return null;
        }

        return new DelInfo(currentNode, dataIndex);
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

                    Node p = currentNode.parent;
                    String nodeSeperator = printWithParent ? "p" + (p == null ? "{root}" : p.print()) + "  " : "<-   ";

                    res.append(currentNode.toString()).append(nodeSeperator);
                }
            }
            res.append(System.lineSeparator());

            pchilds = nextpclids;
        }

        return res.toString();
    }

    @Override
    public boolean isDebug() {
        return isDebug;
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

public interface Config {

    boolean isDebug();
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

public class DelInfo {

    public Node delNode;

    public int dataIndex;

    public DelInfo(Node delNode, int dataIndex) {
        super();
        this.delNode = delNode;
        this.dataIndex = dataIndex;
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

import java.util.List;

public class LinkInfo {
    public int index;

    public Data link;

    public List<Data> datas;

    public LinkInfo(int index, Data link, List<Data> datas) {
        super();
        this.index = index;
        this.link = link;
        this.datas = datas;
    }

}

/////////////////////////////////////////////////////////////////////////////////////////////


package com.huawei.it.demo.ari.btree.v2;

import java.util.ArrayList;
import java.util.List;
import java.util.Locale;

public class LogUtil {

    static Config config;

    /**
     * 用于记录性能消耗
     * @author c84112937
     * @since 2019年10月18日
     * @param tag 标签，用于查找日志
     * @param startTime 开始时间的毫秒值
     */
    public static void timeLog(String tag, long startTime) {
        StringBuilder sb = new StringBuilder();
        sb.append("consumeLog#").append(tag).append(" spend: ").append(System.currentTimeMillis() - startTime).append(
            " ms");
        System.out.println(sb.toString());
    }

    /**
     * 自定义日志输出
     * <blockquote>For example,
     * <pre>{
     *     String s = null;
     *     LogUtil.batch()
     *         .newLine("w3TaskHandler#deleteToDoTask")
     *         .newLine("deleteUrl: %s", "http://www.xxxx.com")
     *         .newLine("taskId: %s", UUID.randomUUID().toString())
     *         .newLine("null-print: %s", s)
     *         .newLine("list-print: %s", Arrays.asList("CN", "EN", false))
     *         .print();
     * 日志打印: 2019-10-24 10:49:14 [main] INFO  LogUtil:107 -  
     *                w3TaskHandler#deleteToDoTask 
     *                deleteUrl: http://www.xxxx.com 
     *                taskId: e57f57ed-8a12-4707-8680-70f15c499a0c 
     *                null-print: null 
     *                list-print: [CN, EN, false]
     * }</pre></blockquote>
     * @author c84112937
     * @since 2019年10月24日
     * @return 日志输出实体类。
     */
    public static StrLog batch() {
        return new StrLog();
    }

    /**
     * 自定义日志输出
     * <blockquote>For example,
     * <pre>{
     *     print("time consume: %s; age: %s; list: %s", "99s", 99, Arrays.asList("tom", "jerry"));
     * // 日志输出: 2019-10-24 10:37:13 [main] INFO  LogUtil:65 -  
     * //              time consume: 99s; age: 99; list: [tom, jerry]
     * }</pre></blockquote>
     * @author c84112937
     * @since 2019年10月24日
     * @param s 消息主体
     * @param args 消息变量
     */
    public static void a(String s, Object... args) {
        int len = args.length;
        for (int i = 0; i < len; i++) {
            Object obj = args[i];
            args[i] = (obj == null) ? "null" : obj.toString();
        }
        if (config.isDebug())
            System.out.println(String.format(Locale.ENGLISH, s, args));
    }

    /**
     * 批量输出日志的实体类。 
     * @author c84112937
     * @since 2019年10月24日
     */
    public static class StrLog {
        private StringBuilder sb = new StringBuilder();

        private List<Object> args = new ArrayList<>();

        /**
         * 新起一行日志输出。
         * @author c84112937
         * @since 2019年10月24日
         * @param s 消息主体
         * @param objects 消息参数
         * @return 日志实体对象
         */
        public StrLog newLine(String s, Object... objects) {
            if (sb.length() > 0) {
                sb.append(" %n ");
            }
            sb.append(s);

            for (Object obj : objects) {
                obj = (obj == null) ? "null" : obj.toString();
                args.add(obj);
            }
            return this;
        }

        /**
         * 进行日志打印。
         * @author c84112937
         * @since 2019年10月24日
         */
        public void print() {
            if (config.isDebug())
                System.out.println(String.format(Locale.ENGLISH, sb.toString(), args.toArray()));
        }
    }
}

/////////////////////////////////////////////////////////////////////////////////////////////


package com.huawei.it.demo.ari.btree.v2;

import java.util.LinkedList;
import java.util.List;

public class Node {

    public Node parent;

    public Data parentLeftData;

    public Data parentRightData;

    public int maxKeyNum;

    public int commonMinKeyNum;

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

        fixEdgeData();
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
                break;
            } else if (((isLeaf && size - 1 == i) || (!isLeaf && size - 2 == i)) && newKey > key) {
                datasInsert(size, newData);
                break;
            } else {
                continue;
            }
        }

        fixEdgeData();
    }

    private void datasInsert(int index, IData d) {
        LogUtil.a("datasInsert#isLeaf: %s  index: %s  newData: %s  old datas: %s", isLeaf(), index, d, datas);
        if (d.isSingle()) {
            Data data = (Data) d;
            datas.add(index, data);
        } else {

            // eg.
            // link data link data link
            // 0___ 1___ 2___ 3___ 4___
            // index=3
            int i = index - 1;
            boolean isEnd = datas.size() == index;
            boolean isStart = 0 == i;

            Data removeData = datas.remove(i);
            TupleData tupleData = (TupleData) d;

            // eg.
            // link data data link
            // 0___ 1___ 2___ 3___
            // index=3
            LogUtil.a("datasInsert#removeData.rightBrother: %s", removeData.rightBrother);
            if (isEnd) {
                tupleData.tail.rightBrother = null;
            } else {
                tupleData.tail.rightBrother = removeData.rightBrother;
                Data data = datas.get(index);
                data.leftBrother = tupleData.tail.child;
            }
            LogUtil.a("datasInsert#removeData.leftBrother: %s", removeData.leftBrother);
            if (isStart) {
                tupleData.head.leftBrother = null;
            } else {
                tupleData.head.leftBrother = removeData.leftBrother;
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
     * 关键字数量小于等于最小值
     * <p>用于当前节点删除节点后，判断关键字移动或节点合并(节点合并也是关键字移动的一种)
     * @author c84112937
     * @since Apr 24, 2020
     * @return
     */
    public boolean isPoor() {
        return countKey() <= commonMinKeyNum;
    }

    public Data popupLeafDataTail() {
        Data data = datas.remove(datas.size() - 1);
        fixEdgeData();
        return data;
    }

    public Data popupLeafDataHead() {
        Data data = datas.remove(0);
        return data;
    }

    public DoubleData popupDoubleDataTail() {
        DoubleData res = new DoubleData();
        res.setLeft(datas.remove(datas.size() - 1));
        res.setRight(datas.remove(datas.size() - 1));
        fixEdgeData();
        return res;
    }

    public DoubleData popupDoubleDataHead() {
        DoubleData res = new DoubleData();
        res.setLeft(datas.remove(0));
        res.setRight(datas.remove(0));
        fixEdgeData();
        return res;
    }

    public TupleData popupTupleDataTail() {
        TupleData res = new TupleData();
        res.setHead(datas.remove(datas.size() - 1));
        res.setMiddle(datas.remove(datas.size() - 1));
        res.setTail(datas.remove(datas.size() - 1));
        fixEdgeData();
        return res;
    }

    public void popupDoubleData(int index) {
        datas.remove(index);
        datas.remove(index);
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

        fixEdgeData();
    }

    @Override
    public String toString() {
        return printWithCurlyBraces();
    }

    public String printWithCurlyBraces() {
        return print4Internal(true);
    }

    public String print() {
        return print4Internal(false);
    }

    private String print4Internal(boolean hasCurlyBraces) {
        StringBuilder res = new StringBuilder();
        if (hasCurlyBraces)
            res.append("[");

        boolean isInit = false;
        for (int i = 0; i < datas.size(); i++) {
            Data data = datas.get(i);
            if (data.isLink()) {
                continue;
            }
            if (isInit) {
                res.append(hasCurlyBraces ? "," : "_");
            } else {
                isInit = true;
            }
            res.append(datas.get(i).key);
        }
        if (hasCurlyBraces)
            res.append("]");
        return res.toString();
    }

    public void fixEdgeData() {
        if (isLeaf() && datas.size() < 1) {
            return;
        }
        datas.get(0).leftBrother = null;
        datas.get(datas.size() - 1).rightBrother = null;
    }

    public void replace(int i, Data newData) {
        Data oldData = datas.get(i);
        oldData.key = newData.key;
        oldData.value = newData.value;
    }

    public Data delete(int index) {
        return datas.remove(index);
    }

    public void headInsert(Data newData) {
        datas.add(0, newData);
    }

    public void headInsert(List<Data> newDatas) {
        datas.addAll(0, newDatas);
    }

    public void tailInsert(Data newData) {
        datas.add(newData);
    }

    public void tailInsert(List<Data> newDatas) {
        datas.addAll(newDatas);
    }

}

/////////////////////////////////////////////////////////////////////////////////////////////


package com.huawei.it.demo.ari.btree.v2;

import java.util.HashSet;
import java.util.Set;

public class testBTreeV2 {

    public static void main(String[] args) {
        randomTest();
        // reShowTest();
    }

    static void reShowTest() {
        int[] inputs = {250, 49, 674, 654, 901, 334, 263, 434, 467, 929, 183, 404, 469, 395, 697, 35, 339, 672, 435,
            149, 949, 279};
        Btree tree = new Btree(5);
        // tree.setPrintWithParent(true);
        // tree.setDebug(true);
        for (int i = 0; i < inputs.length; i++) {
            tree.insert(inputs[i]);
            System.out.println(tree);
            System.out.println("===========================");
            // if (inputs[i] == 307)
            // break;
        }
    }

    static void randomTest() {
        Btree tree = new Btree(5);
        // tree.setPrintWithParent(true);
        // tree.setDebug(true);

        // StringBuilder res = new StringBuilder();
        // res.append("inputs: ");
        Set<Integer> cache = new HashSet<Integer>();
        try {
            long start = System.currentTimeMillis();
            for (int i = 1; i < 100000; i++) {
                int input = getRandom(cache);
                tree.insert(input);
                // res.append(input + ", ");
                // System.out.println(i + " times: ");
            }
            LogUtil.timeLog("tree insert", start);
            System.out.println(tree);
            System.out.println("===========================");

        } finally {
            System.out.println(" len: " + cache.size() + ",  input: " + cache);
        }
    }

    static int getRandom(Set<Integer> cache) {
        int input;
        do {
            input = (int) (Math.random() * 1000000);
        } while (cache.contains(input));
        cache.add(input);
        return input;
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

    @Override
    public String toString() {
        return "" + list;
    }

}

/////////////////////////////////////////////////////////////////////////////////////////////
```
