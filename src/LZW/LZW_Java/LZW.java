package LZW.LZW_Java;

import java.util.ArrayList;
import java.util.List;

public class LZW {
  public List<Object> encode(String pattern) {
    List<String> alphabetDict = createAlphabetDict(pattern);
    List<String> dict = new ArrayList<>(alphabetDict);
    List<Integer> output = new ArrayList<>();
    String currentPatternToCheck = String.valueOf(pattern.charAt(0));
    int latestFoundIdx = 0;

    for (int i = 1;; i++) {
      int foundIdx = findIn(dict, currentPatternToCheck);

      if (foundIdx == 0) {
        output.add(latestFoundIdx);
        dict.add(currentPatternToCheck);
        currentPatternToCheck = String.valueOf(pattern.charAt(i - 1));
        i--;
      } else if (i == pattern.length()) {

        // in the case were we reach the end of the pattern
        // and the currentPatternToCheck is in the dict so we
        // add its index to the output and add nothing to the dict
        output.add(foundIdx);
        break;
      } else {
        currentPatternToCheck += String.valueOf(pattern.charAt(i));
        latestFoundIdx = foundIdx;
      }
    }

    return List.of(output, alphabetDict);
  }

  public List<String> createAlphabetDict(String pattern) {
    List<String> dict = new ArrayList<>();

    for (int i = 0; i < pattern.length(); i++) {
      String charStr = String.valueOf(pattern.charAt(i));
      if (!dict.contains(charStr)) {
        dict.add(charStr);
      }
    }

    return dict;
  }

  public int findIn(List<String> a, String element) {
    for (int i = 0; i < a.size(); i++) {
      if (a.get(i).equals(element)) {
        return i + 1;
      }
    }
    return 0;
  }

  public String decode(List<Integer> output, List<String> alphabetDict) {
    List<String> dict = alphabetDict;
    String result = alphabetDict.get(output.get(0) - 1);
    String previousStep = result;
    String currentStep = "";

    for (int i = 1; i < output.size(); i++) {
      int index = output.get(i) - 1;
      if (index >= dict.size()) {
        // if the index is not known in the dict yet we construct the
        // unknown by concatenating the previous step and the first
        // symbol in the previous step
        currentStep = String.valueOf(previousStep.charAt(0));
        result += previousStep + currentStep;
        dict.add(previousStep + currentStep.charAt(0));
        previousStep += currentStep;
      } else {
        currentStep = dict.get(index);
        result += currentStep;
        dict.add(previousStep + currentStep.charAt(0));
        previousStep = currentStep;
      }
    }

    return result;
  }
}