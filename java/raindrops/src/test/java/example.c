#include<stdio.h>
void main()
{
	int num1, num2;
	printf("Enter two numbers:\n");
	scanf("%d \n %d", &num1, &num2)
	findOuccurence(num1, num2);
}

void findOuccurence(int num1, int num2)
{
	int count = 0;
	char number1 = (char) num1;
	char number2 = (char) num2;
	for(int i = 0; i < strlen(number1); i++){
		if(number1[i] == number2){
			count++;
		}
	}
	printf("%d", count)
}