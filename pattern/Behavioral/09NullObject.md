# 在空对象模式（Null Object Pattern）中，一个空对象取代 NULL 对象实例的检查。Null 对象不是检查空值，而是反应一个不做任何动作的关系。这样的 Null 对象也可以在数据不可用的时候提供默认的行为。

在空对象模式中，我们创建一个指定各种要执行的操作的抽象类和扩展该类的实体类，还创建一个未对该类做任何实现的空对象类，该空对象类将无缝地使用在需要检查空值的地方。


```java
import org.junit.Test;

public class s22NullObject {
    static abstract class AbstractCustomer {
        protected String name;
        public abstract boolean isNil();
        public abstract String getName();
    }

    static class RealCustomer extends AbstractCustomer {

        public RealCustomer(String name) {
            this.name = name;
        }

        @Override
        public String getName() {
            return name;
        }

        @Override
        public boolean isNil() {
            return false;
        }
    }

    static class NullCustomer extends AbstractCustomer {

        @Override
        public String getName() {
            return "Not Available in Customer Database";
        }

        @Override
        public boolean isNil() {
            return true;
        }
    }

    static class CustomerFactory {

        public static final String[] names = {"Rob", "Joe", "Julie"};

        public static AbstractCustomer getCustomer(String name){
            for (int i = 0; i < names.length; i++) {
                if (names[i].equalsIgnoreCase(name)){
                    return new RealCustomer(name);
                }
            }
            return new NullCustomer();
        }
    }

    @Test
    public void test() throws Exception{
        AbstractCustomer customer1 = CustomerFactory.getCustomer("Rob");
        AbstractCustomer customer2 = CustomerFactory.getCustomer("Bob");
        AbstractCustomer customer3 = CustomerFactory.getCustomer("Julie");
        AbstractCustomer customer4 = CustomerFactory.getCustomer("Laura");

        System.out.println("Customers");
        System.out.println(customer1.getName());
        System.out.println(customer2.getName());
        System.out.println(customer3.getName());
        System.out.println(customer4.getName());
    }
}

```
