public class Triple {
  private int offset;
  private int length;
  private char codeword;

  public Triple(int offset, int length, char codeword) {
    this.offset = offset;
    this.length = length;
    this.codeword = codeword;
  }

  public int getOffset() {
    return offset;
  }

  public int getLength() {
    return length;
  }

  public char getCodeword() {
    return codeword;
  }

  public void setOffset(int offset) {
    this.offset = offset;
  }

  public void setLength(int length) {
    this.length = length;
  }

  public void setCodeword(char codeword) {
    this.codeword = codeword;
  }

  // Optional: Override toString for easier debugging
  @Override
  public String toString() {
    return "{" + offset + " " + length + " " + codeword + "}";
  }
}
