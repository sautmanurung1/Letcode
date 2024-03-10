package code_java;
public class soal6 {
    public static void main(String[] args) {
        int[] amount = {0,0,0,0,0,0,0,0,0,0};
        int[] number = {1,2,3,4,2,9,3};
        for(int i = 0; i < number.length; i++) {
            amount[number[i]]++;
        }
        System.out.println(amount[1]);
    }
}