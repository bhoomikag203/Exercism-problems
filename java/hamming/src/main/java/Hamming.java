class Hamming {
	String leftStrand, rightStrand;

    Hamming(String leftStrand, String rightStrand) {
    	this.leftStrand = leftStrand;
    	this.rightStrand = rightStrand;
    	validateInputs(leftStrand, rightStrand);
    }

    void validateInputs(String leftStrand, String rightStrand){
    	if( leftStrand.length() != rightStrand.length())
    		throw new IllegalArgumentException("leftStrand and rightStrand must be of equal length.");
    }

    int getHammingDistance() {
    	int hammingDistance = 0;
    	for(int i =  0; i < leftStrand.length(); i++){
    		if(leftStrand.charAt(i) != rightStrand.charAt(i))
    			hammingDistance++;
    	}
    	return hammingDistance;    
    }

}
