import java.util.*;
public class Test1 {
    public static void main(String[] args) {
        Set set = new HashSet();
        set.add("abc");
        set.add("cde");
        set.add("efg");
        set.add("abc");
        set.add(5);
        System.out.println("size="+ set.size());
        List list = new ArrayList();
        list.add("abc");
        list.add("aaa");
        list.add("fff");
        set.addAll(list);
        System.out.println("size="+ set.size());
        for (Iterator it = set.iterator(); it.hasNext();) {
            System.out.println("value="+it.next().toString());
        }
        Map m = new HashMap<String, String>();
        m.put("a","b");
        System.out.println("size="+m.size());


    }


}
