package code_java;

public class soal2 {
    public static void main(String[] args) {
        int a = 2, b = 5, c = 7;
        for(;;) {
            if(a > 2 * b) {
                break;
            }
            a += (c-b);
            b--;
            c++;
        }
        System.out.println(c-a);
    }
}
