public class EnumDemo {
    public static void main(String[] args) {
        Day day = Day.MONDAY;
        //System.out.println("day is :%d",day.MONDAY);
        System.out.println(day.MONDAY);
        //day.MONDAY
    }

}

enum Day {
    MONDAY, TUESDAY, WENDESDAY, THURSDAY, FRIDAY, STATURDAY,SUNDAY
}
