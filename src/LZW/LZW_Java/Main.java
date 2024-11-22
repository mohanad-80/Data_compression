package LZW.LZW_Java;

import java.io.BufferedReader;
import java.io.BufferedWriter;
import java.io.FileReader;
import java.io.FileWriter;
import java.io.IOException;
import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;
import java.util.stream.Collectors;

public class Main {
  public static void main(String[] args) throws IOException {
    // Test encoding and decoding separately using file I/O
    testEncodingWithFileIO();
    testDecodingWithFileIO();

    // Testing with console output
    consoleTest();
  }

  private static void testEncodingWithFileIO() throws IOException {
    LZW lzw = new LZW();

    // Read input from file
    BufferedReader reader = new BufferedReader(new FileReader("original.txt"));
    String input = reader.readLine();
    reader.close();

    // Encode input
    List<Object> result = lzw.encode(input);
    List<Integer> output = (List<Integer>) result.get(0);
    List<String> alphabetDict = (List<String>) result.get(1);

    // Write encoded data to file
    BufferedWriter writer = new BufferedWriter(new FileWriter("encoded.txt"));
    writer.write(String.join(",", alphabetDict));
    writer.write("\n");
    writer.write(output.stream().map(String::valueOf).collect(Collectors.joining(",")));
    writer.close();
  }

  private static void testDecodingWithFileIO() throws IOException {
    LZW lzw = new LZW();

    // Read encoded data from file
    BufferedReader reader = new BufferedReader(new FileReader("encoded.txt"));
    String alphabetDictStr = reader.readLine();
    String outputStr = reader.readLine();
    reader.close();

    // Parse encoded data
    List<String> alphabetDict = new ArrayList<>(Arrays.asList(alphabetDictStr.split(",")));
    List<Integer> output = Arrays.stream(outputStr.split(","))
        .map(Integer::parseInt)
        .collect(Collectors.toList());

    // Decode output
    String decodedStr = lzw.decode(output, alphabetDict);

    // Write decoded data to file
    BufferedWriter writer = new BufferedWriter(new FileWriter("decoded.txt"));
    writer.write(decodedStr);
    writer.close();
  }

  private static void consoleTest() {
    LZW lzw = new LZW();

    String[] testStrings = {
        "ffcabracadabrarrarradff", "rararbcrarbc", "xyxyxyxyxyxyzzzzzzzz",
        "wabba/wabba/wabba/wabba/woo/woo/woo", "a/bar/array/by/barrayar/bay", "barrayar/bar/by/barrayar/bay"
    };

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

      System.out.println("decoded:  " + lzw.decode(output, alphabetDict));
      System.out.println("original: " + testString);
      System.out.println("=====================================");
    }
  }
}