import java.util.ArrayList;
import java.util.List;

public class LZ77 {
  private static final int LOOKAHEAD_SIZE = 6;
  private static final int SEARCH_WINDOW_SIZE = 7;

  public List<Triple> encode(String s) {
    List<Triple> triples = new ArrayList<>();
    int lookaheadPtr = 0;
    int searchPtr = 0;

    while (lookaheadPtr < s.length()) {

      // use this if you want to have look-ahead limit
      List<Object> values = findTheLongestMatch(searchPtr, lookaheadPtr, LZ77.LOOKAHEAD_SIZE, s);

      // use this if you do not want look-ahead limit
      // List<Object> values = findTheLongestMatch(searchPtr, lookaheadPtr, s.length()
      // - lookaheadPtr, s);

      int offset = (int) values.get(0);
      int length = (int) values.get(1);
      char codeword = (char) values.get(2);
      triples.add(new Triple(offset, length, codeword));

      lookaheadPtr += length + 1;
      searchPtr += Math.max(0, lookaheadPtr - LZ77.SEARCH_WINDOW_SIZE); // use this if you want to have search limit
    }

    return triples;
  }

  private List<Object> findTheLongestMatch(int searchPtr, int lookaheadPtr, int lookaheadSize, String s) {
    int maxLength = 0;
    int offset = 0;
    char nextCharacter = s.charAt(lookaheadPtr);

    // looking for a match in the search window
    for (; searchPtr < lookaheadPtr; searchPtr++) {

      // not a match, skip
      if (s.charAt(searchPtr) != s.charAt(lookaheadPtr)) {
        continue;
      }

      List<Object> values = findMatchLengthAndCodeword(searchPtr, lookaheadPtr, lookaheadSize, s);
      int length = (int) values.get(0);
      char character = (char) values.get(1);

      if (length >= maxLength) {
        maxLength = length;
        offset = lookaheadPtr - searchPtr;
        nextCharacter = character;
      }
    }

    return List.of(offset, maxLength, nextCharacter);
  }

  private List<Object> findMatchLengthAndCodeword(int p1, int p2, int lookaheadSize, String s) {
    char nextCharacter = '\0'; // empty char
    int length = 0;
    int lookAheadPtr = p2;

    for (int searchWindowPtr = p1;; searchWindowPtr++) {
      // case were the pointer is in both the pattern and look-ahead buffers
      // AND the pointed at chars are equal
      if ((lookAheadPtr < Math.min(s.length() - 1, p2 + lookaheadSize)
          && s.charAt(searchWindowPtr) == s.charAt(lookAheadPtr))) {
        length++;
        lookAheadPtr++;
        continue;
      }

      // case were the pointer is in the pattern buffer
      // but hit the end of the look-ahead buffer so we
      // still can count this char in the length
      // AND the pointed at chars are equal
      if ((lookAheadPtr < s.length() - 1) && s.charAt(searchWindowPtr) == s.charAt(lookAheadPtr)) {
        length++;
      }

      // last case were we return
      // when the pointed at chars are not equal
      // OR we hit the end of the pattern or the look-ahead buffers

      // we check if we are still in the buffer
      // we encode the current char in the triple otherwise we leave it
      // empty since the last char is encoded in the offset and length.
      if (lookAheadPtr <= s.length() - 1) {
        nextCharacter = s.charAt(p2 + length);
      }
      break;
    }

    return List.of(length, nextCharacter);
  }

  public String decode(List<Triple> t) {
    StringBuilder stringBuilder = new StringBuilder();

    for (Triple triple : t) {
      int i = stringBuilder.length() - triple.getOffset();
      int j = triple.getLength();

      while (i >= 0 && j > 0 && stringBuilder.length() > 0) {
        stringBuilder.append(stringBuilder.charAt(i));
        i++;
        j--;
      }
      stringBuilder.append(triple.getCodeword());
    }

    return stringBuilder.toString();
  }
}
