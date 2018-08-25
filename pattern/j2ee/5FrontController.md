# 前端控制器模式（Front Controller Pattern）是用来提供一个集中的请求处理机制，所有的请求都将由一个单一的处理程序处理。该处理程序可以做认证/授权/记录日志，或者跟踪请求，然后把请求传给相应的处理程序。以下是这种设计模式的实体。

前端控制器（Front Controller） - 处理应用程序所有类型请求的单个处理程序，应用程序可以是基于 web 的应用程序，也可以是基于桌面的应用程序。
调度器（Dispatcher） - 前端控制器可能使用一个调度器对象来调度请求到相应的具体处理程序。
视图（View） - 视图是为请求而创建的对象。



```java

public class j2ee05FrontController {
    class HomeView {
        public void show(){
            System.out.println("Displaying Home Page");
        }
    }

    class StudentView {
        public void show(){
            System.out.println("Displaying Student Page");
        }
    }

    class Dispatcher {
        private StudentView studentView;
        private HomeView homeView;
        public Dispatcher(){
            studentView = new StudentView();
            homeView = new HomeView();
        }

        public void dispatch(String request){
            if(request.equalsIgnoreCase("STUDENT")){
                studentView.show();
            }else{
                homeView.show();
            }
        }
    }

    class FrontController {

        private Dispatcher dispatcher;

        public FrontController(){
            dispatcher = new Dispatcher();
        }

        private boolean isAuthenticUser(){
            System.out.println("User is authenticated successfully.");
            return true;
        }

        private void trackRequest(String request){
            System.out.println("Page requested: " + request);
        }

        public void dispatchRequest(String request){
            //记录每一个请求
            trackRequest(request);
            //对用户进行身份验证
            if(isAuthenticUser()){
                dispatcher.dispatch(request);
            }
            System.out.println("=====================> request finished.");
        }
    }

    @Test
    public void test() throws Exception{
        FrontController frontController = new FrontController();
        frontController.dispatchRequest("HOME");
        frontController.dispatchRequest("STUDENT");
    }
}
```
