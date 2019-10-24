class ReverseString {
	String reverseString;
    String reverse(String inputString) {
    	if (inputString == null || inputString.isEmpty()) {
    		return inputString;
    	}
    	reverseString = "";
    	for(int i = inputString.length() - 1; i >= 0 ; i--){
    		reverseString = reverseString + inputString.charAt(i); 
    	}
    	return reverseString;
    }
  
}