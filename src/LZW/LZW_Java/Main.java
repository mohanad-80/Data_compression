package LZW.LZW_Java;

import java.util.List;

public class Main {
  public static void main(String[] args) {
    LZW lzw = new LZW();

    String[] testStrings = { "ffcabracadabrarrarradff", "rararbcrarbc", "xyxyxyxyxyxyzzzzzzzz",
        "wabba/wabba/wabba/wabba/woo/woo/woo", "a/bar/array/by/barrayar/bay", "barrayar/bar/by/barrayar/bay" };

    for (String testString : testStrings) {
      List<Object> result = lzw.encode(testString);
      List<Integer> output = (List<Integer>) result.get(0);
      List<String> alphabetDict = (List<String>) result.get(1);
      for (String c : alphabetDict) {
        System.out.print(c + ", ");
      }
      System.out.println();
      for (int i : output) {
        System.out.print(i + ", ");
      }
      System.out.println();
      System.out.println("decoded: " + lzw.decode(output, alphabetDict));
      System.out.println("oiginal: " + testString);
      System.out.println("=====================================");
    }
  }
}