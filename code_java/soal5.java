package code_java;

public class soal5 {
    public static void main(String[] args) {
        char[] ch = {'Z', 'X', 'C'};
        char tmp;
        for(int i = 0; i < 5; i++) {
            tmp = ch[0];
            ch[0] = ch[2];
            ch[1] = tmp;
            ch[2] = ch[1];
        }
        System.out.println(ch);
    }
}
