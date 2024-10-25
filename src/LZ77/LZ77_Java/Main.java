import java.util.List;

public class Main {

  public static void main(String[] args) {
    LZ77 lz77 = new LZ77();

    String[] testStrings = { "ffcabracadabrarrarradff", "rararbcrarbc", "xyxyxyxyxyxyzzzzzzzz" };

    for (String testString : testStrings) {
      List<Triple> result = lz77.encode(testString);
      for (Triple triple : result) {
        System.out.println(triple.toString());
      }
      System.out.println("decoded: " + lz77.decode(result));
      System.out.println("oiginal: " + testString);
      System.out.println("=====================================");
    }
  }
}
