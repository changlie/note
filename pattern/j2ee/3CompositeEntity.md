# 组合实体模式（Composite Entity Pattern）用在 EJB 持久化机制中。一个组合实体是一个 EJB 实体 bean，代表了对象的图解。当更新一个组合实体时，内部依赖对象 beans 会自动更新，因为它们是由 EJB 实体 bean 管理的。以下是组合实体 bean 的参与者。

组合实体（Composite Entity） - 它是主要的实体 bean。它可以是粗粒的，或者可以包含一个粗粒度对象，用于持续生命周期。  
粗粒度对象（Coarse-Grained Object） - 该对象包含依赖对象。它有自己的生命周期，也能管理依赖对象的生命周期。  
依赖对象（Dependent Object） - 依赖对象是一个持续生命周期依赖于粗粒度对象的对象。  
策略（Strategies） - 策略表示如何实现组合实体。  



```java
import org.junit.Test;

public class j2ee03CompositeEntity {
    class DependentObject1 {

        private String data;

        public void setData(String data){
            this.data = data;
        }

        public String getData(){
            return data;
        }
    }

    class DependentObject2 {

        private String data;

        public void setData(String data){
            this.data = data;
        }

        public String getData(){
            return data;
        }
    }

    class CoarseGrainedObject {
        DependentObject1 do1 = new DependentObject1();
        DependentObject2 do2 = new DependentObject2();

        public void setData(String data1, String data2){
            do1.setData(data1);
            do2.setData(data2);
        }

        public String[] getData(){
            return new String[] {do1.getData(),do2.getData()};
        }
    }

    class CompositeEntity {
        private CoarseGrainedObject cgo = new CoarseGrainedObject();

        public void setData(String data1, String data2){
            cgo.setData(data1, data2);
        }

        public String[] getData(){
            return cgo.getData();
        }
    }


    class Client {
        private CompositeEntity compositeEntity = new CompositeEntity();

        public void printData(){
            for (int i = 0; i < compositeEntity.getData().length; i++) {
                System.out.print(", Data: " + compositeEntity.getData()[i]);
            }
            System.out.println();
        }

        public void setData(String data1, String data2){
            compositeEntity.setData(data1, data2);
        }
    }

    @Test
    public void test() throws Exception{
        Client client = new Client();
        client.setData("Test", "Data");
        client.printData();
        client.setData("Second Test22", "Data222");
        client.printData();
    }
}

```
