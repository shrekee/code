class Person {
    public int age = 10;
    public Person() {
        System.out.println("初始化年龄: "+ age);
    }
    public int GetAge(int age) {
        this.age = age;
        return this.age;
    }
}

public class test1 {
    public static void main(String[] args) {
        Person Harry = new Person();
        System.out.println("Harry's age is " + Harry.GetAge(12));
        System.out.println("this.age : " + Harry.age);
    }
}
