class RaindropConverter {

    String convert(int number) {
    	int factors[] = {};
    	for(int i = 1; i <= number; i++){
    		if(number % i == 0){
    			factors[i] = i;
    		}
    	}

    	for(int j = 0; j < factors.length; j++){
    		if(factors[j] == 3)
    			return "Pling";
    		else if (factors[j] == 5)
   				return "Plang";
   			else if (factors[j] == 7)
   				return "Plong";  
   			else 
   				continue;	
    	}
    	String num = Integer.toString(number);
   		return num;	
    }

}
