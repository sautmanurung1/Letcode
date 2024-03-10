package code_java;

import java.util.Scanner;

public class soal9 {
    public static void main(String[] args) {
        Scanner scan = new Scanner(System.in);
        String first = scan.nextLine();
        if (first.charAt(0) == 'A'){
            System.out.print("A ");
            if (first.charAt(1) == 'B') {
                System.out.print("B");
            }
        } else if(first.charAt(0) == 'B') {
            System.out.print("B ");
            if (first.charAt(1) == 'C') {
                System.out.print("C");
            }
        } else if(first == "AB") {
            System.out.print("AB");
        } else {
            System.out.print("-");
        }
        scan.close();
    }
}
