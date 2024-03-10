package code_java;

class Square {
    private int side;

    public void setSide(int s) {
        this.side = s;
    }

    public int getArea() {
        return this.side * this.side;
    }
}

public class soal7 {
    public static void main(String[] args) {
        int side = 5;
        Square sq = new Square();
        sq.setSide(side);
        System.out.println(sq.getArea());
    }
}
