import java.util.List;

public class Main {

  public static void main(String[] args) {
    LZ77 lz77 = new LZ77();

    String testString = "ffcabracadabrarrarradff";
    List<Triple> result = lz77.encode(testString);
    for (Triple triple : result) {
      System.out.println(triple.toString());
    }
    System.out.println("decoded: " + lz77.decode(result));
    System.out.println("oiginal: " + testString);
    System.out.println("=====================================");
    testString = "rararbcrarbc";
    result = lz77.encode(testString);
    for (Triple triple : result) {
      System.out.println(triple.toString());
    }
    System.out.println("decoded: " + lz77.decode(result));
    System.out.println("oiginal: " + testString);
    System.out.println("=====================================");
    testString = "xyxyxyxyxyxyzzzzzzzz";
    result = lz77.encode(testString);
    for (Triple triple : result) {
      System.out.println(triple.toString());
    }
    System.out.println("decoded: " + lz77.decode(result));
    System.out.println("oiginal: " + testString);
    System.out.println("=====================================");
  }
}
