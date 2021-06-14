package main

import "fmt"

func main() {
	// conditional statements
	// if else 
	// if else if
	// if with relational operators
	// switch statement

	var age int;
	fmt.Print("Enter Your Age: ");
	fmt.Scanln(&age);

	// if-else
	if age >= 18 {
		fmt.Println("You are eligible for driving..");
	} else {
		fmt.Println("You are not  eligible for driving..");
	}

	fmt.Println();
	// if-else-if
	var choose string;
	fmt.Print("Enter Rock, Paper or Scissors: ");
	fmt.Scanln(&choose);

	if choose == "Rock" {
		fmt.Println("You Choosed: " + choose);
	} else if choose == "Paper" {
		fmt.Println("You Choosed: " + choose);
	} else if choose == "Scissors" {
		fmt.Println("You Choosed: " + choose);
	}

	fmt.Println();
	// if else with relational operators 
	var occupation string;
	var salary int;

	fmt.Print("Enter Your Occupation: ");
	fmt.Scanln(&occupation);

	fmt.Print("Enter Your Salary: ");
	fmt.Scanln(&salary);


	if occupation == "Manager" || occupation == "Director" {
		fmt.Println("You are part of the Management..");
	}
	
	if salary >= 9000 && salary <=100000 {
		fmt.Println("You deserve a raise!");

	}

	fmt.Println();
	// switch statement
	var position string;
	fmt.Print("Enter Your Position(north, south, west, east): ");
	fmt.Scanln(&position);

	switch position {
		case "north":
			fmt.Println("You are moving to the north..");
			break;
		case "south":
			fmt.Println("You are moving to the south..");
			break;

		case "west":
			fmt.Println("You are moving to the west..");
			break;

		case "east":
			fmt.Println("You are moving to the east..");
			break;

		default:
			fmt.Println("Please Enter A Valid Position..")
		
	}




}