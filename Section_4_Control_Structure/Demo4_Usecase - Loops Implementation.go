package main
import "fmt"
func demo4(){
	var numTeams, numMembers int
	fmt.Print("Enter the number of teams: ")
	fmt.Scan(&numTeams)
	fmt.Print("Enter the number of members in each team: ")
	fmt.Scan(&numMembers)
	for i:=1;i<=numTeams;i++{
		var teamName string
		fmt.Printf("Enter name of team %d: ", i)
		fmt.Scan(&teamName)
		fmt.Printf("Team: %s\n", teamName)
		for j := 1 ; j<= numMembers;j++{
			var memberName string
			fmt.Printf("Enter member %d name for team%s: ", j, teamName)
			fmt.Scan(&memberName)
			fmt.Printf("-Member: %s\n", memberName)
		}
	}	
}